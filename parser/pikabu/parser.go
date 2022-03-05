package pikabu

import (
	"epos_v3/configuration"
	"epos_v3/domain"
	"epos_v3/domain/data"
	"epos_v3/parser/common"
	"epos_v3/utils"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"time"
)

type Parser struct{}

func getEntryId(URL string) string {
	id := "PKB0-ENT-"
	paramsDelim := strings.Index(URL, "?")
	if paramsDelim >= 0 {
		urlWithoutParams := strings.Split(URL, "?")[0]
		urlArr := strings.Split(urlWithoutParams, "_")
		id = id + urlArr[len(urlArr)-1]
	} else {
		urlArr := strings.Split(URL, "_")
		id = id + urlArr[len(urlArr)-1]
	}
	return id
}

func getEntryTitle(doc *goquery.Document) string {
	header := doc.Find(".story__header").First()
	title := header.Find(".story__title-link").Nodes[0].FirstChild.Data
	return utils.EncodeFromCp1251(title)
}

func getParamStoryAuthor(doc *goquery.Document, config configuration.BaseConfig, logger domain.AppLogger) data.EntryParam {
	header := doc.Find(".story__header").First()
	storyUser := header.Find(".story__user-link")

	author, isExist := storyUser.Attr("data-name")
	if !isExist {
		logger.ErrorLogger.Println("[data-name] attribute not found")
	}

	return data.EntryParam{
		Id: "",
		Code: data.ParamCode{
			Code:  "PCODE-ATR",
			Title: "author",
			Alias: "Author",
		},
		Type: data.ParamType{
			Code:  "PTYPE-STR",
			Title: "string",
			Alias: "String",
		},
		Value: author,
		Extra: "n/a",
	}
}

func getParamsGraphAuthor(doc *goquery.Document, config configuration.BaseConfig, logger domain.AppLogger) data.EntryParam {
	var author string

	storyContent := doc.Find(".story__content-inner").First()
	storyContent.Find(".story-block_type_text").Each(func(i int, selection *goquery.Selection) {
		if selection != nil && len(selection.Nodes) == 1 {
			authorLink := selection.Find("a")
			if authorLink != nil && authorLink.Nodes != nil {
				tAuthor, isExist := authorLink.Attr("href")
				if !isExist {
					logger.ErrorLogger.Println("[href] attribute not found")
				} else {
					author += tAuthor
				}
			} else {
				authorLink = selection.Find("p")
				author += authorLink.Text()
			}
		}
	})

	return data.EntryParam{
		Id: "",
		Code: data.ParamCode{
			Code:  "PCODE-ART",
			Title: "artist",
			Alias: "Artist",
		},
		Type: data.ParamType{
			Code:  "PTYPE-STR",
			Title: "string",
			Alias: "String",
		},
		Value: utils.EncodeFromCp1251(author),
		Extra: "n/a",
	}
}

func getParamStoryDate(doc *goquery.Document, config configuration.BaseConfig, logger domain.AppLogger) data.EntryParam {
	header := doc.Find(".story__header").First()
	date := header.Find("time")

	dateTime, isExist := date.Attr("datetime")
	if !isExist {
		logger.ErrorLogger.Println("[datetime] attribute not found")
	}

	return data.EntryParam{
		Id: "",
		Code: data.ParamCode{
			Code:  "PCODE-PDT",
			Title: "postDate",
			Alias: "Post date",
		},
		Type: data.ParamType{
			Code:  "PTYPE-STR",
			Title: "string",
			Alias: "String",
		},
		Value: dateTime,
		Extra: "n/a",
	}
}

func getParamTags(doc *goquery.Document, config configuration.BaseConfig, logger domain.AppLogger) data.EntryParam {
	storyContent := doc.Find(".story__content-inner").First()
	imageWithTags := storyContent.Find("img").First()
	tags, isExist := imageWithTags.Attr("alt")
	if !isExist {
		logger.ErrorLogger.Println("[alt] attribute not found")
	}

	return data.EntryParam{
		Id: "",
		Code: data.ParamCode{
			Code:  "PCODE-TAG",
			Title: "tags",
			Alias: "Tags",
		},
		Type: data.ParamType{
			Code:  "PTYPE-STR",
			Title: "string",
			Alias: "String",
		},
		Value: utils.EncodeFromCp1251(tags),
		Extra: "n/a",
	}
}

func fillEntryParams(doc *goquery.Document, config configuration.BaseConfig, logger domain.AppLogger) []data.EntryParam {
	var params []data.EntryParam

	graphAuthor := getParamsGraphAuthor(doc, config, logger)

	params = append(params, getParamStoryAuthor(doc, config, logger))

	if graphAuthor.Value != "" {
		params = append(params, graphAuthor)
	}

	params = append(params, getParamStoryDate(doc, config, logger))
	params = append(params, getParamTags(doc, config, logger))

	return params
}

func fillEntryContent(doc *goquery.Document, title string, config configuration.BaseConfig, logger domain.AppLogger) []data.Content {
	var content []data.Content

	storyContent := doc.Find(".story__content-inner").First()
	storyContent.Find("img").Each(func(i int, selection *goquery.Selection) {
		largeImgLink, isExist := selection.Attr("data-large-image")
		if isExist {
			content = append(content, data.Content{
				Id:    "",
				Entry: "",
				Title: title,
				Type: data.ContentType{
					Code:  "CTYPE-IMG",
					Title: "image",
					Alias: "Image path",
				},
				Data:   largeImgLink,
				Access: 0,
			})
		} else {
			logger.ErrorLogger.Println("[data-large-image] not found")
		}
	})

	return content
}

func (Parser) ParseEntry(URL string, config configuration.BaseConfig, logger domain.AppLogger) (bool, data.Entry) {
	sourceConfig := utils.GetSourceConfig("configuration/source/pikabuConfig.json")
	result, doc := common.GetHtmlFormUrl(URL, sourceConfig.Secure.AuthHeader)
	if result {
		var entry data.Entry
		entry.Id = getEntryId(URL)
		entry.Title = getEntryTitle(doc)
		entry.Added = time.Now()
		entry.URL = URL
		entry.Completed = false
		entry.Ignored = false
		entry.Access = 3
		entry.Params = fillEntryParams(doc, config, logger)
		entry.Content = fillEntryContent(doc, entry.Title, config, logger)
		entry.Objects = []data.Object{}
		return true, entry
	}

	return false, data.Entry{}
}

func (Parser) ValidateEntry(checkType int, entry data.Entry, logger domain.AppLogger) bool {
	var result bool

	if entry.Title == "" {
		result = false
	}
	if entry.URL == "" {
		result = false
	}
	if !strings.Contains(entry.Id, "PKB0-ENT-") {
		result = false
	}
	for _, content := range entry.Content {
		if content.Data == "" {
			result = false
		}
		if content.Title == "" {
			result = false
		}
	}
	if len(entry.Params) != checkType {
		result = false
	}
	for _, param := range entry.Params {
		if param.Value == "" {
			result = false
		}
	}

	result = true
	return result
}
