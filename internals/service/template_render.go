package service

import (
	"encoding/json"
	"strings"

	"github.com/manasss0508/notifyx-go/internals/repository/queries"
)

func TemplateRender(template queries.Template, variables string) (string, string) {
	var vars map[string]string
	json.Unmarshal([]byte(variables), &vars)

	// render subject
	subject := templateInterpolation(template.Subject, vars)

	// render body
	body := templateInterpolation(template.Body, vars)

	return subject, body
}

func templateInterpolation(template string, variables map[string]string) string {

	for {
		start := strings.Index(template, "{{")
		if start == -1 {
			break
		}

		afterOpen := start + 2

		relativeEnd := strings.Index(template[afterOpen:], "}}")
		if relativeEnd == -1 {
			break
		}

		end := afterOpen + relativeEnd

		variable := template[afterOpen:end]

		value, ok := variables[variable]
		if !ok {
			break
		}

		template = template[:start] + value + template[end+2:]
	}

	return template
}
