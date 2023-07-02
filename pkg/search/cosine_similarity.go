package search

import (
	"errors"
	"github.com/MassBank/MassBank3/pkg/database"
	"github.com/MassBank/MassBank3/pkg/massbank"
	"gonum.org/v1/gonum/mat"
	"reflect"
)

type cosineParameters struct {
	massTolerance float64
	theshold      int
}

type cosineResult struct {
	parameters cosineParameters
	result     SearchResult
	status     SearchStatus
}

type CosineSearch struct {
	database   database.MB3Database
	parameters cosineParameters
	results    map[SearchId]cosineResult
}

func NewSearchCosine(db database.MB3Database) CosineSearch {
	cs := CosineSearch{database: db,
		parameters: cosineParameters{
			massTolerance: 0.3,
			theshold:      100,
		},
		results: map[SearchId]cosineResult{},
	}
	return cs
}

func (cs CosineSearch) SetParameters(par map[string]any) error {
	if mt, ok := par["MassTolerance"]; ok {
		if reflect.TypeOf(mt) == reflect.TypeOf(cs.parameters.massTolerance) {
			cs.parameters.massTolerance = mt.(float64)
		}
	}
	if th, ok := par["Threshold"]; ok {
		if reflect.TypeOf(th) == reflect.TypeOf(cs.parameters.theshold) {
			cs.parameters.theshold = th.(int)
		}
	}
	return nil
}

func (cs CosineSearch) SetDatabase(mb3Database database.MB3Database) {

}

func (cs CosineSearch) Search(spectrum massbank.MsSpectrum, filters database.Filters) (*SearchId, error) {
	spectra, err := cs.database.GetSpectra(filters)
	if err != nil {
		return nil, err
	}
	id := cs.newId()
	sp1 := spectrum.ToMatrix()
	for acc, sp := range spectra {
		sp2 := sp.ToMatrix()
		sp1 = align(sp1, sp2)
		score := cosineSearch(sp1, sp2)
		cs.results[id].result.Scores[acc] = score
	}
	return &id, nil
}

func (cs CosineSearch) GetStatus(searchId SearchId) (SearchStatus, string) {
	return cs.results[searchId].status, ""
}

func (cs CosineSearch) GetResult(id SearchId) (*SearchResult, error) {
	if st, msg := cs.GetStatus(id); st != Finished {
		return nil, errors.New("Task not finished: " + msg)
	}
	result := cs.results[id]
	return &result.result, nil
}

func (cs CosineSearch) newId() SearchId {
	return ""
}

func align(sp1 *mat.Dense, sp2 *mat.Dense) *mat.Dense {
	return sp1
}

func cosineSearch(sp1 *mat.Dense, sp2 *mat.Dense) float64 {
	return 0
}
