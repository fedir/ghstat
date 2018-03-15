package main

import (
	"strconv"
)

func addPoints(s [][]string, columnIndex int, totalPointsColumnIndex int) [][]string {
	for i := range s {
		currentValue, _ := strconv.Atoi(s[i][totalPointsColumnIndex])
		currentValue = currentValue + i + 1
		s[i][totalPointsColumnIndex] = strconv.Itoa(currentValue)
	}
	return s
}

func assignPlaces(s [][]string, totalPointsColumnIndex int) [][]string {
	for i := range s {
		place := i + 1
		s[i][totalPointsColumnIndex] = strconv.Itoa(place)
	}
	return s
}
