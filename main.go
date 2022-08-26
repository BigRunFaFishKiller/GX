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

	e.Run(":8001")
}
