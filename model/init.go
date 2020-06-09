package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Database struct {
	mysql *gorm.DB
}

var DB *Database

func connectDB(username, password, addr, dbname string) *gorm.DB {
	config := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		dbname,
		true,
		"Local",
	)
	db, err := gorm.Open("mysql", config)
	if err != nil {
		zap.S().Errorf("database connection failed.database name: %s", dbname)
	}

	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	db.DB().SetMaxOpenConns(100) // 设置最大连接数
	db.DB().SetMaxIdleConns(50)  // 设置最大空闲连接数
}

func InitSelfDB() *gorm.DB {
	return connectDB(
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.dbname"),
	)
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func (db *Database) Init() {
	DB = &Database{GetSelfDB()}
}

func (db *Database) Close() {
	db.mysql.Close()
}
