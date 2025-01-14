// Code generated by ogen, DO NOT EDIT.

package techempower

import (
	"context"
	"net/url"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric/instrument/syncint64"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
)

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	cfg       config
	requests  syncint64.Counter
	errors    syncint64.Counter
	duration  syncint64.Histogram
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
	if c.requests, err = c.cfg.Meter.SyncInt64().Counter(otelogen.ClientRequestCount); err != nil {
		return nil, err
	}
	if c.errors, err = c.cfg.Meter.SyncInt64().Counter(otelogen.ClientErrorsCount); err != nil {
		return nil, err
	}
	if c.duration, err = c.cfg.Meter.SyncInt64().Histogram(otelogen.ClientDuration); err != nil {
		return nil, err
	}
	return c, nil
}

// Caching invokes Caching operation.
//
// GET /cached-worlds
func (c *Client) Caching(ctx context.Context, params CachingParams) (res WorldObjects, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("Caching"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "Caching",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1, otelAttrs...)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
		}
		span.End()
	}()
	c.requests.Add(ctx, 1, otelAttrs...)
	u := uri.Clone(c.serverURL)
	u.Path += "/cached-worlds"

	q := u.Query()
	{
		// Encode "count" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.Count))
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["count"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeCachingResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DB invokes DB operation.
//
// GET /db
func (c *Client) DB(ctx context.Context) (res WorldObject, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("DB"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "DB",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1, otelAttrs...)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
		}
		span.End()
	}()
	c.requests.Add(ctx, 1, otelAttrs...)
	u := uri.Clone(c.serverURL)
	u.Path += "/db"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeDBResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// JSON invokes json operation.
//
// GET /json
func (c *Client) JSON(ctx context.Context) (res HelloWorld, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("json"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "JSON",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1, otelAttrs...)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
		}
		span.End()
	}()
	c.requests.Add(ctx, 1, otelAttrs...)
	u := uri.Clone(c.serverURL)
	u.Path += "/json"

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeJSONResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// Queries invokes Queries operation.
//
// GET /queries
func (c *Client) Queries(ctx context.Context, params QueriesParams) (res WorldObjects, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("Queries"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "Queries",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1, otelAttrs...)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
		}
		span.End()
	}()
	c.requests.Add(ctx, 1, otelAttrs...)
	u := uri.Clone(c.serverURL)
	u.Path += "/queries"

	q := u.Query()
	{
		// Encode "queries" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.Queries))
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["queries"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeQueriesResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// Updates invokes Updates operation.
//
// GET /updates
func (c *Client) Updates(ctx context.Context, params UpdatesParams) (res WorldObjects, err error) {
	startTime := time.Now()
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("Updates"),
	}
	ctx, span := c.cfg.Tracer.Start(ctx, "Updates",
		trace.WithAttributes(otelAttrs...),
		trace.WithSpanKind(trace.SpanKindClient),
	)
	defer func() {
		if err != nil {
			span.RecordError(err)
			c.errors.Add(ctx, 1, otelAttrs...)
		} else {
			elapsedDuration := time.Since(startTime)
			c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
		}
		span.End()
	}()
	c.requests.Add(ctx, 1, otelAttrs...)
	u := uri.Clone(c.serverURL)
	u.Path += "/updates"

	q := u.Query()
	{
		// Encode "queries" parameter.
		e := uri.NewQueryEncoder(uri.QueryEncoderConfig{
			Style:   uri.QueryStyleForm,
			Explode: true,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.Queries))
		}(); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
		q["queries"] = e.Result()
	}
	u.RawQuery = q.Encode()

	r := ht.NewRequest(ctx, "GET", u, nil)
	defer ht.PutRequest(r)

	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	result, err := decodeUpdatesResponse(resp, span)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
