//Declare a main package (a package is a way to group functions)
package main

//Import the popular fmt package, which contains functions for formatting text, including printing to the console. This package is one of the standard library packages you got when you installed Go.
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// GLOBAL VARIABLES
var ()

const ()

func readFullFile() {
	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
	fmt.Println(len(content))

}

func readLineByLine() []string {
	//Go compiler doesn’t allow creating variables that are never used. You can fix this by using an _ (underscore) in place of index

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var array []string

	for scanner.Scan() {
		str := scanner.Text()
		array = append(array, str)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return array
}

func inout() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println("The double of your number is:", output)
}

func checkRules(line string) bool {
	fmt.Println("line", line)
	firstSplit := strings.Split(line, " ")
	fmt.Println("firstSplit", firstSplit)
	secondSplit := strings.Split(firstSplit[0], "-")
	fmt.Println("secondSplit", secondSplit)
	pos2, err := strconv.ParseInt(secondSplit[1], 10, 64)
	if err != nil {
		// built-in function called panic which causes a run time error.
		panic(fmt.Sprintf("Incorrect value for float64 '%val'", pos2))
	}
	pos1, err := strconv.ParseInt(secondSplit[0], 10, 64)
	if err != nil {
		// built-in function called panic which causes a run time error.
		panic(fmt.Sprintf("Incorrect value for float64 '%val'", pos1))
	}

	char := strings.Split(firstSplit[1], "")[0]

	password := strings.Split(firstSplit[2], "")
	fmt.Println("char", char, "pos1", pos1, "pos2", pos2, "password", password)

	result := false
	if char == password[pos1-1] && char != password[pos2-1] {
		result = true
	}
	if char == password[pos2-1] && char != password[pos1-1] {
		result = true
	}

	fmt.Println("result", result)

	return result
}

//A main function executes by default when you run code in the file.
func main() {
	var array []string
	array = readLineByLine()
	var count int
	count = 0
	for _, value := range array {

		if checkRules(value) == true {
			//fmt.Println(index)
			count++
		}
	}
	fmt.Println(count)
}
