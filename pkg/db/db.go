package db

import (
	"database/sql"
	"fmt"

	r "github.com/Pallinder/go-randomdata"
	"github.com/gofiber/fiber/v2/log"
	"github.com/natholdallas/templates/fibergorm/pkg/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Ctx gorm.DB

// Auto Create Database
func init() {
	db, err := sql.Open("mysql", dsn(false))
	if err != nil {
		log.Panic("Failed to connect database, so we gonna be create it.")
	} else {
		conf.Ctx.DevFunc(func() {
			log.Info("DevMode is enabled, so we reset the database to default.")
			db.Exec("DROP DATABASE IF EXISTS " + conf.Ctx.DBName)
		})
	}
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + conf.Ctx.DBName)
	if err != nil {
		log.Panic("Failed to create database")
	}
	db.Close()
}

// DB Ctx Init
func init() {
	db, err := gorm.Open(mysql.Open(dsn(true)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
	Ctx = *db

	// Migrations
	db.AutoMigrate(&User{}, &Media{})
}

// Init Data
func init() {
	if Emptied(&User{}) {
		Ctx.Create(&User{Name: "NatholDallas", Avatar: randomImg(), Username: "505050728@qq.com", Password: "ar1234", IsAdmin: true})
		conf.Ctx.DevFunc(func() {
			for range 49 {
				Ctx.Create(&User{Name: r.SillyName(), Avatar: randomImg(), Username: r.Email(), Password: "ar1234", IsAdmin: true})
			}
			for range 50 {
				Ctx.Create(&User{Name: r.SillyName(), Avatar: randomImg(), Username: r.Email(), Password: "ar1234"})
			}
		})
	}
}

func dsn(usedb bool) string {
	c := conf.Ctx
	if usedb {
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", c.Username, c.Password, c.Hostname, c.DBPort, c.DBName, c.DBQuery)
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", c.Username, c.Password, c.Hostname, c.DBPort)
}

func randomImg() string {
	return "https://avatar.iran.liara.run/public/girl?username=Scott" + r.SillyName()
}
