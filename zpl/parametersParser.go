package zpl

import (
	"strings"
)

var ZplPdfParsers = map[string]ParametersParserFunc{
	"^FD": fdParser,
	"^LH": lhParser,
}

type ParametersParserFunc func(cToken string, buffer []byte) []string

func defaultParser(cToken string, buffer []byte) []string {
	var result []string

	v := removeZplCommandToken(string(buffer), cToken)

	vals := strings.Split(v, ",")

	for _, coord := range vals {
		result = append(result, coord)
	}

	return result
}

// Removing zpl command token from parameters strig, the text field in this case.
func removeZplCommandToken(str string, token string) string {
	return strings.Replace(string(str), strings.Replace(token, "^", "", -1), "", -1)
}
