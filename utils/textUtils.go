package utils

import (
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strings"
)

func EncodeFromCp1251(text string) string {

	sr := strings.NewReader(text)
	tr := transform.NewReader(sr, charmap.Windows1251.NewDecoder())
	buf, err := ioutil.ReadAll(tr)
	if err != err {
		// обработка ошибки
	}

	return string(buf) // строка в UTF-8
}
