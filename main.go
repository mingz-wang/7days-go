package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"7days-web/middlewares"
	"gee"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gee.New()
	r.Use(middlewares.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "Gee", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})
	r.GET("/date", func(c *gee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title": "gee",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	v1 := r.Group("/v1")

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
