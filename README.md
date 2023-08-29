# Mozhi
![Matrix](https://img.shields.io/matrix/%23mozhi%3Aprojectsegfau.lt)
Mozhi (spelt moḻi) is an alternative-frontend for many translation engines.

It was initially made as a maintained fork/rewrite of [simplytranslate](https://codeberg.org/SimpleWeb/SimplyTranslate-Web), but has grown to have a lot more features as well!

I'm initially focusing on the api and engines, but eventually Mozhi will have a functioning CLI and webapp.

## Supported Engines:
- Google
- Reverso
- DeepL
- LibreTranslate
- Yandex
- IBM Watson
- MyMemory
- DuckDuckGo (almost 1-1 with Bing Translate)

## Building
```
go mod download
go run github.com/swaggo/swag/cmd/swag@latest init
go build -o mozhi
```

## API Docs
Mozhi makes use of swagger (with the swagger fiber middleware) to manage the documentation of the API.

You can find it in /api/swagger of any instance.

## Features
- An all mode where the responses of all supported engines will be shown.
- Autodetect which will show the language that was detected
- Text-To-Speech for multiple engines
- A good API (subjective :P)
- All the stuff you expect from a translation utility :)

## Etymology
Mozhi is the word in Tamil for language. Simple as that :P
