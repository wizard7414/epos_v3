package common

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/proxy"
	"net/http"
)

func GetHtmlFormUrl(url string, authHeader string) (bool, *goquery.Document) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("Cookie", authHeader)
	if err != nil {
		return false, nil
	}
	response, _ := client.Do(request)
	if response.StatusCode != 200 {
		return false, nil
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return false, nil
	}

	return true, doc
}

func GetHtmlFromUrlAuth(pathUrl string, authHeader string) (bool, *goquery.Document) {
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, nil)
	if err != nil {
		return false, nil
	}
	tbTransport := &http.Transport{Dial: dialer.Dial}

	request, requestErr := http.NewRequest("GET", pathUrl, nil)
	if requestErr != nil {
		return false, nil
	}

	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
	request.Header.Add("Accept-Encoding", "gzip, deflate, br")
	request.Header.Add("Accept-Language", "ru-RU,ru;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Cache-Control", "max-age=0")
	request.Header.Add("Cookie", authHeader)
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("Host", "pikabu.ru")
	request.Header.Add("Referer", "https://pikabu.ru/story/anime_art_8522982?utm_source=andpostshare&utm_medium=sharing")
	request.Header.Add("Sec-Fetch-Dest", "document")
	request.Header.Add("Sec-Fetch-Mode", "navigate")
	request.Header.Add("Sec-Fetch-Site", "same-origin")
	request.Header.Add("Sec-Fetch-User", "?1")
	request.Header.Add("TE", "?trailers")
	request.Header.Add("Upgrade-Insecure-Requests", "1")
	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:94.0) Gecko/20100101 Firefox/94.0")

	client := &http.Client{Transport: tbTransport}
	response, responseErr := client.Do(request)
	if responseErr != nil {
		return false, nil
	}

	if response.StatusCode != 200 {
		return false, nil
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return false, nil
	}

	return true, doc
}
