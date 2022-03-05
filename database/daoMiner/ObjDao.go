package dbMiner

import (
	"epos_v3/domain/miner"
	"github.com/jinzhu/gorm"
)

type ObjDao struct {
	DB *gorm.DB
}

func (dao *ObjDao) Create(object miner.ObjectS) error {
	return dao.DB.Save(&object).Error
}

func (dao *ObjDao) Delete(objectId int64) error {
	result := dao.DB.Delete(&miner.ObjectS{}, objectId)
	return result.Error
}

func (dao *ObjDao) GetById(objectId int64) (miner.ObjectS, error) {
	object := miner.ObjectS{}
	result := dao.DB.First(&object, objectId)
	return object, result.Error
}

func (dao *ObjDao) GetByTitle(objectTitle string) ([]miner.ObjectS, error) {
	var objects []miner.ObjectS
	result := dao.DB.Where("title = ?", objectTitle).Find(&objects)
	return objects, result.Error
}

func (dao *ObjDao) GetByUrl(objectUrl string) ([]miner.ObjectS, error) {
	var objects []miner.ObjectS
	result := dao.DB.Where("url = ?", objectUrl).Find(&objects)
	return objects, result.Error
}

func (dao *ObjDao) GetAllObjects() ([]miner.ObjectS, error) {
	var objects []miner.ObjectS
	result := dao.DB.Find(&objects)
	return objects, result.Error
}
