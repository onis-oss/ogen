// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"net/http"

	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/validate"
)

func decodeDeletePetResponse(resp *http.Response, span trace.Span) (res DeletePetRes, err error) {
	switch resp.StatusCode {
	case 204:
		return &DeletePetNoContent{}, nil
	default:
		switch ct := resp.Header.Get("Content-Type"); ct {
		case "application/json":
			buf := getBuf()
			defer putBuf(buf)
			if _, err := io.Copy(buf, resp.Body); err != nil {
				return res, err
			}

			d := jx.GetDecoder()
			defer jx.PutDecoder(d)
			d.ResetBytes(buf.Bytes())

			var response Error
			if err := func() error {
				if err := response.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return res, err
			}

			return &ErrorStatusCode{
				StatusCode: resp.StatusCode,
				Response:   response,
			}, nil
		default:
			return res, validate.InvalidContentType(ct)
		}
	}
}
