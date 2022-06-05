package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/GolangStudiy/go-users-postgres-rest-api/src/infrastructure"
	"github.com/GolangStudiy/go-users-postgres-rest-api/src/server"
	"github.com/GolangStudiy/go-users-postgres-rest-api/tests"
	"github.com/google/uuid"
)

var connection *sql.DB

func beforeTests() {
	db := tests.MountDatabaseContainer()
	dbPort := db.GetPort()

	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", strconv.Itoa(dbPort))
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "root")
	os.Setenv("DB_NAME", "users")

	var err error
	connection, err = infrastructure.GetDbConnection()

	if err != nil {
		log.Fatal(err)
	}
}

func TestShouldBeSaveUserAndGetIdAfter(t *testing.T) {
	var userJson = []byte(`{"email":"john@doe.com"}`)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(userJson))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)
	b, _ := json.Marshal(m["Email"])
	email := string(b)

	if !strings.Contains(email, "john@doe.com") {
		t.Errorf("Expected response code %s. Got %s", "john@doe.com", email)
	}

	req, _ = http.NewRequest("GET", "/users/get-id-by-email/john@doe.com", nil)
	req.Header.Set("Content-Type", "application/json")
	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var n string
	json.Unmarshal(response.Body.Bytes(), &n)
	b, _ = json.Marshal(n)
	id := string(b)

	uuid, err := uuid.Parse(id)

	if err != nil {
		t.Errorf(err.Error())
	} else if uuid.String() == "" {
		t.Errorf("id cannot be null")
	}
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	server.Router().ServeHTTP(rr, req)

	return rr
}

func TestMain(m *testing.M) {
	beforeTests()
	a := server.App{}
	a.Initialize()
	code := m.Run()
	os.Exit(code)
}
