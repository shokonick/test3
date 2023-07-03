package utils

import (
	"github.com/otiai10/gosseract/v2"
)

func ImgTxt(file string) string {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage(file)
	text, _ := client.Text()
	return text
}
