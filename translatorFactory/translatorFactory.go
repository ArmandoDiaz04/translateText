package translatorFactory

import "TraductorTexto/translatorFactory/translator"

type translatorFactory struct {
	translators map[string]translator.ITranslator
}

func NewTranslatorFactory(translators map[string]translator.ITranslator) ITranslatorFactory {
	return &translatorFactory{
		translators: translators,
	}
}
func (t *translatorFactory) GetTranslator(id string) translator.ITranslator {
	return t.translators[id]
}
