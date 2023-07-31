package app

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	gorm2 "github.com/zutim/gorm"
	"github.com/zutim/log"
	"github.com/zutim/mongo"
	"github.com/zutim/pool"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type apps struct {
	log      *zap.SugaredLogger
	orderLog *zap.SugaredLogger
	userLog  *zap.SugaredLogger
	pool     *pool.ScalableGroutinePool
	db       *gorm.DB
	mongoDb  *mongo.MongoClient
	redisDb  *redis.Client
}

var (
	app *apps
)

func newApps() *apps {
	app = &apps{
		log:      log.NewLogMap().WithOptionPath(""),
		orderLog: log.NewLogMap().WithOptionPath("log/order.log"),
		userLog:  log.NewLogMap().WithOptionPath("log/user.log"),

		pool: pool.NewGoroutinePool(),
		db: gorm2.NewDb(&gorm2.Conf{
			Dsn:         viper.GetString("mysql.dsn"),
			MaxIdle:     viper.GetInt("mysql.maxIdle"),
			MaxOpen:     viper.GetInt("mysql.maxOpen"),
			MaxLifeTime: viper.GetInt("mysql.maxLife"),
		}),
		mongoDb: mongo.NewMongoClient(&mongo.Conf{
			Dsn:  viper.GetString("mongo.dsn"),
			Pool: viper.GetInt("mongo.pool"),
		}),
		redisDb: redis.NewClient(&redis.Options{
			Addr:         viper.GetString("redis"),
			DialTimeout:  10 * time.Second,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			PoolSize:     10,
			PoolTimeout:  30 * time.Second,
			DB:           1,
		}),
	}
	return app
}

func Log() *zap.SugaredLogger {
	return app.log
}

func OrderLog() *zap.SugaredLogger {
	return app.orderLog
}

func UserLog() *zap.SugaredLogger {
	return app.userLog
}

func Pool() *pool.ScalableGroutinePool {
	return app.pool
}

func DB() *gorm.DB {
	return app.db
}

func Mongo() *mongo.MongoClient {
	return app.mongoDb
}

func Redis() *redis.Client {
	return app.redisDb
}
