package api

import (
	"github.com/manasss0508/notifyx-go/internals/queue"
	"github.com/manasss0508/notifyx-go/internals/repository/queries"
	"github.com/manasss0508/notifyx-go/internals/service"
	"github.com/manasss0508/notifyx-go/internals/template_engine"
	"go.opentelemetry.io/otel/trace"
)

type AppState struct {
	Db            *queries.Queries
	Rbmq          *queue.QueueConn
	TemplateCache *template_engine.TemplateCache
	EmailService  *service.EmailService
	OtelTracer    trace.Tracer
}

func NewAppState(db *queries.Queries, rbmq *queue.QueueConn, otelTracer trace.Tracer,
	templateCache *template_engine.TemplateCache,
	emailService *service.EmailService) *AppState {
	return &AppState{
		Db:            db,
		Rbmq:          rbmq,
		TemplateCache: templateCache,
		OtelTracer:    otelTracer,
		EmailService:  emailService,
	}
}
