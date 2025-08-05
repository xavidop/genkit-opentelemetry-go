package main

import (
	"context"
	"log"
	"time"

	"github.com/firebase/genkit/go/genkit"
	opentelemetry "github.com/xavierportillaedo/genkit-opentelemetry-go"
)

func main() {
	ctx := context.Background()

	basicExample(ctx)

	time.Sleep(120 * time.Minute) // Allow time for exporters to flush
}

func basicExample(ctx context.Context) {
	// Initialize Genkit
	g, err := genkit.Init(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize OpenTelemetry plugin with default settings
	plugin := opentelemetry.New(opentelemetry.Config{
		ServiceName: "my-genkit-app",
		ForceExport: true, // Export even in development
	})

	if err := plugin.Init(ctx, g); err != nil {
		log.Fatal(err)
	}

	log.Println("Basic OpenTelemetry setup completed")
}
