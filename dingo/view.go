package dingo

import (
	"html/template"
	"net/http"
)

type Viewer interface {
	Render(http.ResponseWriter, interface{})
}

func RenderTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	templates := template.Must(template.ParseFiles(conf.TemplateDir + templateName))
	err := templates.ExecuteTemplate(w, templateName, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
