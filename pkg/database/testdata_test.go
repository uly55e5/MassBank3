package database

var TestDbConfigs = map[string]DBConfig{
	"workingPostgres": {
		Database:  Postgres,
		DbUser:    "mbtestuser",
		DbPwd:     "mbtestpwd",
		DbHost:    "testpostgres",
		DbName:    "mbtestdb",
		DbPort:    5432,
		DbConnStr: "",
	},
	"wrongPostgres": {
		Database:  Postgres,
		DbUser:    "mbtestuser",
		DbPwd:     "mbtestpwd",
		DbHost:    "wronghost",
		DbName:    "mbtestdb",
		DbPort:    5432,
		DbConnStr: "",
	},
	"workingPostgresConnString": {
		Database:  Postgres,
		DbUser:    "",
		DbPwd:     "",
		DbHost:    "",
		DbName:    "",
		DbPort:    0,
		DbConnStr: "host=testpostgres port=5432 user=mbtestuser password=mbtestpwd dbname=mbtestdb sslmode=disable",
	},
	"emptyPostgres": {
		Database: Postgres,
	},
	"onlyHostPostgres": {
		Database: Postgres,
		DbHost:   "wronghost",
	},
	"workingMongo": {
		Database:  MongoDB,
		DbUser:    "mbtestuser",
		DbPwd:     "mbtestpwd",
		DbHost:    "testmongo",
		DbName:    "mbtestdb",
		DbPort:    27017,
		DbConnStr: "",
	},
}

var TestDbConfigPostgres = map[string]*PostgresSQLDB{
	"working": {
		user:       "mbtestuser",
		dbname:     "mbtestdb",
		password:   "mbtestpwd",
		host:       "testpostgres",
		port:       5432,
		connString: "host=testpostgres port=5432 user=mbtestuser password=mbtestpwd dbname=mbtestdb sslmode=disable",
		database:   nil,
	},
	"wrongHost": {
		user:       "mbtestuser",
		dbname:     "mbtestdb",
		password:   "mbtestpwd",
		host:       "wronghost",
		port:       5432,
		connString: "host=wronghost port=5432 user=mbtestuser password=mbtestpwd dbname=mbtestdb sslmode=disable",
		database:   nil,
	},
	"workingConnString": {
		user:       "",
		dbname:     "",
		password:   "",
		host:       "",
		port:       0,
		connString: "host=testpostgres port=5432 user=mbtestuser password=mbtestpwd dbname=mbtestdb sslmode=disable",
		database:   nil,
	},
}
