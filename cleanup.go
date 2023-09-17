package crud_module

var restricted_words = []string{
	"api", "apis", "root", "admin", "v1", "v2",
}

func isWordRestricted(element string) bool {
	return Contains(element)
}

func Contains(name string) bool {
	for _, value := range restricted_words {
		if value == name {
			return true
		}
	}
	return false
}
