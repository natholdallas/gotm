package conf

import (
	"flag"

	"github.com/fsnotify/fsnotify"
	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

var (
	Ctx  *Conf
	Flag *FlagConf
)

func init() {
	Ctx = new(Conf)
	Flag = new(FlagConf)
	LoadCtx()
	LoadFlag()

	viper.WatchConfig()
	viper.OnConfigChange(ConfigChangeEvent)
}

func LoadCtx() {
	viper.SetConfigName("config")
	viper.AddConfigPath("conf")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal error config file: %s", err)
	}
	viper.Unmarshal(Ctx)
}

func LoadFlag() {
	flag.BoolVar(&Flag.Sync, "sync", false, "")
	flag.Parse()
}

func ConfigChangeEvent(e fsnotify.Event) {
	log.Info("Config Change: ", e.Name)
	LoadCtx()
}
