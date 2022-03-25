package zpl

import (
	"strings"
)

var ZplPdfHandlers = map[string]ParametersHandlerFunc{
	"^FD": fdHandler,
	"^FO": foHandler,
}

type ParametersHandlerFunc func(c string, b []byte) []CommandParameter

func fdHandler(cToken string, buffer []byte) []CommandParameter {
	var result []CommandParameter

	// Removing zpl command token from parameters strig, the text field in this case.
	v := strings.Replace(string(buffer), strings.Replace(cToken, "^", "", -1), "", -1)
	result = append(result, CommandParameter{Value: v})

	return result
}

func foHandler(cToken string, buffer []byte) []CommandParameter {
	var result []CommandParameter

	// Removing zpl command token from parameters strig, the text field in this case.
	v := strings.Replace(string(buffer), strings.Replace(cToken, "^", "", -1), "", -1)
	// Explode coordinates to separate params
	vals := strings.Split(v, ",")

	for _, coord := range vals {
		result = append(result, CommandParameter{Value: coord})
	}

	return result
}
