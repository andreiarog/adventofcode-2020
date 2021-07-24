//Declare a main package (a package is a way to group functions)
package main

//Import the popular fmt package, which contains functions for formatting text, including printing to the console. This package is one of the standard library packages you got when you installed Go.
import (
	"bufio"
	"fmt"
	//"io/ioutil"
	//"github.com/yourbasic/graph"
	"log"
	"os"
	//"regexp"
	//"sort"
	"strconv"
	//"strings"
)

func readLineByLine() []int {
	//Go compiler doesn’t allow creating variables that are never used. You can fix this by using an _ (underscore) in place of index

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	array := []int{}

	for scanner.Scan() {
		str := scanner.Text()
		//strSplit := strings.Split(str, " ")
		digit, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		array = append(array, digit)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return array
}

func isSumFound(slice []int, value int) bool {

	sumFound := false

	for _, first := range slice {
		for _, second := range slice {
			sum := first + second
			if sum == value {
				sumFound = true
				break
			}
		}
	}

	return sumFound

}

//A main function executes by default when you run code in the file.
func main() {
	array := readLineByLine()

	//instructionsSwitched := make(map[int]instruction)

	preemble := []int{}
	//start is inclusive
	start := 0
	//end is not inclusive
	end := 25
	var sumFound bool
	var result int

	for _, entry := range array[25:] {
		preemble = array[start:end]
		sumFound = isSumFound(preemble, entry)
		if !sumFound {
			result = entry
			break
		}
		if sumFound {
			start++
			end++
		}

	}

	fmt.Println(result)
	//fmt.Println(g, nodesMap)
}
