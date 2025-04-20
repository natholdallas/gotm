package db

import (
	"github.com/natholdallas/templates/fibergorm/pkg/conf"
	"gorm.io/gorm"
)

// NOTE: User

func (v *User) BeforeUpdate(tx *gorm.DB) error {
	conf.Ctx.DevPrint(v)
	tx.Set("old_avatar", v.Avatar)
	return nil
}

func (v *User) AfterSave(tx *gorm.DB) error {
	if v.Avatar == "" {
		return nil
	}
	if avatar, ok := tx.Get("old_avatar"); ok {
		if avatar != v.Avatar {
			InactivateMedia(avatar.(string))
		}
	}
	return ActivateMedia(v.Avatar)
}

// NOTE: Media

func (v *Media) AfterDelete(tx *gorm.DB) error {
	return conf.Ctx.RemoveMedia(v.Value)
}
