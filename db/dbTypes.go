package db

func (Tb_co_log) TableName() string {
	return "TB_CO_LOG_HIST"
}

type Tb_co_log struct {
	BasDt string `json:"basDt"`
	Id    int    `json:"id"`
	Msg   string `json:"msg"`
}
