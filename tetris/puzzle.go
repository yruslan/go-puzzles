package main

import "fmt"

func solveMatrix(m *Matrix, elements []int, currentElement int) bool {
	if len(elements) == 0 {
		fmt.Println("Found solution!")
		m.Display()
		return true
	} else {
		//m.Display()
		r := getAllRotations(Elements[elements[0]].CopyN(currentElement))

		for _, e := range r {
			//e.CopyN(currentElement).Display()
			for i := 0; i < m.Rows-e.Rows+1; i++ {
				for j := 0; j < m.Cols-e.Cols+1; j++ {
					if m.Fits(j, i, e) {
						newM := m.Fit(j, i, e)
						solveMatrix(newM, elements[1:], currentElement+1)
					}
				}

			}
		}

	}
	return false
}

func getAllRotations(m *Matrix) []*Matrix {
	r := make([]*Matrix, 0, 4)

	r = append(r, m)

	m1 := m.Rotate()
	r = addIfNotEqual(r, m1)

	m2 := m1.Rotate()
	r = addIfNotEqual(r, m2)

	m3 := m2.Rotate()
	r = addIfNotEqual(r, m3)

	return r
}

func addIfNotEqual(r []*Matrix, m *Matrix) []*Matrix {
	found := false

	for _, v := range r {
		if v.Equals(m) {
			found = true
		}
	}

	if !found {
		return append(r, m)
	}
	return r
}

func SolvePuzzle(cols int, rows int, elements []int) {
	m0 := NewMatrix(cols, rows)
	solveMatrix(m0, elements, 1)
}
