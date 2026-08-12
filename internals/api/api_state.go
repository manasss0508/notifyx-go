package api

import (
	"github.com/manasss0508/notifyx-go/internals/queue"
	"github.com/manasss0508/notifyx-go/internals/repository/queries"
	"go.opentelemetry.io/otel/trace"
)

type AppState struct {
	Db         *queries.Queries
	Rbmq       *queue.QueueConn
	OtelTracer trace.Tracer
}

func NewAppState(db *queries.Queries, rbmq *queue.QueueConn, otelTracer trace.Tracer) *AppState {
	return &AppState{
		Db:         db,
		Rbmq:       rbmq,
		OtelTracer: otelTracer,
	}
}
