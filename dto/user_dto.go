package dto

type CreateUserRequest struct {
	Id string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Level int `json:"level"`
	Password string `json:"password" binding:"required"`
}

type CreateUserResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Level int `json:"level"`
}

type SearchUserResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Level int `json:"level"`
}
