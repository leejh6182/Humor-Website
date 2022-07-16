package main

import (
    "github.com/gin-gonic/gin"
    . "src/handler"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

func initializeSchema(db *DB)
{
    db.AutoMigrate(&User{})
}

func main() {

    //DB connection
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("Fail to connect to the DB")
    }

    initializeSchema(db)

    //initialize object
    userRepository := UserRepository{db}
    userHandler := UserHandler{userRepository}

    router := gin.Default()
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    
    router.GET("/users", userHandler.SearchUser)
    router.POST("/users", userHandler.CreateUser)
    
    router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

