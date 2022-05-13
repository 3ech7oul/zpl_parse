package zpl

import (
	"strings"
)

func lhParser(cToken string, buffer []byte) []string {
	var result []string

	v := removeZplCommandToken(string(buffer), cToken)
	// Explode coordinates to separate params
	vals := strings.Split(v, ",")

	result = append(result, vals...)

	return result
}
