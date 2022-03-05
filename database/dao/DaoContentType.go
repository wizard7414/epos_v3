package dao

import "epos_v3/domain/data"

func (dao *EposDbDao) CreateContentType(contentType data.ContentType) error {
	return dao.DB.Save(&contentType).Error
}

func (dao *EposDbDao) DeleteContentType(code string) error {
	result := dao.DB.Delete(&data.ContentType{}, code)
	return result.Error
}

func (dao *EposDbDao) GetContentTypeByCode(code string) (data.ContentType, error) {
	contentType := data.ContentType{}
	result := dao.DB.Where("code = ?", code).First(&contentType)
	return contentType, result.Error
}

func (dao *EposDbDao) GetContentTypeByTitle(title string) (data.ContentType, error) {
	contentType := data.ContentType{}
	result := dao.DB.Where("title = ?", title).First(&contentType)
	return contentType, result.Error
}

//----------------------------------------------------------------------------------
