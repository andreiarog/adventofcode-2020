//Declare a main package (a package is a way to group functions)
package main

//Import the popular fmt package, which contains functions for formatting text, including printing to the console. This package is one of the standard library packages you got when you installed Go.
import (
	"bufio"
	"fmt"
	//"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// GLOBAL VARIABLES
var ()

const ()

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

func checkElements(doc []string) bool {
	var listEntries []string
	for _, entry := range doc {
		strSplit := strings.Split(entry, ":")
		listEntries = append(listEntries, strSplit[0])
	}

	elements := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	checked := true
	for _, element := range elements {
		_, found := Find(listEntries, element)
		if !found {
			checked = false
		}
	}

	return checked
}

func validateElements(doc []string) bool {
	var listEntries []string
	var listFields []string
	for _, entry := range doc {
		strSplit := strings.Split(entry, ":")
		if len(strSplit) == 2 {
			listFields = append(listFields, strSplit[0])
			listEntries = append(listEntries, strSplit[1])
		}
	}

	validCounter := 0
	validArray := []string{}

	for i, field := range listFields {
		entry := listEntries[i]
		if field == "byr" {
			if digit, err := strconv.Atoi(entry); err == nil {
				if len(entry) == 4 && digit > 1919 && digit < 2003 {
					validCounter++
					validArray = append(validArray, "byr")
				}
			}
		}
		if field == "iyr" {
			if digit, err := strconv.Atoi(entry); err == nil {
				if len(entry) == 4 && digit > 2009 && digit < 2021 {
					validCounter++
					validArray = append(validArray, "iyr")
				}
			}

		}
		if field == "eyr" {
			if digit, err := strconv.Atoi(entry); err == nil {
				if len(entry) == 4 && digit > 2019 && digit < 2031 {
					validCounter++
					validArray = append(validArray, "eyr")
				}
			}
		}
		if field == "hgt" {

			reEyr2 := regexp.MustCompile(`[0-9]{2}in$`)
			reEyr1 := regexp.MustCompile(`[0-9]{3}cm$`)
			alreadyChecked := false

			if reEyr1.Match([]byte(entry)) {
				// get the numbers
				re := regexp.MustCompile("[0-9]+")
				digitStr := re.FindAllString(entry, -1)[0]
				digit, _ := strconv.Atoi(digitStr)
				if digit > 149 && digit < 194 {
					validCounter++
					validArray = append(validArray, "hgt1")
					alreadyChecked = true
				}
			}
			if !alreadyChecked {
				if reEyr2.Match([]byte(entry)) {
					// get the numbers
					re := regexp.MustCompile("[0-9]+")
					digitStr := re.FindAllString(entry, -1)[0]
					digit, _ := strconv.Atoi(digitStr)
					if digit > 58 && digit < 77 {
						validCounter++
						validArray = append(validArray, "hgt2")
					}
				}
			}
		}
		if field == "hcl" {
			reHcl := regexp.MustCompile(`#[0-9a-f]{6}$`)
			if reHcl.Match([]byte(entry)) {
				validCounter++
				validArray = append(validArray, "hcl")
			}
		}
		if field == "ecl" {
			validValues := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
			for _, value := range validValues {
				if entry == value {
					validCounter++
					validArray = append(validArray, "ecl")
				}
			}
		}
		if field == "pid" {
			_, err := strconv.Atoi(entry)
			if err == nil {
				if len(entry) == 9 {
					validCounter++
					validArray = append(validArray, "pid")
				}
			}
		}
	}
	fmt.Println("validCounter", validCounter)
	fmt.Println("validArray", validArray)
	valid := false
	if validCounter == 7 {

		valid = true
	}

	return valid
}

func checkRules(doc []string) bool {

	checked := checkElements(doc)
	var valid bool
	if checked {
		valid = validateElements(doc)

	}

	return valid
}

//A main function executes by default when you run code in the file.
func main() {
	var matrix [][]string
	matrix = readLineByLine()
	var count int
	count = 0

	//fmt.Println(len(matrix))

	for i, doc := range matrix {
		fmt.Println(i)
		if checkRules(doc) == true {
			fmt.Println(i)
			count++
		}
	}
	fmt.Println(count)
}
