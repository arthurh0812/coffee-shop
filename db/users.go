package db

import (
	"errors"

	"github.com/arthurh0812/coffee-shop/models"
)

type QueryOptions struct {
}

var ErrNotFound = errors.New("no matching user could be found")

type UserDB struct{}

func GetAllUsers(opts ...QueryOptions) models.UserList {
	return userList
}

func GetUserByID(id models.UserID) *models.User {
	for _, user := range userList {
		if user.ID == id {
			return user
		}
	}
	return nil
}

func GetUserByUsername(name string) (*models.User, error) {
	for _, user := range userList {
		if user.Username == name {
			return user, nil
		}
	}
	return nil, ErrNotFound
}

func CreateUser(u *models.User) *models.User {
	return nil
}

func UpdateUserByID(id models.UserID, update *models.User) *models.User {
	for _, user := range userList {
		if user.ID == id {
			user = update
			return update
		}
	}
	return nil
}

func DeleteUserByID(id models.UserID) (deleted *models.User) {
	toDelete, deleted := -1, (*models.User)(nil)
	for i, user := range userList {
		if user.ID == id {
			toDelete = i
			deleted = user
			break
		}
	}
	if toDelete == -1 {
		return nil
	}
	userList = append(userList[:toDelete], userList[toDelete+1:]...)
	return deleted
}

var userList = models.UserList{
	{
		Username: "max45",
	},
}
