// main.go
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/api/houses", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "houses will go here",
        })
    })

    r.Run(":5000")
}