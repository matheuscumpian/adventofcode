package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main(){
	file, _ := os.Open("file.txt")
	reader := bufio.NewReader(file)
	matrix, _ := readMatrix(reader)

	entries := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	product := 1
	for i := 0; i < len(entries); i++ {
		product *= countTrees(matrix, entries[i][1], entries[i][0])
	}
	fmt.Println(product)
}


func countTrees(matrix [][]string, patternX int, patternY int) int {
	var xPos = patternX
	var yPos int
	var trees int
	for xPos < len(matrix) {
		if yPos + patternY < len(matrix[0]) {
			yPos += patternY
		} else {
			aux := (yPos + patternY) - len(matrix[0])
			yPos = aux
		}
		if matrix[xPos][yPos] == "#" {
			trees++
		}
		xPos += patternX
	}
	return trees
}


func readMatrix(r io.Reader) ([][]string, error) {
	scanner := bufio.NewScanner(r)
	var result [][]string
	var counter int
	for scanner.Scan() {
		x := scanner.Text()
		line := strings.Split(x, "")
		result = append(result, line)
		counter += 1
	}
	return result, scanner.Err()
}