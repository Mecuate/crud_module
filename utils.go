package crud_module

import (
	"net/http"
)

func FindMethod(v ReqVerb) ([]string, error) {
	switch v {
	case GET:
		return []string{http.MethodGet}, nil
	case READ:
		return []string{"READ"}, nil
	case CREATE:
		return []string{"CREATE"}, nil
	case POST:
		return []string{http.MethodPost}, nil
	case UPDATE:
		return []string{"UPDATE"}, nil
	case PUT:
		return []string{http.MethodPut}, nil
	case DELETE:
		return []string{http.MethodDelete}, nil
	case CONNECT:
		return []string{http.MethodConnect}, nil
	case PATCH:
		return []string{http.MethodPatch}, nil
	case TRACE:
		return []string{http.MethodTrace}, nil
	case CRUD:
		return []string{"CREATE", "READ", "UPDATE", http.MethodDelete}, nil
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
