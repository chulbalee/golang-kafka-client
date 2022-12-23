package db

import "fmt"

type TbCoLogDao struct {
	table Tb_co_log
	conn  *DB
}

func (t TbCoLogDao) Create(data Tb_co_log) {

	result := t.conn.db.Create(&data)

	if result.Error != nil {
		fmt.Println("TbCoLogDao did not insert")
	}
	fmt.Println("TbCoLogDao data inserted : ", result.RowsAffected)

}

func (t TbCoLogDao) Select() {

}

func (t TbCoLogDao) Update() {

}
