package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

type password struct {
	min   int
	max   int
	alpha string
	pass  string
}

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

func newPassword(min int, max int,
	alpha string, pass string) *password {

	thePassword := password{
		min:   min,
		max:   max,
		alpha: alpha,
		pass:  pass,
	}

	return &thePassword
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func StringToInt(t []string) []int {
	var t2 = []int{}

	for _, i := range t {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		t2 = append(t2, j)
	}
	return t2
}

func main() {

	//Read the file in as different objects
	inputs, err := ioutil.ReadFile("passwords.txt")
	check(err)
	tf := string(inputs)
	passwords, err := ReadStrings(strings.NewReader(tf))

	println(passwords[0])
	result := 0
	//set passwords
	for x := 0; x < len(passwords); x = x + 3 {

		//get min
		indexes := StringToInt(strings.Split(passwords[x], "-"))
		min := indexes[0]
		max := indexes[1]
		alpha := passwords[x+1][0]
		password := passwords[x+2]
		if min > len(password) {
			break
		} else if max > len(password) {
			//only check the first index if the second index is out of bounds
			if password[min-1] == alpha {
				result++
			}
		} else {
			//only exactly one must match
			if password[min-1] == alpha && !(password[max-1] == alpha) {
				result++
			} else if !(password[min-1] == alpha) && password[max-1] == alpha {
				result++
			}
		}
	}

	println("The valid passwords are ", result)
}
