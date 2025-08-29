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

	basicExample(ctx)

	time.Sleep(120 * time.Minute) // Allow time for exporters to flush
}

func basicExample(ctx context.Context) {

	// Initialize OpenTelemetry plugin with default settings
	plugin := opentelemetry.New(opentelemetry.Config{
		ServiceName: "my-genkit-app",
		ForceExport: true, // Export even in development
	})
	// Initialize Genkit
	genkit.Init(ctx,
		genkit.WithPlugins(plugin),
	)

	log.Println("Basic OpenTelemetry setup completed")
}
