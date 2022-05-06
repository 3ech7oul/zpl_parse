package zpl

import (
	"strings"
)

/*
	Format: ^FPd,g
	d =direction
		Values:
			H = horizontal printing (left to right) 
			V = vertical printing (top to bottom) 
			R = reverse printing (right to left)
	g = additional inter- character gap (in dots)
		Values: 
			0 to 9999
			Default: 0 if no value is entered
*/
func fpParser(cToken string, buffer []byte) []CommandParameter {
	var result []CommandParameter

	v := removeZplCommandToken(string(buffer), cToken)

	vals := strings.Split(v, ",")

	for _, params := range vals {
		result = append(result, CommandParameter{Value: params})
	}

	return result
}