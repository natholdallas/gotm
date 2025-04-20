package db

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/natholdallas/templates/fibergorm/pkg/enum"
	"gorm.io/gorm"
)

// NOTE: Custom

func InactivateMedia(v string) error {
	result := Ctx.Where(&Media{Preview: v}).Update("activated", false)
	return result.Error
}

func ActivateMedia(v string) error {
	result := Ctx.Model(&Media{Preview: v}).Update("activated", true)
	return result.Error
}

// NOTE: Common

func Count(model any, count *int64) *int64 {
	Ctx.Model(model).Count(count)
	return count
}

func Emptied(model any) bool {
	var count int64
	Ctx.Model(model).Count(&count)
	return count <= 0
}

func Exists(model, where any, args ...any) bool {
	var count int64
	Ctx.Model(model).Where(where, args).Count(&count)
	return count > 0
}

func FindByID(id any, value ...any) *gorm.DB {
	len := len(value)
	if len < 1 {
		log.Error(enum.InvalidParams)
	} else if len == 1 {
		return Ctx.First(value, id)
	}
	return Ctx.Model(value[0]).First(value[1], id)
}

func UpdatesAllByID(id, model, value any) *gorm.DB {
	return Ctx.
		Model(model).
		Select("*").
		Where("id = ?", id).
		Updates(value)
}

func UpdatesByID(id, model, value any) *gorm.DB {
	return Ctx.
		Model(model).
		Where("id = ?", id).
		Updates(value)
}

// NOTE: Scopes

type GormFunc = func(*gorm.DB) *gorm.DB

func PaginateScope(page, size int) GormFunc {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		if size > 100 && size <= 0 {
			size = 20
		}
		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}
