package kv

import (
	"fmt"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
)

var store map[string]string

func extractKey(c *gin.Context) string {
	path := c.Param("path")
	path = strings.Replace(path, "//", "/", 1) // filter accidental "/key" instead of "key" calls
	if strings.HasPrefix(path, "/") {
		path = strings.Replace(path, "/", "", 1)
	}
	return path
}

func GetHandler(c *gin.Context) {
	path := extractKey(c)
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
}

func PutHandler(c *gin.Context) {
	path := extractKey(c)

	val, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to read request body"})
		return
	}

	fmt.Printf("\nPUT called with key: [ %s ], value: [ %s ]\n", path, val)
	fmt.Printf("Query params: [ %v ]\n", c.Request.URL.Query())

	store[path] = string(val)
	c.JSON(200, gin.H{"success": true})
}

func DelHandler(c *gin.Context) {
	path := extractKey(c)

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
}

func DumpHandler(c *gin.Context) {
	c.JSON(200, store)
}

func ResetHandler(c *gin.Context) {
	ResetStore()
	c.JSON(200, gin.H{"result": "ok"})
}

func ResetStore() {
	store = make(map[string]string)
}
