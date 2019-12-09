package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, err := readInput("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	var total int
	for _, v := range input {
		total += calcFuel(v)
	}
	fmt.Println(total)
}

func calcFuel(input int) int {
	return (input / 3) - 2
}

func readInput(path string) (input []int, err error) {
	file, err := os.Open(path)
	if err != nil {
		return []int{}, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)

	strSlice := strings.Split(string(data), "\r\n")

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
