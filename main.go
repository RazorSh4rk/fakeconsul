package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/gin-gonic/gin"
)

var store map[string]string

func main() {
	warning :=
		figure.NewColorFigure("THIS IS A FAKE SERVICE", "", "red", true)
	warning.Print()
	warning =
		figure.NewColorFigure("ONLY USE FOR TESTING", "", "red", true)
	warning.Print()
	warning =
		figure.NewColorFigure("thanks :3c", "", "purple", true)
	warning.Print()

	store = make(map[string]string)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/v1/kv/*path", func(c *gin.Context) {
		path := c.Param("path")
		path = strings.Replace(path, "//", "/", 1)
		recurse := c.Query("recurse")

		if recurse == "" || recurse == "false" {
			// only get exact path
			val, ok := store[path]
			fmt.Printf("\nGET called with key: [ %s ], value is: [ %s ]\n", path, val)

			if !ok {
				c.JSON(404, gin.H{
					"error": "Key not found",
				})
				return
			}

			c.JSON(200, gin.H{
				"Key":   path,
				"Value": val,
			})
		} else {
			// get all keys matching path
			var values []gin.H
			for k, v := range store {
				if strings.HasPrefix(k, path) {
					values = append(values, gin.H{
						"Key":   k,
						"Value": v,
					})
					fmt.Print(". ")
				}
			}
			fmt.Printf("\nGET recursive called with key: [ %s ], values: [ %v ]\n", path, values)

			c.JSON(200, values)
		}
	})

	r.PUT("/v1/kv/*path", func(c *gin.Context) {
		path := c.Param("path")
		path = strings.Replace(path, "//", "/", 1)

		val, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(400, gin.H{"error": "Failed to read request body"})
			return
		}

		fmt.Printf("\nPUT called with key: [ %s ], value: [ %s ]\n", path, val)
		fmt.Printf("Query params: [ %v ]\n", c.Request.URL.Query())

		store[path] = string(val)
		c.JSON(200, gin.H{"success": true})
	})

	r.DELETE("/v1/kv/*path", func(c *gin.Context) {
		path := c.Param("path")
		path = strings.Replace(path, "//", "/", 1)

		fmt.Printf("\nDELETE called with key: [ %s ]\n", path)
		fmt.Printf("Query params: [ %v ]\n", c.Request.URL.Query())

		for k := range store {
			if strings.HasPrefix(k, path) {
				delete(store, k)
				fmt.Print(". ")
			}
		}
		fmt.Println()

		c.JSON(200, gin.H{"success": true})
	})

	r.GET("/debug/dump", func(c *gin.Context) {
		c.JSON(200, store)
	})

	r.GET("/debug/reset", func(c *gin.Context) {
		store = make(map[string]string)
		c.JSON(200, gin.H{"result": "ok"})
	})

	r.Run()
}
