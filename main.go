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
	token := os.Getenv("TOKEN")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()
	router.POST("/props", func(c *gin.Context) {
		user_id := c.PostForm("user_id")
		user_name := c.PostForm("user_name")

		resp, err := http.Get("https://slack.com/api/users.info?user=" + user_id + "&token=" + token)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		s := string(body)
		fmt.Println(s)
		if err != nil {
			fmt.Println(err.Error())
		}

		c.JSON(200, gin.H{
			"status":  "posted",
			"user_id": user_id,
			"nick":    user_name,
		})
	})
	router.Run(":" + port)
}
