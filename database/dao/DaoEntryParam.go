package dao

import "epos_v3/domain/data"

func (dao *EposDbDao) CreateEntryParam(entryParam data.DbEntryParam) error {
	return dao.DB.Save(&entryParam).Error
}

func (dao *EposDbDao) DeleteEntryParam(id string) error {
	result := dao.DB.Delete(&data.DbEntryParam{}, id)
	return result.Error
}

func (dao *EposDbDao) GetEntryParamById(id string) (data.DbEntryParam, error) {
	entryParam := data.DbEntryParam{}
	result := dao.DB.Where("id = ?", id).First(&entryParam)
	return entryParam, result.Error
}

func (dao *EposDbDao) GetEntryParamByValue(value string) (data.DbEntryParam, error) {
	entryParam := data.DbEntryParam{}
	result := dao.DB.Where("value = ?", value).First(&entryParam)
	return entryParam, result.Error
}

func (dao *EposDbDao) GetEntryParamCount() (int, error) {
	var entryParam []data.DbEntryParam
	result := dao.DB.Find(&entryParam)
	return len(entryParam), result.Error
}

//----------------------------------------------------------------------------------

func (dao *EposDbDao) CreateRtEntryParam(rtEntryParam data.DbRtEntryParam) error {
	return dao.DB.Save(&rtEntryParam).Error
}

func (dao *EposDbDao) DeleteRtEntryParam(entryId string, paramId string) error {
	entryParam := data.DbRtEntryParam{}
	result := dao.DB.Where("entry = ? AND param = ?", entryId, paramId).Delete(&entryParam)
	return result.Error
}

func (dao *EposDbDao) GetRtEntryParamById(entryId string, paramId string) (data.DbRtEntryParam, error) {
	entryParam := data.DbRtEntryParam{}
	result := dao.DB.Where("entry = ? AND param = ?", entryId, paramId).First(&entryParam)
	return entryParam, result.Error
}

func (dao *EposDbDao) GetRtEntryParamByEntry(entryId string) ([]data.DbRtEntryParam, error) {
	var entryParam []data.DbRtEntryParam
	result := dao.DB.Where("entry = ?", entryId).Find(&entryParam)
	return entryParam, result.Error
}

func (dao *EposDbDao) GetRtEntryParamByParam(paramId string) ([]data.DbRtEntryParam, error) {
	var entryParam []data.DbRtEntryParam
	result := dao.DB.Where("param = ?", paramId).First(&entryParam)
	return entryParam, result.Error
}

//----------------------------------------------------------------------------------
