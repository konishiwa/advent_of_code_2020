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
	//accumulated amount
	acc := 0
	//exit
	exit := false
	//starting index
	var index int

	for exit == false {

		println("index is ", index)

		if _, ok := lineCall[index]; ok {
			println("we have a duplicate!")
			println(acc)
			os.Exit(3)
		} else {
			lineCall[index] = 1
		}

		if commands[index].function == "acc" {
			acc += commands[index].steps
			index++
		} else if commands[index].function == "jmp" {
			index = index + commands[index].steps
		} else {
			index++
		}
	}

	//println(total_inputs)
	//fmt.Print(string(inputs))

	//create a two-dimensional array for the input

}
