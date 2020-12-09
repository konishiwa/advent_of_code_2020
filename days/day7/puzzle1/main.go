package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type command struct {
	function string
	steps    int
}

//checks if file exists
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInFile() [594]command {

	//read the values in
	file, err := os.Open("input.txt")
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	//store in multidimensional array
	var commands [594]command

	//map that holds how many times a line has been called
	bags := make(map[string]string)

	//example lines
	// 	shiny indigo bags contain 4 vibrant lime bags.
	// clear lime bags contain 1 dotted lime bag, 2 clear gold bags.
	// dotted turquoise bags contain 2 shiny green bags, 5 striped magenta bags, 3 muted green bags.

	for i, each_ln := range text {

		//clean the data
		//if its: contain no other bags - we're not interested
		if !strings.Contains(each_ln, "contain no other bags.") {

			//split string
			s := strings.Split(each_ln, " ")
			//first two words are the type of bag
			key := s[0] + s[1]
			//add to bags 
			
			//depending on the value x add x times
			for a = append(a, iter.Val)

		}

		//first two lines are what it can be kept in

		//the function
		function := s[0]

		if function != "nop" {
			//the steps
			steps, _ := strconv.Atoi(s[1])
			//set command struct
			commands[i] = command{function, steps}
		} else {
			commands[i] = command{function, 0}
		}
	}
	return commands
}

func main() {

	//get commands
	commands := readInFile()
	//map that holds how many times a line has been called
	lineCall := make(map[int]int)

}
