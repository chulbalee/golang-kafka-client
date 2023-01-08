package db

import (
	"fmt"
	"golang-kafka-client/conf"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var once sync.Once
var database *gorm.DB

func Init(config conf.Config) {

	once.Do(func() {
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

		database = db
	})
}

func GetInstance() *gorm.DB {
	return database
}

func Create(data interface{}) {

	switch data.(type) {
	case Tb_co_log:
		convData := data.(Tb_co_log)
		result := database.Select("BasDt", "Id", "Msg").Create(&convData)
		if result.Error != nil {
			fmt.Println("did not insert")
		}
		fmt.Println("data inserted : ", result.RowsAffected)
	default:
		fmt.Println("DB Casting Failed")
	}

}

func RawQuery(q string) {
	database.Exec(q)
}

func Select() {

}

func Update() {

}
