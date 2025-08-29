package main

import (
	"context"
	"log"
	"time"

	"github.com/firebase/genkit/go/genkit"
	opentelemetry "github.com/xavidop/genkit-opentelemetry-go"
)

func main() {
	ctx := context.Background()

	otelExample(ctx)

	time.Sleep(120 * time.Minute) // Allow time for exporters to flush
}

func otelExample(ctx context.Context) {
	// Initialize Genkit
	genkit.Init(ctx)

	// Example: Using Jaeger preset (commented out to avoid interference)
	plugin := opentelemetry.NewWithPreset(opentelemetry.PresetOTLP, opentelemetry.Config{
		ServiceName:    "my-genkit-app",
		ForceExport:    true, // Export even in development
		MetricInterval: 15 * time.Second,
	})

	if err := plugin.Init(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Preset examples completed")
}
