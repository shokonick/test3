package utils

import (
	"github.com/otiai10/gosseract/v2"
)

func ImgTxt(file string) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage(file)
	text, err := client.Text()
	return text, err
}
