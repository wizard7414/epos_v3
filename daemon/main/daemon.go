package main

import (
	"epos_v3/configuration"
	"epos_v3/database/dao"
	"epos_v3/domain"
	"epos_v3/domain/data"
	"epos_v3/domain/resource"
	logger2 "epos_v3/logger"
	"epos_v3/parser"
	"epos_v3/utils"
	"fmt"
	"os"
	"strconv"
)

var appLogger = domain.AppLogger{
	GeneralLogger: logger2.GeneralLogger,
	ErrorLogger:   logger2.ErrorLogger,
}

var appChan = domain.AppChan{
	ProcessChan: make(chan string, 10),
	SuccessChan: make(chan string, 10),
	FailedChan:  make(chan string, 10),
	EntryChan:   make(chan data.Entry, 10),
	ObjectChan:  make(chan data.Object, 10),
}

func loadData(workList chan string, data []string, logger domain.AppLogger) {
	for url := range data {
		logger.GeneralLogger.Println("[Stage 0][Loaded] " + data[url])
		workList <- data[url]
	}

	close(workList)
}

func process(channels domain.AppChan, config configuration.BaseConfig, logger domain.AppLogger) {
	for url := range channels.ProcessChan {
		logger.GeneralLogger.Println("[Stage 1][Start parsing process] " + url)
		parseResult, parsedEntry := parser.ParseUrl(url, config, logger)
		if parseResult {
			logger.GeneralLogger.Println("[Stage 1][Parsing process success] " + url)
			channels.SuccessChan <- url
			if parsedEntry.Id != "" {
				logger.GeneralLogger.Println("[Stage 1][Successfully parsed entry] " + parsedEntry.URL)
				channels.EntryChan <- parsedEntry
			}
		} else {
			logger.GeneralLogger.Println("[Stage 1][Parsing process failed] " + url)
			channels.FailedChan <- url
		}
	}
}

func saveDataToList(channel chan string, config configuration.BaseConfig, fileName string, logger domain.AppLogger) {
	for url := range channel {
		err := utils.CreateFileFromStrings(config.Parse.ResultListTarget, fileName, []string{url})
		if err != nil {
			logger.GeneralLogger.Println("[Stage 2][Error saving URL into " + fileName + ".txt file] " + url)
		} else {
			logger.GeneralLogger.Println("[Stage 2][Save URL into " + fileName + ".txt file] " + url)
		}
	}
}

func getEntryAuthor(entry data.Entry) string {
	author := ""

	for _, param := range entry.Params {
		if param.Code.Code == "PCODE-ATR" {
			author = param.Value
		}
	}
	return author
}

func saveEntryContent(entry data.Entry, config configuration.BaseConfig, logger domain.AppLogger) {
	contentPath := "/home/wizard/EPOS-LC/DMP/P/PKB0_-_pikabu"

	author := getEntryAuthor(entry)

	entryAuthorPath := contentPath + "/" + author
	entryImgPath := contentPath + "/" + author + "/" + entry.Id + "/"

	_, err := os.Stat(entryAuthorPath)
	if os.IsNotExist(err) {
		dirErr := os.Mkdir(entryAuthorPath, 0755)
		if dirErr != nil {
			logger2.ErrorLogger.Println("Unable to create directory for entry: " + entry.Id)
		}
	}

	dirErr := os.Mkdir(entryImgPath, 0755)
	if dirErr != nil {
		logger2.ErrorLogger.Println("Unable to create directory for entry: " + entry.Id)
	}

	for _, image := range entry.Content {
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

}

func processEntries(channel chan data.Entry, config configuration.BaseConfig, logger domain.AppLogger) {
	generalBase := dao.EposDao{
		DB: dao.EposDbDao{},
	}
	err := generalBase.InitDb("/home/wizard/Документы/epos_v3_db/epos-db-3.2.db")
	if err != nil {
		panic("Unable to start test set!")
	}
	defer generalBase.CloseDb()

	entryCount := 0
	paramCounter := 27199
	contentCounter := 32360

	for entry := range channel {
		entryCount++
		_, entryErr := generalBase.DB.GetEntryByUrl(entry.URL)
		if entryErr != nil {
			if entryErr.Error() == "record not found" {

				for c := range entry.Content {
					contentCounter++
					entry.Content[c].Entry = entry.Id
					entry.Content[c].Id = utils.GenerateContentId(contentCounter)
				}
				for p := range entry.Params {
					stored, err := generalBase.DB.GetEntryParamByValue(entry.Params[p].Value)
					if err != nil {
						if err.Error() == "record not found" {
							paramCounter++
							entry.Params[p].Id = utils.GenerateEntryParamId(paramCounter)
						}
					} else {
						entry.Params[p].Id = stored.Id
					}
				}
				fmt.Println("[" + strconv.Itoa(entryCount) + "]" + entry.Id + " added")
				saveEntryContent(entry, config, logger)

				err := generalBase.EntryAdd(entry)
				if err != nil {
					return
				}
			}
		} else {
			fmt.Println("[" + strconv.Itoa(entryCount) + "]" + entry.Id + " skipped")
		}
	}

}

func main() {
	conf := utils.InitConfiguration("./configuration/baseConfig.json")

	parseList := utils.ReadParseListFromFile(conf.Parse.ParseList, appLogger)
	go loadData(appChan.ProcessChan, parseList, appLogger)

	for i := 0; i < conf.ProcessPoolSize; i++ {
		go process(appChan, *conf, appLogger)
	}

	go processEntries(appChan.EntryChan, *conf, appLogger)

	go saveDataToList(appChan.SuccessChan, *conf, "success", appLogger)
	go saveDataToList(appChan.FailedChan, *conf, "error", appLogger)

	for {
	}
}
