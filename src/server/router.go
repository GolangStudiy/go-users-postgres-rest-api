package server

import (
	entrypoint "github.com/GolangStudiy/go-users-postgres-rest-api/src/entrypoint/user"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/users", entrypoint.PostUser).Methods("POST")
	r.HandleFunc("/users/get-id-by-email/{email}", entrypoint.GetUserIdByEmail).Methods("GET")

	return r
}
