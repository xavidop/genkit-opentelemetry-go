package opentelemetry

import (
	"context"
	"crypto/tls"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// hasScheme returns true if the endpoint contains a URL scheme.
func hasScheme(endpoint string) bool {
	return strings.Contains(endpoint, "://")
}

// stripScheme removes the scheme (http:// or https://) from an endpoint.
func stripScheme(endpoint string) string {
	if strings.HasPrefix(endpoint, "https://") {
		return endpoint[8:]
	}
	if strings.HasPrefix(endpoint, "http://") {
		return endpoint[7:]
	}
	return endpoint
}

// createDefaultTraceExporter creates the default OTLP trace exporter.
func (ot *OpenTelemetry) createDefaultTraceExporter(ctx context.Context) (trace.SpanExporter, error) {
	// If OTEL_EXPORTER_OTLP_TRACES_ENDPOINT is "stdout", use stdout exporter
	if endpoint := os.Getenv("OTEL_EXPORTER_OTLP_TRACES_ENDPOINT"); endpoint == "stdout" {
		return stdouttrace.New(stdouttrace.WithPrettyPrint())
	}

	if ot.config.OTLPUseHTTP {
		// For HTTP, the endpoint should not include the scheme for WithEndpoint()
		endpoint := ot.config.OTLPEndpoint

		// Determine if we should use TLS based on the original endpoint
		useTLS := false
		if hasScheme(endpoint) {
			useTLS = strings.HasPrefix(endpoint, "https://")
			endpoint = stripScheme(endpoint)
		}

		opts := []otlptracehttp.Option{
			otlptracehttp.WithEndpoint(endpoint),
			otlptracehttp.WithTimeout(30 * time.Second), // Add 30 second timeout
		}

		if ot.config.OTLPHeaders != nil {
			opts = append(opts, otlptracehttp.WithHeaders(ot.config.OTLPHeaders))
		}

		// Configure TLS based on the original scheme
		if useTLS {
			opts = append(opts, otlptracehttp.WithTLSClientConfig(&tls.Config{}))
		} else {
			opts = append(opts, otlptracehttp.WithInsecure())
		}

		return otlptracehttp.New(ctx, opts...)
	} else {
		// For gRPC, strip any scheme from the endpoint
		endpoint := stripScheme(ot.config.OTLPEndpoint)

		opts := []otlptracegrpc.Option{
			otlptracegrpc.WithEndpoint(endpoint),
			otlptracegrpc.WithTimeout(30 * time.Second), // Add 30 second timeout
		}

		if ot.config.OTLPHeaders != nil {
			opts = append(opts, otlptracegrpc.WithHeaders(ot.config.OTLPHeaders))
		}

		// Configure gRPC connection
		dialOpts := []grpc.DialOption{}

		// If original endpoint starts with https, use TLS
		if len(ot.config.OTLPEndpoint) > 8 && ot.config.OTLPEndpoint[:8] == "https://" {
			dialOpts = append(dialOpts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
		} else {
			opts = append(opts, otlptracegrpc.WithInsecure())
		}

		if len(dialOpts) > 0 {
			opts = append(opts, otlptracegrpc.WithDialOption(dialOpts...))
		}

		return otlptracegrpc.New(ctx, opts...)
	}
}

// createDefaultMetricExporter creates the default OTLP metric exporter.
func (ot *OpenTelemetry) createDefaultMetricExporter(ctx context.Context) (metric.Exporter, error) {
	// If OTEL_EXPORTER_OTLP_METRICS_ENDPOINT is "stdout", use stdout exporter
	if endpoint := os.Getenv("OTEL_EXPORTER_OTLP_METRICS_ENDPOINT"); endpoint == "stdout" {
		return stdoutmetric.New(stdoutmetric.WithPrettyPrint())
	}

	if ot.config.OTLPUseHTTP {
		// For HTTP, the endpoint should not include the scheme for WithEndpoint()
		endpoint := ot.config.OTLPEndpoint

		// Determine if we should use TLS based on the original endpoint
		useTLS := false
		if hasScheme(endpoint) {
			useTLS = strings.HasPrefix(endpoint, "https://")
			endpoint = stripScheme(endpoint)
		}

		opts := []otlpmetrichttp.Option{
			otlpmetrichttp.WithEndpoint(endpoint),
			otlpmetrichttp.WithTimeout(30 * time.Second), // Add 30 second timeout
		}

		if ot.config.OTLPHeaders != nil {
			opts = append(opts, otlpmetrichttp.WithHeaders(ot.config.OTLPHeaders))
		}

		// Configure TLS based on the original scheme
		if useTLS {
			opts = append(opts, otlpmetrichttp.WithTLSClientConfig(&tls.Config{}))
		} else {
			opts = append(opts, otlpmetrichttp.WithInsecure())
		}

		return otlpmetrichttp.New(ctx, opts...)
	} else {
		// For gRPC, strip any scheme from the endpoint
		endpoint := stripScheme(ot.config.OTLPEndpoint)

		opts := []otlpmetricgrpc.Option{
			otlpmetricgrpc.WithEndpoint(endpoint),
			otlpmetricgrpc.WithTimeout(30 * time.Second), // Add 30 second timeout
		}

		if ot.config.OTLPHeaders != nil {
			opts = append(opts, otlpmetricgrpc.WithHeaders(ot.config.OTLPHeaders))
		}

		// Configure gRPC connection
		dialOpts := []grpc.DialOption{}

		// If original endpoint starts with https, use TLS
		if len(ot.config.OTLPEndpoint) > 8 && ot.config.OTLPEndpoint[:8] == "https://" {
			dialOpts = append(dialOpts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
		} else {
			opts = append(opts, otlpmetricgrpc.WithInsecure())
		}

		if len(dialOpts) > 0 {
			opts = append(opts, otlpmetricgrpc.WithDialOption(dialOpts...))
		}

		return otlpmetricgrpc.New(ctx, opts...)
	}
}

// createDefaultLogHandler creates the default structured log handler.
func (ot *OpenTelemetry) createDefaultLogHandler() slog.Handler {
	opts := &slog.HandlerOptions{
		Level: ot.config.LogLevel,
	}

	// Use JSON handler for structured logging
	return slog.NewJSONHandler(os.Stdout, opts)
}

// createStdoutMetricExporter creates a stdout metric exporter for development/Jaeger preset.
func createStdoutMetricExporter() metric.Exporter {
	exporter, _ := stdoutmetric.New(stdoutmetric.WithPrettyPrint())
	return exporter
}

// setupPrometheusMetrics creates a Prometheus metric exporter.
func (ot *OpenTelemetry) setupPrometheusMetrics(ctx context.Context) error {
	if ot.config.MetricExporter != nil {
		// Custom exporter provided, use parent logic
		return ot.setupMetrics(ctx)
	}

	// Create Prometheus exporter
	exporter, err := prometheus.New()
	if err != nil {
		return err
	}

	meterProvider := metric.NewMeterProvider(metric.WithReader(exporter))
	otel.SetMeterProvider(meterProvider)

	// Start HTTP server for /metrics endpoint if enabled
	if ot.config.EnablePrometheusEndpoint {
		// Ensure serverWg is initialized (safety check)
		if ot.serverWg == nil {
			ot.serverWg = &sync.WaitGroup{}
		}

		port := ot.config.PrometheusPort
		if port == 0 {
			port = 9090
		}

		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler())

		// Create server context for graceful shutdown
		serverCtx, serverCancel := context.WithCancel(context.Background())
		ot.serverCancel = serverCancel

		ot.server = &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: mux,
		}

		// Use a channel to signal when the server has started listening
		serverStarted := make(chan error, 1)

		// Increment WaitGroup before starting goroutine
		ot.serverWg.Add(1)

		go func() {
			defer ot.serverWg.Done()

			slog.Info("Starting Prometheus metrics server", "port", port, "endpoint", "/metrics")

			// Create a listener first to ensure we can bind to the port
			listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
			if err != nil {
				serverStarted <- err
				return
			}

			// Signal that we've successfully bound to the port
			serverStarted <- nil

			// Start serving with context cancellation support
			go func() {
				<-serverCtx.Done()
				// Context cancelled, initiate shutdown
				shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()
				ot.server.Shutdown(shutdownCtx)
			}()

			// Start serving
			if err := ot.server.Serve(listener); err != nil && err != http.ErrServerClosed {
				slog.Error("Prometheus metrics server failed", "error", err)
			}
		}()

		// Wait for the server to start or fail
		select {
		case err := <-serverStarted:
			if err != nil {
				return fmt.Errorf("failed to start Prometheus metrics server on port %d: %w", port, err)
			}
			slog.Info("Prometheus metrics server started successfully", "port", port)
		case <-time.After(5 * time.Second):
			return fmt.Errorf("timeout waiting for Prometheus metrics server to start on port %d", port)
		}
	}

	return nil
}
