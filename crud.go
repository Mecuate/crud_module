package crud_module

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	GET ReqVerb = iota
	READ
	CREATE
	POST
	UPDATE
	PUT
	DELETE
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
	"PUT",
	"PATCH",
	"CONNECT",
}

func CreateSingleHandlerCRUD(r MuxRouter, path string, handler HandleFunc, method ReqVerb) {
	methodSelection, err := FindMethod(method)
	if err != nil {
		panic(err)
	}
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
