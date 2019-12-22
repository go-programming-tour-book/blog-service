package middleware

import (
	"github.com/uber/jaeger-client-go"

	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	opentracing "github.com/opentracing/opentracing-go"
)

func Tracing() func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		span := opentracing.SpanFromContext(ctx)
		if span != nil {
			span = global.Tracer.StartSpan(c.Request.URL.Path, opentracing.ChildOf(span.Context()))
		} else {
			span = global.Tracer.StartSpan(c.Request.URL.Path)
		}
		defer span.Finish()

		var traceID string
		var SpanID string
		var spanContext = span.Context()
		switch spanContext.(type) {
		case jaeger.SpanContext:
			traceID = spanContext.(jaeger.SpanContext).TraceID().String()
			SpanID = spanContext.(jaeger.SpanContext).SpanID().String()
		}
		c.Set("X-Trace-ID", traceID)
		c.Set("X-Span-ID", SpanID)

		c.Next()
	}
}
