package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	memory, err := readInput("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	index := 0
	for i := 0; i < len(memory)-4; i += 4 {

		instructions := getInstruction(index, memory)
	
		opCode := instructions[0]
		parameters := instructions[1:4]

		switch opCode {
		case 1:
			add(parameters, &memory)
		case 2:
			mult(parameters, &memory)
		case 99:
			fmt.Println("stopping")
			fmt.Println(memory)
		}
		index += 4
		
	}

}

func add(params[]int, memory *[]int) {
	mem := *memory
	param1 := mem[params[0]]
	param2 := mem[params[1]]
	result := param1 + param2
	mem[params[2]] = result
	*memory = mem
}

func mult(params[]int, memory *[]int) {
	mem := *memory
	param1 := mem[params[0]]
	param2 := mem[params[1]]
	result := param1 * param2
	mem[params[2]] = result
	*memory = mem
}

func getInstruction(startIndex int, memory []int) []int {
	instruction := []int{}
	for i := 0; i < 4; i++ {
		instruction = append(instruction, memory[i+startIndex])
	}
	return instruction
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
