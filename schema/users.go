package schema

import (
	"encoding/json"
	"io"
	"time"

	"github.com/arthurh0812/coffee-shop/models"
)

type UserResponse struct {
	User    *models.User `json:"user"`
	Message string       `json:"message"`
	Status  int          `json:"status"`
}

func EncodeUserResponse(w io.Writer, res *UserResponse) error {
	return json.NewEncoder(w).Encode(res)
}

type UsersResponse struct {
	Users   models.UserList `json:"users"`
	Count   int             `json:"count"`
	Message string          `json:"message"`
	Status  int             `json:"status"`
}

func EncodeUsersResponse(w io.Writer, res *UsersResponse) error {
	return json.NewEncoder(w).Encode(res)
}

type CreateUserRequest struct {
	Username    string    `json:"username"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

func DecodeCreateUserRequest(r io.Reader) (*CreateUserRequest, error) {
	req := &CreateUserRequest{}
	err := json.NewDecoder(r).Decode(req)
	return req, err
}

type CreateUserResponse struct {
	CreatedUser *models.User `json:"createdUser"`
	Message     string       `json:"message"`
	Status      int          `json:"status"`
}

func EncodeCreateUserResponse(w io.Writer, res *CreateUserResponse) error {
	return json.NewEncoder(w).Encode(res)
}

type UpdateUserRequest struct {
	Username    string    `json:"username"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}

func DecodeUpdateUserRequest(r io.Reader) (*UpdateUserRequest, error) {
	req := &UpdateUserRequest{}
	err := json.NewDecoder(r).Decode(req)
	return req, err
}

type UpdateUserResponse struct {
	UpdatedUser *models.User `json:"updatedUser"`
	Message     string       `json:"message"`
	Status      int          `json:"status"`
}

func EncodeUpdateUserResponse(w io.Writer, res *UpdateUserResponse) error {
	return json.NewEncoder(w).Encode(res)
}

type DeleteUserResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func EncodeDeleteUserResponse(w io.Writer, res *DeleteUserResponse) error {
	return json.NewEncoder(w).Encode(res)
}
