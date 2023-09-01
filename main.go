package main

import (
	"client_administration/router"
	"fmt"
	"os"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"

	// gofiber/corss
	"github.com/joho/godotenv"
)




func main() {
	//Load Env File
	envErr := godotenv.Load(".env")
	if(envErr != nil){
		fmt.Print("Env file is not")
	}
	PORT := os.Getenv("PORT")
    app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET, PUT, POST, DELET"},
       
	}))
	router.SetupRoute(app)

    app.Listen(PORT)
}
