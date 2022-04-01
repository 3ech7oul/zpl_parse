package zpl

import (
	"strings"
)

var ZplPdfParsers = map[string]ParametersParserFunc{
	"^FD": fdParser,
	"^FO": foParser,
}

type ParametersParserFunc func(c string, b []byte) []CommandParameter

// Removing zpl command token from parameters strig, the text field in this case.
func removeZplCommandToken(str string, token string) string {
	return strings.Replace(string(str), strings.Replace(token, "^", "", -1), "", -1)
}
