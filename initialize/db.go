package initialize

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	db, _ := gorm.Open(mysql.Open("root@(127.0.0.1:3306)/user_go?charset=utf8mb4&parseTime=True&loc=Local"))
	mysqlDB, _ := db.DB()
	mysqlDB.SetMaxIdleConns(1)
	mysqlDB.SetMaxOpenConns(10)
	return db
}

func Cache() *redis.Client {
	cache := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return cache
}
