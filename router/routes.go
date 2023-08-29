package router

import (
	"client_administration/controllers"

	"github.com/gofiber/fiber"
)

func SetupRoute(app *fiber.App){

	/*
	* @api {POST} /api/register
	* @access Public
	* @dec Register a User
	*/
	app.Post("/api/register",  controllers.RegisterUser)

	/*
	* @api {POST} /api/login
	* @access Public
	* @dec Login a User
	*/
	app.Post("/api/login",  controllers.LoginUser)

	
}
