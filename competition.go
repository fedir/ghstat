// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"sort"
)

func rateAndPrintGreetings(ghData []Repository) {
	greetings := rateGhData(ghData)
	fmt.Println(greetings)
}

func rateGhData(ghData []Repository) string {

	greetings := ""

	// Add points by repository total popularity (more popular is better)
	sort.Slice(ghData[:], func(i, j int) bool {
		return ghData[i].Watchers > ghData[j].Watchers
	})
	greetings += fmt.Sprintf("* The most popular project is `%s`\n", ghData[0].Name)
	for i := range ghData {
		ghData[i].PlacementPopularity = ghData[i].PlacementPopularity + i
		ghData[i].PlacementOverall = ghData[i].PlacementOverall + i
	}

	greetings += fmt.Sprintf("* The newest project is `%s`\n", ghData[0].Name)
	for i := range ghData {
		ghData[i].PlacementAge = ghData[i].PlacementAge + i
		ghData[i].PlacementOverall = ghData[i].PlacementOverall + i
	}

	// Add points by number of commits (more commits is better)
	sort.Slice(ghData[:], func(i, j int) bool {
		return ghData[i].TotalCommits > ghData[j].TotalCommits
	})
	greetings += fmt.Sprintf("* The project with more commits is `%s`\n", ghData[0].Name)
	for i := range ghData {
		ghData[i].PlacementTotalCommits = ghData[i].PlacementTotalCommits + i
		ghData[i].PlacementOverall = ghData[i].PlacementOverall + i
	}

	// Add points by number of tags (more tags is better)
	sort.Slice(ghData[:], func(i, j int) bool {
		return ghData[i].TotalTags > ghData[j].TotalTags
	})
	greetings += fmt.Sprintf("* The project with more tags is `%s`\n", ghData[0].Name)
	for i := range ghData {
		ghData[i].PlacementTotalTags = ghData[i].PlacementTotalTags + i
		ghData[i].PlacementOverall = ghData[i].PlacementOverall + i
	}

	// Add points by Top10 contributors followers
	sort.Slice(ghData[:], func(i, j int) bool {
		return ghData[i].Top10ContributorsFollowers > ghData[j].Top10ContributorsFollowers
	})
	greetings += fmt.Sprintf("* The project made by most notable top contributors is `%s`\n", ghData[0].Name)
	for i := range ghData {
		ghData[i].PlacementTop10ContributorsFollowers = ghData[i].PlacementTop10ContributorsFollowers + i
		ghData[i].PlacementOverall = ghData[i].PlacementOverall + i
	}

	// Add points by Top10 contributors followers
	sort.Slice(ghData[:], func(i, j int) bool {
		return ghData[i].ClosedIssuesPercentage > ghData[j].ClosedIssuesPercentage
	})
	greetings += fmt.Sprintf("* The project with best errors resolving rate is `%s`\n", ghData[0].Name)
	for i := range ghData {
		ghData[i].PlacementClosedIssuesPercentage = ghData[i].PlacementClosedIssuesPercentage + i
		ghData[i].PlacementOverall = ghData[i].PlacementOverall + i
	}

	// Add points by commits by day (more commits shows good healthy community)
	sort.Slice(ghData[:], func(i, j int) bool {
		return ghData[i].CommitsByDay > ghData[j].CommitsByDay
	})
	greetings += fmt.Sprintf("* The project with more commits by day is `%s`\n", ghData[0].Name)
	for i := range ghData {
		ghData[i].PlacementCommitsByDay = ghData[i].PlacementCommitsByDay + i
		ghData[i].PlacementOverall = ghData[i].PlacementOverall + i
	}

	// Add points by active forkers (more active forkers shows good open source spirit of the community)
	sort.Slice(ghData[:], func(i, j int) bool {
		return ghData[i].ActiveForkersPercentage > ghData[j].ActiveForkersPercentage
	})
	greetings += fmt.Sprintf("* The project with the most active community is `%s`\n", ghData[0].Name)
	for i := range ghData {
		ghData[i].PlacementActiveForkersColumn = ghData[i].PlacementActiveForkersColumn + i
		ghData[i].PlacementOverall = ghData[i].PlacementOverall + i
	}

	// Assign places to projects by all metrics
	sort.Slice(ghData[:], func(i, j int) bool {
		return ghData[i].PlacementOverall < ghData[j].PlacementOverall
	})
	greetings += fmt.Sprintf("* The best project (taking in account placements in all competitions) is `%s`\n", ghData[0].Name)

	return greetings
}
