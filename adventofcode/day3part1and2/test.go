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

func buildMatrix() [][]string {
	//Go compiler doesn’t allow creating variables that are never used. You can fix this by using an _ (underscore) in place of index

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var matrix [][]string

	for scanner.Scan() {
		str := scanner.Text()
		array := strings.Split(str, "")
		matrix = append(matrix, array)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return matrix
}

func countTrees(matrix [][]string, stepRight int, stepDown int) int {
	count := 0
	startingRow := 0
	startingCol := 0

	for startingRow < len(matrix) {
		trees, nextStartingRow, nextStartingCol := countSubTrees(matrix, startingRow, startingCol, stepRight, stepDown)
		count += trees
		startingRow = nextStartingRow
		startingCol = nextStartingCol
		if startingRow > len(matrix) {
			break
		}

	}
	return count
}

func countSubTrees(matrix [][]string, startingRow int, startingCol int, stepRight int, stepDown int) (int, int, int) {
	var next string
	trees := 0

	// while loop alternative
	// length is more than index so is out of bounds

	for startingCol < len(matrix[0]) {
		fmt.Println("Row:", startingRow, "Column", startingCol, "count:", trees)
		// the first "next" is in fact the first cell
		if startingRow >= len(matrix) {
			break
		}
		next = matrix[startingRow][startingCol]
		if next == "#" {
			trees++
		}
		startingRow = startingRow + stepDown
		startingCol = startingCol + stepRight

	}
	// if we have 10 cells in one array and we stop at last cell index=9 then the next cell would be index=12 which doesn't exist, so what we want is the third cell with index=2
	nextStartingCol := startingCol - (len(matrix[0]))
	return trees, startingRow, nextStartingCol

}

//A main function executes by default when you run code in the file.
func main() {
	matrix := buildMatrix()
	//fmt.Println(countTrees(matrix))

	// rule is 3 to right 1 down
	//fmt.Println(len(matrix))
	//fmt.Println(len(matrix[0]))
	fmt.Println(countTrees(matrix, 1, 2))
}
