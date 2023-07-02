package database

import (
	"github.com/Code-Hex/dd"
	"github.com/MassBank/MassBank3/pkg/massbank"
	"log"
	"math"
	"time"
)

// Filters is the abstract description of filters used to find MassBank records
// in the database
type Filters struct {
	InstrumentType    *[]string
	Splash            string
	MsType            *[]massbank.MsType
	IonMode           massbank.IonMode
	CompoundName      string //regex
	Mass              *float64
	MassEpsilon       *float64
	Formula           string // regex
	Peaks             *[]float64
	PeakDifferences   *[]float64
	InchiKey          string
	Contributor       *[]string
	IntensityCutoff   *int64
	Limit             int64
	Offset            int64
	IncludeDeprecated bool
}

// DatabaseType is an enum containing the database type
type DatabaseType int

// The list of supported databases
const (
	MongoDB  DatabaseType = 0
	Postgres              = 1
)

// DBConfig is the abstract database configuration which should be used when working
// with [MB3Database].
type DBConfig struct {
	Database  DatabaseType
	DbUser    string
	DbPwd     string
	DbHost    string
	DbName    string
	DbPort    uint
	DbConnStr string
}

var DefaultValues = struct {
	MassEpsilon     float64
	IntensityCutoff int64
	Limit           int64
	Offset          int64
}{0.3, 100, math.MaxInt64, 0}

// MBErrorType is an enum for the error types during database operations
type MBErrorType int

// The list of error types
const (
	DatabaseNotReady MBErrorType = iota
	CouldNotReachHost
	InternalError
	NotFoundError
	ConversionError
)

// MBDatabaseError is an error type specific for the database interactions
type MBDatabaseError struct {
	InnerError error  // inner error mostly from the database backend
	Message    string // the error message
}

// Implements the error interface for [MBDatabaseError]
func (err *MBDatabaseError) Error() string {
	var msg = err.Message
	if err != nil && len(err.Message) > 0 && err.InnerError != nil {
		msg += ": "
	}
	if err.InnerError != nil {
		return msg + err.Error()
	}
	return msg
}

type MBCountValues struct {
	Val   string
	Count int
}
type MBMinMaxValues struct {
	Min float64
	Max float64
}

type MB3Values struct {
	Contributor    []MBCountValues
	InstrumentType []MBCountValues
	MSType         []MBCountValues
	IonMode        []MBCountValues
	Intensity      MBMinMaxValues
	Mass           MBMinMaxValues
	Peak           MBMinMaxValues
}

type MB3StoredMetaData struct {
	Version   string
	TimeStamp time.Time
	GitCommit string
}

type MB3MetaData struct {
	StoredMetadata MB3StoredMetaData
	SpectraCount   int
	CompoundCount  int
	IsomerCount    int
}

type SearchResult struct {
	SpectraCount int
	ResultCount  int
	Data         map[string]SearchResultData
}

type SpectrumMetaData struct {
	Id    string
	Title string
}

type SearchResultData struct {
	Names   []string
	Formula string
	Mass    float64
	Smiles  string
	Spectra []SpectrumMetaData
}

// MB3Database This is the Interface which has to be implemented for databases using MassBank3
//
// Any database can be used as in the backend as long as it defines the interface.
type MB3Database interface {

	// Connect to the database.
	Connect() error

	// Disconnect from the database.
	Disconnect() error

	Ping() error

	// Count MassBank records in the database.
	Count() (int64, error)

	// DropAllRecords drops all MassBank records in the Database.
	DropAllRecords() error

	// GetRecord gets a single MassBank record by the Accession string.
	// It should return nil and a [NotFoundError] if the record is not in the
	// database.
	GetRecord(*string) (*massbank.MassBank2, error)

	// GetRecords Get an array of MassBank records by filtering
	//
	// Will return an empty list if the filter does not match any records.
	GetRecords(filters Filters) (*SearchResult, error)

	// GetUniqueValues is used to get the values for filter frontend
	GetUniqueValues(filters Filters) (MB3Values, error)

	GetMetaData() (*MB3MetaData, error)

	// UpdateMetadata updates the metadata describing the MassBank version.
	// Provides the database id of an existing entry if it is already in the
	// database.
	//
	// Returns the id of the database entry as string.
	UpdateMetadata(meta *massbank.MbMetaData) (string, error)

	// AddRecord adds a new MassBank record to the database. If the Accession
	// id already exists it will return an error.
	//
	// The second parameter is the database id of the version information. You
	// can get it from [UpdateMetadata].
	AddRecord(record *massbank.MassBank2, metaDataId string) error

	// AddRecords adds a list of MassBank records given as an array to the
	// database. If one of the Accession ids  exists the  function should roll
	// back the transaction and return an error.
	//
	// The second parameter is the database id of the version information. You
	// can get it from [UpdateMetadata].
	AddRecords(records []*massbank.MassBank2, metaDataId string) error

	// UpdateRecord will replace an existing MassBank record. Depending on the
	// upsert parameter it also inserts the record if it not exists.
	//
	// The second parameter is the database id of the version information. You
	// can get it from [UpdateMetadata].
	//
	// This should return number of  modified and inserted records, but this is
	// not implemented for all databases.
	UpdateRecord(record *massbank.MassBank2, metaDataId string, upsert bool) (uint64, uint64, error)

	// UpdateRecords will replace existing MassBank record. Depending on the
	// upsert parameter it also inserts the record if it not exists. This should
	// roll back the whole transaction if the there is an error.
	//
	// The second parameter is the database id of the version information. You
	// can get it from [UpdateMetadata].
	//
	// This should return number of  modified and inserted records, but this is
	// not implemented for all databases.
	UpdateRecords(records []*massbank.MassBank2, metaDataId string, upsert bool) (uint64, uint64, error)

	GetSmiles(accession *string) (*string, error)

	GetSpectra(filters Filters) (map[string]massbank.MsSpectrum, error)
}

var db MB3Database

func InitDb(dbConfig DBConfig) (MB3Database, error) {
	if db == nil {
		var err error
		if dbConfig.Database == MongoDB {
			db, err = NewMongoDB(dbConfig)
			if err != nil {
				panic(err)
			}
		} else if dbConfig.Database == Postgres {
			db, err = NewPostgresSQLDb(dbConfig)
			log.Println(dd.Dump(db))
			if err != nil {
				panic(err)
			}
		}
		if err = db.Connect(); err != nil {
			panic(err)
		}
	}
	err := db.Ping()
	if err != nil {
		db = nil
	}
	return db, err
}
