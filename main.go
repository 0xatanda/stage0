package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type CatFactResponse struct {
	Fact string `json:"fact"`
}

func main() {
	r := gin.Default()

	r.GET("/catfact", func(c *gin.Context) {
		client := resty.New()
		client.SetTimeout(5 * time.Second)

		// Fetch cat fact from external API
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
				"stack": "Go/Gin, Python/Django, JavaScript/Node.js, JAVA/Spring Boot, C#/.NET",
			},
			"timestamp": time.Now().Format(time.RFC3339Nano),
			"fact":      fact,
		}
		c.JSON(http.StatusOK, response)
	})
	r.Run(":8080")
}
