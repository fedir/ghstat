// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/fedir/ghstat/sorting"
)

func rateGhData(ghData [][]string, columnsIndexes map[string]int) string {
	greetings := ""
	// Add points by repository total popularity (more popular is better)
	sorting.SortSliceByColumnIndexIntDesc(ghData, columnsIndexes["stargazersColumn"])
	greetings += fmt.Sprintf("* The most popular project is `%s`\n", ghData[0][0])
	addPoints(ghData, columnsIndexes["stargazersColumn"], columnsIndexes["totalPointsColumnIndex"])
	// Add points by project age (we like fresh ideas)
	sorting.SortSliceByColumnIndexIntAsc(ghData, columnsIndexes["ageColumn"])
	greetings += fmt.Sprintf("* The newest project is `%s`\n", ghData[0][0])
	addPoints(ghData, columnsIndexes["ageColumn"], columnsIndexes["totalPointsColumnIndex"])
	// Add points by number of commits (more commits is better)
	sorting.SortSliceByColumnIndexIntDesc(ghData, columnsIndexes["totalCommitsColumn"])
	greetings += fmt.Sprintf("* The project with more commits is `%s`\n", ghData[0][0])
	addPoints(ghData, columnsIndexes["totalCommitsColumn"], columnsIndexes["totalPointsColumnIndex"])
	// Add points by number of tags (more tags is better)
	sorting.SortSliceByColumnIndexIntDesc(ghData, columnsIndexes["totalTagsColumn"])
	greetings += fmt.Sprintf("* The project with more tags is `%s`\n", ghData[0][0])
	addPoints(ghData, columnsIndexes["totalCommitsColumn"], columnsIndexes["totalPointsColumnIndex"])
	// Add points by Top10 contributors followers
	sorting.SortSliceByColumnIndexIntDesc(ghData, columnsIndexes["top10ContributorsFollowersColumn"])
	greetings += fmt.Sprintf("* The project made by most notable top contributors is `%s`\n", ghData[0][0])
	addPoints(ghData, columnsIndexes["top10ContributorsFollowersColumn"], columnsIndexes["totalPointsColumnIndex"])
	// Add points by proportion of total and resolved issues (less opened issues is better)
	sorting.SortSliceByColumnIndexFloatDesc(ghData, columnsIndexes["closedIssuesPercentageColumn"])
	greetings += fmt.Sprintf("* The project with best errors resolving rate is `%s`\n", ghData[0][0])
	addPoints(ghData, columnsIndexes["closedIssuesPercentageColumn"], columnsIndexes["totalPointsColumnIndex"])
	// Add points by commits by day (more commits shows good healthy community)
	sorting.SortSliceByColumnIndexFloatDesc(ghData, columnsIndexes["commitsByDayColumn"])
	greetings += fmt.Sprintf("* The project with more commits by day is `%s`\n", ghData[0][0])
	addPoints(ghData, columnsIndexes["activeForkersColumn"], columnsIndexes["totalPointsColumnIndex"])
	// Add points by active forkers (more active forkers shows good open source spirit of the community)
	sorting.SortSliceByColumnIndexFloatDesc(ghData, columnsIndexes["activeForkersColumn"])
	greetings += fmt.Sprintf("* The project with the most active community is `%s`\n", ghData[0][0])
	addPoints(ghData, columnsIndexes["activeForkersColumn"], columnsIndexes["totalPointsColumnIndex"])
	// Assign places to projects by all metrics
	sorting.SortSliceByColumnIndexIntAsc(ghData, columnsIndexes["totalPointsColumnIndex"])
	greetings += fmt.Sprintf("* The best project (taking in account placements in all competitions) is `%s`\n", ghData[0][0])
	assignPlaces(ghData, columnsIndexes["totalPointsColumnIndex"])
	return greetings
}

func addPoints(s [][]string, columnIndex int, totalPointsColumnIndex int) [][]string {
	if totalPointsColumnIndex == 0 {
		log.Fatalf("Error occurred. Please check map of columns indexes")
	}
	for i := range s {
		currentValue, _ := strconv.Atoi(s[i][totalPointsColumnIndex])
		currentValue = currentValue + i + 1
		s[i][totalPointsColumnIndex] = strconv.Itoa(currentValue)
	}
	return s
}

func firstPlaceGreeting(s [][]string, message string) string {
	return fmt.Sprintf("* %s `%s`\n", message, s[0][0])
}

func assignPlaces(s [][]string, totalPointsColumnIndex int) [][]string {
	if totalPointsColumnIndex == 0 {
		log.Fatalf("Error occurred. Please check map of columns indexes")
	}
	for i := range s {
		place := i + 1
		s[i][totalPointsColumnIndex] = strconv.Itoa(place)
	}
	return s
}
