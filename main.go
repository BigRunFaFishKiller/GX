package main

import (
	"GX"
	"fmt"
	"net/http"
)

func main() {
	e := GX.New()

	e.GET("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pone")
	})

	e.Run(":8001")
}
