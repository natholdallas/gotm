package conf

import (
	"os"

	"github.com/gofiber/fiber/v2/log"
	t "github.com/natholdallas/templates/fibergorm/pkg/tools"
)

type Conf struct {
	AppName   string
	AppPort   string
	DevMode   bool
	SecretKey string
	WebPath   string
	MediaPath string
	MediaCron string
	CacheCron string

	DBName   string
	DBQuery  string
	DBPort   string
	Hostname string
	Username string
	Password string
}

func (cfg *Conf) RemoveMedia(filename string) error {
	return os.Remove(cfg.MediaPath + "/" + filename)
}

func (cfg *Conf) SaveMediaPath(filename string) string {
	return cfg.MediaPath + "/" + filename
}

func (cfg *Conf) DevFunc(block func()) {
	if Ctx.DevMode {
		block()
	}
}

func (cfg *Conf) DevPrint(v ...any) {
	if Ctx.DevMode {
		for i := range v {
			log.Info(t.JSONStringify(v[i]))
		}
	}
}

type FlagConf struct {
	Sync bool
}
