package db

import "gorm.io/gorm"

type Tb_co_log struct {
	gorm.Model
	Tx  string `json:"tx"`
	Id  string `json:"id"`
	Msg string `json:"msg"`
}
