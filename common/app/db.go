package app

import (
	"github.com/spf13/viper"
	"github.com/zutim/gorm"
	"gorm.io/gorm"
	"time"
)

func newDb() *gorm.DB {

	conf := newMysqlConfig()

	instance := db.New()

	if err := instance.Connect(conf.Dsn, &gorm.Config{}); err != nil {
		panic(err)
	}

	if err := instance.EnableConnectionPool(conf.MaxIdle, conf.MaxOpen, time.Duration(conf.MaxLifeTime)); err != nil {
		panic(err)
	}

	return instance.DB
}

type dbConf struct {
	Dsn         string
	MaxIdle     int
	MaxOpen     int
	MaxLifeTime int
}

func DB() *gorm.DB {
	return app.Db
}
func newMysqlConfig() *dbConf {
	return &dbConf{
		MaxIdle:     viper.GetInt("mysql.maxIdle"),
		MaxOpen:     viper.GetInt("mysql.maxOpen"),
		MaxLifeTime: viper.GetInt("mysql.maxLife"),
		Dsn:         viper.GetString("mysql.dsn"),
	}
}
