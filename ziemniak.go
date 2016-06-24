// Package main provides main
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProfileJson struct {
	Email string `json:"email"`
}
type Member struct {
	Profile ProfileJson `json:"profile"`
	Id      string      `json:"id"`
	Name    string      `json:"name"`
}
type User struct {
	Members []Member `json:"members"`
}

func main() {
	port := os.Getenv("PORT")
	token := os.Getenv("TOKEN")
	separator := "->"

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()
	router.POST("/props", func(c *gin.Context) {
		// user_id := c.PostForm("user_id")
		// user_name := c.PostForm("user_name")
		text := c.PostForm("text")
		fmt.Println(text)

		resp, err := http.Get("https://slack.com/api/users.list?token=" + token)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		var dat User
		if err := json.Unmarshal(body, &dat); err != nil {
			panic(err)
		}
		if err != nil {
			fmt.Println(err.Error())
		}

		paramsArray := strings.Split(text, separator)
		mentions := paramsArray[0]
		props := paramsArray[1]

		if mentions == "" || props == "" {
			c.JSON(400, gin.H{"message": "Nope"})
			return
		}

		mentionsArray := strings.Split(mentions, ",")
		for _, mention := range mentionsArray {
			mention := strings.TrimSpace(mention)
			if mention[0] == '@' {
				mention = strings.TrimLeft(mention, "@")
			}
			fmt.Println(mention)
		}

		c.JSON(200, gin.H{"message": "Yup"})
	})
	router.Run(":" + port)
}
