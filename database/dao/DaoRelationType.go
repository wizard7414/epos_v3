package dao

import "epos_v3/domain/data"

func (dao *EposDbDao) CreateRelationType(relationType data.RelationType) error {
	return dao.DB.Save(&relationType).Error
}

func (dao *EposDbDao) DeleteRelationType(code string) error {
	result := dao.DB.Delete(&data.RelationType{}, code)
	return result.Error
}

func (dao *EposDbDao) GetRelationTypeByCode(code string) (data.RelationType, error) {
	relationType := data.RelationType{}
	result := dao.DB.First(&relationType, code)
	return relationType, result.Error
}

func (dao *EposDbDao) GetRelationTypeByTitle(title string) (data.RelationType, error) {
	relationType := data.RelationType{}
	result := dao.DB.Where("title = ?", title).First(&relationType)
	return relationType, result.Error
}

//----------------------------------------------------------------------------------
