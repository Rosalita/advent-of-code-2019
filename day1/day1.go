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

	var total1 int
	for _, v := range input {
		total1 += calcFuel(v)
	}
	fmt.Println(total1) //part 1

	var total2 int
	for _, v := range input {

		f := calcFuel(v)
		fmt.Println(f)
		total2 += f

		if f >= 0 {
			f2 := calcFuel(f)
			total2 += f2
			if f2 >= 0 {
				f3 := calcFuel(f2)
				total2 += f3
				if f3 >= 0 {
					f4 := calcFuel(f3)
					total2 += f4
					if f4 >= 0 {
						f5 := calcFuel(f4)
						total2 += f5
						if f5 >= 0 {
							f6 := calcFuel(f5)
							total2 += f6
							if f6 >= 0 {
								f7 := calcFuel(f6)
								total2 += f7
								if f7 >= 0 {
									f8 := calcFuel(f7)
									total2 += f8
									if f8 >= 0 {
										f9 := calcFuel(f8)
										total2 += f9
										if f9 >= 0 {
											f10 := calcFuel(f9)
											total2 += f10
											if f10 >= 0 {
												f11 := calcFuel(f10)
                        total2 += f11
                        if f11 >= 0 {
                          f12 := calcFuel(f11)
                          total2 += f12
                        }
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(total2) // part 2
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

	strSlice := strings.Split(string(data), "\n")

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
