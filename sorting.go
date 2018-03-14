package main

import (
	"sort"
	"strconv"
)

func sortSliceByColumnIndexIntAsc(s [][]string, columnIndex int) [][]string {
	sort.Slice(s, func(i, j int) bool {
		firstCellValue, _ := strconv.Atoi(s[i][columnIndex])
		secondCellValue, _ := strconv.Atoi(s[j][columnIndex])
		return firstCellValue < secondCellValue
	})
	return s
}

func sortSliceByColumnIndexIntDesc(s [][]string, columnIndex int) [][]string {
	sort.Slice(s, func(i, j int) bool {
		firstCellValue, _ := strconv.Atoi(s[i][columnIndex])
		secondCellValue, _ := strconv.Atoi(s[j][columnIndex])
		return firstCellValue > secondCellValue
	})
	return s
}

func sortSliceByColumnIndexFloatAsc(s [][]string, columnIndex int) [][]string {
	sort.Slice(s, func(i, j int) bool {
		firstCellValue, _ := strconv.ParseFloat(s[i][columnIndex], 32)
		secondCellValue, _ := strconv.ParseFloat(s[j][columnIndex], 32)
		return firstCellValue < secondCellValue
	})
	return s
}

func sortSliceByColumnIndexFloatDesc(s [][]string, columnIndex int) [][]string {
	sort.Slice(s, func(i, j int) bool {
		firstCellValue, _ := strconv.ParseFloat(s[i][columnIndex], 32)
		secondCellValue, _ := strconv.ParseFloat(s[j][columnIndex], 32)
		return firstCellValue > secondCellValue
	})
	return s
}
