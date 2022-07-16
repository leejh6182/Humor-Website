package repository

import (
	"github.com/gin-gonic/gin"
	. "src/dto"
	. "src/domain"
	"errors"
)

type UserRepository struct {
	db *DB
}

func (UserRepository *UserRepository) save(user User) string {

	return true
}