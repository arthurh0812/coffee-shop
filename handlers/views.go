package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arthurh0812/coffee-shop/templates"
)

type Views struct {
	handler
}

var views *Views

func NewViews(l *log.Logger) *Views {
	if views == nil {
		views = &Views{handler: newHandler("Views", l)}
	}
	return views
}

func (v *Views) Homepage(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	templ := templates.Get()
	err := templ.ExecuteTemplate(w, "index.html", templates.Data{Title: "Homepage"})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to send HTML: %v", err), http.StatusInternalServerError)
	}
}
