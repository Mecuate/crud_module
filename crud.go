package crud_module

import (
	"github.com/gorilla/mux"
)

func NewGroup(base string, r *mux.Router) {
	r.PathPrefix(base)
}
