package models

import (
	"encoding/json"
	"io"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type UserID string

func NewUserID() UserID {
	return UserID(uuid.NewString())
}

func ToUserID(source string) (UserID, error) {
	bytes := source[:]
	if len(bytes) != 16 {
		return "", ErrCannotConvertToObjectID
	}
	var id = make([]byte, 16, 16)
	copy(id, bytes)
	return UserID(id), nil
}

type User struct {
	ID UserID `json:"id"`

	Username    string    `json:"username" validate:""`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Email       string    `json:"email" validate:"email"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

func NewUser(username, email string) *User {
	return &User{
		ID:        NewUserID(),
		CreatedAt: time.Now(),

		Username: username,
		Email:    email,
	}
}

func (u *User) SetFirstName(name string) *User {
	u.FirstName = name
	return u
}

func (u *User) SetLastName(name string) *User {
	u.LastName = name
	return u
}

func (u *User) SetDateOfBirth(t time.Time) *User {
	u.DateOfBirth = t
	return u
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func (u *User) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(u)
}

func (u *User) Update(update *User) {
	if update.Username != "" {
		u.Username = update.Username
	}
	if update.FirstName != "" {
		u.FirstName = update.FirstName
	}
	if update.LastName != "" {
		u.LastName = update.LastName
	}
	if update.Email != "" {
		u.Email = update.Email
	}
	if !update.DateOfBirth.IsZero() {
		u.DateOfBirth = update.DateOfBirth
	}
}

type UserList []*User

func (u UserList) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(u)
}
