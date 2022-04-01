package zpl

func fdParser(cToken string, buffer []byte) []CommandParameter {
	var result []CommandParameter

	v := removeZplCommandToken(string(buffer), cToken)
	result = append(result, CommandParameter{Value: v})

	return result
}
