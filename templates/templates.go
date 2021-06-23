package templates

import (
	"html/template"
)

type Data struct {
	Title string
}

var templ *template.Template

func Get() *template.Template {
	if templ == nil {
		templ = template.Must(template.ParseFiles("templates/index.html"))
	}
	return templ
}
