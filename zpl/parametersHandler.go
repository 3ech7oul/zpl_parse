package zpl

import (
	"strings"
)

var ZplPdfHandlers = map[string]ParametersHandlerFunc{
	"^FD": fdHandler,
}

type ParametersHandlerFunc func(c string, b []byte) []CommandParameter

func fdHandler(cToken string, buffer []byte) []CommandParameter {
	var result []CommandParameter

	// Removing zpl command token from parameters strig, the text field in this case.
	v := strings.Replace(string(buffer), strings.Replace(cToken, "^", "", -1), "", -1)
	result = append(result, CommandParameter{Value: v})

	return result
}
