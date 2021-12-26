package main

var Elements = []*Matrix{
	// 0.
	// ****
	FromSlice(4, 1, &[]int{1, 1, 1, 1}),

	// 1.
	// **
	// **
	FromSlice(2, 2, &[]int{1, 1, 1, 1}),

	// 2.
	// ***
	// *
	FromSlice(3, 2, &[]int{1, 1, 1, 1, 0, 0}),

	// 3.
	// *
	// ***
	FromSlice(3, 2, &[]int{1, 0, 0, 1, 1, 1}),

	// 4.
	//  **
	// **
	FromSlice(3, 2, &[]int{0, 1, 1, 1, 1, 0}),

	// 5.
	// **
	//  **
	FromSlice(3, 2, &[]int{1, 1, 0, 0, 1, 1}),

	// 6.
	// ***
	//  *
	FromSlice(3, 2, &[]int{1, 1, 1, 0, 1, 0}),
}
