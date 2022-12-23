package db

import "gorm.io/gorm"

type Tb_co_log struct {
	gorm.Model
	Tx  string
	Id  string
	Msg string
}
