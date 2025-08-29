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

	prometheusExample(ctx)

	time.Sleep(120 * time.Minute) // Allow time for exporters to flush
}

func prometheusExample(ctx context.Context) {
	// Example: Using Jaeger preset (commented out to avoid interference)
	plugin := opentelemetry.NewWithPreset(opentelemetry.PresetPrometheus, opentelemetry.Config{
		ServiceName:    "my-genkit-app",
		ForceExport:    true, // Export even in development
		PrometheusPort: 8081, // Custom port for the Prometheus exporter with endpoint /metrics
	})

	// Initialize Genkit
	genkit.Init(ctx,
		genkit.WithPlugins(plugin),
	)

	log.Println("Preset examples completed")
}
