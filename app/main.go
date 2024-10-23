package main

import (
	"github.com/parthvinchhi/url-shortner/pkg/routes"
)

func main() {
	r := routes.Routes()
	r.Run(":8080")
}
