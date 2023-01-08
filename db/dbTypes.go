package db

func (Tb_co_log) TableName() string {
	return "TB_CO_LOG_HIST"
}

type Tb_co_log struct {
	BasDt string `gorm:"primaryKey;column:basDt;not null"`
	Id    int    `gorm:"primaryKey;column:id;not null;default:18"`
	Msg   string `gorm:"column:msg"`
}
