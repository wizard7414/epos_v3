package logger

import (
	"epos_v3/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// GeneralLogger exported
var GeneralLogger *log.Logger

// ErrorLogger exported
var ErrorLogger *log.Logger

func init() {
	conf := utils.InitConfiguration("./configuration/baseConfig.json")

	absPath, err := filepath.Abs(conf.LogPath)
	if err != nil {
		fmt.Println("Error reading given path:", err)
	}

	generalLog, err := os.OpenFile(absPath+"/"+utils.GetIsoDate()+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	GeneralLogger = log.New(generalLog, "General:", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(generalLog, "Error:", log.Ldate|log.Ltime|log.Lshortfile)
}
