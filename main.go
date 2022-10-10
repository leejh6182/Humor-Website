package main

import (
    "github.com/gin-gonic/gin"
    . "src/handler"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "gorm.io/driver/postgres"
    "log"
    "os"
    "time"
)

func CORSMiddleware() gin.HandlerFunc {

    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
   
        if c.Request.URL.Path != "/login"{

            tokenString, err := c.Cookie("token")
            if err != nil {
                c.JSON(400, err)
                return
            }

            //Validate token
            _, err = ValidateToken(tokenString)
            if err != nil {
                c.JSON(401, err)
                return
            }
        }

        c.Next()
    }
}

func main() {

    //
    newLogger := logger.New(
                    log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
                    logger.Config{
                        SlowThreshold:              time.Second,   // Slow SQL threshold
                        LogLevel:                   logger.Info, // Log level
                        IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
                        Colorful:                  false,          // Disable color
                    },
                )

    //DB connection
    conn := "host=localhost user=jhlee password=jhlee dbname=mydb sslmode=disable"
    db, err := gorm.Open(postgres.Open(conn), &gorm.Config{Logger: newLogger,})
    if err != nil {
        panic("Fail to connect to the DB")
    }

    //Dependency injection
    rootHandler := RootHandler{db}
    rootHandler.Init()

    router := gin.Default()
//  router.LoadHTMLGlob("templates/*.html")
    router.Use(CORSMiddleware())


    router.GET("/users", rootHandler.SearchUser)
    router.POST("/users", rootHandler.CreateUser)
    router.PUT("/users", rootHandler.UpdateUser)

    router.POST("/posts", rootHandler.CreatePost)
    router.GET("/posts", rootHandler.SearchAllPost)
    router.PUT("/posts/:id", rootHandler.UpdateContent)
    
//    router.POST("/posts/:id/comments", rootHandler.AddComment)
//    router.POST("/posts/:id/like", rootHandler.UpdateContent)    
//    router.PUT("/posts/:id/dislike", rootHandler.UpdateContent)

    router.POST("/login", rootHandler.Login)
    
    router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

