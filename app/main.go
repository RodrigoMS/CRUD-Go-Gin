package main

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    _ "github.com/lib/pq"
    "github.com/gin-gonic/gin"
    "github.com/RodrigoMS/CRUD-Go-Gin/controllers"
)

const (
    DB_USER     = "youruser"
    DB_PASSWORD = "yourpassword"
    DB_NAME     = "yourdbname"
    DB_HOST     = "db"
    DB_PORT     = "5432"
)

func main() {
    var db *sql.DB
    var err error

    for i := 0; i < 10; i++ {
        db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME))
        if err != nil {
            log.Printf("Failed to connect to database: %v", err)
            time.Sleep(2 * time.Second)
            continue
        }

        err = db.Ping()
        if err == nil {
            break
        }

        log.Printf("Failed to ping database: %v", err)
        time.Sleep(2 * time.Second)
    }

    if err != nil {
        log.Fatalf("Could not connect to database: %v", err)
    }

    defer db.Close()

    r := gin.Default()

    r.GET("/items", func(c *gin.Context) {
        controllers.GetItems(c, db)
    })

    r.POST("/items", func(c *gin.Context) {
        controllers.CreateItem(c, db)
    })

    r.GET("/items/:id", func(c *gin.Context) {
        controllers.GetItem(c, db)
    })

    r.PUT("/items/:id", func(c *gin.Context) {
        controllers.UpdateItem(c, db)
    })

    r.DELETE("/items/:id", func(c *gin.Context) {
        controllers.DeleteItem(c, db)
    })

    r.Run()
}
