package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//checks if file exists
func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Bag struct {
	num     int
	bagType string
}

func readInFile() []string {

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
	return text
}

func parseData(text []string) map[string]Bag {

	//map for all bags
	bagDef := make(map[string]Bag)

	//map that holds how many times a line has been called
	//bags := make(map[string]contents)

	//set all bag contents
	for _, each_ln := range text {

		//get the key
		s := strings.Split(each_ln, "bags contain")
		key := s[0]
		bag := Bag{
			num:     0,
			bagType: "nothing",
		}

		//if its: contain no other bags - we're not interested
		if !strings.Contains(each_ln, "contain no other bags.") {
			//clean the string up - remove whats not needed
			replacer := strings.NewReplacer(key, " ", "bags", " ", "bag", " ", "contain", " ", ".", " ")
			cleanLn := replacer.Replace(each_ln)
			cleanLn = strings.TrimSpace(cleanLn)
			//split string
			contents := strings.Split(cleanLn, ",")
			for x := 0; x < len(contents); x++ {

				contents[x] = strings.TrimSpace(contents[x])
				//first character is the number
				num := string(contents[x][0])
				numBags, _ := strconv.Atoi(num)

				//println("numBags ", numBags)
				bag = Bag{
					bagType: contents[x][1:],
					num:     numBags,
				}
			}
		}
		bagDef[key] = bag
	}
	return bagDef
}

func main() {

	//get commands
	text := readInFile()
	total := 0
	bagDef := parseData(text)

	fmt.Print(bagDef)
	for key, _ := range bagDef {

		println(key)
		currentBag := bagDef[key]
		fmt.Println(currentBag)
		// for x := 0; x < currentBag.num; x++ {
		// 	total += getContents(key, bagDef)
		// }

	}

	println(total)

}

func getContents(key string, bag map[string]Bag) int {

	total := 0
	//base case
	if key == "nothing" {
		return 0
	}

	currentBag := bag[key]
	if key == "shiny gold" {
		total++
	}
	key = currentBag.bagType
	println("the key is ", key)
	return total + getContents(key, bag)
}
