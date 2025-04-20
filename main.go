package main

import (
	"os"

	"github.com/natholdallas/templates/fibergorm/pkg/conf"
	_ "github.com/natholdallas/templates/fibergorm/pkg/db"
	"github.com/natholdallas/templates/fibergorm/pkg/handler"
	"github.com/natholdallas/templates/fibergorm/pkg/router"
	"github.com/natholdallas/templates/fibergorm/pkg/task"

	"github.com/gofiber/fiber/v2"
)

func init() {
	os.MkdirAll(conf.Ctx.MediaPath, 0o777)
}

func main() {
	if conf.Flag.Sync {
		task.Sync()
	}

	app := fiber.New(fiber.Config{
		AppName:      conf.Ctx.AppName,
		ErrorHandler: handler.ErrorHandler,
	})
	router.Setup(app)
	app.Listen(conf.Ctx.AppPort)
}
