package api

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func serverMux(state *AppState) http.Handler {
	mux := http.NewServeMux()                                                // creating mux
	mux.HandleFunc("POST /notification", (*state).createNotificationHandler) // setup route
	mux.HandleFunc("GET /notification/{id}", (*state).getNotificationHandler)

	return otelhttp.NewHandler(mux, "notifyx")
}

func ConfigureServer(state *AppState) *http.Server {
	return &http.Server{
		Addr:    ":3000",
		Handler: serverMux(state),
	}
}
