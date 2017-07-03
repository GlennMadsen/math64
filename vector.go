package math64

import (
	"math"
)

type Vector struct {
	X, Y, Z float64
}

func (v *Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vector) Normal() Vector {
	inverseLength := 1 / v.Magnitude()
	return Vector{inverseLength * v.X, inverseLength * v.Y, inverseLength * v.Z}
}

func (v *Vector) ToPolarVector() PolarVector {
	normal := v.Normal()
	phi := math.Atan2(normal.Y, normal.X)
	if phi < 0 {
		phi = phi + 2*math.Pi
	}
	return PolarVector{math.Acos(v.Normal().Z), phi, v.Magnitude()}
}

func (a *Vector) DotProduct(b *Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (a *Vector) CrossProduct(b *Vector) Vector {
	return Vector{a.Y*b.Z - a.Z*b.Y, a.Z*b.X - a.X*b.Z, a.X*b.Y - a.Y*b.X}
}

func (a *Vector) ScalarMultiply(b float64) Vector {
	return Vector{a.X * b, a.Y * b, a.Z * b}
}

func (a *Vector) ScalarAddition(b float64) Vector {
	return Vector{a.X + b, a.Y + b, a.Z + b}
}

func (a *Vector) VectorAddition(b *Vector) Vector {
	return Vector{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func (a *Vector) Negation() Vector {
	return Vector{-a.X, -a.Y, -a.Z}
}

func (a *Vector) VectorSubtraction(b *Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}
