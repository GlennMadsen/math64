package math64

type Matrix struct {
	Matrix [4][4]float64
}

func (m *Matrix) Transposed() Matrix {
	return Matrix{[4][4]float64{
		{m.Matrix[0][0], m.Matrix[1][0], m.Matrix[2][0], m.Matrix[3][0]},
		{m.Matrix[0][1], m.Matrix[1][1], m.Matrix[2][1], m.Matrix[3][1]},
		{m.Matrix[0][2], m.Matrix[1][2], m.Matrix[2][2], m.Matrix[3][2]},
		{m.Matrix[0][3], m.Matrix[1][3], m.Matrix[2][3], m.Matrix[3][3]},
	}}
}

func (m *Matrix) Transpose() {
	m.Matrix = [4][4]float64{
		{m.Matrix[0][0], m.Matrix[1][0], m.Matrix[2][0], m.Matrix[3][0]},
		{m.Matrix[0][1], m.Matrix[1][1], m.Matrix[2][1], m.Matrix[3][1]},
		{m.Matrix[0][2], m.Matrix[1][2], m.Matrix[2][2], m.Matrix[3][2]},
		{m.Matrix[0][3], m.Matrix[1][3], m.Matrix[2][3], m.Matrix[3][3]},
	}
}

func (a *Matrix) MatrixProduct(b Matrix) Matrix {
	product := Matrix{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			product.Matrix[i][j] = a.Matrix[i][0]*b.Matrix[0][j] +
				a.Matrix[i][1]*b.Matrix[1][j] +
				a.Matrix[i][2]*b.Matrix[2][j] +
				a.Matrix[i][3]*b.Matrix[3][j]
		}
	}
	return product
}

func (m *Matrix) VectorProduct(v Vector) Vector {
	product := Vector{
		v.X*m.Matrix[0][0] + v.Y*m.Matrix[1][0] + v.Z*m.Matrix[2][0] + m.Matrix[3][0],
		v.X*m.Matrix[0][1] + v.Y*m.Matrix[1][1] + v.Z*m.Matrix[2][1] + m.Matrix[3][1],
		v.X*m.Matrix[0][2] + v.Y*m.Matrix[1][2] + v.Z*m.Matrix[2][2] + m.Matrix[3][2]}
	w := v.X*m.Matrix[0][3] + v.Y*m.Matrix[1][3] + v.Z*m.Matrix[2][3] + m.Matrix[3][3]
	if w != 1 && w != 0 {
		product.X /= w
		product.Y /= w
		product.Z /= w
	}
	return product
}

func (m *Matrix) ScalarProduct(scalar float64) Matrix {
	return Matrix{[4][4]float64{
		{m.Matrix[0][0] * scalar, m.Matrix[0][1] * scalar, m.Matrix[0][2] * scalar, m.Matrix[0][3] * scalar},
		{m.Matrix[1][0] * scalar, m.Matrix[1][1] * scalar, m.Matrix[1][2] * scalar, m.Matrix[1][3] * scalar},
		{m.Matrix[2][0] * scalar, m.Matrix[2][1] * scalar, m.Matrix[2][2] * scalar, m.Matrix[2][3] * scalar},
		{m.Matrix[3][0] * scalar, m.Matrix[3][1] * scalar, m.Matrix[3][2] * scalar, m.Matrix[3][3] * scalar},
	}}
}

func (m *Matrix) Inverse() (bool, Matrix) {
	determinant :=
		m.Matrix[0][0]*m.Matrix[1][1]*m.Matrix[2][2]*m.Matrix[3][3] +
			m.Matrix[0][0]*m.Matrix[1][2]*m.Matrix[2][3]*m.Matrix[3][1] +
			m.Matrix[0][0]*m.Matrix[1][3]*m.Matrix[2][1]*m.Matrix[3][2] +

			m.Matrix[0][1]*m.Matrix[1][0]*m.Matrix[2][3]*m.Matrix[3][2] +
			m.Matrix[0][1]*m.Matrix[1][2]*m.Matrix[2][0]*m.Matrix[3][3] +
			m.Matrix[0][1]*m.Matrix[1][3]*m.Matrix[2][2]*m.Matrix[3][0] +

			m.Matrix[0][2]*m.Matrix[1][0]*m.Matrix[2][1]*m.Matrix[3][3] +
			m.Matrix[0][2]*m.Matrix[1][1]*m.Matrix[2][3]*m.Matrix[3][0] +
			m.Matrix[0][2]*m.Matrix[1][3]*m.Matrix[2][0]*m.Matrix[3][1] +

			m.Matrix[0][3]*m.Matrix[1][0]*m.Matrix[2][2]*m.Matrix[3][1] +
			m.Matrix[0][3]*m.Matrix[1][1]*m.Matrix[2][0]*m.Matrix[3][2] +
			m.Matrix[0][3]*m.Matrix[1][2]*m.Matrix[2][1]*m.Matrix[3][0] -

			m.Matrix[0][0]*m.Matrix[1][1]*m.Matrix[2][3]*m.Matrix[3][2] -
			m.Matrix[0][0]*m.Matrix[1][2]*m.Matrix[2][1]*m.Matrix[3][3] -
			m.Matrix[0][0]*m.Matrix[1][3]*m.Matrix[2][2]*m.Matrix[3][1] -

			m.Matrix[0][1]*m.Matrix[1][0]*m.Matrix[2][2]*m.Matrix[3][3] -
			m.Matrix[0][1]*m.Matrix[1][2]*m.Matrix[2][3]*m.Matrix[3][0] -
			m.Matrix[0][1]*m.Matrix[1][3]*m.Matrix[2][0]*m.Matrix[3][2] -

			m.Matrix[0][2]*m.Matrix[1][0]*m.Matrix[2][3]*m.Matrix[3][1] -
			m.Matrix[0][2]*m.Matrix[1][1]*m.Matrix[2][1]*m.Matrix[3][3] -
			m.Matrix[0][2]*m.Matrix[1][3]*m.Matrix[2][1]*m.Matrix[3][0] -

			m.Matrix[0][3]*m.Matrix[1][0]*m.Matrix[2][1]*m.Matrix[3][2] -
			m.Matrix[0][3]*m.Matrix[1][1]*m.Matrix[2][2]*m.Matrix[3][0] -
			m.Matrix[0][3]*m.Matrix[1][2]*m.Matrix[2][0]*m.Matrix[3][1]
	if determinant == 0 {
		return false, Matrix{}
	}
	b := Matrix{[4][4]float64{
		{
			m.Matrix[1][1]*m.Matrix[2][2]*m.Matrix[3][3] +
				m.Matrix[1][2]*m.Matrix[2][3]*m.Matrix[3][1] +
				m.Matrix[1][3]*m.Matrix[2][1]*m.Matrix[3][2] -
				m.Matrix[1][1]*m.Matrix[2][3]*m.Matrix[3][2] -
				m.Matrix[1][2]*m.Matrix[2][1]*m.Matrix[3][3] -
				m.Matrix[1][3]*m.Matrix[2][2]*m.Matrix[3][1],

			m.Matrix[0][1]*m.Matrix[2][3]*m.Matrix[3][2] +
				m.Matrix[0][2]*m.Matrix[2][1]*m.Matrix[3][3] +
				m.Matrix[0][3]*m.Matrix[2][2]*m.Matrix[3][1] -
				m.Matrix[0][1]*m.Matrix[2][2]*m.Matrix[3][3] -
				m.Matrix[0][2]*m.Matrix[2][3]*m.Matrix[3][1] -
				m.Matrix[0][3]*m.Matrix[2][1]*m.Matrix[3][2],

			m.Matrix[0][1]*m.Matrix[1][2]*m.Matrix[3][3] +
				m.Matrix[0][2]*m.Matrix[1][3]*m.Matrix[3][1] +
				m.Matrix[0][3]*m.Matrix[1][1]*m.Matrix[3][2] -
				m.Matrix[0][1]*m.Matrix[1][3]*m.Matrix[3][2] -
				m.Matrix[0][2]*m.Matrix[1][1]*m.Matrix[3][3] -
				m.Matrix[0][3]*m.Matrix[1][2]*m.Matrix[3][1],

			m.Matrix[0][1]*m.Matrix[1][3]*m.Matrix[2][2] +
				m.Matrix[0][2]*m.Matrix[1][1]*m.Matrix[2][3] +
				m.Matrix[0][3]*m.Matrix[1][2]*m.Matrix[2][1] -
				m.Matrix[0][1]*m.Matrix[1][2]*m.Matrix[2][3] -
				m.Matrix[0][2]*m.Matrix[1][3]*m.Matrix[1][3] -
				m.Matrix[0][3]*m.Matrix[1][1]*m.Matrix[2][2],
		},
		{
			m.Matrix[1][0]*m.Matrix[2][3]*m.Matrix[3][2] +
				m.Matrix[1][2]*m.Matrix[2][0]*m.Matrix[3][3] +
				m.Matrix[1][3]*m.Matrix[2][2]*m.Matrix[3][0] -
				m.Matrix[1][0]*m.Matrix[2][2]*m.Matrix[3][3] -
				m.Matrix[1][2]*m.Matrix[2][3]*m.Matrix[3][0] -
				m.Matrix[1][3]*m.Matrix[2][0]*m.Matrix[3][2],

			m.Matrix[0][0]*m.Matrix[2][2]*m.Matrix[3][3] +
				m.Matrix[0][2]*m.Matrix[2][3]*m.Matrix[3][0] +
				m.Matrix[0][3]*m.Matrix[2][0]*m.Matrix[3][2] -
				m.Matrix[0][0]*m.Matrix[2][3]*m.Matrix[3][2] -
				m.Matrix[0][2]*m.Matrix[2][0]*m.Matrix[3][3] -
				m.Matrix[0][3]*m.Matrix[2][2]*m.Matrix[3][0],

			m.Matrix[0][0]*m.Matrix[1][3]*m.Matrix[3][2] +
				m.Matrix[0][2]*m.Matrix[1][0]*m.Matrix[3][3] +
				m.Matrix[0][3]*m.Matrix[1][2]*m.Matrix[3][0] -
				m.Matrix[0][0]*m.Matrix[1][2]*m.Matrix[3][3] -
				m.Matrix[0][2]*m.Matrix[1][3]*m.Matrix[3][0] -
				m.Matrix[0][3]*m.Matrix[1][0]*m.Matrix[3][2],

			m.Matrix[0][0]*m.Matrix[1][2]*m.Matrix[2][3] +
				m.Matrix[0][2]*m.Matrix[1][3]*m.Matrix[2][0] +
				m.Matrix[0][3]*m.Matrix[1][0]*m.Matrix[2][2] -
				m.Matrix[0][0]*m.Matrix[1][3]*m.Matrix[2][2] -
				m.Matrix[0][2]*m.Matrix[1][0]*m.Matrix[2][3] -
				m.Matrix[0][3]*m.Matrix[1][2]*m.Matrix[2][0],
		},
		{
			m.Matrix[1][0]*m.Matrix[2][1]*m.Matrix[3][3] +
				m.Matrix[1][1]*m.Matrix[2][3]*m.Matrix[3][0] +
				m.Matrix[1][3]*m.Matrix[2][0]*m.Matrix[3][1] -
				m.Matrix[1][0]*m.Matrix[2][3]*m.Matrix[3][1] -
				m.Matrix[1][1]*m.Matrix[2][0]*m.Matrix[3][3] -
				m.Matrix[1][3]*m.Matrix[2][1]*m.Matrix[3][0],

			m.Matrix[0][0]*m.Matrix[2][3]*m.Matrix[3][1] +
				m.Matrix[0][1]*m.Matrix[2][0]*m.Matrix[3][3] +
				m.Matrix[0][3]*m.Matrix[2][1]*m.Matrix[3][0] -
				m.Matrix[0][0]*m.Matrix[2][1]*m.Matrix[3][3] -
				m.Matrix[0][1]*m.Matrix[2][3]*m.Matrix[3][0] -
				m.Matrix[0][3]*m.Matrix[2][0]*m.Matrix[3][1],

			m.Matrix[0][0]*m.Matrix[1][1]*m.Matrix[3][3] +
				m.Matrix[0][1]*m.Matrix[1][3]*m.Matrix[3][0] +
				m.Matrix[0][3]*m.Matrix[1][0]*m.Matrix[3][1] -
				m.Matrix[0][0]*m.Matrix[1][3]*m.Matrix[3][1] -
				m.Matrix[0][1]*m.Matrix[1][0]*m.Matrix[3][3] -
				m.Matrix[0][3]*m.Matrix[1][1]*m.Matrix[3][0],

			m.Matrix[0][0]*m.Matrix[1][3]*m.Matrix[2][1] +
				m.Matrix[0][1]*m.Matrix[1][0]*m.Matrix[2][3] +
				m.Matrix[0][3]*m.Matrix[1][1]*m.Matrix[2][0] -
				m.Matrix[0][0]*m.Matrix[1][1]*m.Matrix[2][3] -
				m.Matrix[0][1]*m.Matrix[1][3]*m.Matrix[2][0] -
				m.Matrix[0][3]*m.Matrix[1][0]*m.Matrix[2][1],
		},
		{
			m.Matrix[1][0]*m.Matrix[2][2]*m.Matrix[3][1] +
				m.Matrix[1][1]*m.Matrix[2][0]*m.Matrix[3][2] +
				m.Matrix[1][2]*m.Matrix[2][1]*m.Matrix[3][0] -
				m.Matrix[1][0]*m.Matrix[2][1]*m.Matrix[3][2] -
				m.Matrix[1][1]*m.Matrix[2][2]*m.Matrix[3][0] -
				m.Matrix[1][2]*m.Matrix[2][0]*m.Matrix[3][1],

			m.Matrix[0][0]*m.Matrix[2][1]*m.Matrix[3][2] +
				m.Matrix[0][1]*m.Matrix[2][2]*m.Matrix[3][0] +
				m.Matrix[0][2]*m.Matrix[2][0]*m.Matrix[3][1] -
				m.Matrix[0][0]*m.Matrix[2][2]*m.Matrix[3][1] -
				m.Matrix[0][1]*m.Matrix[2][0]*m.Matrix[3][2] -
				m.Matrix[0][2]*m.Matrix[2][1]*m.Matrix[3][0],

			m.Matrix[0][0]*m.Matrix[1][2]*m.Matrix[3][1] +
				m.Matrix[0][1]*m.Matrix[1][0]*m.Matrix[3][2] +
				m.Matrix[0][2]*m.Matrix[1][1]*m.Matrix[3][0] -
				m.Matrix[0][0]*m.Matrix[1][1]*m.Matrix[3][2] -
				m.Matrix[0][1]*m.Matrix[1][2]*m.Matrix[3][0] -
				m.Matrix[0][2]*m.Matrix[1][0]*m.Matrix[3][1],

			m.Matrix[0][0]*m.Matrix[1][1]*m.Matrix[2][2] +
				m.Matrix[0][1]*m.Matrix[1][2]*m.Matrix[2][0] +
				m.Matrix[0][2]*m.Matrix[1][0]*m.Matrix[2][1] -
				m.Matrix[0][0]*m.Matrix[1][2]*m.Matrix[2][1] -
				m.Matrix[0][1]*m.Matrix[1][0]*m.Matrix[2][2] -
				m.Matrix[0][2]*m.Matrix[1][1]*m.Matrix[2][0],
		},
	}}
	return true, b.ScalarProduct(1 / determinant)
}
