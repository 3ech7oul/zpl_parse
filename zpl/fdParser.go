package zpl

func fdParser(params []string) []CommandParameter {
	var result []CommandParameter

	for _, coord := range params {
		result = append(result, CommandParameter{Value: coord})
	}
	return result
}
