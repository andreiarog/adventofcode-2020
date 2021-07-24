//Declare a main package (a package is a way to group functions)
package main

//Import the popular fmt package, which contains functions for formatting text, including printing to the console. This package is one of the standard library packages you got when you installed Go.
import (
	"bufio"
	"fmt"
	//"io/ioutil"
	"log"
	"os"
	//"strconv"
	"strings"
)

// GLOBAL VARIABLES
var ()

const ()

// type doc struct {
// 	ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
// 	byr:1937 iyr:2017 cid:147 hgt:183cm
// }

func readLineByLine() [][]string {
	//Go compiler doesn’t allow creating variables that are never used. You can fix this by using an _ (underscore) in place of index

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var matrix [][]string

	i := 0
	for scanner.Scan() {
		i++
		str := scanner.Text()
		strSplit := strings.Split(str, "")
		row := []string{}
		for _, k := range strSplit {
			//fmt.Println("appended to doc on line ", i)
			row = append(row, k)
		}
		matrix = append(matrix, row)
	}
	// // add last one
	// matrix = append(matrix, row)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return matrix
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func getAdjacents(m [][]string, indi int, indj int) []string {

	maxi := len(m) - 1
	max := len(m[0]) - 1

	adjacents := []string{}

	if indi == 0 {
		// add the cell imediately below
		adjacents = append(adjacents, m[indi+1][indj])

		if indj == 0 {
			// add the cell to the right and bottom right
			adjacents = append(adjacents, m[indi][indj+1])
			adjacents = append(adjacents, m[indi+1][indj+1])
		}
		if indj == max {
			// add the cell to the left and bottom left
			adjacents = append(adjacents, m[indi][indj-1])
			adjacents = append(adjacents, m[indi+1][indj-1])
		}
		if indj != 0 && indj != max {
			// add both right, left, bottom right and bottom left
			adjacents = append(adjacents, m[indi][indj+1])
			adjacents = append(adjacents, m[indi+1][indj+1])
			adjacents = append(adjacents, m[indi][indj-1])
			adjacents = append(adjacents, m[indi+1][indj-1])
		}

	}

	if indi == maxi {
		// add the cell imediately above
		adjacents = append(adjacents, m[indi-1][indj])

		if indj == 0 {
			// add the cell to the right and top right
			adjacents = append(adjacents, m[indi][indj+1])
			adjacents = append(adjacents, m[indi-1][indj+1])
		}
		if indj == max {
			// add the cell to the left and top left
			adjacents = append(adjacents, m[indi][indj-1])
			adjacents = append(adjacents, m[indi-1][indj-1])
		}
		if indj != 0 && indj != max {
			// add both right, left, bottom right and bottom left
			adjacents = append(adjacents, m[indi][indj+1])
			adjacents = append(adjacents, m[indi-1][indj+1])
			adjacents = append(adjacents, m[indi][indj-1])
			adjacents = append(adjacents, m[indi-1][indj-1])
		}
	}

	if indi != 0 && indi != maxi {
		// add the cells imediately below and above
		adjacents = append(adjacents, m[indi+1][indj])
		adjacents = append(adjacents, m[indi-1][indj])

		if indj == 0 {
			// add the cells to the right and bottom right and top right
			adjacents = append(adjacents, m[indi][indj+1])
			adjacents = append(adjacents, m[indi+1][indj+1])
			adjacents = append(adjacents, m[indi-1][indj+1])
		}
		if indj == max {
			// add the cell to the left and bottom left and top left
			adjacents = append(adjacents, m[indi][indj-1])
			adjacents = append(adjacents, m[indi+1][indj-1])
			adjacents = append(adjacents, m[indi-1][indj-1])
		}
		if indj != 0 && indj != max {
			// add both right, left, bottom right and bottom left
			adjacents = append(adjacents, m[indi][indj+1])
			adjacents = append(adjacents, m[indi+1][indj+1])
			adjacents = append(adjacents, m[indi-1][indj+1])
			adjacents = append(adjacents, m[indi][indj-1])
			adjacents = append(adjacents, m[indi+1][indj-1])
			adjacents = append(adjacents, m[indi-1][indj-1])
		}
	}

	return adjacents
}

func changeMatrix(m [][]string) (int, [][]string) {
	changedMatrix := [][]string{}
	changes := 0

	for indi := range m {
		changedMatrixIndi := []string{}
		for indj, j := range m[indi] {
			adjacents := getAdjacents(m, indi, indj)
			occupied := 0
			if j == "L" {
				for _, adj := range adjacents {
					if adj == "#" {
						occupied++
					}
				}
				if occupied == 0 {
					changedMatrixIndi = append(changedMatrixIndi, "#")
					changes++
				} else {
					changedMatrixIndi = append(changedMatrixIndi, j)
				}
			}
			if j == "#" {
				for _, adj := range adjacents {
					if adj == "#" {
						occupied++
					}
				}
				if occupied > 3 {
					changedMatrixIndi = append(changedMatrixIndi, "L")
					changes++
				} else {
					changedMatrixIndi = append(changedMatrixIndi, j)
				}

			}
			if j != "L" && j != "#" {
				changedMatrixIndi = append(changedMatrixIndi, j)
			}

		}
		changedMatrix = append(changedMatrix, changedMatrixIndi)

	}

	return changes, changedMatrix
}

func countOccupiedSeats(m [][]string) int {
	occupied := 0

	for indi := range m {
		for _, j := range m[indi] {
			if j == "#" {
				occupied++
			}
		}
	}

	return occupied
}

//A main function executes by default when you run code in the file.
func main() {
	matrix := readLineByLine()
	var changedMatrix [][]string
	var changes int

	//fmt.Println(len(matrix))

	changes, changedMatrix = changeMatrix(matrix)

	for changes > 0 {
		changes, changedMatrix = changeMatrix(changedMatrix)
		fmt.Println(changes)
	}
	//fmt.Println(changes)
	fmt.Println(changedMatrix)
	occupied := countOccupiedSeats(changedMatrix)
	fmt.Println(occupied)

}
