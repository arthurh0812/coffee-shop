package handlers

import (
	"context"
	//"context"
	"fmt"
	"log"
	"net/http"

	//"github.com/dgrijalva/jwt-go"

	"github.com/gorilla/mux"

	"github.com/arthurh0812/coffee-shop/db"
	"github.com/arthurh0812/coffee-shop/models"

	//"github.com/arthurh0812/coffee-shop/db"
	"github.com/arthurh0812/coffee-shop/schema"
)

type Users struct {
	handler
}

var users *Users

type UserIDKey struct{}

func NewUsers(l *log.Logger) *Users {
	if users == nil {
		users = &Users{handler: newHandler("Auth", l)}
	}
	return users
}

func (u *Users) GetAllUsers(w http.ResponseWriter, req *http.Request) {
	list := db.GetAllUsers()
	err := schema.EncodeUsersResponse(w, &schema.UsersResponse{
		Users:   list,
		Count:   len(list),
		Message: "Successfully queried users from DB!",
		Status:  http.StatusOK,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to send JSON response: %v", err), http.StatusInternalServerError)
	}
}

func (u *Users) ExtractID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		i, ok := mux.Vars(req)["id"]
		if !ok {
			http.Error(w, "please provide an ID parameter in the URI", http.StatusBadRequest)
			return
		}
		id, err := models.ToUserID(i)
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid product ID: %v", err), http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(req.Context(), UserIDKey{}, id)
		req = req.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}

func (u *Users) GetUserByID(w http.ResponseWriter, req *http.Request) {
	id := req.Context().Value(UserIDKey{}).(models.UserID) // assert type
	user := db.GetUserByID(id)
	if user == nil {
		http.Error(w, "failed to query user from DB: not found", http.StatusNotFound)
		return
	}
	err := schema.EncodeUserResponse(w, &schema.UserResponse{
		User:    user,
		Message: "Successfully queried user from DB!",
		Status:  http.StatusOK,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to send JSON response: %v", err), http.StatusInternalServerError)
	}
}

func (u *Users) PreUpdateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, err := schema.DecodeUpdateUserRequest(r.Body)
		if err != nil {
			err = schema.EncodeUpdateUserResponse(w, &schema.UpdateUserResponse{
				UpdatedUser: nil,
				Message:     fmt.Sprintf("failed to decode JSON from request body: %v", err),
				Status:      http.StatusInternalServerError,
			})
			if err != nil {
				http.Error(w, fmt.Sprintf("failed to send JSON response: %v", err), http.StatusInternalServerError)
				return
			}
		}
		ctx := context.WithValue(r.Context(), UserRequestKey{}, req)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func (u *Users) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(UserIDKey{}).(models.UserID)
	req := r.Context().Value(UserRequestKey{}).(*schema.UpdateUserRequest)
	user := db.GetUserByID(id)
	user.Update(models.NewUser(req.Username, "").SetFirstName(req.FirstName).SetLastName(req.LastName).SetDateOfBirth(req.DateOfBirth))
	updated := db.UpdateUserByID(id, user)
	if updated == nil {
		err := schema.EncodeUpdateUserResponse(w, &schema.UpdateUserResponse{
			UpdatedUser: nil,
			Message:     "failed to update user: not found",
			Status:      http.StatusNotFound,
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to send JSON response: %v", err), http.StatusInternalServerError)
			return
		}
	}
	err := schema.EncodeUpdateUserResponse(w, &schema.UpdateUserResponse{
		UpdatedUser: updated,
		Message:     "Successfully updated user from DB!",
		Status:      http.StatusOK,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to send JSON response: %v", err), http.StatusInternalServerError)
	}
}

type UserRequestKey struct{}

func (u *Users) PreCreateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, err := schema.DecodeCreateProductsRequest(r.Body)
		if err != nil {
			err = schema.EncodeCreateUserResponse(w, &schema.CreateUserResponse{
				CreatedUser: nil,
				Message:     fmt.Sprintf("failed to decode JSON from request body: %v", err),
				Status:      http.StatusInternalServerError,
			})
			if err != nil {
				http.Error(w, fmt.Sprintf("failed to send JSON response: %v", err), http.StatusInternalServerError)
				return
			}
		}
		ctx := context.WithValue(r.Context(), UserRequestKey{}, req)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func (u *Users) CreateUser(w http.ResponseWriter, r *http.Request) {
	req := r.Context().Value(UserRequestKey{}).(*schema.CreateUserRequest)

	user := models.NewUser(req.Username, req.Email).SetFirstName(req.FirstName).SetLastName(req.LastName).SetDateOfBirth(req.DateOfBirth)
	err := user.Validate()
	if err != nil {
		err = schema.EncodeCreateUserResponse(w, &schema.CreateUserResponse{
			CreatedUser: nil,
			Message:     fmt.Sprintf("invalid input: %v", err),
			Status:      http.StatusBadRequest,
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to send JSON response: %v", err), http.StatusInternalServerError)
			return
		}
	}
	user = db.CreateUser(user)
	err = schema.EncodeCreateUserResponse(w, &schema.CreateUserResponse{
		CreatedUser: user,
		Message:     "Successfully created user!",
		Status:      http.StatusCreated,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to send JSON response: %v", err), http.StatusInternalServerError)
	}
}

func (u *Users) DeleteUserByID(w http.ResponseWriter, req *http.Request) {
	id := req.Context().Value(UserIDKey{}).(models.UserID)
	deleted := db.DeleteUserByID(id)
	if deleted == nil {
		err := schema.EncodeDeleteUserResponse(w, &schema.DeleteUserResponse{
			Message: "failed to delete user: not found",
			Status:  http.StatusNotFound,
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to send JSON response: %v", err), http.StatusInternalServerError)
			return
		}
	}
	err := schema.EncodeDeleteUserResponse(w, &schema.DeleteUserResponse{
		Message: "Successfully deleted user from DB!",
		Status:  http.StatusOK,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to send JSON response: %v", err), http.StatusInternalServerError)
	}
}
