package zpl

import (
	"strings"
)

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
