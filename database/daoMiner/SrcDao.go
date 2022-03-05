package dbMiner

import (
	"epos_v3/domain/miner"
	"github.com/jinzhu/gorm"
)

type SrcDao struct {
	DB *gorm.DB
}

func (dao *SrcDao) Create(source miner.SourceS) error {
	return dao.DB.Save(&source).Error
}

func (dao *SrcDao) Delete(sourceId int64) error {
	result := dao.DB.Delete(&miner.SourceS{}, sourceId)
	return result.Error
}

func (dao *SrcDao) GetById(sourceId int64) (miner.SourceS, error) {
	source := miner.SourceS{}
	result := dao.DB.First(&source, sourceId)
	return source, result.Error
}

func (dao *SrcDao) GetByTitle(sourceTitle string) ([]miner.SourceS, error) {
	var sources []miner.SourceS
	result := dao.DB.Where("title = ?", sourceTitle).Find(&sources)
	return sources, result.Error
}

func (dao *SrcDao) GetSources() ([]miner.SourceS, error) {
	var sources []miner.SourceS
	result := dao.DB.Find(&sources)
	return sources, result.Error
}

func (dao *SrcDao) UpdateDateById(sourceId int64, newDate int64) error {
	result := dao.DB.Model(miner.SourceS{}).Where("id = ?", sourceId).Update("change_date", newDate)
	return result.Error
}

func (dao *SrcDao) UpdateDateByTitle(sourceTitle string, newDate int64) error {
	result := dao.DB.Model(miner.SourceS{}).Where("title = ?", sourceTitle).Update("change_date", newDate)
	return result.Error
}
