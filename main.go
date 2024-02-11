package main

import (
	"fmt"
	"net/http"

	"github.com/sandeepkumarnayak/routes"
)

func main() {
	fmt.Println("Hello World")
	r := routes.Router()
	fmt.Println("Server is running on port: 4000")
	http.ListenAndServe(":4000", r)
}
