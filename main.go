package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	GIN_PORT := os.Getenv("PORT")
	GIN_MODE := os.Getenv("GIN_MODE")
	PINBOARD_API_TOKEN := os.Getenv("PINBOARD_API_TOKEN")
	PINBOARD_API_URL := os.Getenv("PINBOARD_API_URL")
	//PINBOARD_JSON_DATAFILE := os.Getenv("PINBOARD_JSON_DATAFILE")
	//USE_PINBOARD_API := os.Getenv("USE_PINBOARD_API")

	pbClient := NewPinboardClient(PINBOARD_API_URL, PINBOARD_API_TOKEN)

	r := gin.Default()
	gin.SetMode(GIN_MODE)

	//Test

	r.GET("/recent", func(c *gin.Context) {
		bookmarks, err := pbClient.GetRecentBookmarks("", 10)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, bookmarks)
	})

	//Run the server
	if err := r.Run(":" + GIN_PORT); err != nil {
		log.Fatal(err)
	}
}

//func LoadPinBoardDataFromFile(jsonFilePath string) (*[]Bookmark, error) {
//	var bookmarks []Bookmark
//	// Open our jsonFile
//	jsonFile, err := os.Open(jsonFilePath)
//	// if we os.Open returns an error then handle it
//	if err != nil {
//		return nil, err
//	}
//	defer jsonFile.Close()
//	decoder := json.NewDecoder(jsonFile)
//	err = decoder.Decode(&bookmarks)
//	if err != nil {
//		return nil, err
//	}
//	return &bookmarks, nil
//}
