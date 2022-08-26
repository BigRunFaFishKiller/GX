package main

import (
	"GX"
	"fmt"
)

func main() {
	e := GX.New()

	e.GET("/ping", func(c *GX.Context) {
		fmt.Fprintf(c.Writer, "pone")
	})

	e.Run(":8001")
}
