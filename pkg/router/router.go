package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/natholdallas/templates/fibergorm/pkg/conf"
	"github.com/natholdallas/templates/fibergorm/pkg/handler"
	"github.com/natholdallas/templates/fibergorm/pkg/mid"
)

func Setup(app *fiber.App) {
	app.Static("/", conf.Ctx.WebPath)
	app.Static("/media", conf.Ctx.MediaPath)

	api := app.Group("/api/v1")
	api.Post("/media", mid.IsLogin, handler.UploadImage)

	account := api.Group("/account")
	account.Post("/login", handler.Login)
	account.Post("/login/google", handler.LoginWithGoogle)
	account.Post("/register", handler.Register)
	account.Get("", mid.IsLogin, handler.GetMe)
	account.Put("", mid.IsLogin, mid.NoGoogle, handler.UpdateMe)
	account.Patch("/pwd", mid.IsLogin, mid.NoGoogle, handler.ResetMyPassword)

	user := api.Group("/user")
	user.Get("", mid.IsLogin, mid.IsAdmin, handler.ListUser)
	user.Get("/:id", mid.IsLogin, mid.IsAdmin, handler.FindUser)
	user.Post("", mid.IsLogin, mid.IsAdmin, handler.CreateUser)
	user.Put("/:id", mid.IsLogin, mid.IsAdmin, handler.UpdateUser)
	user.Patch("/pwd", mid.IsLogin, mid.IsAdmin, handler.ResetUserPassword)
	user.Delete("/:id", mid.IsLogin, mid.IsAdmin, handler.RemoveUser)

	test := api.Group("/test")
	test.Get("", handler.TestHandler)
}
