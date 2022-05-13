package zpl

import (
	"strings"
)

func fpParser(cToken string, buffer []byte) []CommandParameter {
	var result []CommandParameter

	v := removeZplCommandToken(string(buffer), cToken)

	vals := strings.Split(v, ",")

	for _, params := range vals {
		result = append(result, CommandParameter{Value: params})
	}

	return result
}