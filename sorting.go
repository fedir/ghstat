// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.
package main

import (
	"log"
	"sort"
	"strconv"
)

func sortSliceByColumnIndexIntAsc(s [][]string, columnIndex int) [][]string {
	if columnIndex == 0 {
		log.Fatalf("Error occurred. Please check map of columns indexes")
	}
	sort.Slice(s, func(i, j int) bool {
		firstCellValue, _ := strconv.Atoi(s[i][columnIndex])
		secondCellValue, _ := strconv.Atoi(s[j][columnIndex])
		return firstCellValue < secondCellValue
	})
	return s
}

func sortSliceByColumnIndexIntDesc(s [][]string, columnIndex int) [][]string {
	if columnIndex == 0 {
		log.Fatalf("Error occurred. Please check map of columns indexes")
	}
	sort.Slice(s, func(i, j int) bool {
		firstCellValue, _ := strconv.Atoi(s[i][columnIndex])
		secondCellValue, _ := strconv.Atoi(s[j][columnIndex])
		return firstCellValue > secondCellValue
	})
	return s
}

func sortSliceByColumnIndexFloatDesc(s [][]string, columnIndex int) [][]string {
	if columnIndex == 0 {
		log.Fatalf("Error occurred. Please check map of columns indexes")
	}
	sort.Slice(s, func(i, j int) bool {
		firstCellValue, _ := strconv.ParseFloat(s[i][columnIndex], 32)
		secondCellValue, _ := strconv.ParseFloat(s[j][columnIndex], 32)
		return firstCellValue > secondCellValue
	})
	return s
}
