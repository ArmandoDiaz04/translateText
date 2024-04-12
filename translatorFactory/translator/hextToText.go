package translator

import "encoding/hex"

type hexToText struct {
}

func NewHexToText() ITranslator {
	return &hexToText{}
}

func (h hexToText) Translate(text string) string {

	decoded, _ := hex.DecodeString(text)

	return string(decoded)
}
