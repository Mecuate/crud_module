package crud_module

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	CREATE ReqVerb = iota
	READ
	UPDATE
	DELETE
	GET
	POST
	PUT
	PATCH
	CONNECT
	TRACE
	CRUD
)

var RegistredMethods = []string{
	"CREATE",
	"READ",
	"UPDATE",
	"DELETE",
	"GET",
	"POST",
	"PUT",
	"PATCH",
	"CONNECT",
	"TRACE",
}

func CreateMultiHandlerCRUD(r MuxRouter, rawPath string, handlers IndividualCRUDHandlers) {
	path := VetPath(rawPath)
	exclusions := []string{}
	for item, itemFunc := range handlers {
		methodSelection, err := FindMethod(item)
		if err != nil {
			panic(err)
		}
		method := methodSelection[0]
		r.Router.HandleFunc(path, itemFunc).Methods(method)
		exclusions = append(exclusions, method)
	}

	LockAllOtherMethods(r, path, exclusions)
	fmt.Println("Multi handlerCRUD: Created.", strings.Join(exclusions, " |"))
}

func CreateSingleHandlerCRUD(r MuxRouter, rawPath string, handler HandleFunc) {
	methodSelection, err := FindMethod(CRUD)
	if err != nil {
		panic(err)
	}

	path := VetPath(rawPath)

	for i := 0; i < len(methodSelection); i++ {
		r.Router.HandleFunc(path, handler).Methods(methodSelection[i])
	}

	LockAllOtherMethods(r, path, methodSelection)
	fmt.Println("CRUD: Created for ", strings.Join(methodSelection, ","))
}

func LockAllOtherMethods(r MuxRouter, path string, excluded []string) {
	methods := RemoveItemsFromArray(RegistredMethods, excluded)
	for i := 0; i < len(methods); i++ {
		r.Router.HandleFunc(path, DefaultLockedMethod).Methods(methods[i])
	}
}

func DefaultLockedMethod(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "405 Method Not Allowed", http.StatusForbidden)
}
