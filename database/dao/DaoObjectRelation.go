package dao

import "epos_v3/domain/data"

func (dao *EposDbDao) CreateRtObjectRelation(rtObjectRelation data.DbRtObjectRelation) error {
	return dao.DB.Save(&rtObjectRelation).Error
}

func (dao *EposDbDao) DeleteRtObjectRelation(objAId string, relationId string, objABId string) error {
	objectRelation := data.DbRtObjectRelation{}
	result := dao.DB.Where("object_a = ? AND relation = ? AND object_b = ?", objAId, relationId, objABId).Delete(&objectRelation)
	return result.Error
}

func (dao *EposDbDao) GetRtObjectRelationById(objAId string, relationId string, objABId string) (data.DbRtObjectRelation, error) {
	objectRelation := data.DbRtObjectRelation{}
	result := dao.DB.Where("object_a = ? AND relation = ? AND object_b = ?", objAId, relationId, objABId).First(&objectRelation)
	return objectRelation, result.Error
}

func (dao *EposDbDao) GetRtObjectRelationByObject(objectId string) ([]data.DbRtObjectRelation, error) {
	var objectRelations []data.DbRtObjectRelation
	result := dao.DB.Where("object_a = ? OR object_b = ?", objectId, objectId).Find(&objectRelations)
	return objectRelations, result.Error
}

func (dao *EposDbDao) GetRtObjectRelationByRelation(relationId string) ([]data.DbRtObjectRelation, error) {
	var objectRelations []data.DbRtObjectRelation
	result := dao.DB.Where("relation = ?", relationId).First(&objectRelations)
	return objectRelations, result.Error
}

//----------------------------------------------------------------------------------
