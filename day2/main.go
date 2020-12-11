package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type line struct {
	minimum int
	maximum int
	char string
	password string
}

func main() {
	file, err := os.Open("file.txt")

	if err != nil {
		fmt.Println("An error occured on file opening process", err)
	}

	reader := bufio.NewReader(file)

	result, err := readPasswords(reader)
	
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(validatePasswords(result))

}


func validatePasswords(passwords []line) int {
	var valids int
	for _, password := range passwords {
		if validatePassword(password) {
			valids += 1
		}
	}

	return valids
}


func validatePassword(password line) bool {


	passwordString := password.password

	if string(passwordString[password.minimum - 1]) == password.char {
		if string(passwordString[password.maximum - 1]) != password.char {
			return true
		}
	}
	if string(passwordString[password.maximum - 1]) == password.char {
		if string(passwordString[password.minimum - 1]) != password.char {
			return true
		}
	}

	return false
}


func readPasswords(r io.Reader) ([]line, error) {

	scanner := bufio.NewScanner(r)

	var result []line
	for scanner.Scan() {
		x := scanner.Text()
		splited := strings.Split(x, " ")
		password := splited[2]
		delimiters := strings.Split(splited[0], "-")
		minimum, _ := strconv.Atoi(delimiters[0])
		maximum, _ := strconv.Atoi(delimiters[1])
		char := strings.Split(splited[1], ":")[0]
		input := line{
			minimum: minimum,
			maximum: maximum,
			char: char,
			password: password,
		}
		result = append(result, input)
	}
	return result, scanner.Err()
}
