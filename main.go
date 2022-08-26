package main

import (
	"GX"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	e := GX.Default()

	e.GET("/ping/:id", func(c *GX.Context) {
		fmt.Fprintf(c.Writer, "pone"+c.Param("id"))
	})

	adminGroup := e.Group("/admin")
	adminGroup.Use(Log)
	adminGroup.GET("/info/:user", func(c *GX.Context) {
		user := c.Param("user")
		c.String(200, "hello admin %s", user)
	})

	e.GET("/panic", func(c *GX.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	e.Run(":8001")
}
func Log(c *GX.Context) {
	log.Print(time.Now())
}
