package translator

import "fmt"

type textToBinary struct{}

func NewtextToBinary() ITranslator {
	return &textToBinary{}
}

func (t *textToBinary) Translate(text string) string {
	res := ""
	for _, c := range text {
		res = fmt.Sprintf("%s%.8b", res, c)
	}
	return res
}
