package main

import (
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Apikey  string `json:"apikey"`
	Numbers string `json:"numbers"`
	Message string `json:"message"`
	Sender  string `json:"sender"`
}

func main() {
	const myurl = "http://localhost:8001"

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "Server is Running"})
	})

	router.POST("/req", func(c *gin.Context) {
		var req Request
		c.BindJSON(&req)
		Url, err := url.Parse("http://https://api.textlocal.in/send/")
		if err != nil {
			panic(err)
		}

		parameters := url.Values{}
		parameters.Add("apiKey", req.Apikey)
		parameters.Add("numbers", req.Numbers)
		parameters.Add("message", req.Message)
		parameters.Add("sender", req.Sender)
		Url.RawQuery = parameters.Encode()

		fmt.Printf("Encoded URL is %q\n", Url.String())

		c.JSON(200, req)
	})
	router.Run()
}
