package app

import (
	"github.com/zutim/mongo"
	"github.com/zutim/pool"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type apps struct {
	Log     map[string]*zap.SugaredLogger
	Pool    *pool.ScalableGroutinePool
	Db      *gorm.DB
	MongoDb *mongo.MongoClient
}

var (
	app *apps
)

func newApps() *apps {
	app = &apps{
		Log: map[string]*zap.SugaredLogger{
			"log":   newDefaultLogger(),
			"order": newOrderLogger(),
			"user":  newUserLogger(),
		},
		Pool:    pool.NewGoroutinePool(),
		Db:      newDb(),
		MongoDb: newMongo(),
	}
	return app
}
