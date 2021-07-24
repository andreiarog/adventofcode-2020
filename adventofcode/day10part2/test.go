//Declare a main package (a package is a way to group functions)
package main

//Import the popular fmt package, which contains functions for formatting text, including printing to the console. This package is one of the standard library packages you got when you installed Go.
import (
	"bufio"
	"fmt"
	//"io/ioutil"
	"github.com/yourbasic/graph"
	"log"
	"os"
	//"regexp"
	"sort"
	"strconv"
	//"strings"
	"math"
)

var (
	array         []int
	graphMain     *graph.Immutable
	graphSkipped  *graph.Immutable
	count         int64
	skipping      = make(map[int]int64)
	skippedNodes  []int
	skippedCounts int64
)

func arrayLineByLine() []int {
	//Go compiler doesn’t allow creating variables that are never used. You can fix this by using an _ (underscore) in place of index

	f, err := os.Open("input2.txt")

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

func createGraph(array []int) *graph.Immutable {
	gm := graph.New(maxIntSlice(array) + 1)

	for ind, i := range array {
		// for _, j := range array[ind : ind+4] {
		for _, j := range array {
			// possible adaptors to connect have between +1 and +3

			if j > i && j < i+4 {
				//fmt.Println("if j-i==1,", j, i, "and array[indj+1]-j<3", array[ind+1])
				if j > 2 && j < maxIntSlice(array) && j-i == 1 && array[ind+2]-j == 1 {
					// skip this node
					fmt.Println("Not adding to this node", i, "this node", j)
					skipping[i]++
					fmt.Println("The number of nodes skipped by", i, "are", skipping[i])
					skippedNodes = append(skippedNodes, j)
				} else {
					gm.Add(i, j)
					fmt.Println("Adding to this node", i, "this node", j)
				}
			}
		}
	}
	g := graph.Sort(gm)

	return g
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

func minIntSlice(slice []int) int {
	var m int

	for i, v := range slice {
		if i == 0 || v < m {
			m = v
		}
	}

	return m
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

func visit(vertice int, skipped int64, g *graph.Immutable) {

	if g.Degree(vertice) > 0 {
		g.Visit(vertice, func(w int, c int64) (skip bool) {
			fmt.Println("visiting node", w, "which was picked by node", vertice)
			if w == maxIntSlice(array) {
				// if zero then just adds 1 path
				skippedFloat := float64(skipped + skipping[vertice])
				fmt.Println("reaching final node with skipping", skippedFloat)
				count += int64(math.Exp2(skippedFloat))
			}
			if w != maxIntSlice(array) {
				fmt.Println("we will now visit node", w, "which was picked by node", vertice, "with accumulated skipped equal to", skipped, "will now be added with", skipping[vertice])
				visit(w, skipped+skipping[vertice], g)

			}

			return
		})
	}

}

//A main function executes by default when you run code in the file.
func main() {
	array = arrayLineByLine()

	graphMain = createGraph(array)

	start := array[0]
	count = 0

	visit(start, 0, graphMain)

	fmt.Println(count)
}
