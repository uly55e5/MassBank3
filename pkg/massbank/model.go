package massbank

import (
	"time"
)

const dateFormat = "2006.01.02"
const deprecatedDateFormat = "2006-01-02"

type MbMetaData struct {
	Commit    string
	Version   string
	Timestamp string
}

type SubtagProperty struct {
	Value  string
	Subtag string
}

type DatabaseProperty struct {
	Database   string
	Identifier string
}

type MbReference string

type Massbank struct {
	Metadata struct {
		FileName   string
		VersionRef MbReference
	}
	Accession   *RecordAccession   `mb2:"ACCESSION"`
	Deprecated  *RecordDeprecated  `mb2:"DEPRECATED" optional:"true" bson:"deprecated,omitempty"`
	RecordTitle *RecordTitle       `mb2:"RECORD_TITLE"`
	Date        *RecordDate        `mb2:"DATE"`
	Authors     *RecordAuthorNames `mb2:"AUTHORS"`
	License     *RecordLicense     `mb2:"LICENSE"`
	Copyright   *RecordCopyright   `mb2:"COPYRIGHT" optional:"true"`
	Publication *RecordPublication `mb2:"PUBLICATION" optional:"true"`
	Project     *RecordProject     `mb2:"PROJECT" optional:"true" bson:"project,omitempty"`
	Comments    []*RecordComment   `mb2:"COMMENT" optional:"true"`
	Compound    struct {
		Names     []*ChName          `mb2:"CH$NAME" json:"name"`
		Classes   *ChCompoundClasses `mb2:"CH$COMPOUND_CLASS" json:"classes"`
		Formula   *ChFormula         `mb2:"CH$FORMULA" json:"formula"`
		CdkDepict []*CdkDepict       `mb2:"CH$CDK_DEPICT" json:"cdk-depict"` // not for productive use
		Mass      *ChMass            `mb2:"CH$EXACT_MASS" json:"mass"`
		Smiles    *ChSmiles          `mb2:"CH$SMILES" json:"smiles"`
		Inchi     *ChInchi           `mb2:"CH$IUPAC" json:"inchi"`
		Link      []*ChLink          `mb2:"CH$LINK" optional:"true" json:"link"`
	} `json:"Compound"`
	Species struct {
		Name    *SpName              `mb2:"SP$SCIENTIFIC_NAME" optional:"true"`
		Lineage *SpLineage           `mb2:"SP$LINEAGE" optional:"true"`
		Link    []*SpLink            `mb2:"SP$LINK" optional:"true"`
		Sample  []*SampleInformation `mb2:"SP$SAMPLE" optional:"true"`
	}
	Acquisition struct {
		Instrument       *AcInstrument         `mb2:"AC$INSTRUMENT"`
		InstrumentType   *AcInstrumentType     `mb2:"AC$INSTRUMENT_TYPE"`
		MassSpectrometry []*AcMassSpectrometry `mb2:"AC$MASS_SPECTROMETRY" optional:"true"`
		Chromatography   []*AcChromatography   `mb2:"AC$CHROMATOGRAPHY" optional:"true"`
		General          []*AcGeneral          `mb2:"AC$GENERAL" optional:"true"`
	}
	MassSpectrometry struct {
		FocusedIon     []*MsFocusedIon     `mb2:"MS$FOCUSED_ION" optional:"true"`
		DataProcessing []*MsDataProcessing `mb2:"MS$DATA_PROCESSING" optional:"true"`
	}
	Peak struct {
		Splash     *PkSplash     `mb2:"PK$SPLASH"`
		Annotation *PkAnnotation `mb2:"PK$ANNOTATION" optional:"true"`
		NumPeak    *PkNumPeak    `mb2:"PK$NUM_PEAK"`
		Peak       *PkPeak       `mb2:"PK$PEAK"`
	}
}

type RecordAccession string

type RecordDeprecated struct {
	Date   time.Time
	Reason string
}

type RecordTitle string

type RecordDate struct {
	Updated  time.Time
	Created  time.Time
	Modified time.Time
}

type RecordAuthorNames []RecordAuthorName

type RecordAuthorName struct {
	Name        string
	MarcRelator string
}

type RecordLicense string

type RecordCopyright string

type RecordPublication string

type RecordProject string

type RecordComment SubtagProperty

type RecordSubtag string

type RecordMbTag string

type ChName string

type ChCompoundClasses []ChCompoundClass

type ChCompoundClass string

type ChFormula string

type ChMass float64

type ChSmiles string

type ChInchi string

type ChLink string

type ExtDatabase string

type CdkDepict string

type SpName string

type SpLineage []SpLineageElement

type SpLineageElement string

type SpLink DatabaseProperty

type SampleInformation string

type AcInstrument string

type Separation string
type Ionization string
type Analyzer string

type AcInstrumentType string

type MsType string

type IonMode string

type AcMassSpectrometry SubtagProperty

type AcChromatography SubtagProperty

type AcGeneral struct {
	SubtagProperty
}

type PkPeak struct {
	Header    []string
	Mz        []float64
	Intensity []float64
	Rel       []uint
}

type MsFocusedIon SubtagProperty

type MsDataProcessing SubtagProperty

type PkSplash string

type PkAnnotation struct {
	Header []string
	Values map[string][]interface{}
}

type PkNumPeak uint
