// Package tracing wraps creating a trace provider
package tracing

import (
	_ "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx" // crdb retries and postgres interface
	_ "github.com/lib/pq"                                   // Register the Postgres driver.
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.uber.org/zap"
)

// New returns an OpenTelemetry TracerProvider configured to use
// the Jaeger exporter that will send spans to the provided url. The returned
// TracerProvider will also use a Resource configured with all the information
// about the application.
func New(url string, logger *zap.SugaredLogger) *tracesdk.TracerProvider {
	// Create the jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		logger.Fatalw("failed to initialize tracing exporter", "error", err)
	}

	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production
		tracesdk.WithBatcher(exp),
		// Record information about this application a resource
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("metadataservice"),
			attribute.String("environment", viper.GetString("tracing.environment")),
		)),
	)

	otel.SetTracerProvider(tp)

	return tp
}
