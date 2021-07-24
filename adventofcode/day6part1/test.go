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

func readLineByLine() [][]string {
	//Go compiler doesn’t allow creating variables that are never used. You can fix this by using an _ (underscore) in place of index

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var doc []string
	var matrix [][]string

	i := 0
	for scanner.Scan() {
		str := scanner.Text()
		strSplit := strings.Split(str, "")
		for _, k := range strSplit {
			//fmt.Println("appended to doc on line ", i)
			doc = append(doc, k)
		}
		if str == "" {
			//fmt.Println("Found empty line on line:", i)
			matrix = append(matrix, doc)
			// fmt.Println("this is what the doc addded looks like:", doc)
			// fmt.Println("this is what the matrix looks like:", matrix)
			//fmt.Println("length matrix", len(matrix))
			// clear doc
			doc = []string{}
		}
		i++
	}
	// add last one
	matrix = append(matrix, doc)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return matrix
}

func collectAnswers(doc []string) int {

	answers := []string{}
	for _, answer := range doc {
		if Find(answers, answer) == -1 {
			answers = append(answers, answer)
		}
	}
	return len(answers)
}

//A main function executes by default when you run code in the file.
func main() {
	var questions int
	var matrix [][]string
	matrix = readLineByLine()

	count := 0
	//idChamp := 0

	for _, group := range matrix {
		questions = collectAnswers(group)
		count += questions
	}

	fmt.Println(count)

}
