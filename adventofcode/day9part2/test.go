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

func isSumFound(slice []int) (bool, []int) {

	goal := 530627549
	sum := 0
	sumFound := false
	subslice := []int{}

	for ind, value := range slice {
		sum += value
		if sum == goal {
			sumFound = true
			subslice = slice[:ind+1]
			break
		}
		if sum > goal {
			break
		}
	}

	return sumFound, subslice

}

func minIntSlice(slice []int) int {
	var m int

	for i, v := range slice {
		if i == 0 || v < m {
			m = v
		}
	}

	return m
}

func maxIntSlice(slice []int) int {
	var m int

	for i, v := range slice {
		if i == 0 || v > m {
			m = v
		}
	}

	return m
}

//A main function executes by default when you run code in the file.
func main() {
	array := readLineByLine()

	//instructionsSwitched := make(map[int]instruction)

	preemble := []int{}
	var preembleSlice []int
	//start is inclusive
	start := 0
	var sumFound bool
	var result int

	i := 0

	for i > -1 {
		preemble = array[start:]
		sumFound, preembleSlice = isSumFound(preemble)
		if sumFound {
			result = maxIntSlice(preembleSlice) + minIntSlice(preembleSlice)
			break
		}
		if !sumFound {
			start++
			if start > len(array)-1 {
				fmt.Println("No sum found")
				break
			}
		}

	}

	fmt.Println(result)
	//fmt.Println(g, nodesMap)
}
