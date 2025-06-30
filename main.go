package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/conzorkingkong/conazon-products/config"
	"github.com/conzorkingkong/conazon-products/controllers"
	authcontrollers "github.com/conzorkingkong/conazon-users-and-auth/controllers"
	"github.com/joho/godotenv"
)

var PORT, PORTExists = "", false
var DatabaseURLEnv, DatabaseURLExists = "", false

func main() {

	godotenv.Load()

	PORT, PORTExists = os.LookupEnv("PORT")
	DatabaseURLEnv, DatabaseURLExists = os.LookupEnv("DATABASEURL")

	config.DatabaseURLEnv = DatabaseURLEnv

	if !DatabaseURLExists {
		log.Fatal("Required environment variable not set")
	}

	if !PORTExists {
		PORT = "8081"
	}

	http.HandleFunc("/", authcontrollers.Root)
	http.HandleFunc("/products/", controllers.Products)
	http.HandleFunc("/products/{id}", controllers.ProductId)

	http.HandleFunc("/healthz", authcontrollers.Healthz)

	fmt.Println("server starting on port", PORT)
	http.ListenAndServe(":"+PORT, nil)
}
