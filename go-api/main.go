package main

import (
	"log"

	"API/database"
	_ "API/docs"
	"API/router"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/storage/redis/v3"

	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

//	@title			API
//	@version		1.0
//	@description	This is an API

//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization

//	@contact.name	g3yk

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath	/api
func main() {
	app := fiber.New(fiber.Config{
		// Prefork:       true,
		AppName:       "Perfumer",
		CaseSensitive: true,
		ServerHeader:  "Fiber",

		// WARNING: Enable following error handler function in production
		// ErrorHandler: func(c *fiber.Ctx, err error) error {
		// 	log.Println(err)
		// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		// 		"error": "Internal Server Error",
		// 	})
		// },
	})

	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(cache.New(cache.Config{
		Storage: redis.New(redis.Config{
			Host: "cache",
		}),
		CacheControl: true,
	}))

	app.Use(swagger.New(swagger.Config{
		BasePath: "/api/",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
	}))

	database.ConnectDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
