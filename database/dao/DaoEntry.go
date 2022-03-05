package dao

import (
	"epos_v3/domain/data"
	"epos_v3/utils"
)

func (dao *EposDao) EntryAdd(entry data.Entry) error {

	for p := range entry.Params {
		param := entry.Params[p]

		_, err := dao.DB.GetEntryParamById(param.Id)
		if err != nil {
			if err.Error() == "record not found" {
				paramErr := dao.DB.CreateEntryParam(data.DbEntryParam{
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
		_, relErr := dao.DB.GetRtEntryParamById(entry.Id, param.Id)
		if relErr != nil {
			if relErr.Error() == "record not found" {
				relErr = dao.DB.CreateRtEntryParam(data.DbRtEntryParam{
					Entry: entry.Id,
					Param: param.Id,
				})
				if relErr != nil {
					return relErr
				}
			}
		}
	}

	for c := range entry.Content {
		content := entry.Content[c]
		_, err := dao.DB.GetContentById(content.Id)
		if err != nil {
			if err.Error() == "record not found" {
				paramErr := dao.DB.CreateContent(data.DbContent{
					Id:     content.Id,
					Entry:  content.Entry,
					Title:  content.Title,
					Type:   content.Type.Code,
					Data:   content.Data,
					Access: content.Access,
				})
				if paramErr != nil {
					return paramErr
				}
			}
		}
	}

	for o := range entry.Objects {
		obj := entry.Objects[o]
		_, relErr := dao.DB.GetRtEntryObjectById(entry.Id, obj.Id)
		if relErr != nil {
			if relErr.Error() == "record not found" {
				relErr = dao.DB.CreateRtEntryObject(data.DbRtEntryObject{
					Entry:  entry.Id,
					Object: obj.Id,
				})
				if relErr != nil {
					return relErr
				}
			}
		}
	}

	return dao.DB.CreateEntry(data.DbEntry{
		Id:        entry.Id,
		Added:     entry.Added.Unix(),
		Title:     entry.Title,
		URL:       entry.URL,
		Completed: utils.BoolToInt(entry.Completed),
		Ignored:   utils.BoolToInt(entry.Ignored),
		Access:    entry.Access,
	})
}

//----------------------------------------------------------------------------------

func (dao *EposDbDao) CreateEntry(entry data.DbEntry) error {
	return dao.DB.Save(&entry).Error
}

func (dao *EposDbDao) DeleteEntry(id string) error {
	result := dao.DB.Delete(&data.DbEntry{}, id)
	return result.Error
}

func (dao *EposDbDao) GetEntryById(id string) (data.DbEntry, error) {
	entry := data.DbEntry{}
	result := dao.DB.Where("id = ?", id).First(&entry)
	return entry, result.Error
}

func (dao *EposDbDao) GetEntryByUrl(url string) (data.DbEntry, error) {
	entry := data.DbEntry{}
	result := dao.DB.Where("url = ?", url).First(&entry)
	return entry, result.Error
}

func (dao *EposDbDao) GetAllEntries() ([]data.DbEntry, error) {
	var entries []data.DbEntry
	result := dao.DB.Find(&entries)
	return entries, result.Error
}

//----------------------------------------------------------------------------------
