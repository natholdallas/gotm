package handler

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/natholdallas/templates/fibergorm/pkg/conf"
	"github.com/natholdallas/templates/fibergorm/pkg/db"
	"github.com/natholdallas/templates/fibergorm/pkg/enum"
	"github.com/natholdallas/templates/fibergorm/pkg/fibers"
	t "github.com/natholdallas/templates/fibergorm/pkg/tools"
)

type MediaOut struct {
	Model
	Value     string `json:"value"`
	Preview   string `json:"preview"`
	Activated bool   `json:"activated"`
}

func UploadImage(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	suffix := strings.ToLower(string(file.Filename[strings.LastIndex(file.Filename, "."):]))
	filename := uuid.New().String() + suffix
	savefile := conf.Ctx.SaveMediaPath(filename)
	media := db.Media{Value: filename, Preview: "/media/" + filename}

	if suffix != ".jpg" && suffix != ".png" && suffix != ".jpeg" {
		return fibers.Err(enum.AvailableImageSuffix + suffix)
	}
	err = c.SaveFile(file, savefile)
	if err != nil {
		return fibers.Err(err, fiber.StatusInternalServerError)
	}

	return c.JSON(t.Copy(&MediaOut{}, media))
}
