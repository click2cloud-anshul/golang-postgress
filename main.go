package main

import (
	"fmt"
	"go-postgres/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("connecting to the port 8080...")

	// connection to server
	log.Fatal(http.ListenAndServe(":8080", r))

}
