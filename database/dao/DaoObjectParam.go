package dao

import "epos_v3/domain/data"

func (dao *EposDao) ObjectParamAdd(objectParam data.ObjectParam) error {
	storedParam := data.DbObjectParam{
		Id:    objectParam.Id,
		Code:  objectParam.Code.Code,
		Type:  objectParam.Type.Code,
		Value: objectParam.Value,
		Extra: objectParam.Extra,
	}

	err := dao.DB.CreateParamType(objectParam.Type)
	if err != nil {
		return err
	}
	err = dao.DB.CreateParamCode(objectParam.Code)
	if err != nil {
		return err
	}
	return dao.DB.CreateObjectParam(storedParam)
}

//----------------------------------------------------------------------------------

func (dao *EposDbDao) CreateObjectParam(objectParam data.DbObjectParam) error {
	return dao.DB.Save(&objectParam).Error
}

func (dao *EposDbDao) DeleteObjectParam(id string) error {
	result := dao.DB.Delete(&data.DbObjectParam{}, id)
	return result.Error
}

func (dao *EposDbDao) GetObjectParamById(id string) (data.DbObjectParam, error) {
	objectParam := data.DbObjectParam{}
	result := dao.DB.Where("id = ?", id).First(&objectParam, id)
	return objectParam, result.Error
}

func (dao *EposDbDao) GetObjectParamByValue(value string) (data.DbObjectParam, error) {
	objectParam := data.DbObjectParam{}
	result := dao.DB.Where("value = ?", value).First(&objectParam)
	return objectParam, result.Error
}

//----------------------------------------------------------------------------------

func (dao *EposDbDao) CreateRtObjectParam(rtObjectParam data.DbRtObjectParam) error {
	return dao.DB.Save(&rtObjectParam).Error
}

func (dao *EposDbDao) DeleteRtObjectParam(objectId string, paramId string) error {
	objectParam := data.DbRtObjectParam{}
	result := dao.DB.Where("object = ? AND param = ?", objectId, paramId).Delete(&objectParam)
	return result.Error
}

func (dao *EposDbDao) GetRtObjectParamById(objectId string, paramId string) (data.DbRtObjectParam, error) {
	objectParam := data.DbRtObjectParam{}
	result := dao.DB.Where("object = ? AND param = ?", objectId, paramId).First(&objectParam)
	return objectParam, result.Error
}

func (dao *EposDbDao) GetRtObjectParamByObject(objectId string) ([]data.DbRtObjectParam, error) {
	var objectParams []data.DbRtObjectParam
	result := dao.DB.Where("object = ?", objectId).Find(&objectParams)
	return objectParams, result.Error
}

func (dao *EposDbDao) GetRtObjectParamByParam(paramId string) ([]data.DbRtObjectParam, error) {
	var objectParams []data.DbRtObjectParam
	result := dao.DB.Where("param = ?", paramId).First(&objectParams)
	return objectParams, result.Error
}

//----------------------------------------------------------------------------------
