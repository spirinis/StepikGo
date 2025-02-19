package main

func contains(string_slice []string, str string) bool {
	for _, val := range string_slice {
		if val == str {
			return true
		}
	}
	return false
}
