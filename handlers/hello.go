package handlers

import (
	"log"
	"net/http"
)

type Hello struct {
	handler
}

var hello *Hello

func NewHello(l *log.Logger) *Hello {
	if hello == nil { // singleton
		hello = &Hello{handler: newHandler("Hello", l)}
	}
	return hello
}

func (h *Hello) Get(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/html")
	w.Write([]byte("Hello World!"))
}
