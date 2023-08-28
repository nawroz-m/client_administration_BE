package router

import (
	"client_administration/controllers"

	"github.com/gofiber/fiber"
)

func SetupRoute(app *fiber.App){

	/*
	* @api {POST} /api/insertdata
	* @access Public
	* @dec Insert user information to DB
	*/
	app.Post("/api/insertdata", controllers.InserUserInfo)

}
