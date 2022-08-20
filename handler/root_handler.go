package handler

import (
    "gorm.io/gorm"
)

type RootHandler struct {
    Db *gorm.DB
}

func (rootHandler *RootHandler) Init() {
    rootHandler.Db.AutoMigrate(&user{})
    rootHandler.Db.AutoMigrate(&post{})
    rootHandler.Db.AutoMigrate(&comment{})
}
