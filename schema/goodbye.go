package schema

import (
	"encoding/json"
	"io"
)

type GoodByeRequest struct {
	Name string `json:"name"`
}

func NewGoodByeRequest(r io.Reader) (*GoodByeRequest, error) {
	req := &GoodByeRequest{}
	err := json.NewDecoder(r).Decode(req)
	return req, err
}

type GoodByeResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func EncodeGoodByeResponse(w io.Writer, res *GoodByeResponse) error {
	return json.NewEncoder(w).Encode(res)
}
