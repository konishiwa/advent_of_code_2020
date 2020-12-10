package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//each node
type Node struct {
	value int
	//because its +=3, there will no more than 3 children at all times
	left   *Node
	mid    *Node
	right  *Node
	parent *Node
	root   bool
}

//printNode to check
func printNode(n *Node) {
	fmt.Println("Node Value: ", n.value)
	if n.left != nil {
		fmt.Println("Node Left: ", n.left.value)
	}
	if n.mid != nil {
		fmt.Println("Node Mid: ", n.mid.value)
	}
	if n.right != nil {
		fmt.Println("Node Right: ", n.right.value)
	}
	if n.root != true {
		fmt.Println("Node Parent: ", n.parent.value)
	}
	//print new line to format prints
	fmt.Println()
}

func main() {

	rawFile := readFile()
	sortFile(rawFile)
	input := fmtFile(rawFile)
	nodes := nodeify(input)
	printNodes(nodes)

	//the node path that we need is the one w/ the child of 164
	getNodePath(nodes)

}

func getNodePath(nodes []Node) {

	//find the parent
	printParent(nodes[102])

}

func printParent(node Node) {

	//base case
	if node.parent == nil {
		println("parent is nil")
		return
	} else {
		println("Node: ", node.value, "Node: ", node.parent)
	}
}

func printNodes(nodes []Node) {

	for _, node := range nodes {
		printNode(&node)
	}
}

func nodeify(filenodes []string) []Node {

	//TOD: fix by implementing Scanf
	n := len(filenodes)
	var nodes []Node = make([]Node, n)
	println("n is: ", n)

	for i := 0; i < n; i++ {
		//set the variables
		nodeSplit := strings.Split(filenodes[i], " ")

		if i == 0 {
			nodes[i].root = true
		}

		if len(nodeSplit) >= 2 {
			nodeLeft, _ := strconv.Atoi(nodeSplit[1])
			nodes[i].left = &Node{
				value:  nodeLeft,
				parent: &nodes[i],
			}
		}
		if len(nodeSplit) >= 3 {
			nodeRight, _ := strconv.Atoi(nodeSplit[2])
			nodes[i].right = &Node{
				value:  nodeRight,
				parent: &nodes[i],
			}
		}
		if len(nodeSplit) == 4 {
			nodeMid, _ := strconv.Atoi(nodeSplit[3])
			nodes[i].right = &Node{
				value:  nodeMid,
				parent: &nodes[i],
			}
		}
		nodes[i].value, _ = strconv.Atoi(nodeSplit[0])
		println("node: ", nodes[i].value, " left: ", nodes[i].left, " right: ",
			nodes[i].right, " mid: ", nodes[i].mid, "parent: ", nodes[i].parent)
	}

	return nodes
}

func readFile() []string {

	//read the values in
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	//append zero at the end
	text = append(text, "0")
	file.Close()
	return text
}

func sortFile(nodes []string) {
	// Pass in our list and a func to compare values
	sort.Slice(nodes, func(i, j int) bool {
		numA, _ := strconv.Atoi(nodes[i])
		numB, _ := strconv.Atoi(nodes[j])
		return numA < numB
	})
}

func fmtFile(nodes []string) []string {

	//format file back to to txt with number on each line
	for i, each := range nodes {

		currentNodeVal, _ := strconv.Atoi(each)
		check := i
		if check == 100 {
			check = 3
		} else if check == 101 {
			check = 2
		} else if check == 102 {
			check = 0
		} else {
			check = 4
		}
		println("the old node line: ", nodes[i])
		//check up to the next 3 lines (since its sorted)
		for x := 1; x < check; x++ {
			checkNodeVal, _ := strconv.Atoi(nodes[i+x])
			min := currentNodeVal
			max := currentNodeVal + 3

			if min <= checkNodeVal && max >= checkNodeVal {
				nodes[i] = nodes[i] + " " + nodes[i+x]
			}

		}
		println("the new node line: ", nodes[i])
	}
	return nodes
}
