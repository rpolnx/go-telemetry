package config

import (
	"context"
	"log"
	"time"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"go.opentelemetry.io/otel/exporters/zipkin"
)

func InitTracer() func(context.Context) error {
	serviceName := "go-telemetry"
	collectorURL := "http://localhost:9412/api/v2/spans"

	zipkinExporter, err := zipkin.New(collectorURL)
	if err != nil {
		logrus.Errorf("Failed to create the Zipkin exporter: %v", err)
		return nil
	}

	if err != nil {
		logrus.Fatal(err)
	}
	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", serviceName),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		log.Printf("Could not set resources: ", err)
	}

	zipkinBatcher := sdktrace.WithBatcher(zipkinExporter, sdktrace.WithBatchTimeout(5*time.Second))

	otel.SetTracerProvider(
		sdktrace.NewTracerProvider(
			zipkinBatcher,
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithResource(resources),
		),
	)
	return zipkinExporter.Shutdown
}
