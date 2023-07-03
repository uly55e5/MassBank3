package database

import (
	"github.com/MassBank/MassBank3/pkg/common"
)

var mongoHost = common.GetEnv("MONGODB_HOST", "testmongo")
var postgresHost = common.GetEnv("POSTGRES_HOST", "testpostgres")
var testDataDir = common.GetEnv("TEST_DATA_DIR", "/go/src/")

var TestDbConfigs = map[string]DBConfig{
	"pg valid": {
		Database:  Postgres,
		DbUser:    "mbtestuser",
		DbPwd:     "mbtestpwd",
		DbHost:    postgresHost,
		DbName:    "mbtestdb",
		DbPort:    5432,
		DbConnStr: "",
	},
	"pg wrong host": {
		Database:  Postgres,
		DbUser:    "mbtestuser",
		DbPwd:     "mbtestpwd",
		DbHost:    "wronghost",
		DbName:    "mbtestdb",
		DbPort:    5432,
		DbConnStr: "",
	},
	"pg valid conn string": {
		Database:  Postgres,
		DbUser:    "",
		DbPwd:     "",
		DbHost:    "",
		DbName:    "",
		DbPort:    0,
		DbConnStr: "host=" + postgresHost + " port=5432 user=mbtestuser password=mbtestpwd dbname=mbtestdb sslmode=disable",
	},
	"pg empty": {
		Database: Postgres,
	},
	"pg only host": {
		Database: Postgres,
		DbHost:   "wronghost",
	},
	"mg valid": {
		Database:  MongoDB,
		DbUser:    "mbtestuser",
		DbPwd:     "mbtestpwd",
		DbHost:    mongoHost,
		DbName:    "mbtestdb",
		DbPort:    27017,
		DbConnStr: "",
	},
	"mg empty": {
		Database: MongoDB,
	},
	"mg wrong host": {
		Database:  MongoDB,
		DbUser:    "mbtestuser",
		DbPwd:     "mbtestpwd",
		DbHost:    "wronghost",
		DbName:    "mbtestdb",
		DbPort:    27017,
		DbConnStr: "",
	},
	"mg valid conn string": {
		Database:  MongoDB,
		DbConnStr: "mongodb://mbtestuser:mbtestpwd@" + mongoHost + ":27017",
		DbName:    "mbtestdb",
	},
	"mg conn string ": {
		Database:  MongoDB,
		DbConnStr: "mongodb://mbtestuser:mbtestpwd@" + mongoHost + ":27017",
	},
}

type DbInitSet int

const (
	All DbInitSet = iota
	Main
	Empty
)

var TestDatabases = map[string]MB3Database{
	"pg valid": &PostgresSQLDB{
		user:       "mbtestuser",
		dbname:     "mbtestdb",
		password:   "mbtestpwd",
		host:       postgresHost,
		port:       5432,
		connString: "host=" + postgresHost + " port=5432 user=mbtestuser password=mbtestpwd dbname=mbtestdb sslmode=disable",
		database:   nil,
	},
	"pg wrong host": &PostgresSQLDB{
		user:       "mbtestuser",
		dbname:     "mbtestdb",
		password:   "mbtestpwd",
		host:       "wronghost",
		port:       5432,
		connString: "host=wronghost port=5432 user=mbtestuser password=mbtestpwd dbname=mbtestdb sslmode=disable",
		database:   nil,
	},
	"pg valid conn string": &PostgresSQLDB{
		user:       "",
		dbname:     "",
		password:   "",
		host:       "",
		port:       0,
		connString: "host=" + postgresHost + " port=5432 user=mbtestuser password=mbtestpwd dbname=mbtestdb sslmode=disable",
		database:   nil,
	},
	"mg valid": &Mb3MongoDB{
		user:     "mbtestuser",
		pwd:      "mbtestpwd",
		host:     mongoHost,
		dbname:   "mbtestdb",
		port:     27017,
		database: nil,
		dirty:    true,
	},
	"mg wrong host": &Mb3MongoDB{
		user:     "mbtestuser",
		pwd:      "mbtestpwd",
		host:     "wronghost",
		dbname:   "mbtestdb",
		port:     27017,
		database: nil,
		dirty:    true,
	},
	"mg valid conn string": &Mb3MongoDB{
		connStr: "mongodb://mbtestuser:mbtestpwd@" + mongoHost + ":27017",
		dbname:  "mbtestdb",
	},
}

var UniqueValueTestData = map[string]MB3Values{
	"all": {
		[]MBCountValues{
			{"AAFC", 1},
			{"Athens_Univ", 1},
			{"Eawag", 1},
			{"Eawag_Additional_Specs", 1},
			{"Fac_Eng_Univ_Tokyo", 1},
			{"Keio_Univ", 1},
			{"MSSJ", 1},
			{"RIKEN", 2},
			{"test", 2},
			{"Washington_State_Univ", 1},
		},
		[]MBCountValues{
			{
				Val:   "EI-B",
				Count: 1,
			},
			{
				Val:   "ESI-QTOF",
				Count: 1,
			},
			{
				Val:   "LC-APCI-QTOF",
				Count: 1,
			},
			{
				Val:   "LC-ESI-IT",
				Count: 1,
			},
			{
				Val:   "LC-ESI-ITFT",
				Count: 3,
			},
			{
				Val:   "LC-ESI-QFT",
				Count: 1,
			},
			{
				Val:   "LC-ESI-QTOF",
				Count: 3,
			},
			{
				Val:   "MALDI-TOF",
				Count: 1,
			},
		},
		[]MBCountValues{
			{
				Val:   "MS",
				Count: 3,
			},
			{
				Val:   "MS2",
				Count: 8,
			},
			{
				Val:   "MS4",
				Count: 1,
			},
		},
		[]MBCountValues{
			{
				Val:   "NEGATIVE",
				Count: 4,
			},
			{
				Val:   "POSITIVE",
				Count: 8,
			},
		},
		MBMinMaxValues{
			Min: 0.016450,
			Max: 235010720.000000,
		},
		MBMinMaxValues{
			Min: 124.034670,
			Max: 1865.00337,
		},
		MBMinMaxValues{
			Min: 15.000000,
			Max: 2136.330000,
		},
	},
}
