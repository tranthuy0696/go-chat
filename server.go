package main

import (
	"fmt"
	"go-chat/keycloak"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// res, err := keycloak.GetTokenFromAdminUser()
	// if err != nil {
	// 	fmt.Println("met ghe ", err)
	// } else {
	// 	fmt.Println("rs ", res.AccessToken)
	// }

	resp, err := keycloak.CreateUser()
	if err != nil {
		fmt.Println("Errorrr")
	}
	fmt.Println("response ", resp)

}
