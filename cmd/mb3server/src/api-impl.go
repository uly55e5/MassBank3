package mb3server

import (
	"github.com/MassBank/MassBank3/pkg/database"
	"log"
	"os"
)

var db database.MB3Database = nil

func initDB() error {
	if db == nil {
		var mongo_uri = os.Getenv("MONGO_URI")
		var mongo_name = os.Getenv("MONGO_DB_NAME")
		log.Println("MongoDB URI: ", mongo_uri)
		log.Println("Database_Name", mongo_name)
		var err error = nil
		var config = database.DBConfig{
			Database:  database.MongoDB,
			DbUser:    "",
			DbPwd:     "",
			DbHost:    "",
			DbName:    os.Getenv("MONGO_DB_NAME"),
			DbPort:    0,
			DbConnStr: os.Getenv("MONGO_URI"),
		}

		db, err = database.NewMongoDB(config)
		if err != nil {
			return err
		}
		err = db.Connect()
		if err != nil {
			return err
		}
	}
	return db.Ping()

}

func GetBrowseOptions() (*BrowseOptions, error) {
	if err := initDB(); err != nil {
		return nil, err
	}
	vals, err := db.GetUniqueValues(database.Filters{})
	if err != nil {
		return nil, err
	}
	var result = BrowseOptions{}
	result.Metadata = Metadata{
		Version:       "",
		Timestamp:     "",
		GitCommit:     "",
		SpectraCount:  0,
		CompoundCount: 0,
		IsomerCount:   0,
		ResultCount:   0,
		Page:          0,
		Limit:         0,
	}
	for _, val := range vals.IonMode {
		result.IonMode = append(result.IonMode, StringCountInner{
			Value: val.Val,
			Count: int32(val.Count),
		})
	}
	for _, val := range vals.MSType {
		result.MsType = append(result.MsType, StringCountInner{
			Value: val.Val,
			Count: int32(val.Count),
		})
	}
	for _, val := range vals.InstrumentType {
		result.InstrumentType = append(result.InstrumentType, StringCountInner{
			Value: val.Val,
			Count: int32(val.Count),
		})
	}

	return &result, nil
}
