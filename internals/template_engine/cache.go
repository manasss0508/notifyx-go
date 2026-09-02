package template_engine

import (
	"sync"

	"github.com/manasss0508/notifyx-go/internals/repository"
	"github.com/manasss0508/notifyx-go/internals/repository/queries"
)

type TemplateCache struct {
	mail map[string]queries.Template
	m    sync.RWMutex
}

func NewTemplateCache() *TemplateCache {
	return &TemplateCache{
		mail: make(map[string]queries.Template),
	}
}

func (T TemplateCache) GetTemplateMail(q *queries.Queries, channelType, templateName string,
) (queries.Template, error) {
	// check if template exits in cache
	// take read lock
	T.m.RLock()

	// check in cache
	value, ok := T.mail[templateName]
	T.m.RUnlock()

	if !ok {
		// if not exits
		// get template from database
		template, err := repository.DbGetTemp(q, channelType, templateName)
		if err != nil {
			return queries.Template{}, err
		}

		// get write lock
		T.m.Lock()

		//inserting template to cache
		T.mail[templateName] = template

		return template, nil
	}

	return value, nil
}
