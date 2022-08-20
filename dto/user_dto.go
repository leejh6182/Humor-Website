package dto

type CreateUserRequest struct {
	UserId string `json:"userId" binding:"required"`
	Name string `json:"name" binding:"required"`
    Address string `json:"address"`
    Email string `json:"email"`
	Level int `json:"level"`
	Password string `json:"password" binding:"required"`
}

type CreateUserResponse struct {
	UserId string `json:"userId"`
	Name string `json:"name"`
    Address string `json:"address"`
    Email string `json:"email"`
	Level int `json:"level"`
}

type SearchUserResponse struct {
	UserId string `json:"userId"`
	Name string `json:"name"`
    Address string `json:"address"`
    Email string `json:"email"`
	Level int `json:"level"`
}

type UpdateUserRequest struct {
    UserId string `json:"userId" binding:"required"`
    Address string `json:"address"`
    Email string `json:"email"`
    Password string `json:"password"`
}
