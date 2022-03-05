package dao

import "epos_v3/domain/data"

func (dao *EposDbDao) CreateSource(source data.DbSource) error {
	return dao.DB.Save(&source).Error
}

func (dao *EposDbDao) DeleteSource(code string) error {
	result := dao.DB.Delete(&data.DbSource{}, code)
	return result.Error
}

func (dao *EposDbDao) GetSourceByCode(code string) (data.DbSource, error) {
	source := data.DbSource{}
	result := dao.DB.First(&source, code)
	return source, result.Error
}

func (dao *EposDbDao) GetSourceByUrl(url string) (data.DbSource, error) {
	source := data.DbSource{}
	result := dao.DB.Where("url = ?", url).First(&source)
	return source, result.Error
}

func (dao *EposDbDao) GetSourceByTitle(title string) (data.DbSource, error) {
	source := data.DbSource{}
	result := dao.DB.Where("title = ?", title).First(&source)
	return source, result.Error
}

//----------------------------------------------------------------------------------
