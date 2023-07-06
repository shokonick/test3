package utils

import (
	"github.com/carlmjohnson/requests"
	"os"
	"context"
	"encoding/base64"
)

type ReversoTTS struct {
	Id string
	Voice string
}

func TTSGoogle(lang string, query string) string {
	params := "?tl="+lang+"&q="+query+"&client=tw-ob"
	var file string
	url := "https://translate.google.com/translate_tts"+params
	err := requests.
		URL(url).
		ToString(&file).
		Fetch(context.Background())
	if err != nil {
		file = ""
	}
	return file
}
func TTSReverso(lang string, query string) string {
	var TTSData = []ReversoTTS{
		// http://voice.reverso.net/RestPronunciation.svc/v1/output=json/GetAvailableVoices with randomized deduplication
		ReversoTTS{
			Id:   "ar",
			Voice: "Mehdi22k",
		},
		ReversoTTS{
			Id:   "zh",
			Voice: "Lulu22k",
		},
		ReversoTTS{
			Id:   "cz",
			Voice: "Eliska22k",
		},
		ReversoTTS{
			Id:   "dk",
			Voice: "Mette22k",
		},
		ReversoTTS{
			Id:   "nl",
			Voice: "Daan22k",
		},
		ReversoTTS{
			Id:   "en",
			Voice: "Will22k",
		},
		ReversoTTS{
			Id:   "fr",
			Voice: "Margaux22k",
		},
		ReversoTTS{
			Id:   "de",
			Voice: "Andreas22k",
		},
		ReversoTTS{
			Id:   "gr",
			Voice: "Dimitris22k",
		},
		ReversoTTS{
			Id:   "heb",
			Voice: "he-IL-Asaf",
		},
		ReversoTTS{
			Id:   "it",
			Voice: "Chiara22k",
		},
		ReversoTTS{
			Id:   "jp",
			Voice: "Sakura22k",
		},
		ReversoTTS{
			Id:   "kr",
			Voice: "Minji22k",
		},
		ReversoTTS{
			Id:   "pl",
			Voice: "Monika22k",
		},
		ReversoTTS{
			Id:   "pt",
			Voice: "Celia22k",
		},
		ReversoTTS{
			Id:   "ro",
			Voice: "ro-RO-Andrei",
		},
		ReversoTTS{
			Id:   "ru",
			Voice: "Alyona22k",
		},
		ReversoTTS{
			Id:   "es",
			Voice: "Antonio22k",
		},
		ReversoTTS{
			Id:   "se",
			Voice: "Erik22k",
		},
		ReversoTTS{
			Id:   "tr",
			Voice: "Ipek22k",
		},
	}
	// https://voice.reverso.net/RestPronunciation.svc/v1/output=json/GetVoiceStream/voiceName=Lulu22k?voiceSpeed=80&inputText=6K%20V6aqM Base64 input text
	text := base64.StdEncoding.EncodeToString([]byte(query))
    var voice string
    for _, s := range TTSData {
        if s.Id == lang {
            voice = s.Voice
            break
        }
    }

	params := "voiceName="+voice+"?voiceSpeed=80&inputText="+text

	var file string

	url := "https://voice.reverso.net/RestPronunciation.svc/v1/output=json/GetVoiceStream/"+params

	UserAgent, ok := os.LookupEnv("SIMPLYTRANSLATE_USER_AGENT")
	if !ok {
		UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"
	}

	err := requests.
		URL(url).
		ToString(&file).
		UserAgent(UserAgent).
		Fetch(context.Background())
	if err != nil {
		file = ""
	}
	return file
}
