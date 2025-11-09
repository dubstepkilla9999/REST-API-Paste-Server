package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type contentWrapper struct {
	Content string `json:"content"`
}

type entry struct {
	ID string `json:"id"`
	contentWrapper
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

var entries = []entry{}

func main() {
	router := gin.Default()
	router.POST("/storage", postEntry)
	router.GET("/storage/:id", getEntryByID)

	router.Run("0.0.0.0:80")
}

func postEntry(c *gin.Context) {
	var newContent contentWrapper

	if err := c.BindJSON(&newContent); err != nil {
		return
	}

	newEntry := entry{
		ID:             generateRandomString(20),
		contentWrapper: newContent,
	}

	entries = append(entries, newEntry)
	c.IndentedJSON(http.StatusCreated, newEntry)
}

func getEntryByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range entries {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "entry not found"})
}

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
