package configuration

import (
	"context"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/manasss0508/notifyx-go/internals/api"
	"github.com/manasss0508/notifyx-go/internals/queue"
	"github.com/manasss0508/notifyx-go/internals/repository"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func Load() *api.AppState {
	// loading environment variable
	err := godotenv.Load()
	if err != nil {
		panic("failed to load environment variables")
	}

	// creating database connection
	queries := repository.CreateConnection()

	// creating rbmq connection
	rbmq := queue.NewQueueConn()

	// creating otel tracer
	tracer := otel.Tracer("notifyx")

	// creating app state
	return api.NewAppState(queries, rbmq, tracer)

}

// OpenTelemetrySetup - it used to setup OpenTelemetry
// OpenTelemetry is used for request tracing
// this
func OpenTelemetrySetup() func(ctx context.Context) error {
	// 1. creating exporter - exporter prints span to terminal
	exporter, err := stdouttrace.New(
		stdouttrace.WithPrettyPrint(),
	)
	if err != nil {
		panic("failed to start open telemetry")
	}

	// 2. creating provider - which process span and send to exporter
	provider := sdktrace.NewTracerProvider( // creating provider
		sdktrace.WithBatcher(exporter), // attaching exporter to provider
	)

	// 3.register provider globally. so every span is process through this provider
	otel.SetTracerProvider(provider)

	return provider.Shutdown
}

func SlogLoggerSetup() {
	logger := slog.New(
		slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelInfo,
			},
		),
	)

	slog.SetDefault(logger)
}
