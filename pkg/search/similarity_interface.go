package search

import (
	"github.com/MassBank/MassBank3/pkg/database"
	"github.com/MassBank/MassBank3/pkg/massbank"
	"time"
)

type SearchId string

type SearchResult struct {
	Scores     map[string]any
	Parameters map[string]any
	SearchTime time.Duration
	DBMetadata massbank.MbMetaData
	Filters    database.Filters
	Limit      int
	Offset     int
}

type SearchStatus int

const (
	InProgress SearchStatus = iota
	Paused
	Error
	Finished
	JobUnknown
)

type SimilaritySearch interface {
	SetParameters(par map[string]any) error
	SetDatabase(db database.MB3Database)
	Search(spectrum massbank.MsSpectrum, filters database.Filters) (SearchId, error)
	GetStatus(searchId SearchId) (SearchStatus, string)
	GetResult(searchId SearchId, Limit int, Offset int) (SearchResult, error)
}
