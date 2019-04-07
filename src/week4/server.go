package main

import "github.com/gin-gonic/gin"

func main()  {
	server := gin.Default()
	server.Handle("GET", "/", func(c *gin.Context) {
		c.String(200, "Hello,  world")	})

	server.Run(":8888")
}
