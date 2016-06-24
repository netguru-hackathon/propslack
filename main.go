// Package main provides main
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type Profile struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
type Member struct {
	Profile Profile `json:"profile"`
	Id      string  `json:"id"`
	Name    string  `json:"name"`
}
type User struct {
	Members []Member `json:"members"`
}

func convertNameToID(s string, token string) (p Profile) {
	resp, err := http.Get("https://slack.com/api/users.list?token=" + token)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		panic(err)
	}

	for i := 0; i < len(user.Members); i++ {
		if user.Members[i].Name == s {
			p.Email = user.Members[i].Profile.Email
			p.FirstName = user.Members[i].Profile.FirstName
			p.LastName = user.Members[i].Profile.LastName
			return
		}
	}
	return
}

func sendPropsInfo(s string, webhook_url string) {
	var jsonStr = []byte(`{"text":"` + s + `"}`)
	http.Post(webhook_url, "application/json", bytes.NewBuffer(jsonStr))
}

func main() {
	port := os.Getenv("PORT")
	// token := os.Getenv("TOKEN")
	webhook_url := os.Getenv("WEBHOOK_URL")
	separator := "->"

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()
	router.POST("/props", func(c *gin.Context) {
		userName := c.PostForm("user_name")
		text := c.PostForm("text")
		fmt.Println(text)

		// fmt.Printf("%+v\n", convertNameToID("wojzag", token))

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

		c.JSON(200, gin.H{
			"status": "posted",
		})
		sendPropsInfo(userName+" dał propsa "+strings.Join(mentionsArray[:], ",")+" za "+props, webhook_url)
	})
	router.Run(":" + port)
}
