package main

import (
	"net/http"

	"gee"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello World!</h1>")
	})
	// TODO: 这里的路由，只是静态路由，不支持/hello/:name这样的动态路由，动态路由我们将在下一次实现。
	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=world
		c.String(http.StatusOK, "hello %s, you are at %s\n", c.Query("name"), c.Path)
	})
	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":8080")
}
