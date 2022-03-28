package zpl

import (
	"strings"
)

var ZplPdfParsers = map[string]ParametersParserFunc{
	"^FD": fdParser,
	"^FO": foParser,
}

type ParametersParserFunc func(c string, b []byte) []CommandParameter

func fdParser(cToken string, buffer []byte) []CommandParameter {
	var result []CommandParameter

	v := removeZplCommandToken(string(buffer), cToken)
	result = append(result, CommandParameter{Value: v})

	return result
}

func foParser(cToken string, buffer []byte) []CommandParameter {
	var result []CommandParameter

	v := removeZplCommandToken(string(buffer), cToken)
	// Explode coordinates to separate params
	vals := strings.Split(v, ",")

	for _, coord := range vals {
		result = append(result, CommandParameter{Value: coord})
	}

	return result
}

// Removing zpl command token from parameters strig, the text field in this case.
func removeZplCommandToken(str string, token string) string {
	return strings.Replace(string(str), strings.Replace(token, "^", "", -1), "", -1)
}
