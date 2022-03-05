package dao

import (
	"epos_v3/domain/data"
	"epos_v3/utils"
)

func (dao *EposDao) ObjectAdd(object data.Object) error {

	for p := range object.Params {
		param := object.Params[p]

		_, err := dao.DB.GetObjectParamById(param.Id)
		if err != nil {
			if err.Error() == "record not found" {
				paramErr := dao.DB.CreateObjectParam(data.DbObjectParam{
					Id:    param.Id,
					Code:  param.Code.Code,
					Type:  param.Type.Code,
					Value: param.Value,
					Extra: param.Extra,
				})
				if paramErr != nil {
					return paramErr
				}
			}
		}

		_, relErr := dao.DB.GetRtObjectParamById(object.Id, param.Id)
		if relErr != nil {
			if relErr.Error() == "record not found" {
				relErr = dao.DB.CreateRtObjectParam(data.DbRtObjectParam{
					Object: object.Id,
					Param:  param.Id,
				})
				if relErr != nil {
					return relErr
				}
			}
		}
	}

	for r := range object.Relations {
		relation := object.Relations[r]

		_, err := dao.DB.GetRelationByCode(relation.Code)
		if err != nil {
			if err.Error() == "record not found" {
				paramErr := dao.RelationAdd(relation)
				if paramErr != nil {
					return paramErr
				}
			}
		}

		_, relErr := dao.DB.GetRtObjectRelationById(object.Id, relation.Code, relation.Object)
		if relErr != nil {
			if relErr.Error() == "record not found" {
				relErr = dao.DB.CreateRtObjectRelation(data.DbRtObjectRelation{
					ObjectA:  object.Id,
					Relation: relation.Code,
					ObjectB:  relation.Object,
				})
				if relErr != nil {
					return relErr
				}
			}
		}
	}

	return dao.DB.CreateObject(data.DbObject{
		Id:        object.Id,
		Added:     object.Added.Unix(),
		Title:     object.Title,
		Completed: utils.BoolToInt(object.Completed),
		Ignored:   utils.BoolToInt(object.Ignored),
		Access:    object.Access,
	})
}

//----------------------------------------------------------------------------------

func (dao *EposDbDao) CreateObject(object data.DbObject) error {
	return dao.DB.Save(&object).Error
}

func (dao *EposDbDao) DeleteObject(id string) error {
	result := dao.DB.Delete(&data.DbObject{}, id)
	return result.Error
}

func (dao *EposDbDao) GetObjectById(id string) (data.DbObject, error) {
	object := data.DbObject{}
	result := dao.DB.Where("id = ?", id).First(&object)
	return object, result.Error
}

func (dao *EposDbDao) GetObjectsByTitle(title string) ([]data.DbObject, error) {
	var objects []data.DbObject
	result := dao.DB.Where("title = ?", title).Find(&objects)
	return objects, result.Error
}

func (dao *EposDbDao) GetAllObjects() ([]data.DbObject, error) {
	var objects []data.DbObject
	result := dao.DB.Find(&objects)
	return objects, result.Error
}

//----------------------------------------------------------------------------------
