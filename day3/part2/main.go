package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type point struct {
	X int
	Y int
	Steps int
}

type coord struct {
	X int
	Y int
}

func main() {
	route1, route2, err := readInput("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	r1points := getPoints(route1)
	r2points := getPoints(route2)

	result := comparePoints(r1points, r2points)

	closest := findClosestPoint(result)
	fmt.Println(closest)
}

func findClosestPoint(points []point) int{
	var closestSteps int
	for i, v := range points {

		if i == 0 { 
			closestSteps = v.Steps
			continue
		}

		if v.Steps < closestSteps{
			closestSteps = v.Steps
		}
	}
	return closestSteps
}

func comparePoints(points1, points2 []point)([]point){
	var matchedPoints []point
	pointMap := make(map[coord]int)

	for _, v := range points1 {
		pointMap[coord{v.X, v.Y}] = v.Steps
	}

	for _, v := range points2 {
		if _, ok := pointMap[coord{v.X, v.Y}]; ok{

			p1steps := pointMap[coord{v.X, v.Y}]
			pointStepTotal := point{v.X, v.Y, v.Steps+p1steps}
			matchedPoints = append(matchedPoints, pointStepTotal)
		}
	}
	return matchedPoints
}

func getPoints(route []string) []point {
	var points []point
	from := point{0, 0, 0}

	var steps int

	for _, v := range route {

		direction := string(v[0])
		distance, err := strconv.Atoi(v[1:])
		if err != nil {
			fmt.Println(err)
		}

		switch direction {
		case "R":
			for i := 0; i < distance; i++ {
				steps ++
				point := point{from.X + (i + 1), from.Y, steps}
				points = append(points, point)
				if i == (distance - 1) {
					from.X += distance
				}
			}
		case "L":
			for i := 0; i < distance; i++ {
				steps ++
				point := point{from.X - (i + 1), from.Y, steps}
				points = append(points, point)
				if i == (distance - 1) {
					from.X -= distance
				}
			}

		case "U":
			for i := 0; i < distance; i++ {
				steps ++
				point := point{from.X, from.Y + (i + 1), steps}
				points = append(points, point)
				if i == (distance - 1) {
					from.Y += distance
				}
			}

		case "D":
			for i := 0; i < distance; i++ {
				steps ++
				point := point{from.X, from.Y - (i + 1), steps}
				points = append(points, point)
				if i == (distance - 1) {
					from.Y -= distance
				}
			}
		}
	}
	return points
}

func readInput(path string) ([]string, []string, error) {
	file, err := os.Open(path)
	if err != nil {
		return []string{}, []string{}, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	routes := strings.Split(string(data), "\n")
	route1 := strings.Split(routes[0], ",")
	route2 := strings.Split(routes[1], ",")

	return route1, route2, nil
}
