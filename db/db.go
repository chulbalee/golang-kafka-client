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

	dsn := fmt.Sprint(config.Database.User, ":", config.Database.Password, "@tcp(", config.Database.Host, ":", config.Database.Port, ")/", config.Database.DB, "?charset=utf8")
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
