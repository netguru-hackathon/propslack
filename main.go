// Package main provides main
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()
	router.POST("/props", func(c *gin.Context) {
		user_id := c.PostForm("user_id")
		user_name := c.PostForm("user_name")

		fmt.Printf("%+v\n", convertNameToID("wojzag"))

		c.JSON(200, gin.H{
			"status":  "posted",
			"user_id": user_id,
			"nick":    user_name,
		})
	})
	router.Run(":" + port)
}

func convertNameToID(s string) (p Profile) {
	token := os.Getenv("TOKEN")
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
