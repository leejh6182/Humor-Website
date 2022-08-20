package handler

import (
	"github.com/gin-gonic/gin"
	. "src/dto"
    //. "fmt"
    "gorm.io/gorm"
    "errors"
)

func (rootHandler *RootHandler)CreateUser(c *gin.Context) {

	createUserRequest := CreateUserRequest{}

	err := c.Bind(&createUserRequest)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

    user := user{ UserId: createUserRequest.UserId,
                  Name: createUserRequest.Name,
                  Address: createUserRequest.Address,
                  Email: createUserRequest.Email,
                  Password: createUserRequest.Password }

	dbErr := rootHandler.Db.Create(&user)
	if dbErr.Error != nil {
		c.JSON(400, dbErr.Error)
		return
	} 

	result := []CreateUserResponse{ CreateUserResponse{ user.UserId, user.Name, user.Address, user.Email, user.Level } }

	c.JSON(200, gin.H{"data": result,})

	return
}

func(rootHandler *RootHandler) SearchUser(c *gin.Context) {

	id := c.Query("userId")

	result := []SearchUserResponse{}

	if id != "" {
        user := user{}
        dbErr := rootHandler.Db.Where("user_id = ?", id).First(&user)
        if dbErr.Error != nil {
            if !errors.Is(dbErr.Error, gorm.ErrRecordNotFound) {           
                c.JSON(400, dbErr.Error)
                return
            } else {
	            c.JSON(200, gin.H{"data": []string{},})
                return
            }
        }  

        searchUserResponse := SearchUserResponse{user.UserId, user.Name, user.Address, user.Email, user.Level}
        result = append(result, searchUserResponse)
	} else {
        users := []user{}
        dbErr := rootHandler.Db.Find(&users)
        if dbErr.Error != nil {
            c.JSON(400, dbErr.Error)
            return  
        }

        for _, user := range users {
	        searchUserResponse := SearchUserResponse{user.UserId, user.Name, user.Address, user.Email, user.Level}
	    	result = append(result, searchUserResponse)
        } 
    }

	c.JSON(200, gin.H{"data": result,})
	return
}

func (rootHandler *RootHandler) UpdateUser(c *gin.Context) {
 	updateUserRequest := UpdateUserRequest{}

	err := c.Bind(&updateUserRequest)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
   
	user := user{}
    
    dbErr := rootHandler.Db.Where("user_id = ?", updateUserRequest.UserId).First(&user)
    if dbErr.Error != nil {
        c.JSON(400, dbErr.Error)
        return
    }

    user.Password = updateUserRequest.Password
    user.Address = updateUserRequest.Address
    user.Email = updateUserRequest.Email

    dbErr = rootHandler.Db.Save(&user)
    if dbErr.Error != nil {
        c.JSON(400, dbErr.Error)
        return
    }

	result := []SearchUserResponse{}
        
	searchUserResponse := SearchUserResponse{user.UserId, user.Name, user.Address, user.Email, user.Level}
    result = append(result, searchUserResponse)
    
    c.JSON(200, gin.H{"data": result,})
    return 	
}

