package main

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/gin-gonic/gin"
	"razorsh4rk.github.io/fakeconsul/kv"
)

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

	kv.ResetStore()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	setupKVRoutes(r)

	r.Run()
}

func setupKVRoutes(r *gin.Engine) {
	r.GET("/v1/kv/*path", kv.GetHandler)
	r.PUT("/v1/kv/*path", kv.PutHandler)
	r.DELETE("/v1/kv/*path", kv.DelHandler)
	r.GET("/debug/dump", kv.DumpHandler)
	r.GET("/debug/reset", kv.ResetHandler)
}
