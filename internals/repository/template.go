package repository

import (
	"context"

	"github.com/manasss0508/notifyx-go/internals/repository/queries"
)

func DbGetTemp(q *queries.Queries, channelType string, templateName string) (queries.Template, error) {
	params := queries.DbGetTemplateParams{
		Channel: channelType,
		Name:    templateName,
	}

	return q.DbGetTemplate(context.Background(), params)
}
