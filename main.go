package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Bookmark struct {
	Href        string    `json:"href"`
	Description string    `json:"description"`
	Extended    string    `json:"extended"`
	Meta        string    `json:"meta"`
	Hash        string    `json:"hash"`
	Time        time.Time `json:"time"`
	Shared      string    `json:"shared"`
	Toread      string    `json:"toread"`
	Tags        string    `json:"tags"`
	Alive       bool
}

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	GIN_PORT := os.Getenv("PORT")
	GIN_MODE := os.Getenv("GIN_MODE")

	r := gin.Default()
	gin.SetMode(GIN_MODE)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	if err := r.Run(":" + GIN_PORT); err != nil {
		log.Fatal(err)
	}
}
