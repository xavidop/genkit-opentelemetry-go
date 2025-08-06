# Monitoring Stack for Genkit Application

This directory contains the monitoring stack configuration for collecting telemetry data from your Genkit application.

## Architecture

The monitoring stack includes:
- **OpenTelemetry Collector**: Collects metrics and traces from your Genkit app at `http://localhost:4033`
- **Prometheus**: Stores and queries metrics data
- **Jaeger**: Handles distributed tracing
- **Grafana**: Provides visualization dashboards

## Quick Start

1. Make sure your Genkit application is running and exposing telemetry at `http://localhost:4033`

2. Start the monitoring stack:
   ```bash
   docker-compose up -d
   ```

3. Access the services:
   - **Grafana**: http://localhost:3000 (admin/admin)
   - **Prometheus**: http://localhost:9090
   - **Jaeger**: http://localhost:16686
   - **OTel Collector Health**: http://localhost:13133

## Configuration Details

### OpenTelemetry Collector
- Scrapes metrics from your Genkit app
- Receives OTLP traces and metrics on ports 4317 (gRPC) and 4318 (HTTP)
- Exports metrics to Prometheus and traces to Jaeger

### Key Endpoints
- **Genkit Telemetry**: scraped by collector
- **Collector Metrics**: `http://localhost:8889/metrics` (Prometheus format)
- **Collector Health**: `http://localhost:13133` (health check)

## Troubleshooting

1. **Check collector health**:
   ```bash
   curl http://localhost:13133
   ```

2. **View collector logs**:
   ```bash
   docker-compose logs otel-collector
   ```

3. **Check Prometheus targets**:
   Visit http://localhost:9090/targets to see if the Genkit endpoint is being scraped successfully.

## Customization

- Modify `otel-collector-config.yml` to adjust scraping intervals or add more endpoints
- Update `prometheus.yml` to add additional scrape targets
- Add Grafana dashboards in the `grafana/provisioning/dashboards/` directory
