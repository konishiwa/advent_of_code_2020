package main

import (
	"bufio"
	"os"
	"strings"
)

//checks if file exists
func check(e error) {
	if e != nil {
		panic(e)
	}
}

type contents struct {
	bags []string
}

func readInFile() {

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

	//map that holds how many times a line has been called
	//bags := make(map[string]contents)

	//set all bag contents
	for _, each_ln := range text {

		//if its: contain no other bags - we're not interested
		if !strings.Contains(each_ln, "contain no other bags.") {

			//clean the string up - remove whats not needed
			cleanLn := strings.NewReplacer(each_ln, "bags", " ",
				each_ln, "bag", " ",
				each_ln, "contains", " ")

			println(cleanLn)

			//split string
			//s := strings.Split(each_ln, " ")
			//first two words are the type of bag
			//key := s[0] + s[1]

			// 	var contents string[]

			// 	//TODO: find a way to prevent string manipulation
			// 	//ignore  next two words: bags contains
			// 	cont := true
			// 	bagIndex := 4
			// 	for cont {

			// 		numBags, _ := strconv.Atoi(s[bagIndex])
			// 		bagLen := len(s[bagIndex+2])

			// 		if strings.Contains(s[bagIndex+3], ".") {
			// 			cont = false
			// 		}

			// 		for x := 0; x < numBags; x++ {

			// 			content_key := s[bagIndex+1] + s[bagIndex+2]
			// 			contents = append(contents, content_key)
			// 		}
			// 		bagIndex += 4
			// 		bags[key] = contents
			// 	}
			// }
		}

	}
}

func main() {

	//get commands
	readInFile()

}
