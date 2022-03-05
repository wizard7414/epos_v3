package dao

import "epos_v3/domain/data"

func (dao *EposDbDao) CreateSourceType(sourceType data.SourceType) error {
	return dao.DB.Save(&sourceType).Error
}

func (dao *EposDbDao) DeleteSourceType(code string) error {
	result := dao.DB.Delete(&data.SourceType{}, code)
	return result.Error
}

func (dao *EposDbDao) GetSourceTypeByCode(code string) (data.SourceType, error) {
	sourceType := data.SourceType{}
	result := dao.DB.First(&sourceType, code)
	return sourceType, result.Error
}

func (dao *EposDbDao) GetSourceTypeByTitle(title string) (data.SourceType, error) {
	sourceType := data.SourceType{}
	result := dao.DB.Where("title = ?", title).First(&sourceType)
	return sourceType, result.Error
}

//----------------------------------------------------------------------------------

func (dao *EposDbDao) CreateRtSourceParam(rtSourceParam data.DbRtSourceParam) error {
	return dao.DB.Save(&rtSourceParam).Error
}

func (dao *EposDbDao) DeleteRtSourceParam(sourceId string, paramId string) error {
	sourceParam := data.DbRtSourceParam{}
	result := dao.DB.Where("source = ? AND param = ?", sourceId, paramId).Delete(&sourceParam)
	return result.Error
}

func (dao *EposDbDao) GetRtSourceParamById(sourceId string, paramId string) (data.DbRtSourceParam, error) {
	sourceParam := data.DbRtSourceParam{}
	result := dao.DB.Where("source = ? AND param = ?", sourceId, paramId).First(&sourceParam)
	return sourceParam, result.Error
}

func (dao *EposDbDao) GetRtSourceParamByEntry(sourceId string) ([]data.DbRtSourceParam, error) {
	var sourceParams []data.DbRtSourceParam
	result := dao.DB.Where("source = ?", sourceId).Find(&sourceParams)
	return sourceParams, result.Error
}

func (dao *EposDbDao) GetRtSourceParamByParam(paramId string) ([]data.DbRtSourceParam, error) {
	var sourceParams []data.DbRtSourceParam
	result := dao.DB.Where("param = ?", paramId).First(&sourceParams)
	return sourceParams, result.Error
}

//----------------------------------------------------------------------------------
