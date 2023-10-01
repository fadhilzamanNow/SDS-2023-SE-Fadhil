package main

import "github.com/gofiber/fiber/v2"

type RequestData struct {
	JariJariLingkaran float64
	sisiPersegi       float64
	alasSegitiga      float64
	tinggiSegitiga    float64
}

type ResponseData struct {
	LuasLingkaran     float64
	luasPersegi       float64
	luasSegitiga      float64
	KelilingLingkaran float64
	kelilingPersegi   float64
	kelilingSegitiga  float64
}

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		request := new(RequestData)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		response := new(ResponseData)

		response.LuasLingkaran = 3.14 * float64(request.JariJariLingkaran*request.JariJariLingkaran)
		response.luasPersegi = float64(request.sisiPersegi * request.sisiPersegi)
		response.luasSegitiga = 0.5 * float64(request.alasSegitiga*request.tinggiSegitiga)
		response.KelilingLingkaran = 2 * float64(request.JariJariLingkaran) * 3.14
		response.kelilingPersegi = 4 * float64(request.sisiPersegi)
		response.kelilingSegitiga = 3 * float64(request.alasSegitiga)

		return c.JSON(response)
	})

	app.Listen(":3000")
}
