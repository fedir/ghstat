// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.
package main

import (
	"fmt"
	"log"
	"strconv"
)

func rateGhData(ghData [][]string, columnsIndexes map[string]int) {
	// Add points by repository total popularity (more popular is better)
	sortSliceByColumnIndexIntDesc(ghData, columnsIndexes["stargazersColumn"])
	firstPlaceGreeting(ghData, "The most popular project is")
	addPoints(ghData, columnsIndexes["stargazersColumn"], columnsIndexes["totalPointsColumnIndex"])
	// Add points by project age (we like fresh ideas)
	sortSliceByColumnIndexIntAsc(ghData, columnsIndexes["ageColumn"])
	firstPlaceGreeting(ghData, "The newest project is")
	addPoints(ghData, columnsIndexes["ageColumn"], columnsIndexes["totalPointsColumnIndex"])
	// Add points by active forkers (more active forkers shows good open source spirit of the community)
	sortSliceByColumnIndexFloatDesc(ghData, columnsIndexes["activeForkersColumn"])
	firstPlaceGreeting(ghData, "The project with the most active community is")
	addPoints(ghData, columnsIndexes["activeForkersColumn"], columnsIndexes["totalPointsColumnIndex"])
	// Add points by proportion of total and resolved issues (less opened issues is better)
	sortSliceByColumnIndexFloatDesc(ghData, columnsIndexes["closedIssuesPercentageColumn"])
	firstPlaceGreeting(ghData, "The project with best errors resolving rate is")
	addPoints(ghData, columnsIndexes["closedIssuesPercentageColumn"], columnsIndexes["totalPointsColumnIndex"])
	// Add points by number of commits (more commits is better)
	sortSliceByColumnIndexIntDesc(ghData, columnsIndexes["totalCommitsColumn"])
	firstPlaceGreeting(ghData, "The project with more commits is")
	addPoints(ghData, columnsIndexes["totalCommitsColumn"], columnsIndexes["totalPointsColumnIndex"])
	// Add points by Top10 contributors followers
	sortSliceByColumnIndexIntDesc(ghData, columnsIndexes["top10ContributorsFollowersColumn"])
	firstPlaceGreeting(ghData, "The project made by most notable top contributors is")
	addPoints(ghData, columnsIndexes["top10ContributorsFollowersColumn"], columnsIndexes["totalPointsColumnIndex"])
	// Assign places to projects by all metrics
	sortSliceByColumnIndexIntAsc(ghData, columnsIndexes["totalPointsColumnIndex"])
	firstPlaceGreeting(ghData, "The best project (taking in account placements in all competitions) is")
	assignPlaces(ghData, columnsIndexes["totalPointsColumnIndex"])
}

func addPoints(s [][]string, columnIndex int, totalPointsColumnIndex int) [][]string {
	if totalPointsColumnIndex == 0 {
		log.Fatalf("Error occured. Please check map of columns indexes")
	}
	for i := range s {
		currentValue, _ := strconv.Atoi(s[i][totalPointsColumnIndex])
		currentValue = currentValue + i + 1
		s[i][totalPointsColumnIndex] = strconv.Itoa(currentValue)
	}
	return s
}

func firstPlaceGreeting(s [][]string, message string) {
	fmt.Printf("* %s `%s`\n", message, s[0][0])
}

func assignPlaces(s [][]string, totalPointsColumnIndex int) [][]string {
	if totalPointsColumnIndex == 0 {
		log.Fatalf("Error occured. Please check map of columns indexes")
	}
	for i := range s {
		place := i + 1
		s[i][totalPointsColumnIndex] = strconv.Itoa(place)
	}
	return s
}
