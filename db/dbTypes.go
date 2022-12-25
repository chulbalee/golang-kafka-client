package db

func (Tb_co_log) TableName() string {
	return "tb_co_log"
}

type Tb_co_log struct {
	Tx  string `json:"tx"`
	Id  int    `json:"id"`
	Msg string `json:"msg"`
}
