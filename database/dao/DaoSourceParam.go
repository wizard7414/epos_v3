package dao

import "epos_v3/domain/data"

func (dao *EposDbDao) CreateSourceParam(sourceParam data.DbSourceParam) error {
	return dao.DB.Save(&sourceParam).Error
}

func (dao *EposDbDao) DeleteSourceParam(id string) error {
	result := dao.DB.Delete(&data.DbSourceParam{}, id)
	return result.Error
}

func (dao *EposDbDao) GetSourceParamById(id string) (data.DbSourceParam, error) {
	sourceParam := data.DbSourceParam{}
	result := dao.DB.First(&sourceParam, id)
	return sourceParam, result.Error
}

func (dao *EposDbDao) GetSourceParamByValue(value string) (data.DbSourceParam, error) {
	sourceParam := data.DbSourceParam{}
	result := dao.DB.Where("value = ?", value).First(&sourceParam)
	return sourceParam, result.Error
}

//----------------------------------------------------------------------------------
