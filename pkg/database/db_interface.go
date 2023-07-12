package database

import (
	"github.com/MassBank/MassBank3/pkg/massbank"
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
	CompoundName      string // regex
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
	MongoDB  DatabaseType = 0 // not supported
	Postgres              = 1
)

// DBConfig is the abstract database configuration which should be used when working
// with [MB3Database].
type DBConfig struct {
	Database  DatabaseType // only Postgres is implemented
	DbUser    string       // the database user
	DbPwd     string       // the password for the database user
	DbHost    string       // the database host
	DbName    string       // the database name
	DbPort    uint
	DbConnStr string
}

// DefaultValues are the default values for the mandatory filter and database parameters
var DefaultValues = struct {
	MassEpsilon     float64
	IntensityCutoff int64
	Limit           int64
	Offset          int64
}{0.3, 100, math.MaxInt64, 0}

// MBcountValues is used for value-count pairs
type MBCountValues struct {
	Val   string
	Count int
}

// MBMinMaxValues is used for minimum-maximum pairs
type MBMinMaxValues struct {
	Min float64
	Max float64
}

// MB3Values are the possible filter values and ranges
type MB3Values struct {
	Contributor    []MBCountValues
	InstrumentType []MBCountValues
	MSType         []MBCountValues
	IonMode        []MBCountValues
	Intensity      MBMinMaxValues
	Mass           MBMinMaxValues
	Peak           MBMinMaxValues
}

// MB3StoredMetaData is static metadata stored in the database
type MB3StoredMetaData struct {
	Version   string
	TimeStamp time.Time
	GitCommit string
}

// MB3MetaData is metadata static and dynanmic metadata of the database
type MB3MetaData struct {
	StoredMetadata MB3StoredMetaData
	SpectraCount   int
	CompoundCount  int
	IsomerCount    int
}

// SearchResult is a result of a search in massbank
type SearchResult struct {
	SpectraCount int
	ResultCount  int
	Data         []SearchResultData
}

// SpectrumMetaData are the metadata defining a spectrum
type SpectrumMetaData struct {
	Id    string
	Title string
}

// SearchResultData is a list of compounds and the asssociated spectrum metadata
type SearchResultData struct {
	Inchi   string
	Names   []string
	Formula string
	Mass    float64
	Smiles  string
	Spectra []SpectrumMetaData
}

// a list of spectra ass search Result
type SpectraList map[string]*massbank.MsSpectrum

// MB3Database This is the Interface which has to be implemented for databases using MassBank3
//
// Any database can be used as in the backend as long as it defines the interface.
type MB3Database interface {

	// Connect to the database.
	Connect() error

	// Disconnect from the database.
	Disconnect() error

	// Ping to check if the database is connected. Returns nil on success.
	Ping() error

	// Count MassBank records in the database.
	Count() (int64, error)

	// IsEmpty returns true if the database is empty or not initialized
	IsEmpty() (bool, error)

	// DropAllRecords drops all MassBank records in the Database.
	DropAllRecords() error

	// GetMetaData returns the database metadata
	GetMetaData() (*MB3MetaData, error)

	// GetRecord gets a single MassBank record by the Accession string.
	// It should return nil and no error if the record is not in the
	// database.
	GetRecord(*string) (*massbank.MassBank2, error)

	// GetRecords Get an array of MassBank records by filtering
	//
	// Will return an empty list if the filter does not match any records.
	GetRecords(filters Filters) (*SearchResult, error)

	// GetSpectra returns a list of spectra matching the filters
	GetSpectra(filters Filters) (SpectraList, error)

	// GetSmiles returns the smiles for a valid accession string
	GetSmiles(accession *string) (*string, error)

	// GetUniqueValues is used to get the values for filter frontend
	GetUniqueValues(filters Filters) (MB3Values, error)

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
}

var db MB3Database

// InitDB initializes the database and tests the connection.
//
// It will panic if the config is not valid or no connection can be established.
func InitDb(dbConfig DBConfig) (MB3Database, error) {
	var err error
	if db == nil {
		// There is only one database type (Postgres)
		db, err = NewPostgresSQLDb(dbConfig)
		if err != nil {
			// panic if the config information is not valid
			panic(err)
		}
		if err = db.Connect(); err != nil {
			// panic if the database connection is not ready
			// TODO implement a failsafe way to react if the database is not responding (i.e. retry)
			panic(err)
		}
	}
	err = db.Ping()
	if err != nil {
		// reset the database if the ping is not successful
		db = nil
	}
	return db, err
}
