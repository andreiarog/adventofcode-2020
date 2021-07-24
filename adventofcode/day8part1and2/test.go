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
	"strings"
)

type instruction struct {
	action string
	move   int
}

func readLineByLine() map[int]instruction {
	//Go compiler doesn’t allow creating variables that are never used. You can fix this by using an _ (underscore) in place of index

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	instructions := make(map[int]instruction)

	i := 0
	for scanner.Scan() {
		str := scanner.Text()
		//strSplit := strings.Split(str, " ")
		move, err := strconv.Atoi(strings.Split(str, " ")[1])
		if err != nil {
			log.Fatal(err)
		}
		action := strings.Split(str, " ")[0]
		instructions[i] = instruction{action, move}
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return instructions
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
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

func computeAcc(instructions map[int]instruction) (int, bool) {

	hasCycle := false

	accumulator := 0
	visited := []int{}
	i := 0

	for i > -1 {
		if i >= len(instructions) {
			hasCycle = false
			break
		}
		if Find(visited, i) {
			hasCycle = true
			break
		}
		if instructions[i].action == "acc" {
			fmt.Println("Found instruction acc in index", i, "and move", instructions[i].move)
			accumulator += instructions[i].move
			visited = append(visited, i)
			i++
		}
		if instructions[i].action == "jmp" {
			fmt.Println("Found instruction jmp in index", i, "and move", instructions[i].move)
			visited = append(visited, i)
			i = i + instructions[i].move
		}
		if instructions[i].action == "nop" {
			fmt.Println("Found instruction nop in index", i, "and move", instructions[i].move)
			visited = append(visited, i)
			i++
		}
	}

	return accumulator, hasCycle
}

//A main function executes by default when you run code in the file.
func main() {
	instructions := readLineByLine()

	var accumulator int
	var hasCycles bool
	instructionsSwitched := make(map[int]instruction)
	i := 0

	for i > -1 {
		for k, v := range instructions {
			instructionsSwitched[k] = v
		}

		if i > len(instructions) {
			break
		}
		// cycle through instructions and if found nop or jmp switch > create new instructions
		// calculate accumulator and if has cycles
		action := instructions[i].action
		if action == "jmp" || action == "nop" {
			if action == "jmp" {
				move := instructions[i].move
				instructionsSwitched[i] = instruction{"nop", move}
				accumulator, hasCycles = computeAcc(instructionsSwitched)
			}
			if action == "nop" {
				move := instructions[i].move
				instructionsSwitched[i] = instruction{"jmp", move}
				accumulator, hasCycles = computeAcc(instructionsSwitched)
			}
			if !hasCycles {
				fmt.Println("no cycle found", i)
				break
			}
			if hasCycles {
				fmt.Println("cycle found", i)
				i++
			}
		}
		if action == "acc" {
			i++
		}
	}

	fmt.Println(accumulator)
	//fmt.Println(g, nodesMap)
}
