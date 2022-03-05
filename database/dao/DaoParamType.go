package dao

import "epos_v3/domain/data"

func (dao *EposDbDao) CreateParamType(paramCode data.ParamType) error {
	return dao.DB.Save(&paramCode).Error
}

func (dao *EposDbDao) DeleteParamType(code string) error {
	result := dao.DB.Delete(&data.ParamType{}, code)
	return result.Error
}

func (dao *EposDbDao) GetParamTypeByCode(code string) (data.ParamType, error) {
	paramType := data.ParamType{}
	result := dao.DB.Where("code = ?", code).First(&paramType)
	return paramType, result.Error
}

func (dao *EposDbDao) GetParamTypeByTitle(title string) (data.ParamType, error) {
	paramType := data.ParamType{}
	result := dao.DB.Where("title = ?", title).First(&paramType)
	return paramType, result.Error
}

//----------------------------------------------------------------------------------
