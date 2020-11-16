package tracing

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func StartSpanFromContext(
	ctx context.Context, operationName string, opts ...opentracing.StartSpanOption,
) (opentracing.Span, context.Context, func()) {
	parentSpan := opentracing.SpanFromContext(ctx)
	if parentSpan != nil {
		opts = append(opts, opentracing.ChildOf(parentSpan.Context()))
	}
	span := opentracing.StartSpan(operationName, opts...)
	return span, opentracing.ContextWithSpan(ctx, span), span.Finish
}

func StartSpanFromGinContext(
	ctx *gin.Context,
	opts ...opentracing.StartSpanOption,
) (opentracing.Span, context.Context, func()) {
	tracer := opentracing.GlobalTracer()
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.Request.Header))

	operationName := fmt.Sprintf("%s %s", ctx.Request.Method, ctx.Request.URL)
	opts = append(opts, ext.RPCServerOption(spanCtx))
	span := tracer.StartSpan(operationName, opts...)

	return span, opentracing.ContextWithSpan(ctx, span), span.Finish
}
