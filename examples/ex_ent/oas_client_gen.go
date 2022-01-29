// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/bits"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = bits.LeadingZeros64
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
)

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	cfg       config
	requests  metric.Int64Counter
	errors    metric.Int64Counter
	duration  metric.Int64Histogram
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...Option) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	c := &Client{
		cfg:       newConfig(opts...),
		serverURL: u,
	}
	if c.requests, err = c.cfg.Meter.NewInt64Counter(otelogen.ClientRequestCount); err != nil {
		return nil, err
	}
	if c.errors, err = c.cfg.Meter.NewInt64Counter(otelogen.ClientErrorsCount); err != nil {
		return nil, err
	}
	if c.duration, err = c.cfg.Meter.NewInt64Histogram(otelogen.ClientDuration); err != nil {
		return nil, err
	}
	return c, nil
}

// CreatePet invokes createPet operation.
//
// Creates a new Pet and persists it to storage.
//
// POST /pets
func (c *Client) CreatePet(ctx context.Context, request CreatePetReq) (res CreatePetRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `CreatePet`,
		trace.WithAttributes(otelogen.OperationID(`createPet`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodeCreatePetRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer jx.PutWriter(buf)
	reqBody = bytes.NewReader(buf.Buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/pets"

	r := ht.NewRequest(ctx, "POST", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeCreatePetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// CreatePetCategories invokes createPetCategories operation.
//
// Creates a new Category and attaches it to the Pet.
//
// POST /pets/{id}/categories
func (c *Client) CreatePetCategories(ctx context.Context, request CreatePetCategoriesReq, params CreatePetCategoriesParams) (res CreatePetCategoriesRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `CreatePetCategories`,
		trace.WithAttributes(otelogen.OperationID(`createPetCategories`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodeCreatePetCategoriesRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer jx.PutWriter(buf)
	reqBody = bytes.NewReader(buf.Buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/pets/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}
	u.Path += "/categories"

	r := ht.NewRequest(ctx, "POST", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeCreatePetCategoriesResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// CreatePetFriends invokes createPetFriends operation.
//
// Creates a new Pet and attaches it to the Pet.
//
// POST /pets/{id}/friends
func (c *Client) CreatePetFriends(ctx context.Context, request CreatePetFriendsReq, params CreatePetFriendsParams) (res CreatePetFriendsRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `CreatePetFriends`,
		trace.WithAttributes(otelogen.OperationID(`createPetFriends`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodeCreatePetFriendsRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer jx.PutWriter(buf)
	reqBody = bytes.NewReader(buf.Buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/pets/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}
	u.Path += "/friends"

	r := ht.NewRequest(ctx, "POST", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeCreatePetFriendsResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// CreatePetOwner invokes createPetOwner operation.
//
// Creates a new User and attaches it to the Pet.
//
// POST /pets/{id}/owner
func (c *Client) CreatePetOwner(ctx context.Context, request CreatePetOwnerReq, params CreatePetOwnerParams) (res CreatePetOwnerRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `CreatePetOwner`,
		trace.WithAttributes(otelogen.OperationID(`createPetOwner`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodeCreatePetOwnerRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer jx.PutWriter(buf)
	reqBody = bytes.NewReader(buf.Buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/pets/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}
	u.Path += "/owner"

	r := ht.NewRequest(ctx, "POST", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeCreatePetOwnerResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DeletePet invokes deletePet operation.
//
// Deletes the Pet with the requested ID.
//
// DELETE /pets/{id}
func (c *Client) DeletePet(ctx context.Context, params DeletePetParams) (res DeletePetRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `DeletePet`,
		trace.WithAttributes(otelogen.OperationID(`deletePet`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/pets/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}

	r := ht.NewRequest(ctx, "DELETE", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeDeletePetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DeletePetOwner invokes deletePetOwner operation.
//
// Delete the attached Owner of the Pet with the given ID.
//
// DELETE /pets/{id}/owner
func (c *Client) DeletePetOwner(ctx context.Context, params DeletePetOwnerParams) (res DeletePetOwnerRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `DeletePetOwner`,
		trace.WithAttributes(otelogen.OperationID(`deletePetOwner`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/pets/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}
	u.Path += "/owner"

	r := ht.NewRequest(ctx, "DELETE", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeDeletePetOwnerResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListPet invokes listPet operation.
//
// List Pets.
//
// GET /pets
func (c *Client) ListPet(ctx context.Context, params ListPetParams) (res ListPetRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `ListPet`,
		trace.WithAttributes(otelogen.OperationID(`listPet`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/pets"

	q := u.Query()
	{
		// Encode "page" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			if val, ok := params.Page.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["page"] = e.Result()
	}
	{
		// Encode "itemsPerPage" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			if val, ok := params.ItemsPerPage.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["itemsPerPage"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeListPetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListPetCategories invokes listPetCategories operation.
//
// List attached Categories.
//
// GET /pets/{id}/categories
func (c *Client) ListPetCategories(ctx context.Context, params ListPetCategoriesParams) (res ListPetCategoriesRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `ListPetCategories`,
		trace.WithAttributes(otelogen.OperationID(`listPetCategories`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/pets/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}
	u.Path += "/categories"

	q := u.Query()
	{
		// Encode "page" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			if val, ok := params.Page.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["page"] = e.Result()
	}
	{
		// Encode "itemsPerPage" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			if val, ok := params.ItemsPerPage.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["itemsPerPage"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeListPetCategoriesResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListPetFriends invokes listPetFriends operation.
//
// List attached Friends.
//
// GET /pets/{id}/friends
func (c *Client) ListPetFriends(ctx context.Context, params ListPetFriendsParams) (res ListPetFriendsRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `ListPetFriends`,
		trace.WithAttributes(otelogen.OperationID(`listPetFriends`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/pets/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}
	u.Path += "/friends"

	q := u.Query()
	{
		// Encode "page" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			if val, ok := params.Page.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["page"] = e.Result()
	}
	{
		// Encode "itemsPerPage" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			if val, ok := params.ItemsPerPage.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["itemsPerPage"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeListPetFriendsResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ReadPet invokes readPet operation.
//
// Finds the Pet with the requested ID and returns it.
//
// GET /pets/{id}
func (c *Client) ReadPet(ctx context.Context, params ReadPetParams) (res ReadPetRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `ReadPet`,
		trace.WithAttributes(otelogen.OperationID(`readPet`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/pets/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeReadPetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ReadPetOwner invokes readPetOwner operation.
//
// Find the attached User of the Pet with the given ID.
//
// GET /pets/{id}/owner
func (c *Client) ReadPetOwner(ctx context.Context, params ReadPetOwnerParams) (res ReadPetOwnerRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `ReadPetOwner`,
		trace.WithAttributes(otelogen.OperationID(`readPetOwner`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	u := uri.Clone(c.serverURL)
	u.Path += "/pets/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}
	u.Path += "/owner"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeReadPetOwnerResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UpdatePet invokes updatePet operation.
//
// Updates a Pet and persists changes to storage.
//
// PATCH /pets/{id}
func (c *Client) UpdatePet(ctx context.Context, request UpdatePetReq, params UpdatePetParams) (res UpdatePetRes, err error) {
	startTime := time.Now()
	ctx, span := c.cfg.Tracer.Start(ctx, `UpdatePet`,
		trace.WithAttributes(otelogen.OperationID(`updatePet`)),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds())
		}
		span.End()
	}()
	c.requests.Add(ctx, 1)
	var (
		contentType string
		reqBody     io.Reader
	)
	contentType = "application/json"
	buf, err := encodeUpdatePetRequestJSON(request, span)
	if err != nil {
		return res, err
	}
	defer jx.PutWriter(buf)
	reqBody = bytes.NewReader(buf.Buf)

	u := uri.Clone(c.serverURL)
	u.Path += "/pets/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.IntToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		u.Path += e.Result()
	}

	r := ht.NewRequest(ctx, "PATCH", u, reqBody)
	defer ht.PutRequest(r)

	r.Header.Set("Content-Type", contentType)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeUpdatePetResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}