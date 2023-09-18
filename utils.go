package crud_module

import (
	"net/http"
)

func Contains(name string) bool {
	for _, value := range restricted_words {
		if value == name {
			return true
		}
	}
	return false
}

func FindMethod(v ReqVerb) ([]string, error) {
	switch v {
	case GET:
		return []string{http.MethodGet}, nil
	case READ:
		return []string{http.MethodGet}, nil
	case CREATE:
		return []string{http.MethodPost}, nil
	case POST:
		return []string{http.MethodPost}, nil
	case UPDATE:
		return []string{http.MethodPatch}, nil
	case PUT:
		return []string{http.MethodPut}, nil
	case DELETE:
		return []string{http.MethodDelete}, nil
	case CRUD:
		return []string{http.MethodPost, http.MethodGet, http.MethodPatch, http.MethodDelete}, nil
	}
	return []string{"UNKNOWN"}, http.ErrNotSupported
}

func RemoveItemsFromArray(arr []string, itemsToRemove []string) []string {
	result := []string{}
	itemsToRemoveMap := make(map[string]bool)
	for _, item := range itemsToRemove {
		itemsToRemoveMap[item] = true
	}
	for _, item := range arr {
		if !itemsToRemoveMap[item] {
			result = append(result, item)
		}
	}

	return result
}
