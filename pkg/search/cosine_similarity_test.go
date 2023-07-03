package search

import (
	"github.com/MassBank/MassBank3/pkg/common"
	"github.com/MassBank/MassBank3/pkg/database"
	"github.com/MassBank/MassBank3/pkg/massbank"
	"github.com/MassBank/MassBank3/pkg/mocks"
	"github.com/golang/mock/gomock"
	"gonum.org/v1/gonum/mat"
	"reflect"
	"testing"
)

func TestCosineSearch_GetResult(t *testing.T) {
	type fields struct {
		database   database.MB3Database
		parameters cosineParameters
		results    map[SearchId]cosineResult
	}
	type args struct {
		id SearchId
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SearchResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := CosineSearch{
				database:   tt.fields.database,
				parameters: tt.fields.parameters,
				results:    tt.fields.results,
			}
			got, err := cs.GetResult(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetResult() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCosineSearch_GetStatus(t *testing.T) {
	type fields struct {
		database   database.MB3Database
		parameters cosineParameters
		results    map[SearchId]cosineResult
	}
	type args struct {
		searchId SearchId
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   SearchStatus
		want1  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := CosineSearch{
				database:   tt.fields.database,
				parameters: tt.fields.parameters,
				results:    tt.fields.results,
			}
			got, got1 := cs.GetStatus(tt.args.searchId)
			if got != tt.want {
				t.Errorf("GetStatus() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetStatus() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCosineSearch_Search(t *testing.T) {
	type args struct {
		spectrum massbank.MsSpectrum
		filters  database.Filters
	}
	mockCtrl := gomock.NewController(t)
	mock_db := mock_database.NewMockMB3Database(mockCtrl)
	mock_db.EXPECT().GetSpectra(database.Filters{}).Return(getTestSpectra(), nil)
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Simple search",
			args: args{
				spectrum: getTestSpectra()["MSBNK-AAFC-AC000005"],
				filters:  database.Filters{},
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := NewSearchCosine(mock_db)
			got, err := cs.Search(tt.args.spectrum, tt.args.filters)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want && got != nil && len(string(*got)) != 32 {
				t.Errorf("Search() got = %v, wanted MD5 Hash", *got)
			}
			if !tt.want && got != nil {
				t.Errorf("Search() got = %v, wanted nil", *got)
			}
		})
	}
}

func getTestSpectra() map[string]massbank.MsSpectrum {
	result := map[string]massbank.MsSpectrum{}
	for k, v := range common.MbTestRecords {
		sp, err := massbank.NewMsSpectrum(v.Peak.Peak.Mz, v.Peak.Peak.Intensity)
		if err != nil {

		} else {
			result[k] = *sp
		}
	}
	return result
}

func TestCosineSearch_SetDatabase(t *testing.T) {
	type fields struct {
		database   database.MB3Database
		parameters cosineParameters
		results    map[SearchId]cosineResult
	}
	type args struct {
		mb3Database database.MB3Database
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := CosineSearch{
				database:   tt.fields.database,
				parameters: tt.fields.parameters,
				results:    tt.fields.results,
			}
			cs.SetDatabase(tt.args.mb3Database)
		})
	}
}

func TestCosineSearch_SetParameters(t *testing.T) {
	type fields struct {
		database   database.MB3Database
		parameters cosineParameters
		results    map[SearchId]cosineResult
	}
	type args struct {
		par map[string]any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := CosineSearch{
				database:   tt.fields.database,
				parameters: tt.fields.parameters,
				results:    tt.fields.results,
			}
			if err := cs.SetParameters(tt.args.par); (err != nil) != tt.wantErr {
				t.Errorf("SetParameters() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCosineSearch_newId(t *testing.T) {
	type fields struct {
		database   database.MB3Database
		parameters cosineParameters
		results    map[SearchId]cosineResult
	}
	tests := []struct {
		name   string
		fields fields
		want   SearchId
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := CosineSearch{
				database:   tt.fields.database,
				parameters: tt.fields.parameters,
				results:    tt.fields.results,
			}
			if got := cs.newId(); got != tt.want {
				t.Errorf("newId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSearchCosine(t *testing.T) {
	type args struct {
		db database.MB3Database
	}
	tests := []struct {
		name string
		args args
		want CosineSearch
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSearchCosine(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSearchCosine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_align(t *testing.T) {
	type args struct {
		sp1 *mat.Dense
		sp2 *mat.Dense
		b   float64
	}
	tests := []struct {
		name string
		args args
		want *mat.Dense
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := align(tt.args.sp1, tt.args.sp2, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("align() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cosineSearch(t *testing.T) {
	type args struct {
		sp1 *mat.VecDense
		sp2 *mat.VecDense
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cosineSearch(tt.args.sp1, tt.args.sp2); got != tt.want {
				t.Errorf("cosineSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_square(t *testing.T) {
	type args struct {
		v *mat.VecDense
	}
	tests := []struct {
		name string
		args args
		want *mat.VecDense
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := square(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("square() = %v, want %v", got, tt.want)
			}
		})
	}
}
