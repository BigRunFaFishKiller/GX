package main

import (
	"GX"
	"fmt"
	"log"
	"time"
)

func main() {
	e := GX.New()

	e.GET("/ping/:id", func(c *GX.Context) {
		fmt.Fprintf(c.Writer, "pone"+c.Param("id"))
	})

	adminGroup := e.Group("/admin")
	adminGroup.Use(Log)
	adminGroup.GET("/info/:user", func(c *GX.Context) {
		user := c.Param("user")
		c.String(200, "hello admin %s", user)
	})

	e.Run(":8001")
}
func Log(c *GX.Context) {
	log.Print(time.Now())
}
