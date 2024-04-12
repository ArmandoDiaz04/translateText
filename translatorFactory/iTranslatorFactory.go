package translatorFactory

import "TraductorTexto/translatorFactory/translator"

type ITranslatorFactory interface {
	GetTranslator(id string) translator.ITranslator
}
