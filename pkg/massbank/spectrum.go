package massbank

import (
	"errors"
	"gonum.org/v1/gonum/mat"
)

type MsSpectrum [2][]float64

func NewMsSpectrum(mzs []float64, intensities []float64) (*MsSpectrum, error) {
	var n = len(mzs)
	if n != len(intensities) {
		return nil, errors.New("Arrays mzs and intensities must have the same size")
	}
	var spec MsSpectrum = MsSpectrum{mzs, intensities}
	return &spec, nil
}

func (sp MsSpectrum) Mz() []float64 {
	return sp[0]
}

func (sp MsSpectrum) Intensities() []float64 {
	return sp[1]
}

func (sp MsSpectrum) Length() int {
	return len(sp[0])
}

func (sp *MsSpectrum) Normalize(norm float64) {
	var max = sp[1][0]
	for _, s := range sp[1] {
		if s > max {
			max = s
		}
	}
	for i, _ := range sp[1] {
		sp[1][i] = sp[1][i] / max * norm
	}
}

func (sp *MsSpectrum) Baseline(threshold float64) {
	var result = [2][]float64{}
	for i, _ := range sp[1] {
		if sp[1][i] > threshold {
			result[0] = append(result[0], sp[0][i])
			result[1] = append(result[1], sp[1][i])
		}
	}
}

func (sp MsSpectrum) ToMatrix() *mat.Dense {
	return mat.NewDense(sp.Length(), 2, append(sp[0], sp[1]...))
}
