package translatorProcess

import (
	"TraductorTexto/translatorFactory"
	"fmt"
)

type translatorProcess struct {
	factory translatorFactory.ITranslatorFactory
}

func NewTranslatorProcess(factory translatorFactory.ITranslatorFactory) ITranslatorProcess {
	return &translatorProcess{
		factory: factory,
	}
}
func (t *translatorProcess) Execute(textOrigen string, textDestiny string, text string) {
	translator := t.factory.GetTranslator(textOrigen + "-" + textDestiny)
	fmt.Println(translator.Translate(text))
}
