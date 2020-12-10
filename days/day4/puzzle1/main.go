package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"strings"
)

func ReadStrings(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []string
	for scanner.Scan() {
		x := scanner.Text()
		result = append(result, x)
	}
	return result, scanner.Err()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//Read the file in as different objects
	inputs, err := ioutil.ReadFile("input.txt")
	check(err)
	tf := string(inputs)
	temp := strings.Split(tf, "\n\n")
	keys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	results := 0

	for x := 0; x < len(temp); x++ {

		valid := true

		for i := 0; i < 7; i++ {

			if !(strings.Contains(temp[x], keys[i])) {
				valid = false
			}
		}

		if valid == true {
			results++
		}
	}

	println("Valid passports ", results)

}
