package main

import (
	"sync"
)

func solveMatrix(m *Matrix, elements []int, currentElement int, solutionChan chan *Matrix, wg *sync.WaitGroup) bool {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	if len(elements) == 0 {
		solutionChan <- m
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
						if currentElement < 4 {
							wg.Add(1)
							go solveMatrix(newM, elements[1:], currentElement+1, solutionChan, wg)
						} else {
							solveMatrix(newM, elements[1:], currentElement+1, solutionChan, nil)
						}
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

func SolvePuzzle(cols int, rows int, elements []int) chan *Matrix {
	var wg sync.WaitGroup
	solutionChan := make(chan *Matrix, 2)
	m0 := NewMatrix(cols, rows)
	wg.Add(1)
	go solveMatrix(m0, elements, 1, solutionChan, &wg)
	go func() {
		wg.Wait()
		close(solutionChan)
	}()
	return solutionChan
}

func contains(m *Matrix, arr []*Matrix) bool {
	for _, v := range arr {
		if m.Equals(v) {
			return true
		}
	}
	return false
}

func Unify(ch chan *Matrix) chan *Matrix {
	outputChan := make(chan *Matrix, 2)

	alreadyProcessed := make([]*Matrix, 0, 1000)

	go func() {
		for matrix := range ch {
			regularized := matrix.Regularize()

			if !contains(regularized, alreadyProcessed) {
				alreadyProcessed = append(alreadyProcessed, regularized)
				outputChan <- regularized
			}
		}
		close(outputChan)
	}()

	return outputChan
}
