package dto

import (
    "time"
)

type CreatePostRequest struct {    
	UserId string `json:"userId" binding:"required"`
    Subject string `json:"subject" binding:"required"`
    Content string `json:"content"`
}

type CreatePostResponse struct {
    Id uint `json:"id"`
	UserId string `json:"userId"`
    Subject string `json:"subject"`
    CreatedAt time.Time `json:"createdAt"`
}

type SearchPostResponse struct {
    Id uint `json:"id"`
    Subject string `json:"subject"`
    Content string `json:"content"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}

type UpdateContentRequest struct {
    Content string `json:"content"`
}

type UpdateContentResponse struct {
    Id uint `json:"id"`
    Subject string `json:"subject"`
    Content string `json:"content"`
    UpdatedAt time.Time `json:"updatedAt"`
}
