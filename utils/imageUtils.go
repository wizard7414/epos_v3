package utils

import (
	"epos_v3/domain/resource"
	"io"
	"net/http"
	"os"
)

func SaveImage(image resource.Image) error {
	response, e := http.Get(image.Url)
	if e != nil {
		return e
	}
	defer response.Body.Close()

	file, err := os.Create(image.TargetUrl + image.TargetName + image.Extension)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}
