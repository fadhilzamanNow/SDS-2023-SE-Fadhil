package main

import (

	//"uts/database"
	"uts/database"

	"uts/route"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("UTS SDS")
	})

	app.Listen(":3000")

	database.Connect()

	//Kumpulan Route Route

	app.Post("/insert", route.InsertData)
	app.Get("/getData", route.GetAllData)
	app.Get("/getDataUser/:id_user", route.GetUserByid)

	app.Get("/delete/:id_user", route.Delete)
	app.Put("/update/:id_user", route.Update)

}
