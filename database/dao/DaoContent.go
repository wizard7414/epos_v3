package dao

import "epos_v3/domain/data"

func (dao *EposDbDao) CreateContent(content data.DbContent) error {
	return dao.DB.Save(&content).Error
}

func (dao *EposDbDao) DeleteContent(id string) error {
	result := dao.DB.Delete(&data.DbContent{}, id)
	return result.Error
}

func (dao *EposDbDao) GetContentById(id string) (data.DbContent, error) {
	content := data.DbContent{}
	result := dao.DB.Where("id = ?", id).First(&content)
	return content, result.Error
}

func (dao *EposDbDao) GetContentByEntryId(id string) ([]data.DbContent, error) {
	var content []data.DbContent
	result := dao.DB.Where("entry = ?", id).Find(&content)
	return content, result.Error
}

func (dao *EposDbDao) GetContentCount() (int, error) {
	var content []data.DbContent
	result := dao.DB.Find(&content)
	return len(content), result.Error
}

//----------------------------------------------------------------------------------
