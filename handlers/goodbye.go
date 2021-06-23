package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arthurh0812/coffee-shop/schema"
)

type GoodBye struct {
	handler
}

var goodbye *GoodBye

func NewGoodBye(l *log.Logger) *GoodBye {
	if goodbye == nil { // singleton
		goodbye = &GoodBye{handler: newHandler("GoodBye", l)}
	}
	return goodbye
}

func (g *GoodBye) Get(w http.ResponseWriter, r *http.Request) {
	req, err := schema.NewGoodByeRequest(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid request: %v", err), http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	err = schema.EncodeGoodByeResponse(w, &schema.GoodByeResponse{
		Message: fmt.Sprintf("Good bye, %s!\n", req.Name),
		Status:  http.StatusOK,
	})
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("failed to send JSON response: %v", err),
			http.StatusInternalServerError,
		)
		return
	}
}
