package entrypoint

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	domain "github.com/GolangStudiy/go-users-postgres-rest-api/src/domain/user"
	services "github.com/GolangStudiy/go-users-postgres-rest-api/src/services/user"
	"github.com/gorilla/mux"
)

func PostUser(w http.ResponseWriter, request *http.Request) {
	reqBody, _ := ioutil.ReadAll(request.Body)
	var user RequestUser
	json.Unmarshal(reqBody, &user)

	responseUser, err := services.Post(domain.User{Email: user.Email})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		respondWithJSON(w, http.StatusCreated, ResponseUser{Email: responseUser.Email})
	}
}

func GetUserIdByEmail(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	email := vars["email"]

	id, err := services.GetIdByEmail(email)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		respondWithJSON(w, http.StatusOK, id)
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
