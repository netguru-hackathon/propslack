// Package main provides main
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()
	router.POST("/props", func(c *gin.Context) {
		user_id := c.PostForm("user_id")
		user_name := c.PostForm("user_name")

		c.JSON(200, gin.H{
			"status":  "posted",
			"user_id": user_id,
			"nick":    user_name,
		})
		resp, err := http.Get("https://slack.com/api/users.info?user=" + user_id + "&token=xoxp-52288129089-52448067687-54001525427-8d1a069a69")
		if err != nil {
			// handle error
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
		}
		fmt.Println(body)
	})
	router.Run(":" + port)
}
