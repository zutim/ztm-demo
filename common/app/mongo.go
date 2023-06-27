package app

import (
	"github.com/spf13/viper"
	"github.com/zutim/mongo"
)

func newMongo() *mongo.MongoClient {
	conf := newMongoConf()
	return mongo.NewMongo(conf.Dsn, conf.Pool)
}

type mongoConf struct {
	Dsn  string
	Pool int
}

//func Mongo() *mongo.MongoClient {
//	return app.MongoDb
//}

func newMongoConf() *mongoConf {
	return &mongoConf{
		Dsn:  viper.GetString("mongo.dsn"),
		Pool: viper.GetInt("mongo.pool"),
	}
}
