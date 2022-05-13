package zpl

func fdParser(cToken string, buffer []byte) []string {
	var result []string

	v := removeZplCommandToken(string(buffer), cToken)
	result = append(result, v)

	return result
}
