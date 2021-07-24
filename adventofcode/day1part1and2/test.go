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
)

// GLOBAL VARIABLES
var (
	array []float64
)

const ()

func readFullFile() {
	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
	fmt.Println(len(content))

}

func readLineByLine() []float64 {
	//Go compiler doesn’t allow creating variables that are never used. You can fix this by using an _ (underscore) in place of index

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		str := scanner.Text()
		value, err := strconv.ParseFloat(str, 64)
		array = append(array, value)
		if err != nil {
			// built-in function called panic which causes a run time error.
			panic(fmt.Sprintf("Incorrect value for float64 '%value'", value))
		}
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

func getProductPair(array []float64) int {

	var product int
	// var first float64
	// var second float64
	// var third float64

	for _, i := range array {
		for _, j := range array {
			for _, k := range array {
				if i+j+k == 2020 {
					// first := i
					// second := j
					// third := k
					product = int(i * j * k)
					break
				}
			}
		}
	}

	return product
}

//A main function executes by default when you run code in the file.
func main() {
	fmt.Println(getProductPair(readLineByLine()))

}
