package main

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/manasss0508/notifyx-go/internals/api"
	"github.com/manasss0508/notifyx-go/internals/configuration"
)

func main() {

	// setupping openTelemetry
	ctx := context.Background()
	shutdown := configuration.OpenTelemetrySetup()
	defer shutdown(ctx) // shutdown openTelemetry

	// setupping logger
	configuration.SlogLoggerSetup()

	// loading all configurations
	appState := configuration.Load() // dont change this because opentelemetry should setuped first for tracer

	//server configurations along with mux
	server := api.ConfigureServer(appState)

	fmt.Println("listening at localhost:3000")
	server.ListenAndServe()
}
