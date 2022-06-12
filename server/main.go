package main

import (
	"fmt"
	"log"
	"net/http"
	"productManagement/database"

	"github.com/gorilla/handlers"

	// "productManagement/controller"
	"productManagement/router"
)

func main() {
	fmt.Println("Everything starts from here")
	database.Connect()
	r := router.Router()
	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8000"
	// }

	// log.Fatal(http.ListenAndServe(":"+port,handlers.CORS(origins,
	// headers := handlers.AllowedHeaders([]string{"content-Type": "application/json", "X-Requested-With"})
	// methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE", "OPTIONS"})
	// origins := handlers.AllowedOrigins([]string{"*"})
	fmt.Println("Sever is getting started...")
	// log.Fatal(http.ListenAndServe(":5000", r))
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(headers, methods, origins)(r)))
	fmt.Println("Server has started at 5000")

}
