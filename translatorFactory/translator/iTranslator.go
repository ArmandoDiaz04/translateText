package translator

type ITranslator interface {
	Translate(text string) string
}
