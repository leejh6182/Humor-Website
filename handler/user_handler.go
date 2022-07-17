package handler

import (
	"github.com/gin-gonic/gin"
	. "src/dto"
	. "src/model"
        . "src/repository"
	"errors"
	"time"
        . "fmt" 
)

var users = []User{}

type UserHandler struct {
	UserRepository UserRepository
}

func (userHandler *UserHandler) ValidateDuplicateUser(id string) bool {
	for _, user := range users {
		if id == user.Id {
			return false
		}
	}
	return true
}

func (userHandler *UserHandler) CreateUser(c *gin.Context) {

	createUserRequest := CreateUserRequest{}

	err := c.Bind(&createUserRequest)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	if userHandler.ValidateDuplicateUser(createUserRequest.Id) == false {
		err := errors.New("User with id '" + createUserRequest.Id + "' already exists")
		c.JSON(400, err.Error())
		return
	}

	user := User {Id: createUserRequest.Id, 
                      Name: createUserRequest.Name, 
                      Level: createUserRequest.Level, 
                      Password: createUserRequest.Password,
                      CreatedAt: time.Now(),
                      UpdatedAt: time.Now()}

	newUser, err := userHandler.UserRepository.Save(user)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	createUserResponse := CreateUserResponse{newUser.Id, newUser.Name, newUser.Level}

	result := []CreateUserResponse{createUserResponse}

	c.JSON(200, gin.H{"data": result,})

	return
}

func (userHandler *UserHandler) SearchUser(c *gin.Context) {

	id := c.Query("id")

	result := []SearchUserResponse{}

	if id != "" {
            user := userHandler.UserRepository.Find(id)
	    searchUserResponse := SearchUserResponse{user.Id, user.Name, user.Level}
            result = append(result, searchUserResponse)
	} else {
	    users := userHandler.UserRepository.FindAll()

	    Println("Complete in fetching")

            for _, user := range users {
	    	searchUserResponse := SearchUserResponse{user.Id, user.Name, user.Level}
	    	result = append(result, searchUserResponse)
            } 
	    
            Println("Complete filling data")
        }

	c.JSON(200, gin.H{"data": result,})
	return
}
