package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/gzip"
)

// Sequence is sequence
type Sequence struct {
	Index int `json:"index" binding:"required"`
	Value int `json:"value"`
}

func main() {
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.GET("/even", evenNosStreamer)
	router.GET("/fibonacci", fibonacciNosStreamer)
	router.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}

func evenNosStreamer(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	chanStream := make(chan int, 5000)
	go func() {
		defer close(chanStream)
		for i := 2; true; i+=2 {
			chanStream <- i
		}
	}()	
	c.String(http.StatusOK, "[")
	i := 1
	for value := range chanStream {
		if i > 1 {
			c.String(http.StatusOK, ",")
		}
		c.JSON(http.StatusOK, Sequence{i, value})
		i++
	}
	c.String(http.StatusOK, "]")
}

func fibonacciNosStreamer(c *gin.Context) {
	var fibs []Sequence
	fibs = append(fibs, Sequence{1, 0}, Sequence{2, 1})
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, fibs)
}
