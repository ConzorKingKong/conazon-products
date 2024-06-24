package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var PORT, PORTExists = "", false
var DatabaseURLEnv, DatabaseURLExists = "", false

func main() {

	godotenv.Load()

	PORT, PORTExists = os.LookupEnv("PORT")
	DatabaseURLEnv, DatabaseURLExists = os.LookupEnv("DATABASEURL")

	if !DatabaseURLExists {
		log.Fatal("Required environment variable not set")
	}

	if !PORTExists {
		PORT = "8081"
	}

	http.HandleFunc("/", Root)
	http.HandleFunc("/products/", Products)
	http.HandleFunc("/products/{id}", ProductId)

	fmt.Println("server starting on port", PORT)
	http.ListenAndServe(":"+PORT, nil)
}
