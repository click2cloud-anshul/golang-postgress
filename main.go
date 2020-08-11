package main

import (
	"fmt"
	router "go-postgres/router"
	"log"
	"net/http"
	"os"
)

func main() {
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)

	//fmt.Println("connecting to the port 8080...")

	// connection to server
	//log.Fatal(http.ListenAndServe(":8080", r))

	//}
	var host string
	host = os.Getenv("DB_User")
	fmt.Printf("%s", host)
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
