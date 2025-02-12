// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
)

// Encode implements json.Marshaler.
func (s Book) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s Book) encodeFields(e *jx.Encoder) {
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.MediaID.Set {
			e.FieldStart("media_id")
			s.MediaID.Encode(e)
		}
	}
	{
		if s.Images.Set {
			e.FieldStart("images")
			s.Images.Encode(e)
		}
	}
	{
		if s.Title.Set {
			e.FieldStart("title")
			s.Title.Encode(e)
		}
	}
	{
		if s.Tags != nil {
			e.FieldStart("tags")
			e.ArrStart()
			for _, elem := range s.Tags {
				elem.Encode(e)
			}
			e.ArrEnd()
		}
	}
	{
		if s.Scanlator.Set {
			e.FieldStart("scanlator")
			s.Scanlator.Encode(e)
		}
	}
	{
		if s.UploadDate.Set {
			e.FieldStart("upload_date")
			s.UploadDate.Encode(e)
		}
	}
	{
		if s.NumPages.Set {
			e.FieldStart("num_pages")
			s.NumPages.Encode(e)
		}
	}
	{
		if s.NumFavorites.Set {
			e.FieldStart("num_favorites")
			s.NumFavorites.Encode(e)
		}
	}
}

var jsonFieldsNameOfBook = [9]string{
	0: "id",
	1: "media_id",
	2: "images",
	3: "title",
	4: "tags",
	5: "scanlator",
	6: "upload_date",
	7: "num_pages",
	8: "num_favorites",
}

// Decode decodes Book from json.
func (s *Book) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Book to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "media_id":
			if err := func() error {
				s.MediaID.Reset()
				if err := s.MediaID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"media_id\"")
			}
		case "images":
			if err := func() error {
				s.Images.Reset()
				if err := s.Images.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"images\"")
			}
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "tags":
			if err := func() error {
				s.Tags = make([]Tag, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Tag
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Tags = append(s.Tags, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"tags\"")
			}
		case "scanlator":
			if err := func() error {
				s.Scanlator.Reset()
				if err := s.Scanlator.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"scanlator\"")
			}
		case "upload_date":
			if err := func() error {
				s.UploadDate.Reset()
				if err := s.UploadDate.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"upload_date\"")
			}
		case "num_pages":
			if err := func() error {
				s.NumPages.Reset()
				if err := s.NumPages.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"num_pages\"")
			}
		case "num_favorites":
			if err := func() error {
				s.NumFavorites.Reset()
				if err := s.NumFavorites.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"num_favorites\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Book")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s Book) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Book) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s Image) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s Image) encodeFields(e *jx.Encoder) {
	{
		if s.T.Set {
			e.FieldStart("t")
			s.T.Encode(e)
		}
	}
	{
		if s.W.Set {
			e.FieldStart("w")
			s.W.Encode(e)
		}
	}
	{
		if s.H.Set {
			e.FieldStart("h")
			s.H.Encode(e)
		}
	}
}

var jsonFieldsNameOfImage = [3]string{
	0: "t",
	1: "w",
	2: "h",
}

// Decode decodes Image from json.
func (s *Image) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Image to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "t":
			if err := func() error {
				s.T.Reset()
				if err := s.T.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"t\"")
			}
		case "w":
			if err := func() error {
				s.W.Reset()
				if err := s.W.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"w\"")
			}
		case "h":
			if err := func() error {
				s.H.Reset()
				if err := s.H.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"h\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Image")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s Image) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Image) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s Images) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s Images) encodeFields(e *jx.Encoder) {
	{
		if s.Pages != nil {
			e.FieldStart("pages")
			e.ArrStart()
			for _, elem := range s.Pages {
				elem.Encode(e)
			}
			e.ArrEnd()
		}
	}
	{
		if s.Cover.Set {
			e.FieldStart("cover")
			s.Cover.Encode(e)
		}
	}
	{
		if s.Thumbnail.Set {
			e.FieldStart("thumbnail")
			s.Thumbnail.Encode(e)
		}
	}
}

var jsonFieldsNameOfImages = [3]string{
	0: "pages",
	1: "cover",
	2: "thumbnail",
}

// Decode decodes Images from json.
func (s *Images) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Images to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "pages":
			if err := func() error {
				s.Pages = make([]Image, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Image
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Pages = append(s.Pages, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"pages\"")
			}
		case "cover":
			if err := func() error {
				s.Cover.Reset()
				if err := s.Cover.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"cover\"")
			}
		case "thumbnail":
			if err := func() error {
				s.Thumbnail.Reset()
				if err := s.Thumbnail.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"thumbnail\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Images")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s Images) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Images) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Image as json.
func (o OptImage) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes Image from json.
func (o *OptImage) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptImage to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptImage) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptImage) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Images as json.
func (o OptImages) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes Images from json.
func (o *OptImages) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptImages to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptImages) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptImages) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes int as json.
func (o OptInt) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Int(int(o.Value))
}

// Decode decodes int from json.
func (o *OptInt) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptInt to nil")
	}
	o.Set = true
	v, err := d.Int()
	if err != nil {
		return err
	}
	o.Value = int(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptInt) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptInt) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes TagType as json.
func (o OptTagType) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes TagType from json.
func (o *OptTagType) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptTagType to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptTagType) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptTagType) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Title as json.
func (o OptTitle) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	o.Value.Encode(e)
}

// Decode decodes Title from json.
func (o *OptTitle) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptTitle to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptTitle) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptTitle) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes SearchByTagIDOKApplicationJSON as json.
func (s SearchByTagIDOKApplicationJSON) Encode(e *jx.Encoder) {
	unwrapped := []SearchResponse(s)
	e.ArrStart()
	for _, elem := range unwrapped {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes SearchByTagIDOKApplicationJSON from json.
func (s *SearchByTagIDOKApplicationJSON) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SearchByTagIDOKApplicationJSON to nil")
	}
	var unwrapped []SearchResponse
	if err := func() error {
		unwrapped = make([]SearchResponse, 0)
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem SearchResponse
			if err := elem.Decode(d); err != nil {
				return err
			}
			unwrapped = append(unwrapped, elem)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = SearchByTagIDOKApplicationJSON(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s SearchByTagIDOKApplicationJSON) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SearchByTagIDOKApplicationJSON) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes SearchOKApplicationJSON as json.
func (s SearchOKApplicationJSON) Encode(e *jx.Encoder) {
	unwrapped := []SearchResponse(s)
	e.ArrStart()
	for _, elem := range unwrapped {
		elem.Encode(e)
	}
	e.ArrEnd()
}

// Decode decodes SearchOKApplicationJSON from json.
func (s *SearchOKApplicationJSON) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SearchOKApplicationJSON to nil")
	}
	var unwrapped []SearchResponse
	if err := func() error {
		unwrapped = make([]SearchResponse, 0)
		if err := d.Arr(func(d *jx.Decoder) error {
			var elem SearchResponse
			if err := elem.Decode(d); err != nil {
				return err
			}
			unwrapped = append(unwrapped, elem)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = SearchOKApplicationJSON(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s SearchOKApplicationJSON) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SearchOKApplicationJSON) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s SearchResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s SearchResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Result != nil {
			e.FieldStart("result")
			e.ArrStart()
			for _, elem := range s.Result {
				elem.Encode(e)
			}
			e.ArrEnd()
		}
	}
	{
		if s.NumPages.Set {
			e.FieldStart("num_pages")
			s.NumPages.Encode(e)
		}
	}
	{
		if s.PerPage.Set {
			e.FieldStart("per_page")
			s.PerPage.Encode(e)
		}
	}
}

var jsonFieldsNameOfSearchResponse = [3]string{
	0: "result",
	1: "num_pages",
	2: "per_page",
}

// Decode decodes SearchResponse from json.
func (s *SearchResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SearchResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "result":
			if err := func() error {
				s.Result = make([]Book, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Book
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Result = append(s.Result, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"result\"")
			}
		case "num_pages":
			if err := func() error {
				s.NumPages.Reset()
				if err := s.NumPages.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"num_pages\"")
			}
		case "per_page":
			if err := func() error {
				s.PerPage.Reset()
				if err := s.PerPage.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"per_page\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SearchResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s SearchResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SearchResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s Tag) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s Tag) encodeFields(e *jx.Encoder) {
	{
		if s.ID.Set {
			e.FieldStart("id")
			s.ID.Encode(e)
		}
	}
	{
		if s.Type.Set {
			e.FieldStart("type")
			s.Type.Encode(e)
		}
	}
	{
		if s.Name.Set {
			e.FieldStart("name")
			s.Name.Encode(e)
		}
	}
	{
		if s.URL.Set {
			e.FieldStart("url")
			s.URL.Encode(e)
		}
	}
	{
		if s.Count.Set {
			e.FieldStart("count")
			s.Count.Encode(e)
		}
	}
}

var jsonFieldsNameOfTag = [5]string{
	0: "id",
	1: "type",
	2: "name",
	3: "url",
	4: "count",
}

// Decode decodes Tag from json.
func (s *Tag) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Tag to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			if err := func() error {
				s.ID.Reset()
				if err := s.ID.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "type":
			if err := func() error {
				s.Type.Reset()
				if err := s.Type.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"type\"")
			}
		case "name":
			if err := func() error {
				s.Name.Reset()
				if err := s.Name.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "url":
			if err := func() error {
				s.URL.Reset()
				if err := s.URL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"url\"")
			}
		case "count":
			if err := func() error {
				s.Count.Reset()
				if err := s.Count.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"count\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Tag")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s Tag) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Tag) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes TagType as json.
func (s TagType) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes TagType from json.
func (s *TagType) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode TagType to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch TagType(v) {
	case TagTypeParody:
		*s = TagTypeParody
	case TagTypeCharacter:
		*s = TagTypeCharacter
	case TagTypeTag:
		*s = TagTypeTag
	case TagTypeArtist:
		*s = TagTypeArtist
	case TagTypeGroup:
		*s = TagTypeGroup
	case TagTypeCategory:
		*s = TagTypeCategory
	case TagTypeLanguage:
		*s = TagTypeLanguage
	default:
		*s = TagType(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s TagType) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *TagType) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s Title) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s Title) encodeFields(e *jx.Encoder) {
	{
		if s.English.Set {
			e.FieldStart("english")
			s.English.Encode(e)
		}
	}
	{
		if s.Japanese.Set {
			e.FieldStart("japanese")
			s.Japanese.Encode(e)
		}
	}
	{
		if s.Pretty.Set {
			e.FieldStart("pretty")
			s.Pretty.Encode(e)
		}
	}
}

var jsonFieldsNameOfTitle = [3]string{
	0: "english",
	1: "japanese",
	2: "pretty",
}

// Decode decodes Title from json.
func (s *Title) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Title to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "english":
			if err := func() error {
				s.English.Reset()
				if err := s.English.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"english\"")
			}
		case "japanese":
			if err := func() error {
				s.Japanese.Reset()
				if err := s.Japanese.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"japanese\"")
			}
		case "pretty":
			if err := func() error {
				s.Pretty.Reset()
				if err := s.Pretty.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"pretty\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Title")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s Title) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Title) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
