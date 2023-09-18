package crud_module

var restricted_words = []string{
	"api", "apis", "root", "admin", "support", "news", "v1", "v2", "v3", "mct", "mecuate",
}

func IsWordRestricted(element string) bool {
	return Contains(element)
}
