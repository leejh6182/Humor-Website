package repository

import (
    "gorm.io/gorm"
    //"gorm.io/driver/postgres"
    . "src/model"
    //"errors"
)

type UserRepository struct {
	Db *gorm.DB
}

func (userRepository *UserRepository) Save(user User) (*User, error) {
    result := userRepository.Db.Create(&user)
    if result.Error != nil {
	return nil, result.Error
    }

    return &user, nil
}

func (userRepository *UserRepository) Find(id string) (*User) {
    user := User{}
    result := userRepository.Db.Where("id = ?", id).First(&user)
    if result.Error != nil {
        return nil
    }

    return &user
}
     
func (userRepository *UserRepository) FindAll() ([]*User) {
    users := []*User{}
    result := userRepository.Db.Find(&users)
    if result.Error != nil {
        return users
    }

    return users
}

func (userRepository *UserRepository) Update(user *User) *User {
    userRepository.Db.Save(user)
    return user
}

