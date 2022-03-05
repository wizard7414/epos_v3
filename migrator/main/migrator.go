package main

import (
	"epos_v3/database/dao"
	"epos_v3/domain"
	"epos_v3/domain/data"
	"epos_v3/domain/miner"
	logger2 "epos_v3/logger"
	"epos_v3/parser/aoi"
	"strconv"
	"strings"
)

var appLogger = domain.AppLogger{
	GeneralLogger: logger2.GeneralLogger,
	ErrorLogger:   logger2.ErrorLogger,
}

func CreateEntryId(object miner.ObjectV, general dao.EposDao) string {
	url := object.Url
	source := strings.Split(url, ".com/")[1]
	source = strings.Split(source, "/art/")[0]
	id := strings.Split(url, "/art/")[1]
	idArr := strings.Split(id, "-")
	src, err := general.DB.GetSourceByTitle(strings.ToLower(source))
	if err != nil {
		return ""
	}
	return "DVA0-" + src.Code + "-" + strings.ReplaceAll(idArr[len(idArr)-1], "$", "")
}

func createEntryParamId(count int) string {
	countStr := strconv.Itoa(count)
	for i := len(countStr); i < 10; i++ {
		countStr = "0" + countStr
	}
	return "ENT-PRM-" + countStr
}

func createEntryContentId(count int) string {
	countStr := strconv.Itoa(count)
	for i := len(countStr); i < 10; i++ {
		countStr = "0" + countStr
	}
	return "ENT-CNT-" + countStr
}

func fillEntryContent(object miner.ObjectV, general dao.EposDao, entryId string) []data.Content {
	var result []data.Content
	count, countErr := general.DB.GetContentCount()
	for i := range object.Attributes {
		param := object.Attributes[i]
		if param.Code.Title == "graphics" {
			if countErr == nil {
				count++
				result = append(result, data.Content{
					Id:    createEntryContentId(count),
					Entry: entryId,
					Title: object.Entry.Title,
					Type: data.ContentType{
						Code:  "CTYPE-IMG",
						Title: "image",
						Alias: "Image path",
					},
					Data:   param.Value,
					Access: 5,
				})
			}
		}
	}
	return result
}

func fillEntryParams(object miner.ObjectV, general dao.EposDao) []data.EntryParam {
	var result []data.EntryParam
	count, countErr := general.DB.GetEntryParamCount()
	for i := range object.Attributes {
		param := object.Attributes[i]
		if param.Code.Title == "graphics" {
			continue
		}
		stored, err := general.DB.GetEntryParamByValue(param.Value)
		if err != nil {
			if err.Error() == "record not found" {
				count++
				pCode, codeErr := general.DB.GetParamCodeByTitle(param.Code.Title)
				pType, typeErr := general.DB.GetParamTypeByTitle(param.AttrType.Title)
				if codeErr == nil && typeErr == nil && countErr == nil {
					result = append(result, data.EntryParam{
						Id:    createEntryParamId(count),
						Code:  pCode,
						Type:  pType,
						Value: param.Value,
						Extra: "n/a",
					})
				}
			}
		} else {
			code, codeErr := general.DB.GetParamCodeByCode(stored.Code)
			pType, typeErr := general.DB.GetParamTypeByCode(stored.Type)
			if codeErr == nil && typeErr == nil {
				result = append(result, data.EntryParam{
					Id:    stored.Id,
					Code:  code,
					Type:  pType,
					Value: stored.Value,
					Extra: stored.Extra,
				})
			}
		}
	}
	return result
}

func migrate(object miner.ObjectV, general dao.EposDao) data.Entry {
	entryId := CreateEntryId(object, general)

	entry := data.Entry{
		Id:        entryId,
		Added:     object.AddDate,
		Title:     object.Entry.Title,
		URL:       object.Entry.Url,
		Completed: false,
		Ignored:   false,
		Access:    5,
		Params:    fillEntryParams(object, general),
		Content:   fillEntryContent(object, general, entryId),
		Objects:   []data.Object{},
	}
	return entry
}

func main() {
	generalBase := dao.EposDao{
		DB: dao.EposDbDao{},
	}
	err := generalBase.InitDb("/home/wizard/Документы/epos_v3_db/epos-db-3.2.db")
	if err != nil {
		panic("Unable to start test set!")
	}
	defer generalBase.CloseDb()

	entries, _ := aoi.ParseInfoFromList("https://ishtaria.fandom.com/wiki/Unit_List", generalBase, appLogger)

	/*imgPath := "/home/wizard/EPOS-LC/DMP/A/AOI0_-_Age_of_Ishtaria/"

	for i := range entries {
		entry := entries[i]

		entryImgPath := imgPath + entry.Id + "/"

		dirErr := os.Mkdir(entryImgPath, 0755)
		if dirErr != nil {
			logger2.ErrorLogger.Println("Unable to create directory for entry: " + entry.Id)
		}

		for c := range entry.Content {
			image := entry.Content[c]
			saveErr := utils.SaveImage(resource.Image{
				Url:        image.Data,
				Extension:  ".png",
				TargetName: image.Id,
				TargetUrl:  entryImgPath,
			})
			if saveErr != nil {
				logger2.ErrorLogger.Println("Unable to save image for entry: " + entry.Id)
			}
		}
	}*/

	for i := range entries {
		entry := entries[i]
		_, storedErr := generalBase.DB.GetEntryById(entry.Id)
		if storedErr == nil {
			logger2.GeneralLogger.Println("[" + strconv.Itoa(i) + "/" + strconv.Itoa(len(entries)) + "] - Entry skipped: " + entry.Id)
		} else {
			logger2.GeneralLogger.Println("[" + strconv.Itoa(i) + "/" + strconv.Itoa(len(entries)) + "] - Entry added: " + entry.Id)
			addErr := generalBase.EntryAdd(entry)
			if addErr != nil {
				logger2.GeneralLogger.Println(addErr)
			}
		}
	}

}
