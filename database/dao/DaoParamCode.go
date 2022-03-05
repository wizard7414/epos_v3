package dao

import "epos_v3/domain/data"

func (dao *EposDbDao) CreateParamCode(paramCode data.ParamCode) error {
	return dao.DB.Save(&paramCode).Error
}

func (dao *EposDbDao) DeleteParamCode(code string) error {
	result := dao.DB.Delete(&data.ParamCode{}, code)
	return result.Error
}

func (dao *EposDbDao) GetParamCodeByCode(code string) (data.ParamCode, error) {
	paramCode := data.ParamCode{}
	result := dao.DB.Where("code = ?", code).First(&paramCode)
	return paramCode, result.Error
}

func (dao *EposDbDao) GetParamCodeByTitle(title string) (data.ParamCode, error) {
	paramCode := data.ParamCode{}
	result := dao.DB.Where("title = ?", title).First(&paramCode)
	return paramCode, result.Error
}

//----------------------------------------------------------------------------------
