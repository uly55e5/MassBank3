package database

var TestDbConfigs = map[string]DBConfig{
	"pg valid": {
		Database:  Postgres,
		DbUser:    "mbtestuser",
		DbPwd:     "mbtestpwd",
		DbHost:    PostgresTestHost,
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
		DbConnStr: "host=" + PostgresTestHost + " port=5432 user=mbtestuser password=mbtestpwd dbname=mbtestdb sslmode=disable",
	},
	"pg empty": {
		Database: Postgres,
	},
	"pg only host": {
		Database: Postgres,
		DbHost:   "wronghost",
	},
}

var TestDatabases = map[string]MB3Database{
	"pg valid": &PostgresSQLDB{
		user:       "mbtestuser",
		dbname:     "mbtestdb",
		password:   "mbtestpwd",
		host:       PostgresTestHost,
		port:       5432,
		connString: "host=" + PostgresTestHost + " port=5432 user=mbtestuser password=mbtestpwd dbname=mbtestdb sslmode=disable",
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
		connString: "host=" + PostgresTestHost + " port=5432 user=mbtestuser password=mbtestpwd dbname=mbtestdb sslmode=disable",
		database:   nil,
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
