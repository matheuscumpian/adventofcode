package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("file.txt")

	if err != nil {
		fmt.Println("An error occured on file opening process", err)
	}

	reader := bufio.NewReader(file)

	numbers, _ := readIntegers(reader)
	result:= findProduct(numbers)
	fmt.Println(result)
}

func findProduct(numbers []int) int {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			for k := 0; k < len(numbers); k++ {
				if numbers[i]+numbers[j]+numbers[k] == 2020 {
					fmt.Println(numbers[i], numbers[j], numbers[k])
					return numbers[i]*numbers[j]*numbers[k]
				}
			}
		}
	}
	return 0
}

func readIntegers(r io.Reader) ([]int, error) {

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}

	return result, scanner.Err()
}
