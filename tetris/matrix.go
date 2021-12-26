package main

import "fmt"

const displayValues = ".123456789ABCDEFGHIJKLMNPQRSTUVWXYZ"

type Matrix struct {
	Cols int
	Rows int

	data []int
}

func NewMatrix(cols int, rows int) *Matrix {
	p := new(Matrix)
	p.Cols = cols
	p.Rows = rows
	p.data = make([]int, cols*rows, cols*rows)
	return p
}

func FromSlice(cols int, rows int, s *[]int) *Matrix {
	if len(*s) != cols*rows {
		panic("Matrix size does not match")
	}

	p := NewMatrix(cols, rows)

	for i := 0; i < p.Rows; i++ {
		for j := 0; j < p.Cols; j++ {
			p.data[i*cols+j] = (*s)[i*cols+j]
		}
	}
	return p
}

func (p *Matrix) Copy() *Matrix {
	out := NewMatrix(p.Cols, p.Rows)

	for i := 0; i < p.Rows; i++ {
		for j := 0; j < p.Cols; j++ {
			out.data[i*p.Cols+j] = p.data[i*p.Cols+j]
		}
	}
	return out
}

func (p *Matrix) CopyN(n int) *Matrix {
	out := NewMatrix(p.Cols, p.Rows)

	for i := 0; i < p.Rows; i++ {
		for j := 0; j < p.Cols; j++ {
			if p.data[i*p.Cols+j] != 0 {
				out.data[i*p.Cols+j] = n
			}
		}
	}
	return out
}

func (p *Matrix) Get(col int, row int) int {
	return p.data[row*p.Cols+col]
}

func (p *Matrix) Set(col int, row int, v int) {
	p.data[row*p.Cols+col] = v
}

func (p *Matrix) Equals(other *Matrix) bool {
	if p.Rows != other.Rows || p.Cols != other.Cols {
		return false
	}

	for i, v := range p.data {
		if v != other.data[i] {
			return false
		}
	}
	return true
}

func (p *Matrix) Fits(col int, row int, m *Matrix) bool {
	sizeOk := p.Cols >= col+m.Cols && p.Rows >= row+m.Rows

	if !sizeOk {
		return false
	}

	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			v1 := m.data[i*m.Cols+j]
			v2 := p.data[(i+row)*p.Cols+j+col]
			if v1 != 0 && v2 != 0 {
				return false
			}
		}
	}

	return true
}

func (p *Matrix) Fit(col int, row int, m *Matrix) *Matrix {
	out := p.Copy()

	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			v := m.data[i*m.Cols+j]
			if (v) != 0 {
				out.data[(row+i)*out.Cols+(col+j)] = v
			}
		}
	}

	return out
}

// Rotate clockwise
func (p *Matrix) Rotate() *Matrix {
	out := NewMatrix(p.Rows, p.Cols)

	for i := 0; i < p.Cols; i++ {
		for j := 0; j < p.Rows; j++ {
			out.data[i*p.Rows+j] = p.data[(p.Rows-j-1)*p.Cols+i]
		}
	}
	return out
}

func (p *Matrix) Display() {
	for i := 0; i < p.Rows; i++ {
		for j := 0; j < p.Cols; j++ {
			fmt.Printf("%3c", displayValues[p.data[i*p.Cols+j]])
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func (p *Matrix) Display2() {
	for i := 0; i < p.Rows; i++ {
		for j := 0; j < p.Cols; j++ {
			fmt.Printf("%3d", p.data[i*p.Cols+j])
		}
		fmt.Println("")
	}
	fmt.Println("")
}
