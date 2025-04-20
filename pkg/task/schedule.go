package task

import (
	"github.com/natholdallas/templates/fibergorm/pkg/conf"
	"github.com/robfig/cron/v3"
)

func init() {
	schedule := cron.New(cron.WithSeconds())
	schedule.AddFunc(conf.Ctx.MediaCron, Media)
	schedule.AddFunc(conf.Ctx.CacheCron, Cache)
	schedule.Start()
}

func Sync() {
	Cache()
	Media()
}
