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
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	weights = make(map[int]int64)
	// because by default the wieights will be 0 we need to validate if it's a real 0
	// this array has all the nodes whose weights have been calculated
	calculatedWeights = []int{}
)

func graphLineByLine() (*graph.Immutable, map[string]int) {
	//Go compiler doesn’t allow creating variables that are never used. You can fix this by using an _ (underscore) in place of index

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var rules []string

	for scanner.Scan() {
		str := scanner.Text()
		//strSplit := strings.Split(str, " ")
		rules = append(rules, str)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	gm := graph.New(len(rules))
	m := getNodesMap(rules)
	//fmt.Println("MAPPING", m)

	var inBagInd int

	for i, rule := range rules {
		inBags, costs := getNodesAndCostConnected(rule)
		//fmt.Println("inbags to add:", inBags)
		for j, inBag := range inBags {
			inBagInd = m[inBag]
			//fmt.Println("Adding to this node", i, "this node", inBagInd, "with this name", inBag, "and this cost", costs[j])
			gm.AddCost(i, inBagInd, costs[j])
		}
	}
	g := graph.Sort(gm)

	return g, m
}

func cleanBag(str string) string {
	// Clean substrings from string

	// Trim bag/bags
	str = strings.ReplaceAll(str, " bags", "")
	str = strings.ReplaceAll(str, " bag", "")

	// Trim numbers
	// Make a Regex to say we only want letters
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}
	str = reg.ReplaceAllString(str, " ")

	// Trim spaces
	str = strings.TrimLeft(str, " ")
	str = strings.TrimRight(str, " ")

	return str
}

func cleanCost(str string) int64 {
	// Clean substrings from string
	digit := 0
	// get the numbers
	re := regexp.MustCompile("[0-9]+")
	digitStr := re.FindAllString(str, -1)
	if len(digitStr) > 0 {
		if n, err := strconv.Atoi(digitStr[0]); err == nil {
			digit = n
		}
	}

	return int64(digit)
}

func getNodesAndCostConnected(rule string) ([]string, []int64) {
	insideBagsStr := strings.Split(rule, "contain")[1]
	insideBagsArray := strings.Split(insideBagsStr, ",")
	var insideBags []string
	var costs []int64
	for _, insideBagStr := range insideBagsArray {
		cost := cleanCost(insideBagStr)
		insideBag := cleanBag(insideBagStr)
		if insideBag != "no other" {
			insideBags = append(insideBags, insideBag)
			costs = append(costs, cost)
		}

	}

	return insideBags, costs
}

func getNodesMap(rules []string) map[string]int {

	m := make(map[string]int)
	for i, rule := range rules {
		outBag := strings.Split(rule, " bags contain")[0]
		m[outBag] = i
	}

	return m
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

func nodeWeight(g *graph.Immutable, v int) {
	//Visit calls the do function for each neighbor w of v
	//You can create an if condition to make skip=true to skip some nodes

	neighbours := []int{}
	costs := []int64{}

	if g.Degree(v) == 0 {
		weights[v] = 0
	}
	if g.Degree(v) != 0 {
		g.Visit(v, func(w int, c int64) (skip bool) {

			neighbours = append(neighbours, w)
			costs = append(costs, c)

			if !Find(calculatedWeights, w) {
				nodeWeight(g, w)
			}

			return
		})
	}

	for n, node := range neighbours {
		if weights[node] != 0 {
			// it's the weight+1 because for each bag we have it's inside bags and the bag itself
			weights[v] = weights[v] + costs[n]*(weights[node]+1)
		}
		if weights[node] == 0 {
			weights[v] = weights[v] + costs[n]*1
		}
	}
	calculatedWeights = append(calculatedWeights, v)

}

//A main function executes by default when you run code in the file.
func main() {
	g, nodesMap := graphLineByLine()

	start := nodesMap["shiny gold"]
	nodeWeight(g, start)
	fmt.Println(weights[start])
	//fmt.Println(g, nodesMap)
}
