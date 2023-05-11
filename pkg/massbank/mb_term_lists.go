package massbank

const (
	MS  MsType = "MS"
	MS2 MsType = "MS2"
	MS3 MsType = "MS3"
	MS4 MsType = "MS4"
)

func (ms MsType) String() string {
	return string(ms)
}

const (
	ANY      IonMode = ""
	POSITIVE IonMode = "POSITIVE"
	NEGATIVE IonMode = "NEGATIVE"
)

var licenseList = []string{
	"CC0",
	"CC BY",
	"CC BY-NC",
	"CC BY-NC-SA",
	"CC BY-SA",
}

var internalCommentTagList = []string{
	"MS2",
	"Merging",
	"Merged",
	"Mass spectrometry",
	"Chromatography",
	"Profile",
}

var databaseList = []string{
	"CAS",
	"CAYMAN",
	"CHEBI",
	"CHEMBL",
	"CHEMPDB",
	"CHEMSPIDER",
	"COMPTOX",
	"HMDB",
	"INCHIKEY",
	"KAPPAVIEW",
	"KEGG",
	"KNAPSACK",
	"LIPIDBANK",
	"LIPIDMAPS",
	"NIKKAJI",
	"PUBCHEM",
	"ZINC",
}

var InstrumentTypeSepList = []string{
	"CE",
	"GC",
	"LC",
}

var InstrumentTypeIonizationList = []string{
	"APCI",
	"APPI",
	"EI",
	"ESI",
	"FAB",
	"MALDI",
	"FD",
	"CI",
	"FI",
	"SIMS",
}

var InstrumentTypeAnalyzerList = []string{
	"B",
	"E",
	"FT",
	"IT",
	"Q",
	"TOF",
}

var tandemAnalyzerList []string

func TandemAnalyzerList() []string {
	if len(tandemAnalyzerList) == 0 {
		for _, s := range InstrumentTypeAnalyzerList {
			for _, ss := range InstrumentTypeAnalyzerList {
				tandemAnalyzerList = append(tandemAnalyzerList, s+ss)
			}
		}
	}
	return tandemAnalyzerList
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

var MSTypeList = []string{
	"MS",
	"MS2",
	"MS3",
	"MS4",
}

var MSIonModeList = []string{
	"POSITIVE",
	"NEGATIVE",
}

var MassSpectrometrySubList = []string{
	"MS_TYPE",
	"ION_MODE",
	"CAPILLARY_VOLTAGE",
	"COLLISION_ENERGY",
	"COLLISION_GAS",
	"DATE",
	"DESOLVATION_GAS_FLOW",
	"DESOLVATION_TEMPERATURE",
	"FRAGMENTATION_MODE",
	"IONIZATION",
	"IONIZATION_ENERGY",
	"LASER",
	"MATRIX",
	"MASS_ACCURACY",
	"MASS_RANGE_M/Z",
	"REAGENT_GAS",
	"RESOLUTION",
	"SCANNING_SETTING",
	"SOURCE_TEMPERATURE",
	"ACTIVATION_PARAMETER",
	"ACTIVATION_TIME",
	"ATOM_GUN_CURRENT",
	"AUTOMATIC_GAIN_CONTROL",
	"BOMBARDMENT",
	"CAPILLARY_TEMPERATURE",
	"CAPILLARY_VOLTAGE",
	"CDL_SIDE_OCTOPOLES_BIAS_VOLTAGE",
	"CDL_TEMPERATURE",
	"DATAFORMAT",
	"DRY_GAS_FLOW",
	"DRY_GAS_TEMP",
	"GAS_PRESSURE",
	"HELIUM_FLOW",
	"INTERFACE_VOLTAGE",
	"ION_GUIDE_PEAK_VOLTAGE",
	"ION_GUIDE_VOLTAGE",
	"ION_SPRAY_VOLTAGE",
	"IT_SIDE_OCTOPOLES_BIAS_VOLTAGE",
	"LENS_VOLTAGE",
	"NEBULIZER",
	"NEBULIZING_GAS",
	"NEEDLE_VOLTAGE",
	"OCTPOLE_VOLTAGE",
	"ORIFICE_TEMP",
	"ORIFICE_TEMPERATURE",
	"ORIFICE_VOLTAGE",
	"PROBE_TIP",
	"RING_VOLTAGE",
	"SAMPLE_DRIPPING",
	"SKIMMER_VOLTAGE",
	"SPRAY_VOLTAGE",
	"TUBE_LENS_VOLTAGE",
}

var MSFragmentationModeList = []string{
	"BIRD",
	"CID",
	"ECD",
	"EDD",
	"ETD",
	"HCD",
	"IRMPD",
	"MPD",
	"NETD",
	"SID",
	"UVPD",
}

var MSIonizationList = []string{
	"APCI",
	"APPI",
	"EI",
	"ESI",
	"FAB",
	"MALDI",
	"FD",
	"CI",
	"FI",
	"SIMS",
}

var commentSubtagList = []string{
	"[MS2]",
	"[Merging]",
	"[Merged]",
	"[Mass spectrometry]",
	"[Chromatography]",
	"[Profile]",
	"CONFIDENCE",
	"INTERNAL_ID",
}
