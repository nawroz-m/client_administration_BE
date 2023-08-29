package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"

	"client_administration/router"
)



func main() {
	//Load Env File
	envErr := godotenv.Load(".env")
	if(envErr != nil){
		fmt.Print("Env file is not")
	}
	PORT := os.Getenv("PORT")
	secretKey := os.Getenv("KEYSECRET")
	fmt.Print(secretKey)
    app := fiber.New()
	router.SetupRoute(app)

    app.Listen(PORT)
}
