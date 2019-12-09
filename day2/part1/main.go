package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	program, err := readInput("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	index := 0
	for i := 0; i < len(program)-4; i += 4 {
		input := getInput(index, program)
		index += 4

		switch input[0] {
		case 1:
			add([]int{input[1], input[2], input[3]}, &program)
		case 2:
			mult([]int{input[1], input[2], input[3]}, &program)
		case 99:
			fmt.Println("stopping")
			fmt.Println(program)
		}
	}

}

func add(values []int, program *[]int) {

	prog := *program
	val1 := prog[values[0]]
	val2 := prog[values[1]]
	result := val1 + val2
	prog[values[2]] = result

	*program = prog
}

func mult(values []int, program *[]int) {

	prog := *program
	val1 := prog[values[0]]
	val2 := prog[values[1]]
	result := val1 * val2
	prog[values[2]] = result

	*program = prog
}

func getInput(startIndex int, program []int) []int {

	output := []int{}

	for i := 0; i < 4; i++ {
		output = append(output, program[i+startIndex])
	}

	return output

}

func readInput(path string) (input []int, err error) {
	file, err := os.Open(path)
	if err != nil {
		return []int{}, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)

	strSlice := strings.Split(string(data), ",")

	var intSlice []int

	for _, value := range strSlice {
		i, err := strconv.Atoi(value)
		if err != nil {
			return []int{}, err
		}
		intSlice = append(intSlice, i)
	}

	return intSlice, nil
}
