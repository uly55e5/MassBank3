package database

import (
	"github.com/MassBank/MassBank3/pkg/common"
	"os"
)

var PostgresTestHost = common.GetEnv("POSTGRES_HOST", "testpostgres")
var testDataDir = common.GetEnv("TEST_DATA_DIR", "/go/src/")

type DbInitSet int

const (
	Test_DS_All DbInitSet = iota
	Test_DS_Main
	Test_DS_Empty
)

func InitPostgresTestDB(set DbInitSet) (MB3Database, error) {
	var filenames = []string{"mb_metadata", "massbank"}
	var files = map[string]string{}
	files["mb_metadata"] = testDataDir + "test-data/metadata.sql"
	switch set {
	case Test_DS_All:
		files["massbank"] = testDataDir + "test-data/massbank-all.sql"
	case Test_DS_Main:
		files["massbank"] = testDataDir + "test-data/massbank.sql"
	case Test_DS_Empty:
	}
	db, err := NewPostgresSQLDb(DBConfig{
		Database:  Postgres,
		DbUser:    "mbtestuser",
		DbPwd:     "mbtestpwd",
		DbHost:    PostgresTestHost,
		DbName:    "mbtestdb",
		DbPort:    5432,
		DbConnStr: "",
	})
	if err != nil {
		return nil, err
	}
	err = db.Connect()
	if err != nil {
		return nil, err
	}
	if _, err = db.database.Exec("DELETE FROM massbank"); err != nil {
		return nil, err
	}
	if _, err = db.database.Exec("ALTER SEQUENCE massbank_id_seq RESTART WITH 1"); err != nil {
		return nil, err
	}
	if _, err = db.database.Exec("DELETE FROM metadata"); err != nil {
		return nil, err
	}
	if _, err = db.database.Exec("ALTER SEQUENCE metadata_id_seq RESTART WITH 1"); err != nil {
		return nil, err
	}
	for _, fn := range filenames {
		if f, ok := files[fn]; ok {

			buf, err := os.ReadFile(f)
			if err != nil {
				return nil, err
			}
			sqlStr := string(buf)
			if _, err = db.database.Exec(sqlStr); err != nil {
				return nil, err
			}
		}
	}
	return db, nil
}
