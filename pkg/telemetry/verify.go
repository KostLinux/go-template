package telemetry

import (
	"context"
	"time"

	"go-template/pkg/logger"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func VerifyConnection(ctx context.Context, tracer trace.Tracer) {
	_, span := tracer.Start(ctx, "ConnectionTest")
	defer span.End()

	span.SetAttributes(
		attribute.String("test.connection", "true"),
		attribute.String("test.timestamp", time.Now().Format(time.RFC3339)),
	)

	logger.Infof("Sent test span to observability provider")
}
