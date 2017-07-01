package math64

import "math"

type PolarVector struct {
	Theta, Phi, Magnitude float64
}

func (pv *PolarVector) ToCartesianVector() Vector {
	return Vector{math.Cos(pv.Phi) * math.Sin(pv.Theta), math.Sin(pv.Phi) * math.Sin(pv.Theta), math.Cos(pv.Theta)}
}
