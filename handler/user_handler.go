package handler

import (
	"github.com/gin-gonic/gin"
	. "src/dto"
	. "src/domain"
	"errors"
)

var users = []User{}

type UserHandler struct {
	userRepository UserRepository
}

func (UserHandler *UserHandler) ValidateDuplicateUser(id string) bool {
	for _, user := range users {
		if id == user.Id() {
			return false
		}
	}
	return true
}

func (UserHandler *UserHandler) CreateUser(c *gin.Context) {

	createUserRequest := CreateUserRequest{}

	err := c.Bind(&createUserRequest)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	if UserHandler.ValidateDuplicateUser(createUserRequest.Id) == false {
		err := errors.New("User with id '" + createUserRequest.Id + "' already exists")
		c.JSON(400, err.Error())
		return
	}

	user := NewUser(createUserRequest.Id, createUserRequest.Name, createUserRequest.Level, createUserRequest.Password)
	users = append(users, *user)

	createUserResponse := CreateUserResponse{user.Id(), user.Name(), user.Level()}

	result := []CreateUserResponse{createUserResponse}

	c.JSON(200, gin.H{"data": result,})

	return
}

func (UserHandler *UserHandler) SearchUser(c *gin.Context) {

	id := c.Query("id")

	result := []SearchUserResponse{}
	for _, user := range users {
		if id == "" {
			searchUserResponse := SearchUserResponse{user.Id(), user.Name(), user.Level()}
			result = append(result, searchUserResponse)
		} else if id == user.Id() {
			searchUserResponse := SearchUserResponse{user.Id(), user.Name(), user.Level()}
			result = append(result, searchUserResponse)
			break
		}
	}

	c.JSON(200, gin.H{"data": result,})
	return
}
