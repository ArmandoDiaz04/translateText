package appConfig

import (
	"TraductorTexto/translatorFactory"
	"TraductorTexto/translatorFactory/translator"
	"TraductorTexto/translatorProcess"
)
//hola
type AppConfig struct {
	TranslatorProcess translatorProcess.ITranslatorProcess
}

func NewAppConfig() *AppConfig {
	fromTexToBinary := translator.NewtextToBinary()
	fromTextToHex := translator.NewTextToHex()
	fromHexToText := translator.NewHexToText()

	translators := map[string]translator.ITranslator{
		"TEXT-BINARY": fromTexToBinary,
		"TEXT-HEX":    fromTextToHex,
		"HEX-TEXT":    fromHexToText,
	}
	translatorFactorSvc := translatorFactory.NewTranslatorFactory(translators)

	translatorProcess := translatorProcess.NewTranslatorProcess(translatorFactorSvc)

	return &AppConfig{
		TranslatorProcess: translatorProcess,
	}
}
