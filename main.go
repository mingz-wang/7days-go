package main

import (
	"log"
	"net/http"
	"time"

	"7days-web/middlewares"
	"gee"
)

func main() {
	r := gee.New()
	r.Use(middlewares.Logger())
	r.GET("/index", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := r.Group("/v1")
	v1.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})
	v1.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=aaa
		c.String(http.StatusOK, "hello %s, you are at %s", c.Query("name"), c.Path)
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	v2.GET("/hello/:name", func(c *gee.Context) {
		// expect /hello/aaa
		c.String(http.StatusOK, "hello %s, you are at %s", c.Param("name"), c.Path)
	})
	v2.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":8080")
}

func onlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
