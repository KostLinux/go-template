package telemetry

import (
	"context"
	"crypto/tls"
	"fmt"
	"go-template/config"
	"go-template/pkg/logger"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"google.golang.org/grpc/credentials"
)

func InitTracer(cfg *config.New) func() {
	ctx := context.Background()
	shutdown, err := newOtelCollector(ctx, cfg)
	if err != nil {
		logger.Errorf("failed to setup OpenTelemetry: %v", err)
	}

	return func() {
		logger.Infof("Shutting down OpenTelemetry")
		if err := shutdown(context.Background()); err != nil {
			logger.Errorf("failed to shutdown OpenTelemetry: %v", err)
		}
	}
}

func newOtelCollector(ctx context.Context, cfg *config.New) (func(context.Context) error, error) {
	// Cleanup functions which need to be executed when the OpenTelemetry SDK is shutting down.
	// When shutdown is called (typically when application is terminating),
	// it executes all these cleanup functions in order and combines their errors
	shutdownHandler := newShutdownHandler()
	telemetry := cfg.Monitoring.Telemetry
	app := cfg.App

	// Setup resource attributes
	resources, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithSchemaURL(semconv.SchemaURL),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(app.Name),
			semconv.ServiceVersionKey.String(app.Version),
			semconv.DeploymentEnvironmentKey.String(app.Env),
			semconv.TelemetrySDKNameKey.String("opentelemetry"),
			semconv.TelemetrySDKLanguageKey.String("go"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create Otel SDK resource: %w", err)
	}

	logger.Infof("Initializing OpenTelemetry with endpoint: %s", telemetry.OTLPEndpoint)
	if apiKey, exists := telemetry.OTLPHeaders["api-key"]; !exists || apiKey == "" {
		return nil, fmt.Errorf("new relic license key not found in configuration")
	}

	// Configure OTEL Exporter with New Relic settings
	opts := []otlptracegrpc.Option{
		otlptracegrpc.WithEndpoint(telemetry.OTLPEndpoint),
		otlptracegrpc.WithHeaders(telemetry.OTLPHeaders),
	}

	// Add TLS configuration unless explicitly disabled
	if !telemetry.OTLPInsecure {
		opts = append(opts, otlptracegrpc.WithTLSCredentials(credentials.NewTLS(&tls.Config{
			MinVersion: tls.VersionTLS12,
		})))
	}

	opts = append(opts,
		otlptracegrpc.WithTimeout(time.Duration(telemetry.OTLPTimeout)*time.Second),
		otlptracegrpc.WithCompressor(telemetry.OTLPCompression),
	)

	client := otlptracegrpc.NewClient(opts...)
	traceExporter, err := otlptrace.New(ctx, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	// Configure batch processor
	batchSpanProcessor := sdktrace.NewBatchSpanProcessor(traceExporter,
		sdktrace.WithMaxQueueSize(telemetry.OTLPQueueSize),
		sdktrace.WithBatchTimeout(time.Duration(telemetry.OTLPBatchTimeout)*time.Second),
		sdktrace.WithMaxExportBatchSize(telemetry.OTLPMaxBatchSize),
	)

	logger.Infof("OTLP Configuration: Compression=%s, Timeout=%ds, QueueSize=%d, BatchSize=%d",
		telemetry.OTLPCompression,
		telemetry.OTLPTimeout,
		telemetry.OTLPQueueSize,
		telemetry.OTLPMaxBatchSize,
	)

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithResource(resources),
		sdktrace.WithSpanProcessor(batchSpanProcessor),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
	)
	shutdownHandler.addFunction(tracerProvider.Shutdown)

	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	return shutdownHandler.shutdown, nil
}
