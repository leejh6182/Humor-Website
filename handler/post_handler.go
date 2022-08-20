package handler

import (
	"github.com/gin-gonic/gin"
	. "src/dto"
    "gorm.io/gorm"
    "errors"
)

func (rootHandler *RootHandler)CreatePost(c *gin.Context) {

	createPostRequest := CreatePostRequest{}

	err := c.Bind(&createPostRequest)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

    user := user{}
    dbErr := rootHandler.Db.Where("user_id = ?", createPostRequest.UserId).First(&user)
    if dbErr.Error != nil {
        c.JSON(400, dbErr.Error)
        return
    }

    post := post{ UserId: user.ID,
                  Subject: createPostRequest.Subject,
                  Content: createPostRequest.Content,
                  Like: 0,
                  Dislike: 0,
                  Comments: []comment{}}

	dbErr = rootHandler.Db.Create(&post)
	if dbErr.Error != nil {
		c.JSON(400, dbErr.Error)
		return
	} 

	result := []CreatePostResponse{ CreatePostResponse{ post.ID, user.UserId, post.Subject, post.CreatedAt } }

	c.JSON(200, gin.H{"data": result,})
	return
}

func(rootHandler *RootHandler) SearchAllPost(c *gin.Context) {

	result := []SearchPostResponse{}

	userId := c.Query("userId")
    if userId != "" {
        user := user{}
        dbErr := rootHandler.Db.Preload("Posts").First(&user, "users.user_id = ?", userId)
        if dbErr.Error != nil {
            if !errors.Is(dbErr.Error, gorm.ErrRecordNotFound) {           
                c.JSON(400, dbErr.Error)
                return
            } else {
    	        c.JSON(200, gin.H{"data": []string{},})
                return
            }
        }  
    
        for _, post := range user.Posts {
            searchPostResponse := SearchPostResponse{Id: post.ID, 
                                                     Subject: post.Subject, 
                                                     Content: post.Content, 
                                                     CreatedAt: post.CreatedAt,
                                                     UpdatedAt: post.UpdatedAt }
            
            result = append(result, searchPostResponse)
        }
                
    } else {
        posts := []post{}
        dbErr := rootHandler.Db.Find(&posts)
        if dbErr.Error != nil {
            c.JSON(400, dbErr.Error)
            return  
        }

        for _, post := range posts {
            searchPostResponse := SearchPostResponse{Id: post.ID, 
                                                     Subject: post.Subject, 
                                                     Content: post.Content, 
                                                     CreatedAt: post.CreatedAt,
                                                     UpdatedAt: post.UpdatedAt }
            
            result = append(result, searchPostResponse)
        }
    }
   
	c.JSON(200, gin.H{"data": result,})
	return
}


func (rootHandler *RootHandler) UpdateContent(c *gin.Context) {

    id := c.Param("id")

 	updateContentRequest := UpdateContentRequest{}

    err := c.Bind(&updateContentRequest)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
   
    post := post{}
    dbErr := rootHandler.Db.Where("id = ?", id).First(&post)
    if dbErr.Error != nil {
        c.JSON(400, dbErr.Error)
        return
    }

    post.Content = updateContentRequest.Content

    dbErr = rootHandler.Db.Save(&post)
    if dbErr.Error != nil {
        c.JSON(400, dbErr.Error)
        return
    }

	result := []UpdateContentResponse{}
        
    updateContentResponse := UpdateContentResponse{post.ID, post.Subject, post.Content, post.UpdatedAt}
    result = append(result, updateContentResponse)
    
    c.JSON(200, gin.H{"data": result,})
    return 	
}


