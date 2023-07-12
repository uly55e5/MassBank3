package search

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/MassBank/MassBank3/pkg/database"
	"github.com/MassBank/MassBank3/pkg/massbank"
	"gonum.org/v1/gonum/mat"
	"math"
	"reflect"
)

type cosineParameters struct {
	massTolerance float64
	theshold      float64
}

type cosineResult struct {
	result SearchResult
	status SearchStatus
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
			cs.parameters.theshold = th.(float64)
		}
	}
	return nil
}

func (cs CosineSearch) SetDatabase(mb3Database database.MB3Database) {

}

func (cs CosineSearch) Search(spectrum *massbank.MsSpectrum, filters database.Filters) (*SearchId, error) {
	spectra, err := cs.database.GetSpectra(filters)
	if err != nil {
		return nil, err
	}
	id := cs.newId()
	spectrum.Normalize(1000)
	spectrum.Baseline(cs.parameters.theshold)
	sp1 := spectrum.ToMatrix()
	cs.results[id] = cosineResult{
		result: SearchResult{
			Scores: map[string]any{},
			Parameters: map[string]any{
				"Threshold":      cs.parameters.theshold,
				"Mass Tolerance": cs.parameters.massTolerance,
			},
			SearchTime: 0,
			DBMetadata: massbank.MbMetaData{},
			Filters:    filters,
			Limit:      0,
			Offset:     0,
		},
		status: InProgress,
	}
	for acc, sp := range spectra {
		sp.Baseline(cs.parameters.theshold)
		sp2 := sp.ToMatrix()
		if sp1 != nil && sp2 != nil {
			sp2 = align(sp1, sp2, cs.parameters.massTolerance)
			if sp2 != nil {
				score := cosineSearch(sp2.ColView(1).(*mat.VecDense), sp2.ColView(3).(*mat.VecDense))
				cs.results[id].result.Scores[acc] = score
			}
		}
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
	j, _ := json.Marshal(&cs)
	var s = md5.Sum(j)
	return SearchId(hex.EncodeToString(s[:]))
}

func align(sp1 *mat.Dense, sp2 *mat.Dense, b float64) *mat.Dense {
	if sp1 == nil || sp2 == nil {
		return nil
	}
	r1, _ := sp1.Caps()
	r2, _ := sp2.Caps()
	data := []float64{}
	for i := 0; i < r2; i++ {
		for j := 0; j < r1; j++ {
			if sp1.At(j, 0) >= sp2.At(i, 0)-b && sp1.At(j, 0) <= sp2.At(i, 0)+b {
				data = append(data, append(sp1.RawRowView(j), sp2.RawRowView(i)...)...)
			}
		}
	}
	if len(data) < 4 {
		return nil
	}
	return mat.NewDense(len(data)/4, 4, data)
}

func cosineSearch(sp1 *mat.VecDense, sp2 *mat.VecDense) float64 {
	result := mat.Dot(sp1, sp2) / (math.Sqrt(mat.Sum(square(sp1))) * math.Sqrt(mat.Sum(square(sp2))))
	return result
}

func square(v *mat.VecDense) *mat.VecDense {
	result := *v
	for i := 0; i < v.Len(); i++ {
		result.SetVec(i, math.Pow(v.AtVec(i), 2))
	}
	return &result
}
