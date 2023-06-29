package search

import (
	"github.com/MassBank/MassBank3/pkg/database"
	"github.com/MassBank/MassBank3/pkg/massbank"
	"time"
)

type SearchResult struct {
	Scores     map[string]any
	Parameters map[string]any
	SearchTime time.Duration
	DBMetadata massbank.MbMetaData
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
	SetParameters(map[string]any) error
	Search(filters database.Filters) (string, error)
	GetStatus(searchId string) (SearchStatus, string)
	GetResult(searchId string) (SearchResult, error)
}
