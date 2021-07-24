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
	//"strconv"
	"sort"
	"strings"
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
		inBags := getNodesConnected(rule)
		//fmt.Println("inbags to add:", inBags)
		for _, inBag := range inBags {
			inBagInd = m[inBag]
			fmt.Println("Adding to this node", i, "this node", inBagInd, "with this name", inBag)
			gm.Add(inBagInd, i)
		}
	}
	g := graph.Sort(gm)

	return g, m
}

func clean(str string) string {
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

func getNodesConnected(rule string) []string {
	insideBagsStr := strings.Split(rule, "contain")[1]
	insideBagsArray := strings.Split(insideBagsStr, ",")
	var insideBags []string
	for _, insideBagStr := range insideBagsArray {
		insideBag := clean(insideBagStr)
		if insideBag != "no other" {
			insideBags = append(insideBags, insideBag)
		}

	}

	return insideBags
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

//A main function executes by default when you run code in the file.
func main() {
	g, nodesMap := graphLineByLine()

	start := nodesMap["shiny gold"]
	count := 0
	graph.BFS(g, start, func(v, w int, _ int64) {
		//fmt.Println(v, "to", w)
		count++

	})

	fmt.Println(count)
}

// get weights on graphs with number of bags needed inside
// keep direction of edges bag --isInside--> insideBag
// contribution[0] = 1
// contribution[w] = weight * contribution[v]
// sum all - contribution[0]
