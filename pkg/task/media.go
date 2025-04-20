package task

import (
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/natholdallas/templates/fibergorm/pkg/db"
)

// Dynamic media cleaner.

func Media() {
	log.Info("Starting clean unactivated & expired media resources.")

	uploads := []db.Media{}
	db.Ctx.Model(&db.Media{}).Where("activated = ?", false).Find(&uploads)
	for _, u := range uploads {
		if time.Now().Compare(u.CreatedAt.Add(time.Hour*5)) >= 0 {
			log.Info("Clean Media File" + u.Value)
			db.Ctx.Delete(&u)
		}
	}
}
