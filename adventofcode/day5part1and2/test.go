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
	"sort"
	"strings"
)

// GLOBAL VARIABLES
var ()

const ()

func readLineByLine() []string {
	//Go compiler doesn’t allow creating variables that are never used. You can fix this by using an _ (underscore) in place of index

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var ticket string
	var array []string

	i := 0
	for scanner.Scan() {
		i++
		ticket = scanner.Text()
		array = append(array, ticket)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return array
}

func isEven(number int) bool {
	return number%2 == 0
}

func decodeRow(ticket string) int {

	var row int
	max := 127
	min := 0
	inputs := strings.Split(ticket, "")[:7]

	for _, input := range inputs {
		if input == "F" {
			if !isEven(max) {
				max--
			}
			max = max - ((max - min) / 2)
		}
		if input == "B" {
			if !isEven(max) {
				min++
			}
			min = min + ((max - min) / 2)

		}
		//fmt.Println("Round", i, "with input", input, "max is ", max, "min is ", min)

	}

	if max != min {
		fmt.Println("rows is not working well")
	}
	if max == min {
		row = max
	}

	return row
}

func decodeColumn(ticket string) int {

	var col int
	max := 7
	min := 0
	inputs := strings.Split(ticket, "")[7:10]

	for _, input := range inputs {
		if input == "L" {
			if !isEven(max) {
				max--
			}
			max = max - ((max - min) / 2)
		}
		if input == "R" {
			if !isEven(max) {
				min++
			}
			min = min + ((max - min) / 2)

		}
		//fmt.Println("Round", i, "with input", input, "max is ", max, "min is ", min)
	}

	if max != min {
		fmt.Println("Columns is not working well")
	}
	if max == min {
		col = max
	}

	return col
}

func Insert(ss []int, s int) []int {
	i := sort.SearchInts(ss, s)
	fmt.Println(i)
	ss = append(ss, 0)
	fmt.Println(ss)
	copy(ss[i+1:], ss[i:])
	fmt.Println(ss)
	ss[i] = s
	return ss
}

func decodeID(ticket string) int {

	row := decodeRow(ticket)
	col := decodeColumn(ticket)
	id := row*8 + col

	//fmt.Println(row)
	//fmt.Println(col)

	return id
}

//A main function executes by default when you run code in the file.
func main() {
	var sortedArray []int
	var array []string
	array = readLineByLine()

	var id int
	//idChamp := 0

	for _, ticket := range array {
		id = decodeID(ticket)
		sortedArray = Insert(sortedArray, id)
	}

	var myTicket int
	for index := range sortedArray[1:] {
		if sortedArray[index]+1 != sortedArray[index+1] {
			myTicket = sortedArray[index] + 1
		}
	}

	fmt.Println(myTicket)

}
