package t

import "github.com/gofiber/fiber/v2/log"

func PrintJSON(v ...any) {
	for _, i := range v {
		d, _ := JSONStringify(i)
		log.Debug(d)
	}
}

func PrintStruct(v ...any) {
	for _, i := range v {
		log.Infof("%#v", i)
	}
}
