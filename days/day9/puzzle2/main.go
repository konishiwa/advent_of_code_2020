package main

import (
	"bufio"
	"os"
	"strconv"
)

//checks if file exists
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInFile() [1000]int {

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

	var commands [1000]int

	//change to integers
	for i, each_ln := range text {

		totals, _ := strconv.Atoi(each_ln)
		//set command struct
		commands[i] = totals

	}
	return commands
}

func main() {

	//get commands
	commands := readInFile()

	for i := 0; i < 1000; i++ {

		found := false

		//start the check at 26
		check := commands[i+25]

		//set the preamble each time - should probably be changed
		var preamble [25]int
		index := 0
		for x := i; x < i+25; x++ {
			preamble[index] = commands[x]
			index++
		}
		//fmt.Printf("%v", preamble)
		//check all possible sums
		for z := 0; z < 25; z++ {

			//set the number to add to
			firstNum := preamble[z]
			println("firstNum ", firstNum)

			//check the numbers after them
			for p := z + 1; p < 25; p++ {

				secondNum := preamble[p]
				total := firstNum + secondNum

				if total == check && firstNum != secondNum {
					found = true
					break
				}
			}
		}
		if found == false {
			println("broke on ", check)

			for i := 0; i < 1000; i++ {

				total := commands[i]
				println(total)

				for x := i + 1; x < 1000; x++ {

					total += commands[x]
					if total > check {
						break
					} else if total == check {
						println("radical")
						println("the total is: ", total)
						println("i is ", i)
						println("x is", x)
						os.Exit(3)
					}
				}

			}
		}
	}
	//println(total_inputs)
	//fmt.Print(string(inputs))

	//create a two-dimensional array for the input

}
