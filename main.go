package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

// Struct for Cat Fact API response
type CatFactResponse struct {
	Fact string `json:"fact"`
}

func main() {
	r := gin.Default()

	// Optional: enable CORS for external access
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Next()
	})

	r.GET("/me", func(c *gin.Context) {
		client := resty.New()
		client.SetTimeout(5 * time.Second)

		var catfact CatFactResponse
		resp, err := client.R().SetResult(&catfact).Get("https://catfact.ninja/fact")

		var fact string
		if err != nil || resp.StatusCode() != http.StatusOK {
			fact = "Could not fetch a cat fact at this time. Try again later!"
		} else {
			fact = catfact.Fact
		}

		response := gin.H{
			"status": "success",
			"user": gin.H{
				"email": "atanda0x@gmail.com",
				"name":  "Atanda Nafiu",
				"stack": "Go/Gin, Python/Django, JavaScript/Node.js, Java/Spring Boot, C#/.NET",
			},
			"timestamp": time.Now().UTC().Format(time.RFC3339Nano),
			"fact":      fact,
		}

		c.JSON(http.StatusOK, response)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
