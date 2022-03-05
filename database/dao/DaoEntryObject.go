package dao

import "epos_v3/domain/data"

func (dao *EposDbDao) CreateRtEntryObject(entryObject data.DbRtEntryObject) error {
	return dao.DB.Save(&entryObject).Error
}

func (dao *EposDbDao) DeleteRtEntryObject(entryId string, objectId string) error {
	entryObject := data.DbRtEntryObject{}
	result := dao.DB.Where("entry = ? AND object = ?", entryId, objectId).Delete(&entryObject)
	return result.Error
}

func (dao *EposDbDao) GetRtEntryObjectById(entryId string, objectId string) (data.DbRtEntryObject, error) {
	entryObject := data.DbRtEntryObject{}
	result := dao.DB.Where("entry = ? AND object = ?", entryId, objectId).First(&entryObject)
	return entryObject, result.Error
}

func (dao *EposDbDao) GetRtEntryObjectByEntry(entryId string) ([]data.DbRtEntryObject, error) {
	var entryObjects []data.DbRtEntryObject
	result := dao.DB.Where("entry = ?", entryId).Find(&entryObjects)
	return entryObjects, result.Error
}

func (dao *EposDbDao) GetRtEntryObjectByObject(objectId string) ([]data.DbRtEntryObject, error) {
	var entryObjects []data.DbRtEntryObject
	result := dao.DB.Where("object = ?", objectId).First(&entryObjects)
	return entryObjects, result.Error
}

//----------------------------------------------------------------------------------
