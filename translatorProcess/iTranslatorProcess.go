package translatorProcess

type ITranslatorProcess interface {
	Execute(textOrigen string, textDestiny string, text string)
}
