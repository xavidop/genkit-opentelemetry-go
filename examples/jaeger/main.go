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

	jaegerExample(ctx)

	time.Sleep(120 * time.Minute) // Allow time for exporters to flush
}

func jaegerExample(ctx context.Context) {
	// Initialize Genkit
	g := genkit.Init(ctx)

	// Example: Using Jaeger preset (commented out to avoid interference)
	jaegerPlugin := opentelemetry.NewWithPreset(opentelemetry.PresetJaeger, opentelemetry.Config{
		ServiceName: "my-genkit-app",
		ForceExport: true, // Export even in development
	})

	if err := jaegerPlugin.Init(ctx, g); err != nil {
		log.Fatal(err)
	}

	log.Println("Preset examples completed")
}
