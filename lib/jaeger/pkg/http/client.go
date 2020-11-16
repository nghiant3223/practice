package http

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/nghiant3223/gopractice/jaeger/pkg/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

type Client interface {
	Get(ctx context.Context, url string, data interface{}) error
}

type client struct {
	client *http.Client
}

func NewClient() Client {
	return &client{
		client: http.DefaultClient,
	}
}

func (c *client) Get(ctx context.Context, url string, data interface{}) error {
	span, ctx, finish := tracing.StartSpanFromContext(ctx, clientGet)
	defer finish()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, http.MethodGet)
	res, err := c.doSpan(span, req)
	if err != nil {
		return err
	}

	j, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(j, data)
}

func (c *client) doSpan(span opentracing.Span, req *http.Request) (*http.Response, error) {
	globalTracer := opentracing.GlobalTracer()
	err := globalTracer.Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)
	if err != nil {
		return nil, err
	}
	return c.client.Do(req)
}
