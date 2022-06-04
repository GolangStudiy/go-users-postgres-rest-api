package server

/// Go fmt import
import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/GolangStudiy/go-users-postgres-rest-api/db"
	"github.com/GolangStudiy/go-users-postgres-rest-api/src/configurations"
	"github.com/gorilla/mux"
)

func (a *App) Initialize() {
	var err error
	a.DatabaseConnection, err = configurations.GetDbConnection()

	if err != nil {
		log.Fatal(err)
	} else if a.DatabaseConnection == nil {
		log.Fatal("Database Connection failed!")
	}

	db.Migrate()

	a.Router = mux.NewRouter()
}

type App struct {
	Router             *mux.Router
	DatabaseConnection *sql.DB
}

func (a *App) Run() {
	apiProtocol := os.Getenv("API_PROTOCOL")
	apiHost := os.Getenv("API_HOST")
	apiPort := os.Getenv("API_PORT")

	a.Router = Router()

	apiAdress := apiProtocol + "://" + apiHost + ":" + apiPort
	if apiProtocol == "https" {
		apiCertFile := os.Getenv("API_CERT_FILE")
		apiKeyFile := os.Getenv("API_KEY_FILE")
		log.Fatal(http.ListenAndServeTLS(apiAdress, apiCertFile, apiKeyFile, a.Router))
	} else {
		log.Fatal(http.ListenAndServe(apiAdress, a.Router))
	}
}

func Main() {
	a := App{}
	a.Initialize()
	a.Run()
}
