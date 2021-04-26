package main

import (
	"go-chat/modules/authentication"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// Init Router
	router := mux.NewRouter()

	//Route Handler /api/login Endpoint

	router.HandleFunc("/api/login", authentication.Login).Methods("POST")
	// router.HandleFunc("/api/users").Methods("POST")

	var httpPort = os.Getenv("HTTP_PORT")
	if len(httpPort) == 0 {
		httpPort = "8000"
	}
	log.Fatal(http.ListenAndServe(":"+httpPort, router))
}
