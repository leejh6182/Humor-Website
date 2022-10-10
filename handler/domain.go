package handler

import (
    "gorm.io/gorm"
)

type user struct {
    gorm.Model
	UserId string `gorm:"unique"`
	Name string
    Address string
    Email string
	Level int
	Password string
    Posts []post 
}

type post struct {
    gorm.Model
    UserId uint
    Subject string
    Content string
    Like int
    Dislike int
    Comments []comment
}

type comment struct{
    gorm.Model
    PostId uint
    Message string    
    CommentId uint    
    Comments []comment
}

type authTokenClaim struct {
    UserId string
    Email string
    Level int
}

