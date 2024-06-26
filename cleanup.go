package crud_module

import "strings"

var restricted_sequences = []string{
	"root", "admin", "//", "\\", "..", "$", "@", "?", "+", "'", "\"",
}

func IsWordRestricted(path string) bool {
	originalString := path

	for _, value := range restricted_sequences {
		if strings.Contains(originalString, value) {
			return true
		}
	}
	return false
}

func VetPath(rawPath string) string {
	if IsWordRestricted(rawPath) {
		panic("Path has characters that are not allowed." + strings.Join(restricted_sequences, " | "))
	} else {
		return rawPath
	}
}
