package parser

import (
	"epos_v3/configuration"
	"epos_v3/domain"
	"epos_v3/domain/data"
	"epos_v3/parser/aoi"
	"epos_v3/parser/pikabu"
	"net/url"
	"time"
)

type EntryParser interface {
	ParseEntry(URL string, config configuration.BaseConfig, logger domain.AppLogger) (bool, data.Entry)

	ValidateEntry(checkType int, entry data.Entry, logger domain.AppLogger) bool
}

type ObjectParser interface {
	ParseObject(URL string, config configuration.BaseConfig, logger domain.AppLogger) data.Object

	ValidateObject(object data.Object, logger domain.AppLogger) bool
}

const (
	Deviant       string = "www.deviantart.com"
	ArtStation           = "www.artstation.com"
	Pikabu               = "pikabu.ru"
	WikiFandom           = "www.fandom.com"
	Danbooru             = "danbooru.donmai.us"
	Gelbooru             = "gelbooru.com"
	AgeOfIstharia        = "ishtaria.fandom.com"
)

func ParseSourceEntry(parser EntryParser, URl string, config configuration.BaseConfig, logger domain.AppLogger) (bool, data.Entry) {
	parseResult, entry := parser.ParseEntry(URl, config, logger)
	if !parseResult {
		return false, data.Entry{}
	}
	validateResult := parser.ValidateEntry(4, entry, logger)
	if validateResult {
		return true, entry
	} else {
		return false, data.Entry{}
	}
}

func ParseUrl(URl string, config configuration.BaseConfig, logger domain.AppLogger) (bool, data.Entry) {
	parsedUrl, err := url.Parse(URl)
	var emptyEntry data.Entry

	if err != nil {
		logger.ErrorLogger.Println("[Stage 1.1][URl path parse error]")
		return false, emptyEntry
	}
	if parsedUrl == nil {
		logger.ErrorLogger.Println("[Stage 1.1][URL host is null]")
		return false, emptyEntry
	} else {
		var pikabuParser = pikabu.Parser{}
		var aoiParser = aoi.Parser{}

		switch parsedUrl.Host {
		case Deviant:
			{
				time.Sleep(time.Second * 3)
				return true, emptyEntry
			}
		case ArtStation:
			{
				time.Sleep(time.Second * 3)
				return true, emptyEntry
			}
		case AgeOfIstharia:
			{
				return ParseSourceEntry(aoiParser, URl, config, logger)
			}
		case Pikabu:
			{
				return ParseSourceEntry(pikabuParser, URl, config, logger)
			}
		case WikiFandom:
			{
				time.Sleep(time.Second * 3)
				return true, emptyEntry
			}
		case Danbooru:
			{
				time.Sleep(time.Second * 3)
				return true, emptyEntry
			}
		case Gelbooru:
			{
				time.Sleep(time.Second * 3)
				return true, emptyEntry
			}
		default:
			{
				logger.ErrorLogger.Println("[Stage 1.1][URL host is unknown]")
				return false, emptyEntry
			}
		}
	}
}
