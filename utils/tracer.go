package utils

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
)

var (
	Tracer = otel.Tracer("api-server")
)

func NewExporter() (sdktrace.SpanExporter, error) {
	// return stdouttrace.New(
	// 	stdouttrace.WithPrettyPrint(),
	// 	stdouttrace.WithWriter(os.Stderr),
	// )

	// Grafana Alloy に送信
	return otlptracehttp.New(
		context.Background(),
		otlptracehttp.WithEndpoint("alloy:4318"),
		otlptracehttp.WithInsecure(),
	)
}

func NewResource(name, version string) *resource.Resource {
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(name),
		semconv.ServiceVersionKey.String(version),
	)
}

func SetupTraceProvider(shutdownTimeout time.Duration) (func(), error) {
	exporter, err := NewExporter()
	if err != nil {
		return nil, err
	}

	resource := NewResource("api-service", "1.0.0")
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(resource),
	)
	otel.SetTracerProvider(tracerProvider)

	cleanup := func() {
		ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		if err := tracerProvider.Shutdown(ctx); err != nil {
			log.Printf("Failed to shutdown tracer provider: %v\n", err)
		}
	}

	return cleanup, nil
}
