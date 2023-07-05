package test

import (
	mb3server "github.com/MassBank/MassBank3/cmd/mb3server/src"
	"github.com/MassBank/MassBank3/pkg/common"
	"github.com/MassBank/MassBank3/pkg/config"
	"github.com/MassBank/MassBank3/pkg/database"
	"github.com/go-chi/chi/v5"
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
)

func initMongoDB() chi.Router {
	apiService := mb3server.NewDefaultApiService()
	apiCtrl := mb3server.NewDefaultApiController(apiService)
	router := mb3server.NewRouter(apiCtrl)
	mb3server.ServerConfig = &config.ServerConfig{
		DBConfig: database.DBConfig{
			Database:  database.MongoDB,
			DbUser:    "mbtestuser",
			DbPwd:     "mbtestpwd",
			DbHost:    "testmongo",
			DbName:    "mbtestdb",
			DbPort:    27017,
			DbConnStr: "",
		},
		ServerPort:   0,
		CdkDepictUrl: "http://cdkdepict:8080",
	}
	var testDataDir = common.GetEnv("TEST_DATA_DIR", "/go/src/")
	var files = map[string]string{
		"mb_metadata": testDataDir + "test-data/mb_metadata.json",
		"massbank":    testDataDir + "test-data/massbank-all.json",
	}

	database.InitMongoDB(mb3server.ServerConfig.DBConfig, files)
	return router
}

func TestGetMessage(t *testing.T) {
	router := initMongoDB()
	apitest.New(). // configuration
			Debug().
			Handler(router).
			Get("/v1/records"). // request
			Expect(t).          // expectations
			Body(responses["allRecords"]).
			Status(http.StatusOK).
			End()
}
