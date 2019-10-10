package unique

func Strings(strings []string) []string {

	var unique []string
	index := make(map[string]bool)

	for _, value := range strings {
		if _, ok := index[value]; !ok {
			index[value] = true
			unique = append(unique, value)
		}
	}

	return unique
}
