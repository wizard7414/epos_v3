package dbMiner

import (
	"epos_v3/domain/miner"
	"github.com/jinzhu/gorm"
)

type AttrCodeDao struct {
	DB *gorm.DB
}

func (dao *AttrCodeDao) Create(attributeCode miner.AttrCodeS) error {
	return dao.DB.Save(&attributeCode).Error
}

func (dao *AttrCodeDao) Delete(attributeCodeId int64) error {
	result := dao.DB.Delete(&miner.AttrCodeS{}, attributeCodeId)
	return result.Error
}

func (dao *AttrCodeDao) GetById(attributeCodeId int64) (miner.AttrCodeS, error) {
	attributeCode := miner.AttrCodeS{}
	result := dao.DB.First(&attributeCode, attributeCodeId)
	return attributeCode, result.Error
}

func (dao *AttrCodeDao) GetByTitle(attributeCodeTitle string) ([]miner.AttrCodeS, error) {
	var attributeCodes []miner.AttrCodeS
	result := dao.DB.Where("title = ?", attributeCodeTitle).Find(&attributeCodes)
	return attributeCodes, result.Error
}
