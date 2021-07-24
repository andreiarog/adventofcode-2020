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
	"sort"
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
		array = Insert(array, digit)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return array
}

//Insert by order
func Insert(ss []int, s int) []int {
	i := sort.SearchInts(ss, s)
	//ln(i)
	ss = append(ss, 0)
	//fmt.Println(ss)
	copy(ss[i+1:], ss[i:])
	//fmt.Println(ss)
	ss[i] = s
	return ss
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

func getDifferences(adaptors []int) (int, int) {

	diff1 := 0
	diff3 := 0

	source := 0

	for _, a := range adaptors {
		fmt.Println("source", source)
		fmt.Println("a", a)

		if a-source == 1 {
			diff1++
		}
		if a-source == 2 {

		}
		if a-source == 3 {
			diff3++
		}

		source = a
	}

	return diff1, diff3
}

//A main function executes by default when you run code in the file.
func main() {
	adaptors := readLineByLine()

	fmt.Println(adaptors)

	diff1, diff3 := getDifferences(adaptors)

	fmt.Println(diff1, diff3+1)
}
