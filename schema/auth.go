package schema

import (
	"encoding/json"
	"io"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func DecodeLoginRequest(r io.Reader) (*LoginRequest, error) {
	req := &LoginRequest{}
	err := json.NewDecoder(r).Decode(req)
	return req, err
}

type LoginResponse struct {
	Token string `json:"token"`
}

func EncodeLoginResponse(w io.Writer, res *LoginResponse) error {
	return json.NewEncoder(w).Encode(res)
}
