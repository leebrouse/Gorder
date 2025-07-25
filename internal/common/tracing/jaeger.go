package tracing

import (
	"context"

	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("default_tracer")

func InitJaegerProvider(jaegerURL, serviceName string) (func(ctx context.Context) error, error) {
	if jaegerURL == "" {
		panic("empty jaeger url")
	}
	// Tracer 获取
	tracer = otel.Tracer(serviceName)

	// New OTel Exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(jaegerURL)))
	if err != nil {
		return nil, err
	}

	// create Tracer factory and config options
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String(serviceName),
		)),
	)
	otel.SetTracerProvider(tp)

	// set context（Propagator）
	b3Propagator := b3.New(b3.WithInjectEncoding(b3.B3MultipleHeader))
	p := propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{}, propagation.Baggage{}, b3Propagator,
	)
	otel.SetTextMapPropagator(p)

	return tp.Shutdown, nil
}

// create span for tracing the total distributed link
func Start(ctx context.Context, name string) (context.Context, trace.Span) {
	return tracer.Start(ctx, name)
}

// Create TraceID for research in the jaeger UI
func TraceID(ctx context.Context) string {
	spanCtx := trace.SpanContextFromContext(ctx)
	return spanCtx.TraceID().String()
}
