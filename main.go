package main

import (
	"fmt"
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

		fmt.Printf("\nGET called with key: [ %s ], value is: [ %s ]\n", key, val)

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

		fmt.Printf("\nPUT called with key: [ %s ], value: [ %s ]\n", key, val)

		store[key] = string(val)
		c.JSON(200, gin.H{"success": true})
	})

	r.DELETE("/v1/kv/:key", func(c *gin.Context) {
		key := c.Param("key")
		_, ok := store[key]

		fmt.Printf("\nDELETE called with key: [ %s ], exists: [ %t ]\n", key, ok)

		if !ok {
			c.JSON(404, gin.H{"error": "Key not found"})
			return
		}

		delete(store, key)
		c.JSON(200, gin.H{"success": true})
	})

	r.Run()
}
