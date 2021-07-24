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

	var doc []string
	var matrix [][]string

	i := 0
	for scanner.Scan() {
		i++
		str := scanner.Text()
		strSplit := strings.Split(str, " ")
		for _, k := range strSplit {
			//fmt.Println("appended to doc on line ", i)
			doc = append(doc, k)
		}
		if str == "" {
			fmt.Println("Found empty line on line:", i)
			matrix = append(matrix, doc)
			// fmt.Println("this is what the doc addded looks like:", doc)
			// fmt.Println("this is what the matrix looks like:", matrix)
			//fmt.Println("length matrix", len(matrix))
			// clear doc
			doc = []string{}
		}
	}
	// add last one
	matrix = append(matrix, doc)

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

func checkRules(doc []string) bool {
	//fmt.Println("doc", doc)
	var listEntries []string
	for _, entry := range doc {
		strSplit := strings.Split(entry, ":")
		//fmt.Println("strSplit", strSplit[0])
		listEntries = append(listEntries, strSplit[0])
	}
	//fmt.Println("listEntries", listEntries)

	elements := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	valid := true
	for _, element := range elements {
		_, found := Find(listEntries, element)
		if !found {
			valid = false
		}
	}

	fmt.Println("result", valid)

	return valid
}

//A main function executes by default when you run code in the file.
func main() {
	var matrix [][]string
	matrix = readLineByLine()
	var count int
	count = 0

	fmt.Println(matrix)
	//fmt.Println(len(matrix))

	for _, doc := range matrix {

		if checkRules(doc) == true {
			count++
		}
	}
	fmt.Println(count)
}
