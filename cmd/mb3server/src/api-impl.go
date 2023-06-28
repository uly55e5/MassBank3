package mb3server

import (
	"github.com/MassBank/MassBank3/pkg/config"
	"github.com/MassBank/MassBank3/pkg/database"
	"github.com/MassBank/MassBank3/pkg/massbank"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

var db database.MB3Database = nil
var ServerConfig *config.ServerConfig = nil

func initDB() error {
	var err error
	if db != nil && db.Ping() == nil {
		return nil
	}
	db, err = database.InitDb(ServerConfig.DBConfig)
	return err
}

func GetBrowseOptions(instrumentTyoe []string, msType []string, ionMode string, contributor []string) (*BrowseOptions, error) {
	if err := initDB(); err != nil {
		return nil, err
	}
	it := &instrumentTyoe
	if len(*it) == 0 || (len(*it) == 1 && (*it)[0] == "") {
		it = nil
	}
	co := &contributor
	if len(*co) == 0 || (len(*co) == 1 && (*co)[0] == "") {
		co = nil
	}
	filters := database.Filters{
		InstrumentType:    it,
		Splash:            "",
		MsType:            getMsTypes(msType),
		IonMode:           getIonMode(ionMode),
		CompoundName:      "",
		Mass:              nil,
		MassEpsilon:       nil,
		Formula:           "",
		Peaks:             nil,
		PeakDifferences:   nil,
		InchiKey:          "",
		Contributor:       co,
		IntensityCutoff:   nil,
		Limit:             0,
		Offset:            0,
		IncludeDeprecated: false,
	}
	vals, err := db.GetUniqueValues(filters)
	if err != nil {
		return nil, err
	}
	var result = BrowseOptions{}
	metadata, err := db.GetMetaData()
	println(metadata)
	time, err := metadata.StoredMetadata.TimeStamp.MarshalText()
	result.Metadata = Metadata{
		Version:       metadata.StoredMetadata.Version,
		Timestamp:     string(time),
		GitCommit:     metadata.StoredMetadata.GitCommit,
		SpectraCount:  int32(metadata.SpectraCount),
		CompoundCount: int32(metadata.CompoundCount),
		IsomerCount:   int32(metadata.IsomerCount),
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
	for _, val := range vals.Contributor {
		result.Contributor = append(result.Contributor, StringCountInner{
			Value: val.Val,
			Count: int32(val.Count),
		})
	}

	return &result, nil
}

func GetRecords(limit int32, page int32, contributor []string, instrumentType []string, msType []string, ionMode string) (*SearchResult, error) {
	if err := initDB(); err != nil {
		return nil, err
	}
	if limit <= 0 {
		limit = 20
	}
	if page <= 0 {
		page = 1
	}

	it := &instrumentType
	if len(*it) == 0 || (len(*it) == 1 && (*it)[0] == "") {
		it = nil
	}
	co := &contributor
	if len(*co) == 0 || (len(*co) == 1 && (*co)[0] == "") {
		co = nil
	}

	var offset = (page - 1) * limit
	if err := initDB(); err != nil {
		return nil, err
	}

	var filters = database.Filters{
		InstrumentType:    it,
		Splash:            "",
		MsType:            getMsTypes(msType),
		IonMode:           getIonMode(ionMode),
		CompoundName:      "",
		Mass:              nil,
		MassEpsilon:       nil,
		Formula:           "",
		Peaks:             nil,
		PeakDifferences:   nil,
		InchiKey:          "",
		Contributor:       co,
		IntensityCutoff:   nil,
		Limit:             int64(limit),
		Offset:            int64(offset),
		IncludeDeprecated: false,
	}
	searchResult, err := db.GetRecords(filters)
	if err != nil {
		return nil, err
	}
	var result = SearchResult{}
	for _, value := range searchResult.Data {
		smiles := (value.Smiles)
		svg, err := getSvgFromSmiles(&smiles)
		svgS := ""
		if err == nil {
			re := regexp.MustCompile("<\\?xml[^>]*>\\n<!DOCTYPE[^>]*>\\n")
			svgS = string(re.ReplaceAll([]byte(*svg), []byte("")))
			re = regexp.MustCompile("\\n")
			svgS = string(re.ReplaceAll([]byte(svgS), []byte(" ")))
		} else {
			log.Println(err)
		}
		var val = SearchResultDataInner{
			Data:    map[string]interface{}{},
			Name:    value.Names,
			Formula: value.Formula,
			Mass:    value.Mass,
			Svg:     svgS,
			Spectra: []SearchResultDataInnerSpectraInner{},
		}
		for _, sp := range value.Spectra {
			val.Spectra = append(val.Spectra, SearchResultDataInnerSpectraInner{sp.Title, sp.Id})
		}
		result.Data = append(result.Data, val)
	}
	result.Metadata.ResultCount = int32(searchResult.ResultCount)
	result.Metadata.SpectraCount = int32(searchResult.SpectraCount)
	return &result, nil
}

func getIonMode(ionMode string) massbank.IonMode {
	switch ionMode {
	case "POSITIVE":
		return massbank.POSITIVE
	case "NEGATIVE":
		return massbank.NEGATIVE
	}
	return massbank.ANY
}

func getMsTypes(msType []string) *[]massbank.MsType {
	result := &[]massbank.MsType{}
	for _, t := range msType {
		switch t {
		case "MS":
			*result = append(*result, massbank.MS)
		case "MS2":
			*result = append(*result, massbank.MS2)
		case "MS3":
			*result = append(*result, massbank.MS3)
		case "MS4":
			*result = append(*result, massbank.MS4)
		}
	}

	if len(*result) == 0 {
		result = nil
	}
	return result
}

func GetSvg(accession string) (*string, error) {
	if err := initDB(); err != nil {
		return nil, err
	}
	smiles, err := db.GetSmiles(&accession)
	if err != nil {
		return nil, err
	}
	svg, err := getSvgFromSmiles(smiles)
	if err != nil {
		return nil, err
	}
	return svg, nil
}

func getSvgFromSmiles(smiles *string) (*string, error) {
	smilesEsc := url.QueryEscape(*smiles)
	resp, err := http.Get(ServerConfig.CdkDepictUrl + "/depict/bot/svg?smi=" + smilesEsc + "&w=50&h=50&abbr=on&hdisp=bridgehead&showtitle=false&zoom=1&annotate=none&r=1")
	if err != nil {
		return nil, err
	}
	svgB, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	svgS := string(svgB)
	return &svgS, nil
}

func GetRecord(accession string) (*MbRecord, error) {
	if err := initDB(); err != nil {
		return nil, err
	}
	record, err := db.GetRecord(&accession)
	if err != nil {
		return nil, err
	}
	result := MbRecord{
		Accession:  *record.Accession,
		Deprecated: MbRecordDeprecated{},
		Title:      *record.RecordTitle,
		Date: MbRecordDate{
			Updated:  record.Date.Updated.String(),
			Created:  record.Date.Created.String(),
			Modified: record.Date.Modified.String(),
		},
		Authors:     nil,
		License:     *record.License,
		Copyright:   "",
		Publication: "",
		Project:     "",
		Comments:    nil,
		Compound: MbRecordCompound{
			Names:     nil,
			Classes:   nil,
			Formula:   *record.Compound.Formula,
			CdkDepict: nil,
			Mass:      *record.Compound.Mass,
			Smiles:    *record.Compound.Smiles,
			Inchi:     *record.Compound.InChI,
			Link:      nil,
		},
		Species: MbRecordSpecies{
			Name:    "",
			Lineage: nil,
			Link:    nil,
			Sample:  nil,
		},
		Acquisition: MbRecordAcquisition{
			Instrument:     *record.Acquisition.Instrument,
			InstrumentType: *record.Acquisition.InstrumentType,
			MassSpectrometry: AcMassSpec{
				MsType:  "",
				IonMode: "",
				Subtags: nil,
			},
			Chromatography: nil,
			General:        nil,
			IonMobility:    nil,
		},
		MassSpectrometry: MbRecordMassSpectrometry{
			FocusedIon:     nil,
			DataProcessing: nil,
		},
		Peak: MbRecordPeak{
			Splash: *record.Peak.Splash,
			Annotation: MbRecordPeakAnnotation{
				Header: nil,
				Values: nil,
			},
			NumPeak: int32(*record.Peak.NumPeak),
			Peak: MbRecordPeakPeak{
				Header: nil,
				Values: nil,
			},
		},
	}
	for _, author := range *record.Authors {
		result.Authors = append(result.Authors, AuthorsInner{
			Name:        author.Name,
			MarcRelator: author.MarcRelator,
		})
	}
	return &result, nil

}
