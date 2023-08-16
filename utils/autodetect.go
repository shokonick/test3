package utils
func AutoDetectWatson(query string) (string, error) {
	json := []byte(`{"text":"` + query + `"}`)
	watsonOut := PostRequest("https://www.ibm.com/demos/live/watson-language-translator/api/translate/detect", json)
	gjsonArr := watsonOut.Get("payload.languages.0.language.language").Array()
	answer := gjsonArr[0].String()
	return answer, nil
}

func AutoDetectLibreTranslate(query string) (string, error) {
	json := []byte(`{"q":"` + query + `"}`)
	libreTranslateOut := PostRequest("https://translate.argosopentech.com/detect", json)
	gjsonArr := libreTranslateOut.Get("0.language").Array()
	answer := gjsonArr[0].String()
	return answer, nil
}
