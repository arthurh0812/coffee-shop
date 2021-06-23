package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type handler struct {
	l    *log.Logger
	name string
}

func newHandler(name string, l *log.Logger) handler {
	return handler{l: l, name: name}
}

func (h handler) SetName(name string) {
	h.name = name
}

func (h handler) Log(v ...interface{}) {
	h.l.SetPrefix(fmt.Sprintf("[%s] ", h.name))
	h.l.Println(v...)
}

func (h handler) Logf(format string, a ...interface{}) {
	h.l.SetPrefix(fmt.Sprintf("[%s] ", h.name))
	h.l.Printf(format, a...)
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
