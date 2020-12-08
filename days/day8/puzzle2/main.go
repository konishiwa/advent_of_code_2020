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

func readInFile() [596]command {

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
	var commands [596]command

	for i, each_ln := range text {

		//split string
		s := strings.Split(each_ln, " ")

		//the function
		function := s[0]

		//the steps
		steps, _ := strconv.Atoi(s[1])
		//set command struct
		commands[i] = command{function, steps}

	}
	return commands
}

func main() {

	//get commands
	commands := readInFile()
	//starting index

	//try to change one line until it hits the end
	for i := 0; i < 596; i++ {

		//create a different array to to revert change
		changedCommands := commands

		if changedCommands[i].function == "jmp" {
			changedCommands[i].function = "nop"
			println("Changed line ", i)
			checkAccumulation(changedCommands)
		} else if changedCommands[i].function == "nop" {
			changedCommands[i].function = "jmp"
			checkAccumulation(changedCommands)
			println("Changed line ", i)
		}
	}

}

func checkAccumulation(commands [596]command) {

	//map that holds how many times a line has been called
	lineCall := make(map[int]int)

	println(len(lineCall))

	var acc int
	var x int

	for x < 596 {
		if _, ok := lineCall[x]; ok && x != 0 {
			println("we have a duplicate at ", x)
			return
		} else {
			lineCall[x] = 1
		}
		if commands[x].function == "acc" {
			acc += commands[x].steps
			x++
		} else if commands[x].function == "jmp" {
			x = x + commands[x].steps
		} else {
			x++
		}
	}

	println("the accumulation is ", acc)
	os.Exit(3)

}
