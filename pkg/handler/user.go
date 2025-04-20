package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/natholdallas/templates/fibergorm/pkg/db"
	"github.com/natholdallas/templates/fibergorm/pkg/enum"
	"github.com/natholdallas/templates/fibergorm/pkg/fibers"
	t "github.com/natholdallas/templates/fibergorm/pkg/tools"
)

type UserIn struct {
	Name     string `json:"name" validate:"max=20"`
	Username string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4,max=20"`
	Avatar   string `json:"avatar"`
	IsAdmin  bool   `json:"isAdmin"`
}

type User struct {
	SoftModel
	Name         string `json:"name"`
	Username     string `json:"username"`
	Avatar       string `json:"avatar"`
	IsAdmin      bool   `json:"isAdmin"`
	IsGoogleUser bool   `json:"isGoogleUser"`
}

type UserUpdateIn struct {
	Name     string `json:"name" validate:"max=20"`
	Username string `json:"username" validate:"required,email"`
	Avatar   string `json:"avatar"`
	IsAdmin  bool   `json:"isAdmin"`
}

type UserPasswordIn struct {
	Password string `json:"password" validate:"required,min=4,max=20"`
}

func ListUser(c *fiber.Ctx) error {
	content := []User{}
	page, size, count := fibers.GetPagination(c)

	result := db.Ctx.
		Model(&db.User{}).
		Where("is_admin = ?", c.Query("type") == "admin").
		Count(&count).
		Scopes(db.PaginateScope(page, size)).
		Find(&content)
	if result.Error != nil {
		return fibers.Err(enum.DataNotFound)
	}

	return c.JSON(Page{Total: count, Page: t.CeilDivide64(count, int64(size)), Content: content})
}

func FindUser(c *fiber.Ctx) error {
	user := User{}
	result := db.Ctx.Model(&db.User{}).First(&user, c.Params("id"))
	if result.Error != nil {
		return fibers.Err(enum.DataNotFound)
	}
	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	d, err := fibers.FormData(c, &UserIn{})
	if err != nil {
		return err
	}
	result := db.Ctx.Model(&db.User{}).Create(d)
	if result.Error != nil {
		return fibers.Err(enum.CreateFailed)
	}
	return nil
}

func UpdateUser(c *fiber.Ctx) error {
	data, err := fibers.FormData(c, &UserUpdateIn{})
	if err != nil {
		return err
	}
	user := db.User{}
	db.Ctx.Where("id = ?", c.Params("id")).First(&user)
	if user.IsGoogleUser {
		return fibers.Err(enum.CantEditGoogleUser)
	}
	t.Copy(&user, data)
	result := db.Ctx.Save(&user)
	if result.Error != nil {
		return fibers.Err(enum.UpdateFailed)
	}
	return nil
}

func ResetUserPassword(c *fiber.Ctx) error {
	data, err := fibers.FormData(c, &UserPasswordIn{})
	if err != nil {
		return err
	}
	user := db.User{}
	db.Ctx.Where("id = ?", c.Params("id")).First(&user)
	if user.IsGoogleUser {
		return fibers.Err(enum.CantEditGoogleUser)
	}
	t.Copy(&user, data)
	result := db.Ctx.Save(&user)
	if result.Error != nil {
		return fibers.Err(enum.ResetPwdFailed)
	}
	return nil
}

func RemoveUser(c *fiber.Ctx) error {
	result := db.Ctx.Delete(&db.User{}, c.Params("id"))
	if result.Error != nil {
		return fibers.Err(enum.RemoveFailed)
	}
	return nil
}
