package common

import (
	"github.com/MassBank/MassBank3/pkg/massbank"
	"time"
)

func parseTimeWoErr(timeStr string) time.Time {
	t, _ := time.Parse(time.RFC3339, timeStr)
	return t
}

func toPtr[T any](v T) *T {
	return &v
}

var MbTestRecords = map[string]massbank.MassBank2{
	"MSBNK-AAFC-AC000005": {
		Metadata: massbank.Metadata{
			FileName:   "/MassBank-data/AAFC/MSBNK-AAFC-AC000005.txt",
			VersionRef: "63f38e47dc8c761f4af1ad51",
		},
		Accession:   toPtr("MSBNK-AAFC-AC000005"),
		Contributor: toPtr("AAFC"),
		Deprecated:  nil,
		RecordTitle: toPtr("Mellein; LC-ESI-ITFT; MS2; CE: 50; R=17500; [M+H]+"),
		Date: &massbank.RecordDate{
			Updated:  parseTimeWoErr("2017-07-07T00:00:00Z"),
			Created:  parseTimeWoErr("2017-07-07T00:00:00Z"),
			Modified: time.Time{},
		},
		Authors: &[]massbank.RecordAuthorName{
			{
				Name:        "Justin B. Renaud",
				MarcRelator: "",
			},
			{
				Name:        " Mark W. Sumarah",
				MarcRelator: "",
			},
			{
				Name:        " Agriculture and Agri-Food Canada",
				MarcRelator: "",
			},
		},
		License:     toPtr("CC BY-SA"),
		Copyright:   toPtr("Copyright (C) 2017"),
		Publication: toPtr("Renaud, J. B.; Sumarah, M. W. Data Independent Acquisition-Digital Archiving Mass Spectrometry: Application to Single Kernel Mycotoxin Analysis of Fusarium Graminearum Infected Maize. Analytical and Bioanalytical Chemistry 2016, 408 (12), 3083–91. DOI:10.1007/s00216-016-9391-5"),
		Project:     nil,
		Comments: &[]massbank.SubtagProperty{{
			Value: "isolated standard",

			Subtag: "CONFIDENCE"},
		},
		Compound: massbank.CompoundProperties{
			Names: &[]string{
				"Mellein",
				"Ochracin",
				"8-hydroxy-3-methyl-3,4-dihydroisochromen-1-one",
			},
			Classes: &[]string{

				"Natural Product",
				"Fungal metabolite",
			},
			Formula:   toPtr("C10H10O3"),
			CdkDepict: nil,
			Mass:      toPtr(178.062990),

			Smiles: toPtr("CC1CC2=C(C(=CC=C2)O)C(=O)O1"),
			InChI:  toPtr("InChI=1S/C10H10O3/c1-6-5-7-3-2-4-8(11)9(7)10(12)13-6/h2-4,6,11H,5H2,1H3"),
			Link: &[]massbank.DatabaseProperty{
				{

					Database:   "INCHIKEY",
					Identifier: "KWILGNNWGSNMPA-UHFFFAOYSA-N",
				},
				{
					Database:   "CAS",
					Identifier: "17397-85-2",
				},
				{

					Database:   "PUBCHEM",
					Identifier: "CID:28516",
				},
				{
					Database:   "CHEMSPIDER",
					Identifier: "26529",
				},
				{
					Database:   "KNAPSACK",
					Identifier: "C00000550",
				},
				{
					Database:   "COMPTOX",
					Identifier: "DTXSID60891794",
				},
			},
		},
		Species: massbank.SpeciesProperties{
			Name:    nil,
			Lineage: nil,
			Link:    nil,
			Sample:  nil,
		},
		Acquisition: massbank.AcquisitionPropterties{
			Instrument: toPtr("Q-Exactive Orbitrap Thermo Scientific"),

			InstrumentType: toPtr("LC-ESI-ITFT"),
			MassSpectrometry: &[]massbank.SubtagProperty{
				{
					Value:  "MS2",
					Subtag: "MS_TYPE",
				},
				{
					Value:  "POSITIVE",
					Subtag: "ION_MODE",
				},
				{
					Value:  "ESI",
					Subtag: "IONIZATION",
				},
				{

					Value: "3.9 kV",

					Subtag: "IONIZATION_VOLTAGE",
				},
				{

					Value:  "HCD",
					Subtag: "FRAGMENTATION_MODE",
				},
				{
					Value:  "50(NCE)",
					Subtag: "COLLISION_ENERGY",
				},
				{

					Value:  "17500",
					Subtag: "RESOLUTION",
				},
			},
			Chromatography: &[]massbank.SubtagProperty{
				{

					Value: "Agilent RRHD Eclipse 50 x 2 mm, 1.8 uM",

					Subtag: "COLUMN_NAME",
				},
				{

					Value:  "100:0 at 0 min, 100:0 at 0.5 min, 0:100 at 3.5 min, 0:100 at 5.5 min, 100:0 at 7 min",
					Subtag: "FLOW_GRADIENT",
				},
				{

					Value:  "0.3 mL min-1",
					Subtag: "FLOW_RATE",
				},
				{

					Value:  "3.44",
					Subtag: "RETENTION_TIME",
				},
				{

					Value:  "1094",
					Subtag: "NAPS_RTI",
				},
				{

					Value:  "A H2O 0.1% FA",
					Subtag: "SOLVENT",
				},
				{

					Value:  "B ACN 0.1% FA",
					Subtag: "SOLVENT",
				},
			},
			General: nil,
		},
		MassSpectrometry: massbank.MassSpecProperties{
			FocusedIon: &[]massbank.SubtagProperty{
				{

					Value: "133.0643",

					Subtag: "BASE_PEAK",
				},
				{

					Value: "179.0697",

					Subtag: "PRECURSOR_M/Z",
				},
				{

					Value: "[M+H]+",

					Subtag: "PRECURSOR_TYPE",
				},
			},
			DataProcessing: &[]massbank.SubtagProperty{
				{

					Value:  "Proteowizard",
					Subtag: "DEPROFILE",
				},
				{
					Value:  "based on Fragment ion formula determination",
					Subtag: "RECALIBRATE",
				},
				{
					Value:  "CUTOFF 0.05 Base Peak",
					Subtag: "INTENSITY",
				},
			},
		},
		Peak: massbank.PeakProperties{
			Splash: toPtr("splash10-01q9-0900000000-f556e9b53553d12b44b5"),
			Annotation: &massbank.PkAnnotation{
				Header: []string{
					"m/z",
					"tentative_formula",
					"mass_error(ppm)",
				},
				Values: map[string][]interface{}{
					"m/z": {
						105.069800,
						115.054000,
						133.064300,
						151.074700,
						161.058900,
						179.069600,
					},
					"mass_error(ppm)": {
						-0.840000,
						-2.050000,
						-3.740000,
						-4.370000,
						-5.020000,
						-3.740000,
					},
					"tentative_formula": {
						"C8H9+",
						"C9H7+",
						"C9H9O1+",
						"C9H11O2+",
						"C10H9O2+",
						"C10H11O3+",
					},
				},
			},
			NumPeak: toPtr(uint(6)),

			Peak: &massbank.PkPeak{
				Header: []string{
					"m/z",
					"int.",
					"rel.int.",
				},
				Mz: []float64{
					105.0698,
					115.054,
					133.0643,
					151.0747,
					161.0589,
					179.0696,
				},
				Intensity: []float64{
					1513707.125,
					134276.65625,
					3576966.75,
					143513.875,
					2536772.25,
					204606.65625,
				},
				Rel: []uint{
					422,
					36,
					999,
					39,
					708,
					56,
				},
			},
		},
	},

	"MSBNK-Athens_Univ-AU229201": {
		Metadata: massbank.Metadata{
			FileName:   "/MassBank-data/Athens_Univ/MSBNK-Athens_Univ-AU229201.txt",
			VersionRef: "63f38e47dc8c761f4af1ad51",
		},
		Accession:   toPtr("MSBNK-Athens_Univ-AU229201"),
		Contributor: toPtr("Athens_Univ"),
		Deprecated:  nil,
		RecordTitle: toPtr("2-Amino-3,8-dimethylimidazo-[4,5-f]quinoxaline (MeIQx); LC-ESI-QTOF; MS2; CE: 10 eV; R=35000; [M+H]+"),
		Date: &massbank.RecordDate{

			Updated:  parseTimeWoErr("2019-04-05T00:00:00Z"),
			Created:  parseTimeWoErr("2019-04-05T00:00:00Z"),
			Modified: time.Time{},
		},
		Authors: &[]massbank.RecordAuthorName{
			{
				Name:        "Nikiforos Alygizakis",
				MarcRelator: "",
			},
			{
				Name:        " Katerina Galani",
				MarcRelator: "",
			},
			{
				Name:        " Nikolaos Thomaidis",
				MarcRelator: "",
			},
			{
				Name:        " University of Athens",
				MarcRelator: "",
			},
		},
		License:     toPtr("CC BY"),
		Copyright:   toPtr("Copyright (C) 2019 Department of Chemistry, University of Athens"),
		Publication: nil,
		Project:     nil,
		Comments: &[]massbank.SubtagProperty{
			{
				Value:  "standard compound",
				Subtag: "CONFIDENCE",
			},
			{
				Value:  "2292",
				Subtag: "INTERNAL_ID",
			},
		},
		Compound: massbank.CompoundProperties{
			Names: &[]string{
				"2-Amino-3,8-dimethylimidazo-[4,5-f]quinoxaline (MeIQx)",

				"MeIQx",

				"3,8-dimethylimidazo[4,5-f]quinoxalin-2-amine",
			},
			Classes: &[]string{

				"N/A",
				"Environmental Standard",
			},
			Formula:   toPtr("C11H11N5"),
			CdkDepict: nil,
			Mass:      toPtr(213.1014454),
			Smiles:    toPtr("CN1C(N)=NC2=C1C=CC1=NC=C(C)N=C21"),
			InChI:     toPtr("InChI=1S/C11H11N5/c1-6-5-13-7-3-4-8-10(9(7)14-6)15-11(12)16(8)2/h3-5H,1-2H3,(H2,12,15)"),
			Link: &[]massbank.DatabaseProperty{
				{
					Database:   "CAS",
					Identifier: "77500-04-0",
				},
				{
					Database:   "CHEBI",
					Identifier: "76604",
				},
				{
					Database:   "KEGG",
					Identifier: "C19255",
				},
				{
					Database:   "PUBCHEM",
					Identifier: "CID:62275",
				},
				{
					Database:   "INCHIKEY",
					Identifier: "DVCCCQNKIYNAKB-UHFFFAOYSA-N",
				},
				{
					Database:   "CHEMSPIDER",
					Identifier: "56076",
				},
				{
					Database:   "COMPTOX",
					Identifier: "DTXSID1020801",
				},
			},
		},
		Species: massbank.SpeciesProperties{
			Name:    nil,
			Lineage: nil,
			Link:    nil,
			Sample:  nil,
		},
		Acquisition: massbank.AcquisitionPropterties{
			Instrument:     toPtr("Bruker maXis Impact"),
			InstrumentType: toPtr("LC-ESI-QTOF"),
			MassSpectrometry: &[]massbank.SubtagProperty{
				{
					Value:  "MS2",
					Subtag: "MS_TYPE",
				},
				{
					Value:  "POSITIVE",
					Subtag: "ION_MODE",
				},
				{
					Value:  "ESI",
					Subtag: "IONIZATION",
				},
				{
					Value:  "CID",
					Subtag: "FRAGMENTATION_MODE",
				},
				{
					Value:  "10 eV",
					Subtag: "COLLISION_ENERGY",
				},
				{
					Value:  "35000",
					Subtag: "RESOLUTION",
				},
			},
			Chromatography: &[]massbank.SubtagProperty{
				{
					Value:  "Acclaim RSLC C18 2.2um, 2.1x100mm, Thermo",
					Subtag: "COLUMN_NAME",
				},
				{
					Value:  "99/1 at 0-1 min, 61/39 at 3 min, 0.1/99.9 at 14-16 min, 99/1 at 16.1-20 min",
					Subtag: "FLOW_GRADIENT",
				},
				{
					Value:  "200 uL/min at 0-3 min, 400 uL/min at 14 min, 480 uL/min at 16-19 min, 200 uL/min at 19.1-20 min",
					Subtag: "FLOW_RATE",
				},
				{
					Value:  "4.075 min",
					Subtag: "RETENTION_TIME",
				},
				{
					Value:  "A 90:10 water:methanol with 0.01% formic acid and 5mM ammonium formate",
					Subtag: "SOLVENT",
				},
				{
					Value:  "B methanol with 0.01% formic acid and 5mM ammonium formate",
					Subtag: "SOLVENT",
				},
			},
			General: nil,
		},
		MassSpectrometry: massbank.MassSpecProperties{
			FocusedIon: &[]massbank.SubtagProperty{
				{
					Value:  "214.1081",
					Subtag: "BASE_PEAK",
				},
				{
					Value:  "214.1087",
					Subtag: "PRECURSOR_M/Z",
				},
				{
					Value:  "[M+H]+",
					Subtag: "PRECURSOR_TYPE",
				},
			},
			DataProcessing: &[]massbank.SubtagProperty{
				{
					Value:  "identity on assigned fragments and MS1",
					Subtag: "RECALIBRATE",
				},
				{
					Value:  "RMassBank 2.10.0",
					Subtag: "WHOLE",
				},
			},
		},
		Peak: massbank.PeakProperties{
			Splash: toPtr("splash10-03di-0090000000-4ca5f2773c0d659574ef"),
			Annotation: &massbank.PkAnnotation{

				Header: []string{
					"m/z",
					"tentative_formula",
					"formula_count",
					"mass",
					"error(ppm)",
				},
				Values: map[string][]interface{}{
					"error(ppm)": {
						-1.22,
						-7.78,
						-13.6,
					},
					"formula_count": {
						1,
						1,
						1,
					},
					"m/z": {
						214.1085,
						215.111,
						216.113,
					},
					"mass": {
						214.1087,
						215.1126,
						216.116,
					},
					"tentative_formula": {
						"C11H12N5+",
						"C10[13]CH12N5+",
						"C9[13]C2H12N5+",
					},
				},
			},
			NumPeak: toPtr(uint(3)),
			Peak: &massbank.PkPeak{

				Header: []string{
					"m/z",
					"int.",
					"rel.int.",
				},
				Mz: []float64{
					214.1085,
					215.111,
					216.113,
				},
				Intensity: []float64{
					2303864,
					289484,
					16528,
				},
				Rel: []uint{
					999,
					125,
					7,
				},
			},
		},
	},
	"MSBNK-Eawag-EA018353": {
		Metadata: massbank.Metadata{
			FileName:   "/MassBank-data/Eawag/MSBNK-Eawag-EA018353.txt",
			VersionRef: "63f38e47dc8c761f4af1ad51",
		},
		Accession:   toPtr("MSBNK-Eawag-EA018353"),
		Contributor: toPtr("Eawag"),
		Deprecated:  nil,
		RecordTitle: toPtr("Sulfadimethoxine; LC-ESI-ITFT; MS2; CE: 30%; R=7500; [M-H]-"),
		Date: &massbank.RecordDate{

			Updated:  parseTimeWoErr("2014-01-14T00:00:00Z"),
			Created:  parseTimeWoErr("2014-01-14T00:00:00Z"),
			Modified: time.Time{},
		},
		Authors: &[]massbank.RecordAuthorName{
			{
				Name:        "Stravs M",
				MarcRelator: "",
			},
			{
				Name:        " Schymanski E",
				MarcRelator: "",
			},
			{
				Name:        " Singer H",
				MarcRelator: "",
			},
			{
				Name:        " Department of Environmental Chemistry",
				MarcRelator: "",
			},
			{
				Name:        " Eawag",
				MarcRelator: "",
			},
		},
		License:     toPtr("CC BY"),
		Copyright:   toPtr("Copyright (C) 2012 Eawag, Duebendorf, Switzerland"),
		Publication: nil,
		Project:     nil,
		Comments: &[]massbank.SubtagProperty{
			{
				Value:  "standard compound",
				Subtag: "CONFIDENCE",
			},
			{
				Value:  "EAWAG_UCHEM_ID 183",
				Subtag: "",
			},
		},
		Compound: massbank.CompoundProperties{
			Names: &[]string{

				"Sulfadimethoxine",

				"4-amino-N-(2,6-dimethoxy-4-pyrimidinyl)benzenesulfonamide",
			},
			Classes: &[]string{

				"N/A",
				"Environmental Standard",
			},
			Formula:   toPtr("C12H14N4O4S"),
			CdkDepict: nil,
			Mass:      toPtr(310.0736),
			Smiles:    toPtr("c1(NS(c2ccc(N)cc2)(=O)=O)cc(nc(n1)OC)OC"),
			InChI:     toPtr("InChI=1S/C12H14N4O4S/c1-19-11-7-10(14-12(15-11)20-2)16-21(17,18)9-5-3-8(13)4-6-9/h3-7H,13H2,1-2H3,(H,14,15,16)"),
			Link: &[]massbank.DatabaseProperty{
				{
					Database:   "CAS",
					Identifier: "122-11-2",
				},
				{
					Database:   "CHEBI",
					Identifier: "32161",
				},
				{
					Database:   "KEGG",
					Identifier: "D01142",
				},
				{
					Database:   "PUBCHEM",
					Identifier: "CID:5323",
				},
				{
					Database:   "INCHIKEY",
					Identifier: "ZZORFUFYDOWNEF-UHFFFAOYSA-N",
				},
				{
					Database:   "CHEMSPIDER",
					Identifier: "5132",
				},
				{
					Database:   "COMPTOX",
					Identifier: "DTXSID1023607",
				},
			},
		},
		Species: massbank.SpeciesProperties{
			Name:    nil,
			Lineage: nil,
			Link:    nil,
			Sample:  nil,
		},
		Acquisition: massbank.AcquisitionPropterties{
			Instrument:     toPtr("LTQ Orbitrap XL Thermo Scientific"),
			InstrumentType: toPtr("LC-ESI-ITFT"),
			MassSpectrometry: &[]massbank.SubtagProperty{
				{
					Value:  "MS2",
					Subtag: "MS_TYPE",
				},
				{
					Value:  "NEGATIVE",
					Subtag: "ION_MODE",
				},
				{
					Value:  "ESI",
					Subtag: "IONIZATION",
				},
				{
					Value:  "HCD",
					Subtag: "FRAGMENTATION_MODE",
				},
				{
					Value:  "30 % (nominal)",
					Subtag: "COLLISION_ENERGY",
				},
				{
					Value:  "7500",
					Subtag: "RESOLUTION",
				},
			},
			Chromatography: &[]massbank.SubtagProperty{
				{
					Value:  "XBridge C18 3.5um, 2.1x50mm, Waters",
					Subtag: "COLUMN_NAME",
				},
				{
					Value:  "90/10 at 0 min, 50/50 at 4 min, 5/95 at 17 min, 5/95 at 25 min, 90/10 at 25.1 min, 90/10 at 30 min",
					Subtag: "FLOW_GRADIENT",
				},
				{
					Value:  "200 ul/min",
					Subtag: "FLOW_RATE",
				},
				{
					Value:  "5.7 min",
					Subtag: "RETENTION_TIME",
				},
				{
					Value:  "A water with 0.1% formic acid",
					Subtag: "SOLVENT",
				},
				{
					Value:  "B methanol with 0.1% formic acid",
					Subtag: "SOLVENT",
				},
			},
			General: nil,
		},
		MassSpectrometry: massbank.MassSpecProperties{
			FocusedIon: &[]massbank.SubtagProperty{
				{
					Value:  "309.0653",
					Subtag: "BASE_PEAK",
				},
				{
					Value:  "309.0663",
					Subtag: "PRECURSOR_M/Z",
				},
				{
					Value:  "[M-H]-",
					Subtag: "PRECURSOR_TYPE",
				},
			},
			DataProcessing: &[]massbank.SubtagProperty{
				{
					Value:  "Spline",
					Subtag: "DEPROFILE",
				},
				{
					Value:  "loess on assigned fragments and MS1",
					Subtag: "RECALIBRATE",
				},
				{
					Value:  "Peaks with additional N2/O included",
					Subtag: "REANALYZE",
				},
				{
					Value:  "RMassBank 1.3.1",
					Subtag: "WHOLE",
				},
			},
		},
		Peak: massbank.PeakProperties{
			Splash: toPtr("splash10-0a4i-0009000000-ccda245feefab7d2bd86"),
			Annotation: &massbank.PkAnnotation{

				Header: []string{
					"m/z",
					"tentative_formula",
					"formula_count",
					"mass",
					"error(ppm)",
				},
				Values: map[string][]interface{}{
					"error(ppm)": {
						3.3,
						0.09,
						-2.23,
						-0.94,
					},
					"formula_count": {
						2,
						1,
						1,
						1,
					},
					"m/z": {
						174.0679,
						195.0234,
						230.0804,
						309.066,
					},
					"mass": {
						174.0673,
						195.0234,
						230.0809,
						309.0663,
					},
					"tentative_formula": {
						"C9H8N3O-",
						"C8H7N2O2S-",
						"C11H10N4O2-",
						"C12H13N4O4S-",
					},
				},
			},
			NumPeak: toPtr(uint(4)),
			Peak: &massbank.PkPeak{

				Header: []string{
					"m/z",
					"int.",
					"rel.int.",
				},
				Mz: []float64{
					174.0679,
					195.0234,
					230.0804,
					309.0660,
				},
				Intensity: []float64{
					3134.2,
					4236.3,
					10982,
					805229.1,
				},
				Rel: []uint{
					3,
					5,
					13,
					999,
				},
			},
		},
	},
	"MSBNK-Eawag_Additional_Specs-ET060401": {
		Metadata: massbank.Metadata{
			FileName:   "/MassBank-data/Eawag_Additional_Specs/MSBNK-Eawag_Additional_Specs-ET060401.txt",
			VersionRef: "63f38e47dc8c761f4af1ad51",
		},
		Accession:   toPtr("MSBNK-Eawag_Additional_Specs-ET060401"),
		Contributor: toPtr("Eawag_Additional_Specs"),
		Deprecated:  nil,
		RecordTitle: toPtr("FEN_246.1101_16.1; LC-ESI-QFT; MS2; CE: 15; R=17500; [M+H]+"),
		Date: &massbank.RecordDate{

			Updated:  parseTimeWoErr("2016-03-17T00:00:00Z"),
			Created:  parseTimeWoErr("2015-09-25T00:00:00Z"),
			Modified: parseTimeWoErr("2016-02-03T00:00:00Z"),
		},
		Authors: &[]massbank.RecordAuthorName{
			{
				Name:        "R. Gulde",
				MarcRelator: "",
			},
			{
				Name:        " E. Schymanski",
				MarcRelator: "",
			},
			{
				Name:        " K. Fenner",
				MarcRelator: "",
			},
			{
				Name:        " Department of Environmental Chemistry",
				MarcRelator: "",
			},
			{
				Name:        " Eawag",
				MarcRelator: "",
			},
		},
		License:     toPtr("CC BY"),
		Copyright:   toPtr("Copyright (C) 2016 Eawag, Duebendorf, Switzerland"),
		Publication: toPtr("Gulde, R.; Meier, U.; Schymanski, E. L.; Kohler, H.-P. E.; Helbling, D. E.; Derrer, S.; Rentsch, D.; Fenner, K. Systematic Exploration of Biotransformation Reactions of Amine-Containing Micropollutants in Activated Sludge. Environmental Science \u0026 Technology 2016, 50 (6), 2908–20. DOI:10.1021/acs.est.5b05186"),
		Project:     nil,
		Comments: &[]massbank.SubtagProperty{
			{
				Value:  "Tentative identification: most likely structure (Level 3)",
				Subtag: "CONFIDENCE",
			},
			{
				Value:  "604",
				Subtag: "INTERNAL_ID",
			},
		},
		Compound: massbank.CompoundProperties{
			Names: &[]string{
				"FEN_246.1101_16.1",

				"N-desethyl-N-acetylfeniramine",

				"N-[1-[3-(trifluoromethyl)phenyl]propan-2-yl]acetamide",
			},
			Classes: &[]string{

				"N/A",
				"Environmental Transformation Products",
			},
			Formula:   toPtr("C12H14F3NO"),
			CdkDepict: nil,
			Mass:      toPtr(245.102700),
			Smiles:    toPtr("CC(CC1=CC(=CC=C1)C(F)(F)F)NC(C)=O"),
			InChI:     toPtr("InChI=1S/C12H14F3NO/c1-8(16-9(2)17)6-10-4-3-5-11(7-10)12(13,14)15/h3-5,7-8H,6H2,1-2H3,(H,16,17)"),
			Link: &[]massbank.DatabaseProperty{
				{
					Database:   "CAS",
					Identifier: "40552-64-5",
				},
				{
					Database:   "PUBCHEM",
					Identifier: "CID:38514",
				},
				{
					Database:   "INCHIKEY",
					Identifier: "ZVKARXLKNIBGIR-UHFFFAOYSA-N",
				},
				{
					Database:   "CHEMSPIDER",
					Identifier: "35298",
				},
				{
					Database:   "COMPTOX",
					Identifier: "DTXSID30891589",
				},
			},
		},
		Species: massbank.SpeciesProperties{
			Name:    nil,
			Lineage: nil,
			Link:    nil,
			Sample:  nil,
		},
		Acquisition: massbank.AcquisitionPropterties{
			Instrument:     toPtr("Q Exactive Plus Orbitrap Thermo Scientific"),
			InstrumentType: toPtr("LC-ESI-QFT"),
			MassSpectrometry: &[]massbank.SubtagProperty{
				{
					Value:  "MS2",
					Subtag: "MS_TYPE",
				},
				{
					Value:  "POSITIVE",
					Subtag: "ION_MODE",
				},
				{
					Value:  "ESI",
					Subtag: "IONIZATION",
				},
				{
					Value:  "HCD",
					Subtag: "FRAGMENTATION_MODE",
				},
				{
					Value:  "15 (nominal)",
					Subtag: "COLLISION_ENERGY",
				},
				{
					Value:  "17500",
					Subtag: "RESOLUTION",
				},
			},
			Chromatography: &[]massbank.SubtagProperty{
				{
					Value:  "Atlantis T3 3um, 3.0x150mm, Waters with guard column",
					Subtag: "COLUMN_NAME",
				},
				{
					Value:  "95/5 at 0 min, 5/95 at 15 min, 5/95 at 20 min, 95/5 at 20.1 min, 95/5 at 25 min",
					Subtag: "FLOW_GRADIENT",
				},
				{
					Value:  "300 uL/min",
					Subtag: "FLOW_RATE",
				},
				{
					Value:  "15.8 min",
					Subtag: "RETENTION_TIME",
				},
				{
					Value:  "A water with 0.1% formic acid",
					Subtag: "SOLVENT",
				},
				{
					Value:  "B methanol with 0.1% formic acid",
					Subtag: "SOLVENT",
				},
			},
			General: nil,
		},
		MassSpectrometry: massbank.MassSpecProperties{
			FocusedIon: &[]massbank.SubtagProperty{
				{
					Value:  "65.0597",
					Subtag: "BASE_PEAK",
				},
				{
					Value:  "246.11",
					Subtag: "PRECURSOR_M/Z",
				},
				{
					Value:  "[M+H]+",
					Subtag: "PRECURSOR_TYPE",
				},
			},
			DataProcessing: &[]massbank.SubtagProperty{
				{
					Value:  "loess on assigned fragments and MS1",
					Subtag: "RECALIBRATE",
				},
				{
					Value:  "Peaks with additional N2/O included",
					Subtag: "REANALYZE",
				},
				{
					Value:  "RMassBank 1.99.7",
					Subtag: "WHOLE",
				},
			},
		},
		Peak: massbank.PeakProperties{
			Splash: toPtr("splash10-052s-1940000000-a515d02f980838cad28d"),
			Annotation: &massbank.PkAnnotation{

				Header: []string{
					"m/z",
					"tentative_formula",
					"formula_count",
					"mass",
					"error(ppm)",
				},
				Values: map[string][]interface{}{
					"error(ppm)": {
						0.830000,
						0.070000,
						-4.770000,
						-2.400000,
						-2.010000,
						0.180000,
						0.740000,
						0.750000,
					},
					"formula_count": {
						1,
						1,
						1,
						1,
						1,
						1,
						1,
						1,
					},
					"m/z": {
						60.0444,
						60.0808,
						123.1162,
						141.0006,
						158.0272,
						159.0416,
						187.073,
						246.1102,
					},
					"mass": {
						60.0444,
						60.0808,
						123.1168,
						141.0009,
						158.0275,
						159.0416,
						187.0729,
						246.11,
					},
					"tentative_formula": {
						"C2H6NO+",
						"C3H10N+",
						"C9H15+",
						"C9FN+",
						"C9H3FN2+",
						"C8H6F3+",
						"C10H10F3+",
						"C12H15F3NO+",
					},
				},
			},
			NumPeak: toPtr(uint(9)),
			Peak: &massbank.PkPeak{

				Header: []string{
					"m/z",
					"int.",
					"rel.int.",
				},
				Mz: []float64{
					60.0444,
					60.0808,
					123.1162,
					141.0006,
					158.0272,
					159.0416,
					187.073,
					214.0897,
					246.1102,
				},
				Intensity: []float64{
					17595.5,
					1571.7,
					990.3,
					4683.9,
					14370,
					43520.4,
					97124.8,
					13043.94434,
					75456.7,
				},
				Rel: []uint{
					180,
					16,
					10,
					48,
					147,
					447,
					999,
					134,
					776,
				},
			},
		},
	},
	"MSBNK-Fac_Eng_Univ_Tokyo-JP009132": {
		Metadata: massbank.Metadata{
			FileName:   "/MassBank-data/Fac_Eng_Univ_Tokyo/MSBNK-Fac_Eng_Univ_Tokyo-JP009132.txt",
			VersionRef: "63f38e47dc8c761f4af1ad51",
		},
		Accession:   toPtr("MSBNK-Fac_Eng_Univ_Tokyo-JP009132"),
		Contributor: toPtr("Fac_Eng_Univ_Tokyo"),
		Deprecated:  nil,
		RecordTitle: toPtr("METHYL PHENYL SULFIDE; EI-B; MS"),
		Date: &massbank.RecordDate{

			Updated:  parseTimeWoErr("2016-01-19T00:00:00Z"),
			Created:  parseTimeWoErr("2008-10-21T00:00:00Z"),
			Modified: parseTimeWoErr("2011-05-06T00:00:00Z"),
		},
		Authors: &[]massbank.RecordAuthorName{
			{
				Name:        "SODA AROMATIC CO.",
				MarcRelator: "",
			},
			{
				Name:        " LTD.",
				MarcRelator: "",
			},
		},
		License:     toPtr("CC BY-NC-SA"),
		Copyright:   nil,
		Publication: nil,
		Project:     nil,
		Comments:    nil,
		Compound: massbank.CompoundProperties{
			Names: &[]string{
				"METHYL PHENYL SULFIDE",
			},
			Classes: &[]string{

				"N/A",
			},
			Formula:   toPtr("C7H8S"),
			CdkDepict: nil,
			Mass:      toPtr(124.034670),
			Smiles:    toPtr("CSc(c1)cccc1"),
			InChI:     toPtr("InChI=1S/C7H8S/c1-8-7-5-3-2-4-6-7/h2-6H,1H3"),
			Link: &[]massbank.DatabaseProperty{
				{
					Database:   "INCHIKEY",
					Identifier: "HNKJADCVZUBCPG-UHFFFAOYSA-N",
				},
				{
					Database:   "COMPTOX",
					Identifier: "DTXSID8059217",
				},
			},
		},
		Species: massbank.SpeciesProperties{
			Name:    nil,
			Lineage: nil,
			Link:    nil,
			Sample:  nil,
		},
		Acquisition: massbank.AcquisitionPropterties{
			Instrument:     toPtr("HITACHI M-80B"),
			InstrumentType: toPtr("EI-B"),
			MassSpectrometry: &[]massbank.SubtagProperty{
				{
					Value:  "MS",
					Subtag: "MS_TYPE",
				},
				{
					Value:  "POSITIVE",
					Subtag: "ION_MODE",
				},
				{
					Value:  "70 eV",
					Subtag: "IONIZATION_ENERGY",
				},
			},
			Chromatography: nil,
			General:        nil,
		},
		MassSpectrometry: massbank.MassSpecProperties{
			FocusedIon: &[]massbank.SubtagProperty{
				{
					Value:  "[M]+*",
					Subtag: "ION_TYPE",
				},
			},
			DataProcessing: nil,
		},
		Peak: massbank.PeakProperties{
			Splash:     toPtr("splash10-0fk9-9500000000-e5cde2928e0b7e33c1cc"),
			Annotation: nil,
			NumPeak:    toPtr(uint(43)),
			Peak: &massbank.PkPeak{

				Header: []string{
					"m/z",
					"int.",
					"rel.int.",
				},
				Mz: []float64{
					15.000000,
					26.000000,
					27.000000,
					37.000000,
					38.000000,
					39.000000,
					45.000000,
					46.000000,
					47.000000,
					49.000000,
					50.000000,
					51.000000,
					52.000000,
					57.000000,
					58.000000,
					61.000000,
					62.000000,
					63.000000,
					65.000000,
					66.000000,
					69.000000,
					70.000000,
					71.000000,
					73.000000,
					74.000000,
					75.000000,
					76.000000,
					77.000000,
					78.000000,
					79.000000,
					81.000000,
					82.000000,
					83.000000,
					91.000000,
					92.000000,
					108.000000,
					109.000000,
					110.000000,
					111.000000,
					123.000000,
					124.000000,
					125.000000,
					126.000000,
				},
				Intensity: []float64{
					5.68, 1.26, 3.11, 1.85, 4.41, 19.29, 28.64, 7.9, 5.09, 1.39, 18.14, 34.82, 3.18, 2.3, 2.04, 1.25, 6.11, 6.12, 22.52, 1.38, 14.67, 1.15, 1.48, 1.16, 6.23, 2.84, 1.81, 12.7, 42.96, 6.05, 2.1, 5.12, 1.61, 31.68, 2.02, 5.64, 46.38, 3.12, 1.97, 8.22, 99.99, 8.31, 4.82,
				},
				Rel: []uint{
					57,
					13,
					31,
					19,
					44,
					193,
					286,
					79,
					51,
					14,
					181,
					348,
					32,
					23,
					20,
					13,
					61,
					61,
					225,
					14,
					147,
					12,
					15,
					12,
					62,
					28,
					18,
					127,
					430,
					61,
					21,
					51,
					16,
					317,
					20,
					56,
					464,
					31,
					20,
					82,
					999,
					83,
					48,
				},
			},
		},
	},
	"MSBNK-Keio_Univ-KO009105": {
		Metadata: massbank.Metadata{
			FileName:   "/MassBank-data/Keio_Univ/MSBNK-Keio_Univ-KO009105.txt",
			VersionRef: "63f38e47dc8c761f4af1ad51",
		},
		Accession:   toPtr("MSBNK-Keio_Univ-KO009105"),
		Contributor: toPtr("Keio_Univ"),
		Deprecated:  nil,
		RecordTitle: toPtr("Desthiobiotin; LC-ESI-IT; MS4; m/z: 215/197/179; [M+H]+"),
		Date: &massbank.RecordDate{

			Updated:  parseTimeWoErr("2012-10-22T00:00:00Z"),
			Created:  parseTimeWoErr("2008-05-12T00:00:00Z"),
			Modified: time.Time{},
		},
		Authors: &[]massbank.RecordAuthorName{
			{
				Name:        "Ojima Y",
				MarcRelator: "",
			},
			{
				Name:        " Kakazu Y",
				MarcRelator: "",
			},
			{
				Name:        " Horai H",
				MarcRelator: "",
			},
			{
				Name:        " Soga T",
				MarcRelator: "",
			},
			{
				Name:        " Institute for Advanced Biosciences",
				MarcRelator: "",
			},
			{
				Name:        " Keio Univ.",
				MarcRelator: "",
			},
		},
		License:     toPtr("CC BY-NC-SA"),
		Copyright:   nil,
		Publication: nil,
		Project:     nil,
		Comments: &[]massbank.SubtagProperty{
			{
				Value:  "KEIO_ID D075",
				Subtag: "",
			},
			{
				Value:  "[MS3] KO009104",
				Subtag: "",
			},
		},
		Compound: massbank.CompoundProperties{
			Names: &[]string{
				"Desthiobiotin",

				"Dethiobiotin",
			},
			Classes: &[]string{

				"N/A",
			},
			Formula:   toPtr("C10H18N2O3"),
			CdkDepict: nil,
			Mass:      toPtr(214.131740),
			Smiles:    toPtr("OC(=O)CCCCCC(N1)C(C)NC(=O)1"),
			InChI:     toPtr("InChI=1S/C10H18N2O3/c1-7-8(12-10(15)11-7)5-3-2-4-6-9(13)14/h7-8H,2-6H2,1H3,(H,13,14)(H2,11,12,15)/t7-,8+/m0/s1"),
			Link: &[]massbank.DatabaseProperty{
				{
					Database:   "CAS",
					Identifier: "533-48-2",
				},
				{
					Database:   "CHEBI",
					Identifier: "16691",
				},
				{
					Database:   "KEGG",
					Identifier: "C01909",
				},
				{
					Database:   "NIKKAJI",
					Identifier: "J9.406D",
				},
				{
					Database:   "PUBCHEM",
					Identifier: "SID:5017",
				},
				{
					Database:   "INCHIKEY",
					Identifier: "AUTOLBMXDDTRRT-JGVFFNPUSA-N",
				},
				{
					Database:   "COMPTOX",
					Identifier: "DTXSID50876136",
				},
			},
		},
		Species: massbank.SpeciesProperties{
			Name:    nil,
			Lineage: nil,
			Link:    nil,
			Sample:  nil,
		},
		Acquisition: massbank.AcquisitionPropterties{
			Instrument:     toPtr("LC/MSD Trap XCT, Agilent Technologies"),
			InstrumentType: toPtr("LC-ESI-IT"),
			MassSpectrometry: &[]massbank.SubtagProperty{
				{
					Value:  "MS4",
					Subtag: "MS_TYPE",
				},
				{
					Value:  "POSITIVE",
					Subtag: "ION_MODE",
				},
				{
					Value:  "0.70/0.70/0.60",
					Subtag: "COLLISION_ENERGY",
				},
			},
			Chromatography: nil,
			General:        nil,
		},
		MassSpectrometry: massbank.MassSpecProperties{
			FocusedIon: &[]massbank.SubtagProperty{
				{
					Value:  "215/197/179",
					Subtag: "PRECURSOR_M/Z",
				},
				{
					Value:  "[M+H]+",
					Subtag: "PRECURSOR_TYPE",
				},
			},
			DataProcessing: &[]massbank.SubtagProperty{
				{
					Value:  "LC/MSD Trap Control and Data Analysis",
					Subtag: "WHOLE",
				},
			},
		},
		Peak: massbank.PeakProperties{
			Splash:     toPtr("splash10-000j-5900000000-8ad8423e2b359b500863"),
			Annotation: nil,
			NumPeak:    toPtr(uint(33)),
			Peak: &massbank.PkPeak{

				Header: []string{
					"m/z",
					"int.",
					"rel.int.",
				},
				Mz: []float64{
					56.2,
					58.2,
					67.2,
					68.2,
					70.2,
					76.7,
					79.2,
					80.0,
					81.1,
					82.0,
					85.0,
					91.1,
					93.1,
					94.0,
					95.1,
					96.1,
					99.1,
					108.1,
					109.1,
					110.1,
					111.0,
					113.1,
					119.1,
					123.1,
					125.1,
					134.1,
					136.1,
					137.1,
					151.1,
					162.0,
					179.1,
					180.1,
					197.2,
				},
				Intensity: []float64{
					7.30,
					8.10,
					66.36,
					1.80,
					6.70,
					4.90,
					14.00,
					2.60,
					64.28,
					18.49,
					3.00,
					10.00,
					49.50,
					6.69,
					28.99,
					10.20,
					342.46,
					6.60,
					33.39,
					12.19,
					41.20,
					17.90,
					180.85,
					11.40,
					12.89,
					52.21,
					498.25,
					13.29,
					142.24,
					3.90,
					39.87,
					54.97,
					2.20,
				},
				Rel: []uint{
					15,
					16,
					133,
					4,
					13,
					10,
					28,
					5,
					129,
					37,
					6,
					20,
					99,
					13,
					58,
					20,
					687,
					13,
					67,
					24,
					83,
					36,
					363,
					23,
					26,
					105,
					999,
					27,
					285,
					8,
					80,
					110,
					4,
				},
			},
		},
	},
	"MSBNK-MSSJ-MSJ00284": {
		Metadata: massbank.Metadata{
			FileName:   "/MassBank-data/MSSJ/MSBNK-MSSJ-MSJ00284.txt",
			VersionRef: "63f38e47dc8c761f4af1ad51",
		},
		Accession:   toPtr("MSBNK-MSSJ-MSJ00284"),
		Contributor: toPtr("MSSJ"),
		Deprecated:  nil,
		RecordTitle: toPtr("Pentoxazone; ESI-QTOF; MS2; NEGATIVE; [M-H]-; CID; 60 V"),
		Date: &massbank.RecordDate{

			Updated:  parseTimeWoErr("2020-03-18T00:00:00Z"),
			Created:  parseTimeWoErr("2020-03-18T00:00:00Z"),
			Modified: time.Time{},
		},
		Authors: &[]massbank.RecordAuthorName{
			{
				Name:        "Atsushi Yamamoto",
				MarcRelator: "",
			},
			{
				Name:        " Faculty of Environmental Studies",
				MarcRelator: "",
			},
			{
				Name:        " Tottori University of Environmental Studies",
				MarcRelator: "",
			},
			{
				Name:        " 1-1",
				MarcRelator: "",
			},
			{
				Name:        " Wakabadai-kita",
				MarcRelator: "",
			},
			{
				Name:        " Tottori City",
				MarcRelator: "",
			},
			{
				Name:        " Tottori 689-1111",
				MarcRelator: "",
			},
			{
				Name:        " Japan.",
				MarcRelator: "",
			},
		},
		License:     toPtr("CC BY"),
		Copyright:   toPtr("Atsushi Yamamoto, Faculty of Environmental Studies, Tottori University of Environmental Studies, 1-1, Wakabadai-kita, Tottori City, Tottori 689-1111, Japan."),
		Publication: nil,
		Project:     nil,
		Comments: &[]massbank.SubtagProperty{
			{
				Value:  "The sample was injected by direct infusion.",
				Subtag: "",
			},
			{
				Value:  "This record was created by the financial support of MEXT/JSPS KAKENHI Grant Number 19HP8024 to the Mass Spectrometry Society of Japan.",
				Subtag: "",
			},
		},
		Compound: massbank.CompoundProperties{
			Names: &[]string{
				"Pentoxazone",
			},
			Classes: &[]string{

				"Non-natural product",
			},
			Formula:   toPtr("C17H17ClFNO4"),
			CdkDepict: nil,
			Mass:      toPtr(353.083000),
			Smiles:    toPtr("CC(=C1C(=O)N(C(=O)O1)c2cc(c(cc2F)Cl)OC3CCCC3)C"),
			InChI:     toPtr("InChI=1S/C17H17ClFNO4/c1-9(2)15-16(21)20(17(22)24-15)13-8-14(11(18)7-12(13)19)23-10-5-3-4-6-10/h7-8,10H,3-6H2,1-2H3"),
			Link: &[]massbank.DatabaseProperty{
				{
					Database:   "CAS",
					Identifier: "110956-75-7",
				},
				{
					Database:   "CHEMSPIDER",
					Identifier: "9888955",
				},
				{
					Database:   "INCHIKEY",
					Identifier: "JZPKLLLUDLHCEL-UHFFFAOYSA-N",
				},
				{
					Database:   "PUBCHEM",
					Identifier: "CID:11714234",
				},
			},
		},
		Species: massbank.SpeciesProperties{
			Name:    nil,
			Lineage: nil,
			Link:    nil,
			Sample:  nil,
		},
		Acquisition: massbank.AcquisitionPropterties{
			Instrument:     toPtr("X500R QTOF (AB Sciex LLC, USA)"),
			InstrumentType: toPtr("ESI-QTOF"),
			MassSpectrometry: &[]massbank.SubtagProperty{
				{
					Value:  "MS2",
					Subtag: "MS_TYPE",
				},
				{
					Value:  "NEGATIVE",
					Subtag: "ION_MODE",
				},
				{
					Value:  "60 V",
					Subtag: "COLLISION_ENERGY",
				},
				{
					Value:  "CID",
					Subtag: "FRAGMENTATION_MODE",
				},
				{
					Value:  "ESI",
					Subtag: "IONIZATION",
				},
			},
			Chromatography: nil,
			General:        nil,
		},
		MassSpectrometry: massbank.MassSpecProperties{
			FocusedIon: &[]massbank.SubtagProperty{
				{
					Value:  "352.0757",
					Subtag: "PRECURSOR_M/Z",
				},
				{
					Value:  "[M-H]-",
					Subtag: "PRECURSOR_TYPE",
				},
			},
			DataProcessing: nil,
		},
		Peak: massbank.PeakProperties{
			Splash:     toPtr("splash10-0006-9400000000-6b99aa683e8584601a75"),
			Annotation: nil,
			NumPeak:    toPtr(uint(9)),
			Peak: &massbank.PkPeak{

				Header: []string{
					"m/z",
					"int.",
					"rel.int.",
				},
				Mz: []float64{
					31.9901,
					34.9691,
					41.0392,
					69.0342,
					97.0289,
					131.001,
					149.9992,
					166.9774,
					185.9761,
				},
				Intensity: []float64{
					0.02542,
					0.02094,
					0.26061,
					0.01645,
					0.02531,
					0.03766,
					0.04276,
					0.04392,
					0.03962,
				},
				Rel: []uint{
					97,
					80,
					999,
					63,
					97,
					144,
					164,
					168,
					152,
				},
			},
		},
	},
	"MSBNK-RIKEN-PR100978": {
		Metadata: massbank.Metadata{
			FileName:   "/MassBank-data/RIKEN/MSBNK-RIKEN-PR100978.txt",
			VersionRef: "63f38e47dc8c761f4af1ad51",
		},
		Accession:   toPtr("MSBNK-RIKEN-PR100978"),
		Contributor: toPtr("RIKEN"),
		Deprecated:  nil,
		RecordTitle: toPtr("Quercetin-3-Glucuronide; LC-ESI-QTOF; MS2; CE:Ramp 5-60 V; [M-H]-"),
		Date: &massbank.RecordDate{

			Updated:  parseTimeWoErr("2016-01-19T00:00:00Z"),
			Created:  parseTimeWoErr("2010-06-21T00:00:00Z"),
			Modified: parseTimeWoErr("2011-05-06T00:00:00Z"),
		},
		Authors: &[]massbank.RecordAuthorName{
			{
				Name:        "Matsuda F",
				MarcRelator: "",
			},
			{
				Name:        " Suzuki M",
				MarcRelator: "",
			},
			{
				Name:        " Sawada Y",
				MarcRelator: "",
			},
			{
				Name:        " Plant Science Center",
				MarcRelator: "",
			},
			{
				Name:        " RIKEN.",
				MarcRelator: "",
			},
		},
		License:     toPtr("CC BY-SA"),
		Copyright:   nil,
		Publication: nil,
		Project:     nil,
		Comments: &[]massbank.SubtagProperty{
			{
				Value:  "Acquisition and generation of the data is financially supported in part by CREST/JST.",
				Subtag: "",
			},
		},
		Compound: massbank.CompoundProperties{
			Names: &[]string{
				"Quercetin-3-Glucuronide",

				"Quer-3-GlcA",

				"Miquelianin",

				"Querciturone",
			},
			Classes: &[]string{

				"Flavonoid",
			},
			Formula:   toPtr("C21H18O13"),
			CdkDepict: nil,
			Mass:      toPtr(478.074740),
			Smiles:    toPtr("OC(=O)C(O1)C(O)C(O)C(O)C1OC(C(=O)3)=C(Oc(c4)c(c(O)cc(O)4)3)c(c2)cc(O)c(O)c2"),
			InChI:     toPtr("InChI=1S/C21H18O13/c22-7-4-10(25)12-11(5-7)32-17(6-1-2-8(23)9(24)3-6)18(13(12)26)33-21-16(29)14(27)15(28)19(34-21)20(30)31/h1-5,14-16,19,21-25,27-29H,(H,30,31)/t14-,15-,16+,19-,21+/m0/s1"),
			Link: &[]massbank.DatabaseProperty{
				{
					Database:   "CAS",
					Identifier: "22688-79-5",
				},
				{
					Database:   "CHEMSPIDER",
					Identifier: "4438874",
				},
				{
					Database:   "KNAPSACK",
					Identifier: "C00005376",
				},
				{
					Database:   "PUBCHEM",
					Identifier: "CID:5274585",
				},
				{
					Database:   "INCHIKEY",
					Identifier: "DUBCCGAQYVUYEU-ZUGPOPFOSA-N",
				},
				{
					Database:   "COMPTOX",
					Identifier: "DTXSID70945358",
				},
			},
		},
		Species: massbank.SpeciesProperties{
			Name:    nil,
			Lineage: nil,
			Link:    nil,
			Sample:  nil,
		},
		Acquisition: massbank.AcquisitionPropterties{
			Instrument:     toPtr("UPLC Q-Tof Premier, Waters"),
			InstrumentType: toPtr("LC-ESI-QTOF"),
			MassSpectrometry: &[]massbank.SubtagProperty{
				{
					Value:  "MS2",
					Subtag: "MS_TYPE",
				},
				{
					Value:  "NEGATIVE",
					Subtag: "ION_MODE",
				},
				{
					Value:  "Ramp 5-60 V",
					Subtag: "COLLISION_ENERGY",
				},
				{
					Value:  "Continuum",
					Subtag: "DATAFORMAT",
				},
				{
					Value:  "600.0 L/Hr",
					Subtag: "DESOLVATION_GAS_FLOW",
				},
				{
					Value:  "400 C",
					Subtag: "DESOLVATION_TEMPERATURE",
				},
				{
					Value:  "LOW-ENERGY CID",
					Subtag: "FRAGMENTATION_MODE",
				},
				{
					Value:  "ESI",
					Subtag: "IONIZATION",
				},
				{
					Value:  "120 C",
					Subtag: "SOURCE_TEMPERATURE",
				},
			},
			Chromatography: &[]massbank.SubtagProperty{
				{
					Value:  "3.0 kV",
					Subtag: "CAPILLARY_VOLTAGE",
				},
				{
					Value:  "23.0 V",
					Subtag: "SAMPLING_CONE",
				},
				{
					Value:  "A CH3CN(0.1%HCOOH)",
					Subtag: "SOLVENT",
				},
				{
					Value:  "B H2O(0.1%HCOOH)",
					Subtag: "SOLVENT",
				},
			},
			General: nil,
		},
		MassSpectrometry: massbank.MassSpecProperties{
			FocusedIon: &[]massbank.SubtagProperty{
				{
					Value:  "477.06692",
					Subtag: "PRECURSOR_M/Z",
				},
				{
					Value:  "[M-H]-",
					Subtag: "PRECURSOR_TYPE",
				},
			},
			DataProcessing: nil,
		},
		Peak: massbank.PeakProperties{
			Splash:     toPtr("splash10-0udi-0609400000-9fd50528da25d66adfc7"),
			Annotation: nil,
			NumPeak:    toPtr(uint(11)),
			Peak: &massbank.PkPeak{

				Header: []string{
					"m/z",
					"int.",
					"rel.int.",
				},
				Mz: []float64{
					107.0143,
					109.0299,
					121.0298,
					151.0035,
					163.0036,
					178.9989,
					255.0311,
					301.0353,
					302.0432,
					477.0669,
					478.0759,
				},
				Intensity: []float64{
					370.8,
					267.7,
					395.4,
					1650,
					295.2,
					793.8,
					333.3,
					4687,
					386,
					2215,
					237.5,
				},
				Rel: []uint{
					79,
					57,
					84,
					352,
					63,
					169,
					71,
					999,
					82,
					472,
					51,
				},
			},
		},
	},
	"MSBNK-RIKEN-PR309089": {
		Metadata: massbank.Metadata{
			FileName:   "/MassBank-data/RIKEN/MSBNK-RIKEN-PR309089.txt",
			VersionRef: "63f38e47dc8c761f4af1ad51",
		},
		Accession:   toPtr("MSBNK-RIKEN-PR309089"),
		Contributor: toPtr("RIKEN"),
		Deprecated:  nil,
		RecordTitle: toPtr("FA 18:2+1O; LC-ESI-QTOF; MS2"),
		Date: &massbank.RecordDate{

			Updated:  parseTimeWoErr("2019-03-28T00:00:00Z"),
			Created:  parseTimeWoErr("2019-03-28T00:00:00Z"),
			Modified: time.Time{},
		},
		Authors: &[]massbank.RecordAuthorName{
			{
				Name:        "Tetsuya Mori",
				MarcRelator: "",
			},
			{
				Name:        " Center for Sustainable Resource Science",
				MarcRelator: "",
			},
			{
				Name:        " RIKEN",
				MarcRelator: "",
			},
		},
		License:     toPtr("CC BY-NC-SA"),
		Copyright:   nil,
		Publication: toPtr("Tsugawa H., Nakabayashi R., Mori T., Yamada Y., Takahashi M., Rai A., Sugiyama R., Yamamoto H., Nakaya T., Yamazaki M., Kooke R., Bac-Molenaar JA., Oztolan-Erol N., Keurentjes JJB., Arita M., Saito K. (2019) \"A cheminformatics approach to characterize metabolomes in stable-isotope-labeled organisms\" Nature Methods 16(4):295-298. [doi:10.1038/s41592-019-0358-2]"),

		Project: nil,
		Comments: &[]massbank.SubtagProperty{
			{
				Value:  "Annotation level-3",
				Subtag: "",
			},
		},
		Compound: massbank.CompoundProperties{
			Names: &[]string{
				"FA 18:2+1O",
			},
			Classes: &[]string{

				"Oxidized fatty acids",
			},
			Formula:   toPtr("C18H32O3"),
			CdkDepict: nil,
			Mass:      toPtr(296.451000),
			Smiles:    toPtr("O=C(O)CCCCCCCC=CC=CC(O)CCCCC"),
			InChI:     toPtr("InChI=1S/C18H32O3/c1-2-3-11-14-17(19)15-12-9-7-5-4-6-8-10-13-16-18(20)21/h7,9,12,15,17,19H,2-6,8,10-11,13-14,16H2,1H3,(H,20,21)"),
			Link: &[]massbank.DatabaseProperty{
				{
					Database:   "INCHIKEY",
					Identifier: "HNICUWMFWZBIFP-UHFFFAOYSA-N",
				},
			},
		},
		Species: massbank.SpeciesProperties{
			Name:    nil,
			Lineage: nil,
			Link:    nil,
			Sample:  nil,
		},
		Acquisition: massbank.AcquisitionPropterties{
			Instrument:     toPtr("LC, Waters Acquity UPLC System; MS, Waters Xevo G2 Q-Tof"),
			InstrumentType: toPtr("LC-ESI-QTOF"),
			MassSpectrometry: &[]massbank.SubtagProperty{
				{
					Value:  "MS2",
					Subtag: "MS_TYPE",
				},
				{
					Value:  "NEGATIVE",
					Subtag: "ION_MODE",
				},
				{
					Value:  "6V",
					Subtag: "COLLISION_ENERGY",
				},
				{
					Value:  "800/h",
					Subtag: "DESOLVATION_GAS_FLOW",
				},
				{
					Value:  "450 C",
					Subtag: "DESOLVATION_TEMPERATURE",
				},
				{
					Value:  "ESI",
					Subtag: "IONIZATION",
				},
			},
			Chromatography: &[]massbank.SubtagProperty{
				{
					Value:  "+3.00 kV",
					Subtag: "CAPILLARY_VOLTAGE",
				},
				{
					Value:  "Acquity bridged ethyl hybrid C18 (1.7 um, 2.1 mm * 100 mm, Waters)",
					Subtag: "COLUMN_NAME",
				},
				{
					Value:  "40 C",
					Subtag: "COLUMN_TEMPERATURE",
				},
				{
					Value:  "A/B = (99.5%/0.5% at 0 min, 99.5%/0.5% at 0.1 min, 20%/80% at 10 min, 0.5%/99.5% at 10.1 min, 0.5%/99.5% at 12.0 min, 99.5%/0.5% at 12.1 min, 99.5%/0.5% at 15.0 min)",
					Subtag: "FLOW_GRADIENT",
				},
				{
					Value:  "0.3 ml/min at 0 min, 0.3 ml/min at 10 min, 0.4 ml/min at 10.1 min, 0.4 ml/min at 14.4 min, 0.3 ml/min at 14.5 min",
					Subtag: "FLOW_RATE",
				},
				{
					Value:  "9.57",
					Subtag: "RETENTION_TIME",
				},
				{
					Value:  "A water including 0.1% formic acid",
					Subtag: "SOLVENT",
				},
				{
					Value:  "B acetonitrile including 0.1% formic acid",
					Subtag: "SOLVENT",
				},
			},
			General: nil,
		},
		MassSpectrometry: massbank.MassSpecProperties{
			FocusedIon: &[]massbank.SubtagProperty{
				{
					Value:  "[M-H]-",
					Subtag: "PRECURSOR_TYPE",
				},
				{
					Value:  "295.2272",
					Subtag: "PRECURSOR_M/Z",
				},
			},
			DataProcessing: nil,
		},
		Peak: massbank.PeakProperties{
			Splash:     toPtr("splash10-0002-0090000000-3b4035f10b24d4c3f5a0"),
			Annotation: nil,
			NumPeak:    toPtr(uint(11)),
			Peak: &massbank.PkPeak{

				Header: []string{
					"m/z",
					"int.",
					"rel.int.",
				},
				Mz: []float64{
					57.08413,
					113.09826,
					172.09279,
					177.12296,
					195.13672,
					196.14641,
					197.18631,
					223.43452,
					233.15977,
					277.21619,
					295.22772,
				},
				Intensity: []float64{
					18.0,
					20.0,
					24.0,
					20.0,
					59.0,
					20.0,
					20.0,
					29.0,
					19.0,
					272.0,
					1553.0,
				},
				Rel: []uint{
					12,
					13,
					15,
					13,
					38,
					13,
					13,
					19,
					12,
					175,
					999,
				},
			},
		},
	},
	"MSBNK-Washington_State_Univ-BML81902": {
		Metadata: massbank.Metadata{

			FileName:   "/MassBank-data/Washington_State_Univ/MSBNK-Washington_State_Univ-BML81902.txt",
			VersionRef: "63f38e47dc8c761f4af1ad51",
		},
		Accession:   toPtr("MSBNK-Washington_State_Univ-BML81902"),
		Contributor: toPtr("Washington_State_Univ"),
		Deprecated:  nil,
		RecordTitle: toPtr("Parthenolide; LC-APCI-QTOF; MS; POSITIVE"),
		Date: &massbank.RecordDate{

			Updated:  parseTimeWoErr("2016-01-19T00:00:00Z"),
			Created:  parseTimeWoErr("2012-10-26T00:00:00Z"),
			Modified: time.Time{},
		},
		Authors: &[]massbank.RecordAuthorName{
			{
				Name:        "Cuthbertson DJ",
				MarcRelator: "",
			},
			{
				Name:        " Johnson SR",
				MarcRelator: "",
			},
			{
				Name:        " Lange BM",
				MarcRelator: "",
			},
			{
				Name:        " Institute of Biological Chemistry",
				MarcRelator: "",
			},
			{
				Name:        " Washington State University",
				MarcRelator: "",
			},
		},
		License:     toPtr("CC BY-SA"),
		Copyright:   nil,
		Publication: nil,
		Project:     nil,
		Comments: &[]massbank.SubtagProperty{
			{
				Value:  "relative retention time with respect to 9-anthracene Carboxylic Acid is 1.002",
				Subtag: "",
			},
		},
		Compound: massbank.CompoundProperties{
			Names: &[]string{
				"Parthenolide",
			},
			Classes: &[]string{

				"N/A",
			},
			Formula:   toPtr("C15H20O3"),
			CdkDepict: nil,
			Mass:      toPtr(248.141245),
			Smiles:    toPtr("C/C/1=C\\CC[C@@]2([C@@H](O2)[C@@H]3[C@@H](CC1)C(=C)C(=O)O3)C"),
			InChI:     toPtr("InChI=1S/C15H20O3/c1-9-5-4-8-15(3)13(18-15)12-11(7-6-9)10(2)14(16)17-12/h5,11-13H,2,4,6-8H2,1,3H3/b9-5+/t11-,12-,13-,15+/m0/s1"),
			Link: &[]massbank.DatabaseProperty{
				{
					Database:   "CAS",
					Identifier: "20554-84-1",
				},
				{
					Database:   "CHEMSPIDER",
					Identifier: "4554343",
				},
				{
					Database:   "PUBCHEM",
					Identifier: "CID:6473881",
				},
				{
					Database:   "INCHIKEY",
					Identifier: "KTEXNACQROZXEV-SLXBATTESA-N",
				},
			},
		},
		Species: massbank.SpeciesProperties{
			Name:    nil,
			Lineage: nil,
			Link:    nil,
			Sample:  nil,
		},
		Acquisition: massbank.AcquisitionPropterties{
			Instrument:     toPtr("Agilent 1200 RRLC; Agilent 6520 QTOF"),
			InstrumentType: toPtr("LC-APCI-QTOF"),
			MassSpectrometry: &[]massbank.SubtagProperty{
				{
					Value:  "MS",
					Subtag: "MS_TYPE",
				},
				{
					Value:  "POSITIVE",
					Subtag: "ION_MODE",
				},
				{
					Value:  "m/z 100-1000",
					Subtag: "SCANNING",
				},
			},
			Chromatography: &[]massbank.SubtagProperty{
				{
					Value:  "Agilent C8 Cartridge Column 2.1X30mm 3.5 micron (guard); Agilent SB-Aq 2.1x50mm 1.8 micron (analytical)",
					Subtag: "COLUMN_NAME",
				},
				{
					Value:  "60 C",
					Subtag: "COLUMN_TEMPERATURE",
				},
				{
					Value:  "linear from 98A/2B at 0 min to 2A/98B at 13 min, hold 6 min at 2A/98B, reequilibration 98A/2B (5 min)",
					Subtag: "FLOW_GRADIENT",
				},
				{
					Value:  "0.6 ml/min",
					Subtag: "FLOW_RATE",
				},
				{
					Value:  "7.394",
					Subtag: "RETENTION_TIME",
				},
				{
					Value:  "A water with 0.2% acetic acid",
					Subtag: "SOLVENT",
				},
				{
					Value:  "B methanol with 0.2% acetic acid",
					Subtag: "SOLVENT",
				},
			},
			General: nil,
		},
		MassSpectrometry: massbank.MassSpecProperties{
			FocusedIon:     nil,
			DataProcessing: nil,
		},
		Peak: massbank.PeakProperties{
			Splash:     toPtr("splash10-0002-0090000000-b81b1903703b01f61d7d"),
			Annotation: nil,
			NumPeak:    toPtr(uint(1)),
			Peak: &massbank.PkPeak{

				Header: []string{
					"m/z",
					"int.",
					"rel.int.",
				},
				Mz: []float64{
					249.1484765,
				},
				Intensity: []float64{
					310073.000000,
				},
				Rel: []uint{
					999,
				},
			},
		},
	},
	"MSBNK-test-TST00001": {
		Metadata: massbank.Metadata{
			FileName:   "/MassBank3/test-data/MSBNK-test-TST00001.txt",
			VersionRef: "63f50065841fc1dec65402fe",
		},
		Accession:   toPtr("MSBNK-test-TST00001"),
		Contributor: toPtr("test"),
		Deprecated:  nil,
		RecordTitle: toPtr("Fiscalin C; LC-ESI-ITFT; MS2; CE: 30; R=17500; [M+H]+"),
		Date: &massbank.RecordDate{

			Updated:  parseTimeWoErr("2017-07-07T00:00:00Z"),
			Created:  parseTimeWoErr("2017-07-07T00:00:00Z"),
			Modified: time.Time{},
		},
		Authors: &[]massbank.RecordAuthorName{
			{
				Name:        "Megan J. Kelman",
				MarcRelator: "",
			},
			{
				Name:        " Justin B. Renaud",
				MarcRelator: "",
			},
			{
				Name:        " Mark W. Sumarah",
				MarcRelator: "",
			},
			{
				Name:        " Agriculture and Agri-Food Canada",
				MarcRelator: "",
			},
		},
		License:     toPtr("CC BY-SA"),
		Copyright:   nil,
		Publication: nil,
		Project:     nil,
		Comments:    nil,
		Compound: massbank.CompoundProperties{
			Names: &[]string{
				"Fiscalin C",
			},
			Classes: &[]string{

				"Natural Product",
			},
			Formula:   toPtr("C27H29N5O4"),
			CdkDepict: nil,
			Mass:      toPtr(487.221940),
			Smiles:    toPtr("CC(C)[C@H]1C2=NC3=CC=CC=C3C(=O)N2[C@@H](C(=O)N1)C[C@@]4([C@@H]5NC(C(=O)N5C6=CC=CC=C64)(C)C)O"),
			InChI:     toPtr("InChI=1S/C27H29N5O4/c1-14(2)20-21-28-17-11-7-5-9-15(17)23(34)31(21)19(22(33)29-20)13-27(36)16-10-6-8-12-18(16)32-24(27)30-26(3,4)25(32)35/h5-12,14,19-20,24,30,36H,13H2,1-4H3,(H,29,33)/t19-,20+,24-,27-/m1/s1"),
			Link:      nil,
		},
		Species: massbank.SpeciesProperties{
			Name:    nil,
			Lineage: nil,
			Link:    nil,
			Sample:  nil,
		},
		Acquisition: massbank.AcquisitionPropterties{
			Instrument:     toPtr("Q-Exactive Orbitrap Thermo Scientific"),
			InstrumentType: toPtr("LC-ESI-ITFT"),
			MassSpectrometry: &[]massbank.SubtagProperty{
				{
					Value:  "MS2",
					Subtag: "MS_TYPE",
				},
				{
					Value:  "POSITIVE",
					Subtag: "ION_MODE",
				},
			},
			Chromatography: nil,
			General:        nil,
		},
		MassSpectrometry: massbank.MassSpecProperties{
			FocusedIon:     nil,
			DataProcessing: nil,
		},
		Peak: massbank.PeakProperties{
			Splash:     toPtr("splash10-03di-0290000000-8035e4fe85235c78b955"),
			Annotation: nil,
			NumPeak:    toPtr(uint(3)),
			Peak: &massbank.PkPeak{

				Header: []string{
					"m/z",
					"int.",
					"rel.int.",
				},
				Mz: []float64{
					185.1073,
					213.1022,
					258.1237,
				},
				Intensity: []float64{
					73653728.0,
					235010720.0,
					52446636.0,
				},
				Rel: []uint{
					312,
					999,
					222,
				},
			},
		},
	},
	"MSBNK-test-TST00002": {
		Metadata: massbank.Metadata{
			FileName:   "/MassBank3/test-data/MSBNK-test-TST00002.txt",
			VersionRef: "63f50065841fc1dec65402fe",
		},
		Accession:   toPtr("MSBNK-test-TST00002"),
		Contributor: toPtr("test"),
		Deprecated:  nil,
		RecordTitle: toPtr("Disialoganglioside GD1a; MALDI-TOF; MS; Pos"),
		Date: &massbank.RecordDate{

			Updated:  parseTimeWoErr("2016-01-19T00:00:00Z"),
			Created:  parseTimeWoErr("2009-11-24T00:00:00Z"),
			Modified: parseTimeWoErr("2011-05-11T00:00:00Z"),
		},
		Authors: &[]massbank.RecordAuthorName{
			{
				Name:        "Wada Y",
				MarcRelator: "",
			},
			{
				Name:        " Osaka Medical Center for Maternal and Child Health",
				MarcRelator: "",
			},
		},
		License:     toPtr("CC BY-SA"),
		Copyright:   toPtr("Copyright (C) 2019 just a test"),
		Publication: toPtr("Beisken S et al (2014) Scientific Data, 1:140029, DOI:10.1038/sdata.2014.29. http://www.ebi.ac.uk/metabolights/MTBLS38"),
		Project:     toPtr("my test project"),
		Comments: &[]massbank.SubtagProperty{
			{
				Value:  "[Chemical] The ceramide part has 38 carbons and 1 double bond (38:1), but its structure cannot be decided.  The displaying structure is GD1a d18:1-20:0.",
				Subtag: "",
			},
			{
				Value:  "Profile spectrum of this record is given as a JPEG file.",
				Subtag: "",
			},
			{
				Value:  "MCH00001.jpg",
				Subtag: "[Profile]",
			},
		},
		Compound: massbank.CompoundProperties{
			Names: &[]string{
				"Disialoganglioside GD1a",
				"another name",
			},
			Classes: &[]string{

				"Natural Product, Glycolipid",
				"Ceramide",
				"Ganglioside",
			},
			Formula:   toPtr("C86H152N4O39"),
			CdkDepict: nil,
			Mass:      toPtr(1865.003370),
			Smiles:    toPtr("CCCCCCCCCCCCCCCCCCCC(=O)N[C@@H](CO[C@H]1[C@@H]([C@H]([C@@H]([C@H](O1)CO)O[C@H]2[C@@H]([C@H]([C@H]([C@H](O2)CO)O[C@H]3[C@@H]([C@H]([C@H]([C@H](O3)CO)O)O[C@H]4[C@@H]([C@H]([C@H]([C@H](O4)CO)O)O[C@@]5(C[C@@H]([C@H]([C@@H](O5)C([C@@H](CO)O)O)NC(=O)C)O)C(=O)O)O)NC(=O)C)O[C@@]6(C[C@@H]([C@H]([C@@H](O6)C([C@@H](CO)O)O)NC(=O)C)O)C(=O)O)O)O)O)[C@@H](/C=C/CCCCCCCCCCCCC)O"),
			InChI:     toPtr("InChI=1S/C86H152N4O39/c1-6-8-10-12-14-16-18-20-21-22-23-25-27-29-31-33-35-37-60(105)90-50(51(100)36-34-32-30-28-26-24-19-17-15-13-11-9-7-2)46-118-80-69(111)68(110)72(58(44-95)121-80)123-82-71(113)78(129-86(84(116)117)39-53(102)62(88-48(4)98)76(127-86)65(107)55(104)41-92)73(59(45-96)122-82)124-79-63(89-49(5)99)74(66(108)56(42-93)119-79)125-81-70(112)77(67(109)57(43-94)120-81)128-85(83(114)115)38-52(101)61(87-47(3)97)75(126-85)64(106)54(103)40-91/h34,36,50-59,61-82,91-96,100-104,106-113H,6-33,35,37-46H2,1-5H3,(H,87,97)(H,88,98)(H,89,99)(H,90,105)(H,114,115)(H,116,117)/b36-34+/t50-,51+,52-,53-,54+,55+,56+,57+,58+,59+,61+,62+,63+,64?,65?,66-,67-,68+,69+,70+,71+,72+,73-,74+,75+,76+,77-,78+,79-,80+,81-,82-,85-,86-/m0/s1"),
			Link: &[]massbank.DatabaseProperty{
				{
					Database:   "INCHIKEY",
					Identifier: "UPMLUBZFFWELOX-IOFMCULOSA-N",
				},
				{
					Database:   "PUBCHEM",
					Identifier: "CID:137628549",
				},
			},
		},
		Species: massbank.SpeciesProperties{
			Name: toPtr("Bos taurus"),
			Lineage: &[]string{
				"cellular organisms",
				"Eukaryota",
				"Fungi/Metazoa group",
				"Metazoa",
				"Eumetazoa",
				"Bilateria",
				"Coelomata",
				"Deuterostomia",
				"Chordata",
				"Craniata",
				"Vertebrata",
				"Gnathostomata",
				"Teleostomi",
				"Euteleostomi",
				"Sarcopterygii",
				"Tetrapoda",
				"Amniota",
				"Mammalia",
				"Theria",
				"Eutheria",
				"Laurasiatheria",
				"Cetartiodactyla",
				"Ruminantia",
				"Pecora",
				"Bovidae",
				"Bovinae",
				"Bos",
			},
			Link: &[]massbank.DatabaseProperty{
				{
					Database:   "NCBI-TAXONOMY",
					Identifier: "9913",
				},
			},
			Sample: &[]string{

				"brain",
			},
		},
		Acquisition: massbank.AcquisitionPropterties{
			Instrument:     toPtr("Voyager DE-PRO, Applied Biosystems"),
			InstrumentType: toPtr("MALDI-TOF"),
			MassSpectrometry: &[]massbank.SubtagProperty{
				{
					Value:  "MS",
					Subtag: "MS_TYPE",
				},
				{
					Value:  "POSITIVE",
					Subtag: "ION_MODE",
				},
				{
					Value:  "UV 337 nm nitrogen lazer, 20 Hz, 10 nsec",
					Subtag: "LASER",
				},
				{
					Value:  "DHB",
					Subtag: "MATRIX",
				},
				{
					Value:  "500 pmol",
					Subtag: "SAMPLE_DRIPPING",
				},
				{
					Value:  "[MS, MS:1003294, electron activated dissociation, ]",
					Subtag: "FRAGMENTATION_MODE",
				},
				{
					Value:  "15 eV",
					Subtag: "KINETIC_ENERGY",
				},
				{
					Value:  "6500 nA",
					Subtag: "ELECTRON_CURRENT",
				},
				{
					Value:  "65 ms",
					Subtag: "REACTION_TIME",
				},
			},
			Chromatography: &[]massbank.SubtagProperty{
				{
					Value:  "Acclaim RSLC C18 2.2um, 2.1x100mm, Thermo",
					Subtag: "COLUMN_NAME",
				},
				{
					Value:  "99/1 at 0-1 min, 61/39 at 3 min, 0.1/99.9 at 14-16 min, 99/1 at 16.1-20 min",
					Subtag: "FLOW_GRADIENT",
				},
				{
					Value:  "200 uL/min at 0-3 min, 400 uL/min at 14 min, 480 uL/min at 16-19 min, 200 uL/min at 19.1-20 min",
					Subtag: "FLOW_RATE",
				},
				{
					Value:  "5.680 min",
					Subtag: "RETENTION_TIME",
				},
				{
					Value:  "A 90:10 water:methanol with 0.01% formic acid and 5mM ammonium formate",
					Subtag: "SOLVENT",
				},
				{
					Value:  "B methanol with 0.01% formic acid and 5mM ammonium formate",
					Subtag: "SOLVENT",
				},
			},
			General: nil,
		},
		MassSpectrometry: massbank.MassSpecProperties{
			FocusedIon: &[]massbank.SubtagProperty{
				{
					Value:  "324.2092",
					Subtag: "BASE_PEAK",
				},
				{
					Value:  "324.207",
					Subtag: "PRECURSOR_M/Z",
				},
				{
					Value:  "[M+H]+",
					Subtag: "PRECURSOR_TYPE",
				},
			},
			DataProcessing: &[]massbank.SubtagProperty{
				{
					Value:  "identity on assigned fragments and MS1",
					Subtag: "RECALIBRATE",
				},
				{
					Value:  "RMassBank 1.99.10",
					Subtag: "WHOLE",
				},
			},
		},
		Peak: massbank.PeakProperties{
			Splash: toPtr("splash10-03dr-0000010098-a1f70871442a6d6662d3"),
			Annotation: &massbank.PkAnnotation{

				Header: []string{
					"m/z",
					"ion",
				},
				Values: map[string][]interface{}{
					"ion": {
						"[LH-2NeuAc+Na]+",
						"[M-2NeuAc+Na]+",
						"[M-NeuAc+Na]+",
						"[LH+Na]+",
						"[LH+K]+",
						"[LH-H+2Na]+",
						"[M+Na]+",
						"[LH-H+K+Na]+",
						"[M+K]+",
						"[M-H+2Na]+",
						"[LH-2H+K+2Na]+",
						"[M-H+K+Na]+",
						"[M-2H+3Na]+",
						"[M-2H+K+2Na]+",
					},
					"m/z": {
						1278.120000,
						1306.210000,
						1597.120000,
						1860.160000,
						1876.140000,
						1882.180000,
						1888.200000,
						1898.150000,
						1904.190000,
						1910.210000,
						1920.180000,
						1926.200000,
						1932.220000,
						1948.240000,
					},
				},
			},
			NumPeak: toPtr(uint(84)),
			Peak: &massbank.PkPeak{

				Header: []string{
					"m/z",
					"int.",
					"rel.int.",
				},
				Mz: []float64{
					1278.12, 1279.11, 1279.18, 1306.21, 1309.13, 1366.03, 1375.23, 1527.60, 1569.08, 1591.14, 1595.21, 1597.12, 1598.11, 1599.07, 1620.30, 1620.80, 1621.18, 1843.21, 1860.16, 1861.17, 1862.18, 1863.19, 1864.22, 1866.26, 1867.23, 1871.18, 1876.14, 1877.14, 1878.13, 1879.19, 1882.18, 1883.18, 1884.18, 1885.20, 1886.19, 1888.20, 1889.21, 1890.21, 1891.22, 1892.24, 1893.23, 1894.24, 1895.21, 1897.14, 1898.15, 1899.16, 1900.17, 1902.16, 1903.20, 1904.19, 1905.20, 1906.20, 1907.23, 1908.19, 1910.21, 1911.21, 1912.21, 1913.22, 1914.21, 1915.20, 1920.18, 1921.21, 1922.19, 1926.20, 1927.19, 1928.20, 1929.21, 1930.20, 1931.32, 1932.22, 1933.23, 1934.23, 1935.22, 1943.20, 1948.24, 1949.22, 1950.22, 1960.26, 1965.23, 1971.32, 2064.29, 2091.32, 2092.30, 2136.33,
				},
				Intensity: []float64{
					205.000000,
					177.000000,
					186.000000,
					242.000000,
					168.000000,
					162.000000,
					291.000000,
					181.000000,
					187.000000,
					198.000000,
					177.000000,
					261.000000,
					228.000000,
					185.000000,
					272.000000,
					266.000000,
					367.000000,
					198.000000,
					905.000000,
					822.000000,
					531.000000,
					284.000000,
					219.000000,
					183.000000,
					162.000000,
					168.000000,
					297.000000,
					280.000000,
					192.000000,
					165.000000,
					497.000000,
					545.000000,
					349.000000,
					242.000000,
					181.000000,
					1311.000000,
					1127.000000,
					844.000000,
					413.000000,
					240.000000,
					185.000000,
					175.000000,
					146.000000,
					165.000000,
					320.000000,
					260.000000,
					218.000000,
					191.000000,
					188.000000,
					574.000000,
					509.000000,
					383.000000,
					222.000000,
					186.000000,
					795.000000,
					674.000000,
					478.000000,
					264.000000,
					194.000000,
					187.000000,
					284.000000,
					215.000000,
					179.000000,
					368.000000,
					363.000000,
					254.000000,
					207.000000,
					187.000000,
					170.000000,
					532.000000,
					495.000000,
					352.000000,
					261.000000,
					194.000000,
					274.000000,
					255.000000,
					226.000000,
					151.000000,
					185.000000,
					168.000000,
					176.000000,
					180.000000,
					206.000000,
					162.000000,
				},
				Rel: []uint{
					156,
					135,
					142,
					184,
					128,
					123,
					222,
					138,
					142,
					151,
					135,
					199,
					174,
					141,
					207,
					203,
					280,
					151,
					690,
					626,
					405,
					216,
					167,
					139,
					123,
					128,
					226,
					213,
					146,
					126,
					379,
					415,
					266,
					184,
					138,
					999,
					859,
					643,
					315,
					183,
					141,
					133,
					111,
					126,
					244,
					198,
					166,
					146,
					143,
					437,
					388,
					292,
					169,
					142,
					606,
					514,
					364,
					201,
					148,
					142,
					216,
					164,
					136,
					280,
					277,
					194,
					158,
					142,
					130,
					405,
					377,
					268,
					199,
					148,
					209,
					194,
					172,
					115,
					141,
					128,
					134,
					137,
					157,
					123,
				},
			},
		},
	},
	"MSBNK-test-TST00003": {
		Metadata: massbank.Metadata{
			FileName:   "/MassBank3/test-data/MSBNK-test-TST00003.txt",
			VersionRef: "63f50065841fc1dec65402fe",
		},
		Accession:   toPtr("MSBNK-test-TST00003"),
		Contributor: toPtr("test"),
		Deprecated: &massbank.RecordDeprecated{
			Date:   parseTimeWoErr("2019-11-25T00:00:00Z"),
			Reason: "Wrong MS measurement assigned",
		},
		RecordTitle: toPtr("11-HDoHE; LC-ESI-QTOF; MS2; CE: 20.0; R=N/A; [M-H]-"),
		Date: &massbank.RecordDate{

			Updated:  parseTimeWoErr("2018-11-21T00:00:00Z"),
			Created:  parseTimeWoErr("2018-11-21T00:00:00Z"),
			Modified: time.Time{},
		},
		Authors: &[]massbank.RecordAuthorName{
			{
				Name:        "Nils Hoffmann",
				MarcRelator: "",
			},
			{
				Name:        " Dominik Kopczynski",
				MarcRelator: "",
			},
			{
				Name:        " Bing Peng",
				MarcRelator: "",
			},
		},
		License:     toPtr("CC BY-SA"),
		Copyright:   toPtr("Copyright (C) 2019, Leibniz Institut fuer Analytische Wissenschaften - ISAS - e.V., Dortmund, Germany"),
		Publication: nil,
		Project:     nil,
		Comments: &[]massbank.SubtagProperty{
			{
				Value:  "standard compound",
				Subtag: "CONFIDENCE",
			},
			{
				Value:  "NATIVE_RUN_ID QExHF03_NM_0001275.mzML",
				Subtag: "",
			},
			{
				Value:  "PROCESSING averaging of repeated ion fragments at 20.0 eV within 5 ppm window [MS, MS:1000575, mean of spectra, ]",
				Subtag: "",
			},
		},
		Compound: massbank.CompoundProperties{
			Names: &[]string{
				"11-HDoHE",
			},
			Classes: &[]string{

				"Natural Product",
				"Lipid Standard",
			},
			Formula:   toPtr("C22H32O3"),
			CdkDepict: nil,
			Mass:      toPtr(344.235140),
			Smiles:    toPtr("CC\\C=C/C\\C=C/C\\C=C/CC(O)\\C=C\\C=C/C\\C=C/CCC(O)=O"),
			InChI:     toPtr("InChI=1S/C22H32O3/c1-2-3-4-5-6-7-9-12-15-18-21(23)19-16-13-10-8-11-14-17-20-22(24)25/h3-4,6-7,10-16,19,21,23H,2,5,8-9,17-18,20H2,1H3,(H,24,25)/b4-3-,7-6-,13-10-,14-11-,15-12-,19-16+"),
			Link: &[]massbank.DatabaseProperty{
				{
					Database:   "CHEBI",
					Identifier: "CHEBI:72794",
				},
				{
					Database:   "LIPIDMAPS",
					Identifier: "LMFA04000028",
				},
				{
					Database:   "INCHIKEY",
					Identifier: "LTERDCBCHFKFRI-BGKMTWLOSA-N",
				},
				{
					Database:   "PUBCHEM",
					Identifier: "CID:11631564",
				},
			},
		},
		Species: massbank.SpeciesProperties{
			Name:    nil,
			Lineage: nil,
			Link:    nil,
			Sample:  nil,
		},
		Acquisition: massbank.AcquisitionPropterties{
			Instrument:     toPtr("Q-Exactive HF, Thermo Scientific [MS:1002523]"),
			InstrumentType: toPtr("LC-ESI-QTOF"),
			MassSpectrometry: &[]massbank.SubtagProperty{
				{
					Value:  "MS2",
					Subtag: "MS_TYPE",
				},
				{
					Value:  "NEGATIVE",
					Subtag: "ION_MODE",
				},
				{
					Value:  "ESI",
					Subtag: "IONIZATION",
				},
				{
					Value:  "HCD",
					Subtag: "FRAGMENTATION_MODE",
				},
				{
					Value:  "20.0 eV",
					Subtag: "COLLISION_ENERGY",
				},
				{
					Value:  "N/A",
					Subtag: "RESOLUTION",
				},
			},
			Chromatography: nil,
			General:        nil,
		},
		MassSpectrometry: massbank.MassSpecProperties{
			FocusedIon: &[]massbank.SubtagProperty{
				{
					Value:  "343.2279",
					Subtag: "BASE_PEAK",
				},
				{
					Value:  "343.2279",
					Subtag: "PRECURSOR_M/Z",
				},
				{
					Value:  "[M-H]-",
					Subtag: "PRECURSOR_TYPE",
				},
			},
			DataProcessing: nil,
		},
		Peak: massbank.PeakProperties{
			Splash: toPtr("splash10-00dl-0914000000-bad728fe323f6890656d"),
			Annotation: &massbank.PkAnnotation{

				Header: []string{
					"m/z",
					"annotation",
					"exact_mass",
					"error(ppm)",
				},
				Values: map[string][]interface{}{
					"annotation": {
						59.013, 95.050, 121.102, 149.133, 165.092, 194.095, 281.227, 299.238, 325.217, "precursor",
					},
					"error(ppm)": {
						2.913263365927882, 0.40954364535633475, 0.29031912309532387, 1.1221931248150303, 0.7078231920609128, 0.7437838072718687, -0.1356156911126373, 0.1124819597449277, -0.2673869905083813, -1.1896746249719279,
					},
					"exact_mass": {
						59.0133, 95.0502, 121.1022, 149.1333, 165.0919, 194.0946, 281.2273, 299.2378, 325.2171, 343.227868554909,
					},
					"m/z": {
						59.013471921284996, 95.0502389272054, 121.1022351582845, 149.13346735636392, 165.09201685587564, 194.09474436442056, 281.22726186116535, 299.2378336588542, 325.2170130411784, 343.22746022542316,
					},
				},
			},
			NumPeak: toPtr(uint(10)),
			Peak: &massbank.PkPeak{

				Header: []string{
					"m/z",
					"int.",
					"rel.int.",
				},
				Mz: []float64{
					59.013471921284996, 95.0502389272054, 121.1022351582845, 149.13346735636392, 165.09201685587564, 194.09474436442056, 281.22726186116535, 299.2378336588542, 325.2170130411784, 343.22746022542316,
				},
				Intensity: []float64{
					47825.663583333335, 123675.15916666668, 2119825.3000000003, 538442.9091666666, 717615.9775, 348560.26666666666, 577177.7541666667, 162688.84041666667, 246360.69083333333, 1749842.5083333335,
				},
				Rel: []uint{
					22,
					58,
					999,
					254,
					338,
					164,
					271,
					76,
					115,
					824,
				},
			},
		},
	},
}
