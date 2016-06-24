// Package main provides main
package main

import (
  "log"
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
  })
  router.Run(":" + port)
}
