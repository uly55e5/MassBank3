package database

import (
	"encoding/json"
	"fmt"
	"github.com/MassBank/MassBank3/pkg/massbank"
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	code, err := run(m)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

var recordnames = []string{
	"MSBNK-AAFC-AC000005",
	"MSBNK-Athens_Univ-AU229201",
	"MSBNK-Eawag-EA018353",
	"MSBNK-Eawag_Additional_Specs-ET060401",
	"MSBNK-Fac_Eng_Univ_Tokyo-JP009132",
	"MSBNK-Keio_Univ-KO009105",
	"MSBNK-MSSJ-MSJ00284",
	"MSBNK-RIKEN-PR100978",
	"MSBNK-RIKEN-PR309089",
	"MSBNK-Washington_State_Univ-BML81902",
	"MSBNK-test-TST00001",
	"MSBNK-test-TST00002",
	"MSBNK-test-TST00003",
}

func testRecord(i uint64) *massbank.MassBank2 {
	var mb = mbTestRecords[recordnames[i]]
	return &mb
}
func testRecords(names []uint64) []*massbank.MassBank2 {
	var records = []*massbank.MassBank2{}
	for _, i := range names {
		var mb = mbTestRecords[recordnames[i]]
		records = append(records, &mb)
	}
	return records
}

func testSearchResults(names []uint64, specCount int, resultCount int) SearchResult {
	var searchResult = SearchResult{}
	searchResult.SpectraCount = specCount
	var data = map[string]SearchResultData{}
	for _, i := range names {
		record := mbTestRecords[recordnames[i]]
		dataset := data[*record.Compound.InChI]
		if dataset.Spectra == nil {
			dataset.Spectra = []SpectrumMetaData{}
		}
		dataset.Spectra = append(dataset.Spectra, SpectrumMetaData{*record.Accession, *record.RecordTitle})
		dataset.Smiles = *record.Compound.Smiles
		dataset.Formula = *record.Compound.Formula
		dataset.Mass = *record.Compound.Mass
		dataset.Names = append(dataset.Names, *record.Compound.Names...)
		data[*record.Compound.InChI] = dataset
	}
	searchResult.Data = data
	searchResult.ResultCount = resultCount
	//println("Want:")
	//println(dd.Dump(searchResult))
	return searchResult
}

func run(m *testing.M) (code int, err error) {
	return m.Run(), nil
}

type testDBinit struct {
	name   string
	initFn func(set DbInitSet) (MB3Database, error)
}

type testDB struct {
	name string
	db   MB3Database
}

func initDBs(set DbInitSet) ([]testDB, error) {
	var testDBs = []testDBinit{
		{"mongodb", initMongoTestDB},
		{"postgres", initPostgresTestDB},
	}
	var result = []testDB{}
	for _, d := range testDBs {
		db, err := d.initFn(set)
		if err != nil {
			return nil, err
		}
		result = append(result, testDB{d.name, db})
	}
	return result, nil
}

func (db testDB) checkCount(t *testing.T, expect int64) {
	if g, err := db.db.Count(); g != expect || err != nil {
		t.Errorf("%s: Count not does not match: expected %d, got %d: %v", db.name, expect, g, err)
	}
}

func TestMB3Database_Connect(t *testing.T) {
	tests := []struct {
		name    string
		db      MB3Database
		wantErr bool
	}{
		{"Postgres valid",
			TestDatabases["pg valid"],
			false,
		},
		{"Postgres with valid connection string",
			TestDatabases["pg valid conn string"],
			false,
		},
		{
			"Postgres with wrong host",
			TestDatabases["pg wrong host"],
			true,
		},
		{"MongoDB valid",
			TestDatabases["mg valid"],
			false,
		},
		{"MongoDB valid with connection string",
			TestDatabases["mg valid conn string"],
			false,
		},
		{"MongoDB with wrong host",
			TestDatabases["mg wrong host"],
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.db.Connect(); (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMB3Database_Disconnect(t *testing.T) {
	tests := []struct {
		name    string
		db      MB3Database
		wantErr bool
	}{
		{"Postgres valid",
			TestDatabases["pg valid"],
			false,
		},
		{"Postgres valid second time",
			TestDatabases["pg valid"],
			true,
		},
		{"Postgres with valid connection string",
			TestDatabases["pg valid conn string"],
			false,
		},
		{"Postgres with valid connection string second time",
			TestDatabases["pg valid conn string"],
			true,
		},
		{
			"Postgres with wrong host",
			TestDatabases["pg wrong host"],
			true,
		},
		{"MongoDB valid",
			TestDatabases["mg valid"],
			false,
		},
		{"MongoDB valid second time",
			TestDatabases["mg valid"],
			true,
		},
		{"MongoDB with wrong host",
			TestDatabases["mg wrong host"],
			true,
		},
		{"MongoDB with connection string",
			TestDatabases["mg valid conn string"],
			false,
		},
		{"MongoDB with connection string second time",
			TestDatabases["mg valid conn string"],
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.db.Disconnect(); (err != nil) != tt.wantErr {
				t.Errorf("Disconnect() error = %v, wantErr %v, db %v", err, tt.wantErr, tt.db)
			}
		})
	}
}

func TestMb3Database_Count(t *testing.T) {
	DBs, err := initDBs(Main)
	if err != nil {
		t.Fatal("Could not init Databases: ", err.Error())
	}
	for _, db := range DBs {
		if c, err := db.db.Count(); c != 8 || err != nil {
			t.Errorf("%s: Could not count expected 8 got %c: %v", db.name, c, err)
		}
		if c, err := db.db.Count(); c != 8 || err != nil {
			t.Errorf("%s: Could not count second time exptected 8 got %v: %v", db.name, c, err)
		}
	}
	DBs, err = initDBs(All)
	if err != nil {
		t.Fatal("Could not init Databases: ", err.Error())
	}
	for _, db := range DBs {
		if c, err := db.db.Count(); c != 13 || err != nil {
			t.Errorf("%s: Could not count expected 13 got %c: %v", db.name, c, err)
		}
		if c, err := db.db.Count(); c != 13 || err != nil {
			t.Errorf("%s: Could not count second time exptected 13 got %v: %v", db.name, c, err)
		}
	}
}

func TestMB3Database_DropAllRecords(t *testing.T) {
	DBs, err := initDBs(Main)
	if err != nil {
		t.Fatal("Could not init Databases: ", err.Error())
	}
	for _, db := range DBs {
		db.checkCount(t, 8)
		if err := db.db.DropAllRecords(); err != nil {
			t.Errorf("%s: Could not drop: %v", db.name, err)
		}
		db.checkCount(t, 0)
		if err := db.db.DropAllRecords(); err != nil {
			t.Errorf("%s: Could not drop on empty db: %v", db.name, err)
		}
		db.checkCount(t, 0)
	}
}

func TestMB3Database_GetRecord(t *testing.T) {
	type args struct {
		s string
	}
	DBs, err := initDBs(All)
	if err != nil {
		t.Fatal("Could not init Databases: ", err.Error())
	}
	for _, db := range DBs {

		type testData struct {
			db      testDB
			name    string
			args    args
			want    massbank.MassBank2
			wantErr bool
		}
		var tests = []testData{}
		for _, record := range recordnames {
			var test = testData{
				db,
				db.name + " " + record,
				args{record},
				mbTestRecords[record],
				false,
			}
			tests = append(tests, test)
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := tt.db.db.GetRecord(&tt.args.s)
				if (err != nil) != tt.wantErr {
					t.Errorf("%s: GetRecord() error = %v, wantErr %v", tt.db.name, err, tt.wantErr)
					return
				}
				got.Metadata.VersionRef = tt.want.Metadata.VersionRef
				bg, errg := json.Marshal(got)
				bw, errw := json.Marshal(tt.want)
				if string(bg) != string(bw) || errg != nil || errw != nil {
					t.Errorf("\nwant: %v \ngot : %v\n", string(bw), string(bg))
				}
			})
		}
	}
}

func ptr[T any](v T) *T {
	return &v
}

func TestMB3Database_GetSmiles(t *testing.T) {
	type args struct {
		s string
	}
	DBs, err := initDBs(All)
	if err != nil {
		t.Fatal("Could not init Databases: ", err.Error())
	}
	for _, db := range DBs {

		type testData struct {
			db      testDB
			name    string
			args    args
			want    *string
			wantErr bool
		}
		var tests = []testData{}
		for _, record := range recordnames {
			var test = testData{
				db,
				db.name + " " + record,
				args{record},
				mbTestRecords[record].Compound.Smiles,
				false,
			}
			tests = append(tests, test)
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := tt.db.db.GetSmiles(&tt.args.s)
				if (err != nil) != tt.wantErr {
					t.Errorf("%s: GetSmiles() error = %v, wantErr %v", tt.db.name, err, tt.wantErr)
					return
				}
				if tt.want != nil && got != nil && *got != *tt.want {
					t.Errorf("\nwant: %v \ngot : %v\n", *tt.want, *got)
				}
				if tt.want == nil && got != nil {
					t.Errorf("\nwant: %v \ngot : %v\n", tt.want, *got)
				}
				if tt.want != nil && got == nil {
					t.Errorf("\nwant: %v \ngot : %v\n", *tt.want, got)
				}
			})
		}
	}
}

func TestMB3Database_GetRecords(t *testing.T) {
	DBs, err := initDBs(All)
	if err != nil {
		t.Fatal("Could not init Databases: ", err.Error())
	}
	type args struct {
		f Filters
	}
	for _, db := range DBs {

		type testData struct {
			db      testDB
			name    string
			args    Filters
			want    SearchResult
			wantErr bool
		}
		var tests = []testData{
			{
				db,
				db.name + " " + "Get all records",
				Filters{},
				testSearchResults([]uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 12, 12),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with deprecated",
				Filters{IncludeDeprecated: true},
				testSearchResults([]uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 13, 13),
				false,
			},
			{
				db,
				db.name + " " + "Get first 3 records",
				Filters{Limit: 3, Offset: 0},
				testSearchResults([]uint64{1, 5, 11}, 12, 12),
				false,
			},
			{
				db,
				db.name + " " + "Get second page with 3 records",
				Filters{Limit: 3, Offset: 3},
				testSearchResults([]uint64{3, 8, 10}, 12, 12),
				false,
			},
			{
				db,
				db.name + " " + "Get all but first  3 records",
				Filters{Limit: 0, Offset: 3},
				testSearchResults([]uint64{0, 2, 3, 4, 6, 7, 8, 9, 10}, 12, 12),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with InstrumentType LC-ESI-ITFT",
				Filters{InstrumentType: &[]string{"LC-ESI-ITFT"}},
				testSearchResults([]uint64{0, 2, 10}, 3, 3),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with InstrumentType LC-ESI-ITFT OR LC-ESI-QFT",
				Filters{InstrumentType: &[]string{"LC-ESI-ITFT", "LC-ESI-QFT"}},
				testSearchResults([]uint64{0, 2, 3, 10}, 4, 4),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with MS type MS",
				Filters{MsType: &[]massbank.MsType{massbank.MS}},
				testSearchResults([]uint64{4, 9, 11}, 3, 3),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with MS type MS2",
				Filters{MsType: &[]massbank.MsType{massbank.MS2}},
				testSearchResults([]uint64{0, 1, 2, 3, 6, 7, 8, 10}, 8, 8),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with MS type MS and MS2",
				Filters{MsType: &[]massbank.MsType{massbank.MS, massbank.MS2}},
				testSearchResults([]uint64{0, 1, 2, 3, 4, 6, 7, 8, 9, 10, 11}, 11, 11),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with ion mode POSITIVE",
				Filters{IonMode: massbank.POSITIVE},
				testSearchResults([]uint64{0, 1, 3, 4, 5, 9, 10, 11}, 8, 8),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with ion mode ANY",
				Filters{IonMode: massbank.ANY},
				testSearchResults([]uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, 12, 12),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with SPLASH ",
				Filters{Splash: "splash10-0udi-0609400000-9fd50528da25d66adfc7"},
				testSearchResults([]uint64{7}, 1, 1),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with non existing SPLASH ",
				Filters{Splash: "splash10-0udi-0609400000-9fd50528a725d66adfc7"},
				testSearchResults([]uint64{}, 0, 0),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with Mass and default epsilon",
				Filters{Mass: ptr(float64(296.251000))},
				testSearchResults([]uint64{8}, 1, 1),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with Mass and default epsilon, no match",
				Filters{Mass: ptr(float64(296.141000))},
				testSearchResults([]uint64{}, 0, 0),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with Mass and custom epsilon",
				Filters{Mass: ptr(float64(296.141000)), MassEpsilon: ptr(float64(0.5))},
				testSearchResults([]uint64{8}, 1, 1),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with methyl in name",
				Filters{CompoundName: "methyl"},
				testSearchResults([]uint64{0, 1, 3, 4}, 4, 4),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with Methyl in name",
				Filters{CompoundName: "Methyl"},
				testSearchResults([]uint64{0, 1, 3, 4}, 4, 4),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with name 11-HDoHE",
				Filters{CompoundName: "11-HDoHE"},
				testSearchResults([]uint64{}, 0, 0),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with contributor RIKEN",
				Filters{Contributor: &[]string{"RIKEN"}},
				testSearchResults([]uint64{7, 8}, 2, 2),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with contributor RIKEN and AAFC",
				Filters{Contributor: &[]string{"RIKEN", "AAFC"}},
				testSearchResults([]uint64{0, 7, 8}, 3, 3),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with formula C11H11",
				Filters{Formula: "C11H11"},
				testSearchResults([]uint64{1}, 1, 1),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with InchiKey ZVKARXLKNIBGIR-UHFFFAOYSA-N",
				Filters{InchiKey: "ZVKARXLKNIBGIR-UHFFFAOYSA-N"},
				testSearchResults([]uint64{3}, 1, 1),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with Peak difference 18.01056 ",
				Filters{PeakDifferences: &[]float64{18.01056}},
				testSearchResults([]uint64{4, 5, 8, 11}, 4, 4),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with Peak difference 27.9857 ",
				Filters{PeakDifferences: &[]float64{27.9857}},
				testSearchResults([]uint64{0, 3, 4, 7, 10, 11}, 6, 6),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with Peak difference 18.01056 or 27.9857",
				Filters{PeakDifferences: &[]float64{18.01056, 27.9857}},
				testSearchResults([]uint64{0, 3, 4, 5, 7, 8, 10, 11}, 8, 8),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with Peak 123.0",
				Filters{Peaks: &[]float64{123.0}},
				testSearchResults([]uint64{3, 4, 5}, 3, 3),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with Peak 123.0 with custom epsilon",
				Filters{Peaks: &[]float64{123.0}, MassEpsilon: ptr(0.1)},
				testSearchResults([]uint64{4, 5}, 2, 2),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with Peak 123.0 and 125.0",
				Filters{Peaks: &[]float64{123.0, 125.0}},
				testSearchResults([]uint64{4, 5}, 2, 2),
				false,
			},
			{
				db,
				db.name + " " + "Get all records with non existing Peak",
				Filters{Peaks: &[]float64{222.0}},
				testSearchResults([]uint64{}, 0, 0),
				false,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := tt.db.db.GetRecords(tt.args)
				if (err != nil) != tt.wantErr {
					t.Errorf("%s: GetRecords() error = %v, wantErr %v", tt.db.name, err, tt.wantErr)
					return
				}
				if got.ResultCount != tt.want.ResultCount {
					t.Errorf("Got wrong result count expected %v, got %v", tt.want.ResultCount, got.ResultCount)
				}
				if got.SpectraCount != tt.want.SpectraCount {
					t.Errorf("Got wrong spectra count expected %v, got %v", tt.want.SpectraCount, got.SpectraCount)
				}
				if (tt.args.Limit != 0 && len(got.Data) > int(tt.args.Limit)) ||
					len(got.Data) != len(tt.want.Data) {
					gotNames := []string{}
					for _, g := range got.Data {
						for _, sp := range g.Spectra {
							gotNames = append(gotNames, sp.Id)
						}
					}
					t.Errorf("Limit was %d, expected %d records, but got %d records: %v", tt.args.Limit, len(tt.want.Data), len(got.Data), gotNames)
				}
				compareDbResults(t, tt.want, *got)
			})
		}

	}
}

func compareDbResults(t *testing.T, want SearchResult, got SearchResult) {
	for wk, w := range want.Data {
		bw, errw := json.Marshal(w)
		found := false
		if g, ok := got.Data[wk]; ok {
			if g.Formula != w.Formula {
				t.Errorf("Result formula does not match: got %v , want %v", g.Formula, w.Formula)
			}
			if g.Smiles != w.Smiles {
				t.Errorf("Result smiles does not match: got %v , want %v", g.Smiles, w.Smiles)
			}
			if g.Mass != w.Mass {
				t.Errorf("Result mass does not match: got %v , want %v", g.Mass, w.Mass)
			}
			if len(g.Names) != len(w.Names) {
				t.Errorf("Count of names differs in Result wanted %v, got %v", w.Names, g.Names)
			}
			bg, errg := json.Marshal(w)
			for _, spw := range w.Spectra {
				for _, sp := range g.Spectra {
					if sp.Id == spw.Id {
						found = true
						if string(bg) != string(bw) || errg != nil || errw != nil {
							t.Errorf("\nwant: %v \ngot : %v\n", string(bw), string(bg))
						}
					}
				}
			}
		} else {
			t.Errorf("Could not find spectra for %v in result", wk)
		}
		if found == false {
			for _, spw := range w.Spectra {
				gotNames := []string{}
				for _, g := range got.Data {
					for _, sp := range g.Spectra {
						gotNames = append(gotNames, sp.Id)
					}
				}
				//println("Want: ", dd.Dump(want))
				//println("Got: ", dd.Dump(got))
				t.Errorf("Expected Accession %v to be in result but it was not found. Result was %v", spw.Id, gotNames)
			}
		}

	}
}

func TestMB3Database_GetUniqueValues(t *testing.T) {
	DBs, err := initDBs(All)
	if err != nil {
		t.Fatal("Could not init Databases: ", err.Error())
	}
	for _, db := range DBs {
		type testData struct {
			db      testDB
			name    string
			want    MB3Values
			wantErr bool
		}

		var tests = []testData{
			{
				db:      db,
				name:    db.name + " Test valid data",
				want:    UniqueValueTestData["all"],
				wantErr: false,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := tt.db.db.GetUniqueValues(Filters{})
				if (err != nil) != tt.wantErr {
					t.Errorf("%s: GetUniqueValues() error = %v, wantErr %v", tt.db.name, err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GetUniqueValues(): \ngot  %v, \nwant %v", got, tt.want)
				}
			})
		}
	}

}

func TestMB3Database_AddRecord(t *testing.T) {
	DBs, err := initDBs(Empty)
	if err != nil {
		t.Fatal("Could not init Databases: ", err.Error())
	}
	type args struct {
		record     massbank.MassBank2
		metaDataId string
	}
	type testData struct {
		name    string
		args    args
		wantErr bool
	}
	for _, db := range DBs {
		db.checkCount(t, 0)
		var tests = []testData{}
		for key, record := range mbTestRecords {
			var versionId string
			if db.name == "mongodb" {
				versionId = "63fcbaf4b9e0e5714f9d623b"
			} else {
				versionId = "1"
			}
			data := testData{
				name:    db.name + "-" + key,
				args:    args{record, versionId},
				wantErr: false,
			}
			tests = append(tests, data)
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := db.db.AddRecord(&tt.args.record, tt.args.metaDataId); (err != nil) != tt.wantErr {
					t.Errorf("AddRecord() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
		db.checkCount(t, int64(len(recordnames)))
	}
}

func TestMB3Database_AddRecords(t *testing.T) {
	DBs, err := initDBs(Empty)
	if err != nil {
		t.Fatal("Could not init Databases: ", err.Error())
	}
	type args struct {
		records    []uint64
		metaDataId string
	}
	type testData struct {
		name      string
		args      args
		wantErr   bool
		wantCount uint64
	}
	for _, db := range DBs {
		var versionId string
		if db.name == "mongodb" {
			versionId = "63fcbaf4b9e0e5714f9d623b"
		} else {
			versionId = "1"
		}
		var tests = []testData{
			{
				db.name + "-" + "Insert first 5",
				args{[]uint64{0, 1, 2, 3, 4}, versionId},
				false,
				5,
			},
			{
				db.name + "-" + "Insert second 5",
				args{[]uint64{5, 6, 7, 8, 9}, versionId},
				false,
				10,
			},
		}

		db.checkCount(t, 0)
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				var records []*massbank.MassBank2
				for _, i := range tt.args.records {
					var mb = mbTestRecords[recordnames[i]]
					records = append(records, &mb)
				}
				if err := db.db.AddRecords(records, tt.args.metaDataId); (err != nil) != tt.wantErr {
					t.Errorf("AddRecords() error = %v, wantErr %v", err, tt.wantErr)
				}
				db.checkCount(t, int64(tt.wantCount))
			})
		}
	}
}

func TestMB3Database_UpdateMetadata(t *testing.T) {
	DBs, err := initDBs(Main)
	if err != nil {
		t.Fatal("Could not init Databases: ", err.Error())
	}
	type args struct {
		metaData massbank.MbMetaData
	}
	type testData struct {
		name    string
		args    args
		wantErr bool
		oldId   bool
	}
	for _, db := range DBs {
		var tests = []testData{
			{
				db.name + " insert existing",
				args{metaData: massbank.MbMetaData{
					Commit:    "1e112e6777e453f54d8e3b3f6cac0f193d53fa38",
					Version:   "2022.12",
					Timestamp: "2023-02-02T13:35:54+01:00",
				}},
				false,
				true,
			},
			{
				db.name + " insert new",
				args{metaData: massbank.MbMetaData{
					Commit:    "1e112e6777e412344d8e3b3f6cac0f193d53fa38",
					Version:   "2023.1",
					Timestamp: "2023-03-02T18:35:54+01:00",
				}},
				false,
				false,
			},
		}
		var versionId string
		if db.name == "mongodb" {
			versionId = "63f746d01247aa361673cb67"
		} else {
			versionId = "1"
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				gotId, err := db.db.UpdateMetadata(&tt.args.metaData)
				if (err != nil) != tt.wantErr {
					t.Errorf("UpdateMetadata() error = %v, wantErr %v", err, tt.wantErr)
				}
				if (tt.oldId && gotId != versionId) || (!tt.oldId && gotId == versionId) {
					t.Errorf("UpdateMetadata() got wrong id: %v", gotId)

				}
			})

		}
	}

}

func TestMB3Database_UpdateRecord(t *testing.T) {
	DBs, err := initDBs(Main)
	if err != nil {
		t.Fatal("Could not init Databases: ", err.Error())
	}
	type args struct {
		record     *massbank.MassBank2
		metaDataId string
		upsert     bool
	}
	type testData struct {
		name      string
		args      args
		wantErr   bool
		wantCount uint64
	}
	for _, db := range DBs {
		db.checkCount(t, 8)
		var versionId string
		if db.name == "mongodb" {
			versionId = "63fcbaf4b9e0e5714f9d623b"
		} else {
			versionId = "1"
		}
		var tests = []testData{
			{
				"Update first",
				args{testRecord(0), versionId, false},
				false,
				8,
			},
			{
				"Update second",
				args{testRecord(1), versionId, true},
				false,
				8,
			},
			{
				"Try Insert with upsert",
				args{testRecord(5), versionId, false},
				false,
				8,
			},
			{
				"Try Insert with upsert",
				args{testRecord(5), versionId, true},
				false,
				9,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if _, _, err := db.db.UpdateRecord(tt.args.record, tt.args.metaDataId, tt.args.upsert); (err != nil) != tt.wantErr {
					t.Errorf("AddRecord() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
			db.checkCount(t, int64(tt.wantCount))

		}
	}
}

func TestMB3Database_UpdateRecords(t *testing.T) {
	DBs, err := initDBs(Main)
	if err != nil {
		t.Fatal("Could not init Databases: ", err.Error())
	}
	type args struct {
		records    []*massbank.MassBank2
		metaDataId string
		upsert     bool
	}
	type testData struct {
		name      string
		args      args
		wantErr   bool
		wantCount uint64
	}
	for _, db := range DBs {
		var versionId string
		if db.name == "mongodb" {
			versionId = "63fcbaf4b9e0e5714f9d623b"
		} else {
			versionId = "1"
		}
		var tests = []testData{
			{
				db.name + "-" + "Update first 5",
				args{testRecords([]uint64{0, 1, 2, 3, 4}), versionId, false},
				false,
				8,
			},
			{
				db.name + "-" + "Try insert 3 w/o upsert",
				args{testRecords([]uint64{5, 6, 7}), versionId, false},
				false,
				8,
			},
			{
				db.name + "-" + "Insert 3 with upsert",
				args{testRecords([]uint64{5, 6, 7}), versionId, true},
				false,
				11,
			},
			{
				db.name + "-" + "Insert/Update second 5 with upsert",
				args{testRecords([]uint64{5, 6, 7, 8, 9}), versionId, true},
				false,
				13,
			},
		}

		db.checkCount(t, 8)
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				var records []*massbank.MassBank2
				for _, r := range tt.args.records {
					if r.Comments == nil {
						r.Comments = &[]massbank.SubtagProperty{}
					}
					*r.Comments = append(*r.Comments, massbank.SubtagProperty{Value: "new comment"})
					records = append(records, r)
				}
				_, _, err := db.db.UpdateRecords(records, tt.args.metaDataId, tt.args.upsert)
				if (err != nil) != tt.wantErr {
					t.Errorf("UpdateRecords() error = %v, wantErr %v", err, tt.wantErr)
				}
				db.checkCount(t, int64(tt.wantCount))
			})
		}
	}
}

func TestMB3Database_GetMetadata(t *testing.T) {
	DBs, err := initDBs(All)
	if err != nil {
		t.Fatal("Could not init Databases: ", err.Error())
	}
	for _, db := range DBs {

		type testData struct {
			db      testDB
			name    string
			want    MB3MetaData
			wantErr bool
		}
		var tests = []testData{
			{
				db,
				db.name + "Get valid metadata",
				MB3MetaData{
					StoredMetadata: MB3StoredMetaData{
						Version:   "2022.12",
						TimeStamp: "2023-02-02T13:35:54+01:00",
						GitCommit: "1e112e6777e453f54d8e3b3f6cac0f193d53fa38",
					},
					SpectraCount:  12,
					CompoundCount: 11,
					IsomerCount:   11,
				},
				false,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := tt.db.db.GetMetaData()
				if (err != nil) != tt.wantErr {
					t.Errorf("%s: GetMetaData() error = %v, wantErr %v", tt.db.name, err, tt.wantErr)
					return
				}
				bg, errg := json.Marshal(got)
				bw, errw := json.Marshal(tt.want)
				if string(bg) != string(bw) || errg != nil || errw != nil {
					t.Errorf("\nwant: %v \ngot : %v\n", string(bw), string(bg))
				}
			})
		}
	}
}
