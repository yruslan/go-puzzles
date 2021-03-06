package main

import "fmt"

//
// Elements:
//           0.   * * * *    *                       1.   * *       * *
//                           *                            * *       * *
//                           *
//                           *
//
//           2.   * * *    * *         *    *        3.  *          * *    * * *      *
//                *          *     * * *    *            * * *      *          *      *
//                           *              * *                     *               * *
//
//           4.     * *    *                         5.  * *          *
//                * *      * *                             * *      * *
//                           *                                      *
//
//           6.   * * *      *      *       *
//                  *      * *    * * *     * *
//                           *              *
func main() {
	// Gehenna - entering the hidden world
	//ch := Unify(SolvePuzzle(5, 8, []int{0, 0, 2, 4, 4, 4, 6, 6, 6, 6}))

	// Gehenna - free Admin
	ch := Unify(SolvePuzzle(7, 4, []int{0, 2, 2, 3, 3, 6, 6}))

	i := 0

	for solution := range ch {
		i++
		fmt.Printf("Solution %d:\n", i)
		solution.Display()
	}
}
