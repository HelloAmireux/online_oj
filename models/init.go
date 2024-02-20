package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

//import (
//	"github.com/go-redis/redis/v8"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"log"
//)
//
var DB = Init()


//var RDB = InitRedisDB()
//
func Init() *gorm.DB {
	//dsn := define.MysqlDNS + "/gin_gorm_oj?charset=utf8mb4&parseTime=True&loc=Local"
	dsn :=  "root:123456@tcp(127.0.0.1:3306)/oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("gorm Init Error : ", err)
	} else{
		log.Println("init is success")
	}
	return db
}
//
//func InitRedisDB() *redis.Client {
//	return redis.NewClient(&redis.Options{
//		Addr:     "localhost:6379",
//		Password: "", // no password set
//		DB:       0,  // use default DB
//	})
//}
