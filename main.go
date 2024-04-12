package main

import (
	"TraductorTexto/appConfig"
)

func main() {
	var textOrigen = "TEXT"
	var textDestination = "BINARY"
	var text = "HOLA"

	app := appConfig.NewAppConfig()

	app.TranslatorProcess.Execute(textOrigen, textDestination, text)
}
