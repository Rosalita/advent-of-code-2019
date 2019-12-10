package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

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

	closest := findClosestManhattan(result)
	fmt.Println(closest)
}

func findClosestManhattan(points []coord) int{
	var closestManhattan int
	for i, v := range points {

		if i == 0 { 
			closestManhattan = v.X + v.Y 
			continue
		}

		manhattan := v.X + v.Y
		if manhattan < closestManhattan{
			closestManhattan = manhattan
		}
	}
	return closestManhattan
}

func comparePoints(points1, points2 []coord)([]coord){
	var matchedPoints []coord
	pointMap := make(map[coord]bool)

	for _, v := range points1 {
		pointMap[v] = true
	}

	for _, v := range points2 {
		if pointMap[v] == true{
			matchedPoints = append(matchedPoints, v)
		}
	}
	return matchedPoints
}

func getPoints(route []string) []coord {
	var points []coord
	from := coord{0, 0}

	for _, v := range route {

		direction := string(v[0])
		distance, err := strconv.Atoi(v[1:])
		if err != nil {
			fmt.Println(err)
		}

		switch direction {
		case "R":
			for i := 0; i < distance; i++ {
				coord := coord{from.X + (i + 1), from.Y}
				points = append(points, coord)
				if i == (distance - 1) {
					from.X += distance
				}
			}
		case "L":
			for i := 0; i < distance; i++ {
				coord := coord{from.X - (i + 1), from.Y}
				points = append(points, coord)
				if i == (distance - 1) {
					from.X -= distance
				}
			}

		case "U":
			for i := 0; i < distance; i++ {
				coord := coord{from.X, from.Y + (i + 1)}
				points = append(points, coord)
				if i == (distance - 1) {
					from.Y += distance
				}
			}

		case "D":
			for i := 0; i < distance; i++ {
				coord := coord{from.X, from.Y - (i + 1)}
				points = append(points, coord)
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
