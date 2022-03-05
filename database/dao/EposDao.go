package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type EposDao struct {
	DB EposDbDao
}

func (dao *EposDao) InitDb(path string) error {
	dataBase, openErr := gorm.Open("sqlite3", path)
	if openErr == nil {
		dao.DB = EposDbDao{
			DB: dataBase,
		}
		return nil
	} else {
		return openErr
	}
}

func (dao *EposDao) CloseDb() error {
	return dao.DB.DB.Close()
}
