package main

import (
	"GX"
	"fmt"
)

func main() {
	e := GX.New()

	e.GET("/ping/:id", func(c *GX.Context) {
		fmt.Fprintf(c.Writer, "pone"+c.Param("id"))
	})

	adminGroup := e.Group("/admin")
	adminGroup.GET("/info/:user", func(c *GX.Context) {
		user := c.Param("user")
		c.String(200, "hello admin %s", user)
	})

	e.Run(":8001")
}
