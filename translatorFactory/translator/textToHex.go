package translator

import "encoding/hex"

type textToHex struct{}

func NewTextToHex() ITranslator {
	return &textToHex{}
}
func (t *textToHex) Translate(text string) string {
	src := []byte(text)
	return hex.EncodeToString(src)
}
