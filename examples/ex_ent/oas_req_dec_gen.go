// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/validate"
)

func decodeCreatePetRequest(r *http.Request, span trace.Span) (req CreatePetReq, err error) {
	switch ct := r.Header.Get("Content-Type"); ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, validate.ErrBodyRequired
		}

		var request CreatePetReq
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, validate.ErrBodyRequired
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode CreatePet:application/json request")
		}
		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}

func decodeCreatePetCategoriesRequest(r *http.Request, span trace.Span) (req CreatePetCategoriesReq, err error) {
	switch ct := r.Header.Get("Content-Type"); ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, validate.ErrBodyRequired
		}

		var request CreatePetCategoriesReq
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, validate.ErrBodyRequired
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode CreatePetCategories:application/json request")
		}
		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}

func decodeCreatePetFriendsRequest(r *http.Request, span trace.Span) (req CreatePetFriendsReq, err error) {
	switch ct := r.Header.Get("Content-Type"); ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, validate.ErrBodyRequired
		}

		var request CreatePetFriendsReq
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, validate.ErrBodyRequired
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode CreatePetFriends:application/json request")
		}
		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}

func decodeCreatePetOwnerRequest(r *http.Request, span trace.Span) (req CreatePetOwnerReq, err error) {
	switch ct := r.Header.Get("Content-Type"); ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, validate.ErrBodyRequired
		}

		var request CreatePetOwnerReq
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, validate.ErrBodyRequired
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode CreatePetOwner:application/json request")
		}
		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}

func decodeUpdatePetRequest(r *http.Request, span trace.Span) (req UpdatePetReq, err error) {
	switch ct := r.Header.Get("Content-Type"); ct {
	case "application/json":
		if r.ContentLength == 0 {
			return req, validate.ErrBodyRequired
		}

		var request UpdatePetReq
		buf := getBuf()
		defer putBuf(buf)
		written, err := io.Copy(buf, r.Body)
		if err != nil {
			return req, err
		}

		if written == 0 {
			return req, validate.ErrBodyRequired
		}

		d := jx.GetDecoder()
		defer jx.PutDecoder(d)
		d.ResetBytes(buf.Bytes())
		if err := func() error {
			if err := request.Decode(d); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return req, errors.Wrap(err, "decode UpdatePet:application/json request")
		}
		return request, nil
	default:
		return req, validate.InvalidContentType(ct)
	}
}
