//Declare a main package (a package is a way to group functions)
package main

//Import the popular fmt package, which contains functions for formatting text, including printing to the console. This package is one of the standard library packages you got when you installed Go.
import (
	"bufio"
	"fmt"
	//"io/ioutil"
	"log"
	"os"
	// "regexp"
	// "strconv"
	//"sort"
	"strings"
)

// GLOBAL VARIABLES
var ()

const ()

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) int {
	for i, item := range slice {
		if item == val {
			return i
		}
	}
	return -1
}

func delete(slice []string, val string) []string {
	// swap the element to delete with the one at the end of the slice and then return the n-1 first elements
	p := Find(slice, val)
	if p == -1 {
		return slice
	}
	//copy the last element of the slice into the position to be deleted
	slice[p] = slice[len(slice)-1]
	//return all but the last entry which was copied
	return slice[:len(slice)-1]
}

func readLineByLine() [][]string {
	//Go compiler doesn’t allow creating variables that are never used. You can fix this by using an _ (underscore) in place of index

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var group []string
	var matrix [][]string

	i := 0
	for scanner.Scan() {
		str := scanner.Text()
		if str != "" {
			group = append(group, str)
		}
		if str == "" {
			matrix = append(matrix, group)

			// clear doc
			group = []string{}
		}
		i++
	}
	// add last one
	matrix = append(matrix, group)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return matrix
}

func interceptSlices(firstSlice []string, secondSlice []string) []string {
	resultSlice := []string{}
	checkMap := map[string]struct{}{}

	for _, addr := range firstSlice {
		checkMap[addr] = struct{}{}
	}
	for _, addr := range secondSlice {
		//if the entry of second slice exists in first slice, add it to result slice
		if _, ok := checkMap[addr]; ok {
			resultSlice = append(resultSlice, addr)
		}
	}

	return resultSlice
}

func collectAnswers(group []string) int {

	// let's start with first person in group
	groupAnswers := strings.Split(group[0], "")

	// now loop through other people's answers
	for _, answersStr := range group[1:] {
		answers := strings.Split(answersStr, "")
		groupAnswers = interceptSlices(groupAnswers, answers)
	}

	return len(groupAnswers)
}

//A main function executes by default when you run code in the file.
func main() {
	var answers int
	var matrix [][]string
	matrix = readLineByLine()
	//fmt.Println(matrix[0][0])
	count := 0
	//idChamp := 0

	for _, group := range matrix {
		answers = collectAnswers(group)
		count += answers
	}

	fmt.Println(count)

}
