package main

import (
	"github.com/gin-gonic/gin"
	"sync"
)


var requestCount int

var mu sync.Mutex

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("index.html")
	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
		incrementRequestCount()
		c.HTML(200, "index.html", gin.H{
			"count":getRequestCount(),
		})
	})

	router.Run("localhost:8080")
}

func incrementRequestCount(){
	mu.Lock()
	requestCount++
	mu.Unlock()
}

func getRequestCount() int{
	mu.Lock()
	defer mu.Unlock()
	return requestCount
}