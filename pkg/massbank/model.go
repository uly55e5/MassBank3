package massbank

import (
	"time"
)

type MbMetaData struct {
	Commit        string `json:"git_commit"`
	Version       string `json:"Version"`
	Timestamp     string `json:"timestamp"`
	SpectraCount  uint   `json:"spectra_count"`
	CompoundCount uint   `json:"compound_count"`
	ResultCount   *uint  `json:"result_count"`
	Page          *uint  `json:"page"`
	Limit         *uint  `json:"limit"`
}

type SubtagProperty struct {
	Subtag string `json:"subtag"`
	Value  string `json:"value"`
}

type DatabaseProperty struct {
	Database   string `json:"database"`
	Identifier string `json:"identifier"`
}

type CompoundProperties struct {
	Names     *[]string           `mb2:"CH$NAME" json:"name"`
	Classes   *[]string           `mb2:"CH$COMPOUND_CLASS" json:"classes"`
	Formula   *string             `mb2:"CH$FORMULA" json:"formula"`
	CdkDepict *[]string           `mb2:"CH$CDK_DEPICT" json:"cdk_depict"` // not for productive use
	Mass      *float64            `mb2:"CH$EXACT_MASS" json:"mass"`
	Smiles    *string             `mb2:"CH$SMILES" json:"smiles"`
	InChI     *string             `mb2:"CH$IUPAC" json:"inchi"`
	Link      *[]DatabaseProperty `mb2:"CH$LINK" optional:"true" json:"link"`
}

type SpeciesProperties struct {
	Name    *string             `mb2:"SP$SCIENTIFIC_NAME" optional:"true" json:"name"`
	Lineage *[]string           `mb2:"SP$LINEAGE" optional:"true" json:"lineage"`
	Link    *[]DatabaseProperty `mb2:"SP$LINK" optional:"true" json:"link"`
	Sample  *[]string           `mb2:"SP$SAMPLE" optional:"true" json:"sample"`
}

type AcquisitionPropterties struct {
	Instrument       *string           `mb2:"AC$INSTRUMENT" json:"instrument"`
	InstrumentType   *string           `mb2:"AC$INSTRUMENT_TYPE" json:"instrument_type"`
	MassSpectrometry *[]SubtagProperty `mb2:"AC$MASS_SPECTROMETRY" optional:"true" json:"mass_spectrometry"`
	Chromatography   *[]SubtagProperty `mb2:"AC$CHROMATOGRAPHY" optional:"true" json:"chromatography"`
	General          *[]SubtagProperty `mb2:"AC$GENERAL" optional:"true" json:"general"`
}

type MassSpecProperties struct {
	FocusedIon     *[]SubtagProperty `mb2:"MS$FOCUSED_ION" optional:"true" json:"focused_ion"`
	DataProcessing *[]SubtagProperty `mb2:"MS$DATA_PROCESSING" optional:"true" json:"data_processing"`
}

type PeakProperties struct {
	Splash     *string       `mb2:"PK$SPLASH" json:"splash"`
	Annotation *PkAnnotation `mb2:"PK$ANNOTATION" optional:"true" json:"annotation"`
	NumPeak    *uint         `mb2:"PK$NUM_PEAK" json:"n_Peak"`
	Peak       *PkPeak       `mb2:"PK$PEAK" json:"peak"`
}

type Metadata struct {
	FileName   string `json:"file_name"`
	VersionRef string `json:"version_ref"`
}

type MassBank2 struct {
	Metadata         Metadata               `json:"metadata"`
	Accession        *string                `mb2:"ACCESSION" json:"accession"`
	Contributor      *string                `json:"contributor"`
	Deprecated       *RecordDeprecated      `mb2:"DEPRECATED" optional:"true" bson:"deprecated,omitempty" json:"deprecated"`
	RecordTitle      *string                `mb2:"RECORD_TITLE" json:"title"`
	Date             *RecordDate            `mb2:"DATE" json:"date"`
	Authors          *[]RecordAuthorName    `mb2:"AUTHORS" json:"authors"`
	License          *string                `mb2:"LICENSE" json:"license"`
	Copyright        *string                `mb2:"COPYRIGHT" optional:"true" json:"copyright"`
	Publication      *string                `mb2:"PUBLICATION" optional:"true" json:"publication"`
	Project          *string                `mb2:"PROJECT" optional:"true" bson:"project,omitempty" json:"project"`
	Comments         *[]SubtagProperty      `mb2:"COMMENT" optional:"true" json:"comments"`
	Compound         CompoundProperties     `json:"compound"`
	Species          SpeciesProperties      `json:"species"`
	Acquisition      AcquisitionPropterties `json:"acquisition"`
	MassSpectrometry MassSpecProperties     `json:"mass_spectrometry"`
	Peak             PeakProperties         `json:"peak"`
}

type RecordDeprecated struct {
	Date   time.Time `json:"date"`
	Reason string    `json:"reason"`
}

type RecordDate struct {
	Updated  time.Time `json:"updated"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

type RecordAuthorName struct {
	Name        string `json:"name"`
	MarcRelator string `json:"marc_relator"`
}

type PkPeak struct {
	Header    []string  `json:"header"`
	Mz        []float32 `json:"mz"`
	Intensity []float64 `json:"intensity"`
	Rel       []uint16  `json:"rel"`
}

type PkAnnotation struct {
	Header []string                 `json:"header"`
	Values map[string][]interface{} `json:"values"`
}
