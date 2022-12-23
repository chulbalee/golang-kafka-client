package db

import (
	"fmt"
	"golang-kafka-client/conf"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	Db *gorm.DB
}

func (d *DB) Init(config conf.Config) {
	fmt.Println(":::db Init")
	fmt.Println(":::db type ", config.Database.Type)
	fmt.Println(":::db host ", config.Database.Host)

	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprint("%s:%s@tcp(%s:%s)/$s?charset=utf8", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	d.Db = db
}
