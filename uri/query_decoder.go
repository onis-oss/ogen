package uri

import (
	"strings"

	"github.com/go-faster/errors"
)

type QueryStyle string

const (
	QueryStyleForm           QueryStyle = "form"
	QueryStyleSpaceDelimited QueryStyle = "spaceDelimited"
	QueryStylePipeDelimited  QueryStyle = "pipeDelimited"
	QueryStyleDeepObject     QueryStyle = "deepObject"
)

type QueryDecoder struct {
	param string
	src   []string // r.URL.Query()["param"]

	style   QueryStyle // immutable
	explode bool       // immutable
}

type QueryDecoderConfig struct {
	Param   string
	Values  []string
	Style   QueryStyle
	Explode bool
}

func NewQueryDecoder(cfg QueryDecoderConfig) *QueryDecoder {
	if len(cfg.Values) == 0 {
		panic("unreachable")
	}

	return &QueryDecoder{
		param:   cfg.Param,
		src:     cfg.Values,
		style:   cfg.Style,
		explode: cfg.Explode,
	}
}

func (d *QueryDecoder) DecodeValue() (string, error) {
	switch d.style {
	case QueryStyleForm:
		if len(d.src) != 1 {
			return "", errors.New("multiple params")
		}
		return d.src[0], nil
	case QueryStyleSpaceDelimited,
		QueryStylePipeDelimited,
		QueryStyleDeepObject:
		return "", errors.Errorf("style %q cannot be used for primitive values", d.style)
	default:
		panic("unreachable")
	}
}

func (d *QueryDecoder) DecodeArray(f func(d Decoder) error) error {
	if len(d.src) < 1 {
		return errors.New("empty array")
	}

	switch d.style {
	case QueryStyleForm:
		if d.explode {
			for _, item := range d.src {
				if err := f(&constval{item}); err != nil {
					return err
				}
			}
			return nil
		}

		if len(d.src) != 1 {
			return errors.New("invalid value")
		}

		for _, item := range strings.Split(d.src[0], ",") {
			if err := f(&constval{item}); err != nil {
				return err
			}
		}

		return nil

	case QueryStyleSpaceDelimited:
		if d.explode {
			for _, item := range d.src {
				if err := f(&constval{item}); err != nil {
					return err
				}
			}
			return nil
		}

		if len(d.src) != 1 {
			return errors.New("invalid value")
		}

		return errors.New("spaceDelimited with explode: false not supported")

	case QueryStylePipeDelimited:
		if d.explode {
			for _, item := range d.src {
				if err := f(&constval{item}); err != nil {
					return err
				}
			}
			return nil
		}

		if len(d.src) != 1 {
			return errors.New("invalid value")
		}

		for _, item := range strings.Split(d.src[0], "|") {
			if err := f(&constval{item}); err != nil {
				return err
			}
		}

		return nil

	case QueryStyleDeepObject:
		return errors.Errorf("style %q cannot be used for arrays", d.style)

	default:
		panic("unreachable")
	}
}

func (d *QueryDecoder) DecodeFields(f func(name string, d Decoder) error) error {
	adapter := func(name, value string) error { return f(name, &constval{value}) }
	switch d.style {
	case QueryStyleForm:
		if d.explode {
			for _, v := range d.src {
				if strings.Count(v, "=") != 1 {
					return errors.Errorf("invalid value: %q", v)
				}

				s := strings.Split(v, "=")
				if err := adapter(s[0], s[1]); err != nil {
					return err
				}
			}
			return nil
		}

		if len(d.src) > 1 {
			return errors.New("multiple values passed")
		}

		cur := &cursor{src: []rune(d.src[0])}
		param, err := cur.readUntil('=')
		if err != nil {
			return err
		}

		if param != d.param {
			return errors.Errorf("invalid param name: %q", param)
		}

		return decodeObject(cur, ',', ',', adapter)

	case QueryStyleSpaceDelimited:
		panic("object cannot have spaceDelimited style")

	case QueryStylePipeDelimited:
		panic("object cannot have pipeDelimited style")

	case QueryStyleDeepObject:
		if !d.explode {
			panic("invalid deepObject style configuration")
		}

		cur := &cursor{}
		for _, v := range d.src {
			cur.src = []rune(v)
			cur.pos = 0

			pname, err := cur.readUntil('[')
			if err != nil {
				return err
			}

			if pname != d.param {
				return errors.Errorf("invalid param name: %q", pname)
			}

			key, err := cur.readUntil(']')
			if err != nil {
				return err
			}

			if !cur.eat('=') {
				return errors.New("invalid value")
			}

			val, err := cur.readAll()
			if err != nil {
				return err
			}

			if err := adapter(key, val); err != nil {
				return err
			}
		}
		return nil

	default:
		panic("unreachable")
	}
}
