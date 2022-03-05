package dao

import "epos_v3/domain/data"

func (dao *EposDao) RelationAdd(relation data.Relation) error {
	storedRelation := data.DbRelation{
		Code:        relation.Code,
		Title:       relation.Title,
		Description: relation.Description,
		Type:        relation.Type.Code,
	}

	err := dao.DB.CreateRelationType(relation.Type)
	if err != nil {
		return err
	}
	return dao.DB.CreateRelation(storedRelation)
}

//----------------------------------------------------------------------------------

func (dao *EposDbDao) CreateRelation(relation data.DbRelation) error {
	return dao.DB.Save(&relation).Error
}

func (dao *EposDbDao) DeleteRelation(id string) error {
	result := dao.DB.Delete(&data.DbRelation{}, id)
	return result.Error
}

func (dao *EposDbDao) GetRelationByCode(code string) (data.DbRelation, error) {
	relation := data.DbRelation{}
	result := dao.DB.First(&relation, code)
	return relation, result.Error
}

func (dao *EposDbDao) GetRelationByTitle(title string) (data.DbRelation, error) {
	relation := data.DbRelation{}
	result := dao.DB.Where("title = ?", title).First(&relation)
	return relation, result.Error
}

//----------------------------------------------------------------------------------
