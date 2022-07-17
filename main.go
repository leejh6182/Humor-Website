package main

import (
    "github.com/gin-gonic/gin"
    . "src/handler"
    . "src/repository"
    . "src/model"
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
)

func initializeSchema(db *gorm.DB){
    db.Migrator().CreateTable(&User{})
}

func main() {

    //DB connection
    conn := "host=localhost user=jhlee password=jhlee dbname=mydb sslmode=disable"
    db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
    if err != nil {
        panic("Fail to connect to the DB")
    }

    //Initialize DB schema
    initializeSchema(db)
    
    //Dependency injection
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

