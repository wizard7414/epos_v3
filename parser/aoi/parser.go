package aoi

import (
	"epos_v3/configuration"
	"epos_v3/database/dao"
	"epos_v3/domain"
	"epos_v3/domain/data"
	"epos_v3/parser/common"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
	"time"
)

type Parser struct{}

func getEntryTitle(doc *goquery.Document, config configuration.BaseConfig, logger domain.AppLogger) string {
	return ""
}

func getIllustrator(doc *goquery.Document) data.EntryParam {
	entryIllustrator := data.EntryParam{
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
		Extra: "n/a",
	}

	doc.Find("td").Each(func(i int, tdSelection *goquery.Selection) {
		colspanOfNode, _ := tdSelection.Attr("colspan")
		styleOfNode, _ := tdSelection.Attr("style")

		if colspanOfNode == "3" && styleOfNode == "height:1em;" {
			if tdSelection.Nodes[0].FirstChild.FirstChild.Data == "Illustrator" {
				tdSelection.Find("a").Each(func(i int, selection *goquery.Selection) {
					val := selection.Nodes[0].Attr[1].Val
					entryIllustrator.Value = strings.ReplaceAll(val, " (page does not exist)", "")
				})
				if entryIllustrator.Value == "" {
					tdSelection.Find("span").Each(func(i int, selection *goquery.Selection) {
						val := selection.Nodes[0].Attr[1].Val
						entryIllustrator.Value = strings.ReplaceAll(val, " (page does not exist)", "")
					})
				}
			}
		}
	})

	return entryIllustrator
}

func (Parser) ParseEntry(URL string, config configuration.BaseConfig, logger domain.AppLogger) (bool, data.Entry) {
	result, doc := common.GetHtmlFormUrl(URL, "")
	if result {
		var entry data.Entry
		entry.Title = getEntryTitle(doc, config, logger)
		return true, entry
	}

	return false, data.Entry{}
}

func (Parser) ValidateEntry(checkType int, graph data.Entry, logger domain.AppLogger) bool {
	return false
}

func ParseInfoFromList(listUrl string, db dao.EposDao, logger domain.AppLogger) ([]data.Entry, bool) {
	result, document := common.GetHtmlFormUrl(listUrl, "")
	contentCount := 22494
	paramCount := 20125
	if result {
		var entries []data.Entry

		document.Find("tbody").Each(func(i int, tbodySelection *goquery.Selection) {
			if i == 1 {
				tbodySelection.Find("tr").Each(func(id int, trSelection *goquery.Selection) {
					if id > 0 {
						var entry data.Entry
						trSelection.Find("td").Each(func(i int, tdSelection *goquery.Selection) {
							switch i {
							case 1:
								{
									title := findEntryTitle(tdSelection)
									url := findEntryUrl(tdSelection)

									entry.Id = generateEntryId(id)
									entry.Title = title
									entry.URL = url
									entry.Added = time.Now()
									entry.Access = 5
									entry.Ignored = false
									entry.Completed = false
								}
							case 2:
								{
									rank := findEntryData(tdSelection)
									if rank != "" {
										paramId := checkParam(db, entries, entry.Params, rank)
										if paramId == "" {
											paramCount++
											paramId = createEntryParamId(paramCount)
										}

										entry.Params = append(entry.Params, data.EntryParam{
											Id: paramId,
											Code: data.ParamCode{
												Code:  "PCODE-RNK",
												Title: "rank",
												Alias: "Rank",
											},
											Type: data.ParamType{
												Code:  "PTYPE-STR",
												Title: "string",
												Alias: "String",
											},
											Value: rank,
											Extra: "n/a",
										})
									}
								}
							case 3:
								{
									entryType := findEntryData(tdSelection)
									if entryType != "" {
										paramId := checkParam(db, entries, entry.Params, entryType)
										if paramId == "" {
											paramCount++
											paramId = createEntryParamId(paramCount)
										}

										entry.Params = append(entry.Params, data.EntryParam{
											Id: paramId,
											Code: data.ParamCode{
												Code:  "PCODE-TYP",
												Title: "type",
												Alias: "Type",
											},
											Type: data.ParamType{
												Code:  "PTYPE-STR",
												Title: "string",
												Alias: "String",
											},
											Value: entryType,
											Extra: "n/a",
										})
									}
								}
							case 4:
								{
									element := findEntryData(tdSelection)
									if element != "" {
										paramId := checkParam(db, entries, entry.Params, element)
										if paramId == "" {
											paramCount++
											paramId = createEntryParamId(paramCount)
										}

										entry.Params = append(entry.Params, data.EntryParam{
											Id: paramId,
											Code: data.ParamCode{
												Code:  "PCODE-ELM",
												Title: "element",
												Alias: "Element",
											},
											Type: data.ParamType{
												Code:  "PTYPE-STR",
												Title: "string",
												Alias: "String",
											},
											Value: element,
											Extra: "n/a",
										})
									}
								}
							}
						})
						result, document := common.GetHtmlFormUrl(entry.URL, "")
						if result {
							entryIllustrator := getIllustrator(document)
							entryImages := getEntryImages(document, entry.Id, entry.Title)
							entryQuotes := getQuotes(document)
							entryInfo := getInfo(document)

							logger.ErrorLogger.Println(entry.Title + " Author: " + entryIllustrator.Value)

							if entryIllustrator.Value != "" {
								paramId := checkParam(db, entries, entry.Params, entryIllustrator.Value)
								if paramId == "" {
									paramCount++
									entryIllustrator.Id = createEntryParamId(paramCount)
								} else {
									entryIllustrator.Id = paramId
								}

								entry.Params = append(entry.Params, entryIllustrator)
							}
							if len(entryImages) > 0 {
								for c := range entryImages {
									image := entryImages[c]
									contentId := checkContent(entries, "CTYPE-IMG", image.Data)
									if contentId == "" {
										contentCount++
										entryImages[c].Id = generateContentId(contentCount)
									} else {
										entryImages[c].Id = contentId
									}
								}
								entry.Content = entryImages
							}
							if len(entryQuotes) > 0 {
								for q := range entryQuotes {
									paramId := checkParam(db, entries, entry.Params, entryQuotes[q].Value)
									if paramId == "" {
										paramCount++
										entryQuotes[q].Id = createEntryParamId(paramCount)
									} else {
										entryQuotes[q].Id = paramId
									}

									entry.Params = append(entry.Params, entryQuotes[q])
								}
							}
							if entryInfo.Value != "" {
								paramId := checkParam(db, entries, entry.Params, entryInfo.Value)
								if paramId == "" {
									paramCount++
									entryInfo.Id = createEntryParamId(paramCount)
								} else {
									entryInfo.Id = paramId
								}

								entry.Params = append(entry.Params, entryInfo)
							}
						}

						entries = append(entries, entry)
					}
				})
			}
		})
		return entries, result
	}

	return []data.Entry{}, result
}

func generateEntryId(id int) string {
	countStr := strconv.Itoa(id)
	for i := len(countStr); i < 10; i++ {
		countStr = "0" + countStr
	}
	return "AOI0-ENT-" + countStr
}

func checkParam(db dao.EposDao, list []data.Entry, params []data.EntryParam, value string) string {
	storedParam, err := db.DB.GetEntryParamByValue(value)
	if err != nil {
		if err.Error() == "record not found" {
			for e := range list {
				entry := list[e]
				for p := range entry.Params {
					param := entry.Params[p]
					if param.Value == value {
						return param.Id
					}
				}
			}

			for p := range params {
				param := params[p]
				if param.Value == value {
					return param.Id
				}
			}
		}
	} else {
		return storedParam.Id
	}

	return ""
}

func checkContent(list []data.Entry, code string, value string) string {
	for e := range list {
		entry := list[e]
		for p := range entry.Content {
			content := entry.Content[p]
			if content.Type.Code == code && content.Data == value {
				return content.Id
			}
		}
	}

	return ""
}

func createEntryParamId(count int) string {
	countStr := strconv.Itoa(count)
	for i := len(countStr); i < 10; i++ {
		countStr = "0" + countStr
	}
	return "ENT-PRM-" + countStr
}

func generateContentId(count int) string {
	countStr := strconv.Itoa(count)
	for i := len(countStr); i < 10; i++ {
		countStr = "0" + countStr
	}
	return "ENT-CNT-" + countStr
}

func findEntryTitle(selection *goquery.Selection) string {
	return selection.Nodes[0].FirstChild.Attr[1].Val
}

func findEntryUrl(selection *goquery.Selection) string {
	return "https://ishtaria.fandom.com" + selection.Nodes[0].FirstChild.Attr[0].Val
}

func findEntryData(selection *goquery.Selection) string {
	rank := selection.Nodes[0].FirstChild.Data
	return strings.ReplaceAll(rank, "\n", "")
}

func getEntryImages(doc *goquery.Document, entryId string, entryTitle string) []data.Content {
	var entryImages []data.Content

	doc.Find("td").Each(func(tdId int, tdSelection *goquery.Selection) {
		colspanOfNode, _ := tdSelection.Attr("colspan")
		rowspanOfNode, _ := tdSelection.Attr("rowspan")

		if colspanOfNode == "3" && (rowspanOfNode == "9" || rowspanOfNode == "7" || rowspanOfNode == "11") {
			tdSelection.Find("a").Each(func(i int, selection *goquery.Selection) {
				if selection.Nodes[0].Attr[0].Val != "" {
					entryImages = append(entryImages, data.Content{
						Entry: entryId,
						Title: entryTitle,
						Type: data.ContentType{
							Code:  "CTYPE-IMG",
							Title: "image",
							Alias: "Image path",
						},
						Data:   selection.Nodes[0].Attr[0].Val,
						Access: 5,
					})
				}
			})
		}
	})

	return entryImages
}

func getQuotes(doc *goquery.Document) []data.EntryParam {
	var entryQuotes []data.EntryParam

	doc.Find("td").Each(func(tdId int, tdSelection *goquery.Selection) {
		colspanOfNode, _ := tdSelection.Attr("colspan")

		if colspanOfNode == "3" {
			if len(tdSelection.Nodes) > 0 {
				if tdSelection.Nodes[0].FirstChild != nil {
					if tdSelection.Nodes[0].FirstChild.FirstChild != nil {
						if tdSelection.Nodes[0].FirstChild.FirstChild.Data == "Quote" {
							tdSelection.Find("div").Each(func(i int, selection *goquery.Selection) {
								if len(selection.Nodes) > 0 {
									if selection.Nodes[0].FirstChild != nil {
										entryQuotes = append(entryQuotes, data.EntryParam{
											Id: "",
											Code: data.ParamCode{
												Code:  "PCODE-QOT",
												Title: "quote",
												Alias: "Quote",
											},
											Type: data.ParamType{
												Code:  "PTYPE-STR",
												Title: "string",
												Alias: "String",
											},
											Value: strings.ReplaceAll(selection.Nodes[0].FirstChild.Data, "\"", ""),
											Extra: "n/a",
										})
									}
								}
							})
						}
					}
				}
			}
		}
	})

	return entryQuotes
}

func getInfo(doc *goquery.Document) data.EntryParam {
	var entryInfo data.EntryParam

	doc.Find("td").Each(func(tdId int, tdSelection *goquery.Selection) {
		colspanOfNode, _ := tdSelection.Attr("colspan")

		if colspanOfNode == "3" {
			tdSelection.Find("div").Each(func(i int, selection *goquery.Selection) {
				styleOfNode, _ := selection.Attr("style")

				if styleOfNode == "overflow:auto;height:180px;" && i == 0 {
					if len(selection.Nodes) > 0 {
						if selection.Nodes[0].FirstChild != nil {
							entryInfo.Id = ""
							entryInfo.Code = data.ParamCode{
								Code:  "PCODE-INF",
								Title: "info",
								Alias: "Info",
							}
							entryInfo.Type = data.ParamType{
								Code:  "PTYPE-TXT",
								Title: "text",
								Alias: "Text",
							}
							entryInfo.Extra = "n/a"
							entryInfo.Value = selection.Nodes[0].FirstChild.Data
						}
					}
				}
			})
		}
	})

	return entryInfo
}
