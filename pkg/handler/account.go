package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/natholdallas/templates/fibergorm/pkg/client"
	"github.com/natholdallas/templates/fibergorm/pkg/conf"
	"github.com/natholdallas/templates/fibergorm/pkg/db"
	"github.com/natholdallas/templates/fibergorm/pkg/enum"
	"github.com/natholdallas/templates/fibergorm/pkg/fibers"
	"github.com/natholdallas/templates/fibergorm/pkg/mid"
	t "github.com/natholdallas/templates/fibergorm/pkg/tools"
)

type LoginIn struct {
	Username string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4,max=20"`
}

type LoginGoogleIn struct {
	Token string `json:"token"`
}

type LoginOut struct {
	Token        string `json:"token"`
	ID           uint   `json:"id"`
	IsAdmin      bool   `json:"isAdmin"`
	IsGoogleUser bool   `json:"isGoogleUser"`
}

type RegisterIn struct {
	Username        string `json:"username" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=4,max=20"`
	ConfirmPassword string `json:"c" validate:"required,eqfield=Password" copier:"-"`
}

type AccountUpdateIn struct {
	Name     string `json:"name" validate:"max=20"`
	Username string `json:"username" validate:"required,email"`
	Avatar   string `json:"avatar"`
}

type AccountPasswordIn struct {
	OriginPassword  string `json:"o" validate:"required,min=4,max=20"`
	Password        string `json:"n" validate:"required,min=4,max=20"`
	ConfirmPassword string `json:"c" validate:"required,eqfield=Password"`
}

func Login(c *fiber.Ctx) error {
	d, err := fibers.FormData(c, &LoginIn{})
	if err != nil {
		return err
	}

	user := db.User{}
	db.Ctx.
		Where("BINARY username = ? AND BINARY password = ? And is_google_user = ?", d.Username, d.Password, false).
		First(&user)
	if user.ID == 0 {
		return fibers.Err(enum.UserNotFound)
	}
	token, err := mid.GenerateJwt(user, conf.Ctx.SecretKey)
	if err != nil {
		return fibers.Err(err, fiber.StatusInternalServerError)
	}

	return c.JSON(LoginOut{
		Token:        token,
		ID:           user.ID,
		IsAdmin:      user.IsAdmin,
		IsGoogleUser: user.IsGoogleUser,
	})
}

func LoginWithGoogle(c *fiber.Ctx) error {
	d, err := fibers.FormData(c, &LoginGoogleIn{})
	if err != nil {
		return err
	}
	g, err := client.GetGoogleUserInfo(d.Token)
	if err != nil {
		return fibers.Err(err, fiber.StatusInternalServerError)
	}
	user := db.User{}
	db.Ctx.Where("BINARY username = ?", g.Email).First(&user)
	t.Copy(&user, g)
	db.Ctx.Save(&user)

	token, err := mid.GenerateJwt(user, conf.Ctx.SecretKey)
	if err != nil {
		return fibers.Err(err, fiber.StatusInternalServerError)
	}

	return c.JSON(LoginOut{
		Token:        token,
		ID:           user.ID,
		IsAdmin:      user.IsAdmin,
		IsGoogleUser: user.IsGoogleUser,
	})
}

func Register(c *fiber.Ctx) error {
	d, err := fibers.FormData(c, &RegisterIn{})
	if err != nil {
		return err
	}
	user := db.User{}
	db.Ctx.Where("BINARY username = ?", d.Username).First(&user)
	if user.ID != 0 {
		return fibers.Err(enum.RegisterFailedSuggest)
	}
	user = t.Copy(&db.User{}, d)
	result := db.Ctx.Create(&user)
	if result.Error != nil {
		return fibers.Err(enum.RegisterFailed)
	}
	return nil
}

func GetMe(c *fiber.Ctx) error {
	claims := mid.GetClaims(c)
	user := User{}
	result := db.Ctx.Model(&db.User{}).First(&user, claims.ID)
	if result.Error != nil {
		return fibers.Err(enum.DataNotFound)
	}
	return c.JSON(user)
}

func UpdateMe(c *fiber.Ctx) error {
	data, err := fibers.FormData(c, &AccountUpdateIn{})
	if err != nil {
		return err
	}
	claims := mid.GetClaims(c)
	result := db.UpdatesAllByID(claims.ID, &db.User{}, data)
	if result.Error != nil {
		log.Error(result.Error)
		return fibers.Err(enum.UpdateFailed)
	}
	return nil
}

func ResetMyPassword(c *fiber.Ctx) error {
	claims := mid.GetClaims(c)
	d, err := fibers.FormData(c, &AccountPasswordIn{})
	if err != nil {
		return err
	}
	user := db.User{}
	db.Ctx.Where("BINARY password = ? AND id = ?", d.OriginPassword, claims.ID).First(&user)
	if user.ID == 0 {
		return fibers.Err(enum.UncorrectPassword)
	}
	result := db.Ctx.Model(&user).Update("password", d.Password)
	if result.Error != nil {
		log.Error(result.Error)
		return fibers.Err(enum.ResetPwdFailed)
	}

	return nil
}
