package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Msg": "Hallo, salam dari antah-berantah",
		})
	})
	r.GET("/bumi", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Msg": "Hallo, salam dari Bumi",
		})
	})
	r.GET("/mars", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Msg": "Hallo, salam dari Mars",
		})
	})
	r.Run(":3000")
}
