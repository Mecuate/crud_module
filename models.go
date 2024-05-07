package crud_module

import (
	"net/http"

	"github.com/gorilla/mux"
)

type versionType = string

type LoggLevel int64

type ReqVerb int64

type MuxRouter struct {
	Router *mux.Router
}

type HandleFunc func(w http.ResponseWriter, r *http.Request)

type IndividualCRUDHandlers map[ReqVerb]HandleFunc
