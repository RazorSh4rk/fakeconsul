package main

import (
	"io"

	"github.com/gin-gonic/gin"
)

var store map[string]string

func main() {
	store = make(map[string]string)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/v1/kv/:key", func(c *gin.Context) {
		key := c.Param("key")
		val, ok := store[key]

		if !ok {
			c.JSON(404, gin.H{
				"error": "Key not found",
			})
			return
		}

		c.JSON(200, gin.H{
			"Key":   key,
			"Value": val,
		})
	})

	r.PUT("/v1/kv/:key", func(c *gin.Context) {
		key := c.Param("key")

		val, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(400, gin.H{"error": "Failed to read request body"})
			return
		}

		store[key] = string(val)
		c.JSON(200, gin.H{"success": true})
	})

	r.DELETE("/v1/kv/:key", func(c *gin.Context) {
		key := c.Param("key")
		_, ok := store[key]

		if !ok {
			c.JSON(404, gin.H{"error": "Key not found"})
			return
		}

		delete(store, key)
		c.JSON(200, gin.H{"success": true})
	})

	r.Run()
}
