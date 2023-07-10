package test

import (
	mb3server "github.com/MassBank/MassBank3/cmd/mb3server/src"
	"github.com/MassBank/MassBank3/pkg/config"
	"github.com/MassBank/MassBank3/pkg/database"
	"github.com/go-chi/chi/v5"
)

func initMongoDB() chi.Router {
	apiService := mb3server.NewDefaultApiService()
	apiCtrl := mb3server.NewDefaultApiController(apiService)
	router := mb3server.NewRouter(apiCtrl)
	mb3server.ServerConfig = &config.ServerConfig{
		DBConfig: database.DBConfig{
			Database:  database.Postgres,
			DbUser:    "mbtestuser",
			DbPwd:     "mbtestpwd",
			DbHost:    database.PostgresTestHost,
			DbName:    "mbtestdb",
			DbPort:    5432,
			DbConnStr: "",
		},
		ServerPort:   0,
		CdkDepictUrl: "http://cdkdepict:8080",
	}
	_, err := database.InitPostgresTestDB(database.Test_DS_All)
	if err != nil {
		println(err.Error())
	}
	return router
}

/*func TestGetMessage(t *testing.T) {
	router := initMongoDB()
	apitest.New(). // configuration
			Debug().
			Handler(router).
			Get("/v1/records"). // request
			Expect(t).
			Assert(func(response *http.Response, request *http.Request) error {

			var resBody = make([]byte, len(responses["allRecords"])+1)
			response.Body.Read(resBody)
			if string(resBody) == responses["allRecords"] {
				return nil
			}
			println("expected: ", responses["allRecords"])
			println("got     : ", string(resBody))
			return errors.New("Body not as expected")
		}).
		Status(http.StatusOK).
		End()
}*/
