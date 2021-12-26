package main

import (
	"reflect"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	m := NewMatrix(2, 3)

	if m.Cols != 2 {
		t.Errorf("Matrix columns does not match. Expected: %d, actual: %d", 2, m.Cols)
	}

	if m.Rows != 3 {
		t.Errorf("Matrix rows does not match. Expected: %d, actual: %d", 3, m.Rows)
	}

	l := len(m.data)

	if l != 6 {
		t.Errorf("Matrix slice is incorrect. Expected: %d, actual: %d", 6, l)
	}
}

func TestInitMatrix(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}

	m := FromSlice(2, 3, &s)

	if m.Cols != 2 {
		t.Errorf("Matrix columns does not match. Expected: %d, actual: %d", 2, m.Cols)
	}

	if m.Rows != 3 {
		t.Errorf("Matrix rows does not match. Expected: %d, actual: %d", 3, m.Rows)
	}

	l := len(m.data)

	if l != 6 {
		t.Errorf("Matrix slice is incorrect. Expected: %d, actual: %d", 6, l)
	}

	if !reflect.DeepEqual(s, m.data) {
		t.Errorf("Matrix content is incorrect. Expected: %d, actual: %d", 6, l)
	}
}

func TestCopyMatrix(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}

	m1 := FromSlice(2, 3, &s)
	m2 := m1.Copy()

	if !m1.Equals(m2) {
		t.Errorf("Should be equal")
	}
}

func TestCopyNMatrix(t *testing.T) {
	s := []int{0, 1, 2, 0, 4, 5}

	m1 := FromSlice(2, 3, &s)
	m2 := m1.CopyN(9)

	mExpected := FromSlice(2, 3, &[]int{0, 9, 9, 0, 9, 9})

	if !mExpected.Equals(m2) {
		m2.Display()

		mExpected.Display()
		t.Errorf("Should be equal")
	}
}

func TestEqualsMatrix(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	m1 := FromSlice(2, 3, &s)
	m2 := FromSlice(2, 3, &s)

	if !m1.Equals(m2) {
		t.Errorf("Expected equal matricies")
	}

	m1.Set(2, 1, 7)
	if m1.Equals(m2) {
		t.Errorf("Expected non-equal matricies")
	}
}

func TestEqualsMatrix2(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	m1 := FromSlice(2, 3, &s)
	m2 := FromSlice(3, 2, &s)

	if m1.Equals(m2) {
		t.Errorf("Expected non-equal matricies")
	}
}

func TestGetSet(t *testing.T) {
	m := NewMatrix(2, 3)

	v := []struct {
		c int
		r int
		v int
	}{
		{0, 0, 1},
		{1, 0, 2},
		{0, 1, 3},
		{1, 1, 4},
		{0, 2, 5},
		{1, 2, 6},
	}

	for _, v := range v {
		m.Set(v.c, v.r, v.v)

		if m.Get(v.c, v.r) != v.v {
			t.Errorf("Unexpected matrix get/set (%d %d %d)", v.c, v.r, v.v)
		}
	}

}

func TestRotation1(t *testing.T) {
	s1 := []int{1, 0, 1, 0, 1, 1}
	m1 := FromSlice(2, 3, &s1)

	m2 := m1.Rotate()
	mExpected := FromSlice(3, 2, &[]int{1, 1, 1, 1, 0, 0})

	if !m2.Equals(mExpected) {
		t.Errorf("Expected equal")
	}
}

func TestRotation2(t *testing.T) {
	s1 := []int{1, 0, 1, 0, 1, 1}
	m1 := FromSlice(2, 3, &s1)

	m2 := m1.Rotate().Rotate()
	mExpected := FromSlice(2, 3, &[]int{1, 1, 0, 1, 0, 1})

	if !m2.Equals(mExpected) {
		t.Errorf("Expected equal")
	}
}

func TestRotation3(t *testing.T) {
	s1 := []int{1, 0, 1, 0, 1, 1}
	m1 := FromSlice(2, 3, &s1)

	m2 := m1.Rotate().Rotate().Rotate()
	mExpected := FromSlice(3, 2, &[]int{0, 0, 1, 1, 1, 1})

	if !m2.Equals(mExpected) {
		t.Errorf("Expected equal")
	}
}

func TestFits(t *testing.T) {
	s1 := []int{1, 0, 1, 0, 1, 1}
	m1 := FromSlice(2, 3, &s1)

	b1 := NewMatrix(4, 4)

	if !b1.Fits(0, 0, m1) {
		t.Errorf("Should fit")
	}

	if !b1.Fits(0, 1, m1) {
		t.Errorf("Should fit")
	}

	if b1.Fits(0, 2, m1) {
		t.Errorf("Should not fit")
	}

	b1.Set(1, 0, 2)

	if !b1.Fits(0, 0, m1) {
		t.Errorf("Should fit")
	}

	if b1.Fits(1, 0, m1) {
		t.Errorf("Should not fit")
	}
}

func TestFits2(t *testing.T) {
	s1 := []int{0, 0, 1, 1, 1, 1}
	m1 := FromSlice(3, 2, &s1)

	b1 := NewMatrix(4, 4)

	//println(b1.Rows - m1.Rows + 1)
	//println(b1.Cols - m1.Cols + 1)

	if !b1.Fits(1, 2, m1) {
		t.Errorf("Should fit")
	}

	//b2 := b1.Fit(1, 2, m1)

	//b2.Display()
}

func TestFit(t *testing.T) {
	s1 := []int{1, 0, 1, 0, 1, 1}

	m1 := FromSlice(2, 3, &s1)
	//m2 := m1.Rotate().Rotate()

	b1 := NewMatrix(4, 4)

	b2 := b1.Fit(0, 0, m1)

	expected1 := FromSlice(4, 4, &[]int{1, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0})

	if !b2.Equals(expected1) {
		t.Errorf("Should be equal")
	}

	if b2.Fits(1, 0, m1) {
		t.Errorf("Should not fit")
	}

	b3 := b2.Fit(1, 1, m1.CopyN(2).Rotate().Rotate())

	expected2 := FromSlice(4, 4, &[]int{1, 0, 0, 0, 1, 2, 2, 0, 1, 1, 2, 0, 0, 0, 2, 0})

	if !b3.Equals(expected2) {
		b3.Display()
		expected2.Display()

		t.Errorf("Should be equal")
	}
}
