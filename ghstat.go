// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var repositoriesKeys = []string{
	"astaxie/beego",
	"gohugoio/hugo",
	"gin-gonic/gin",
	"labstack/echo",
	"revel/revel",
	"gobuffalo/buffalo",
	"go-chi/chi",
}

func main() {
	var (
		debug                  = flag.Bool("d", false, "Debug mode")
		repositoriesKeysManual = flag.String("r", "", "Repositories keys")
	)
	flag.Parse()

	if *repositoriesKeysManual != "" {
		repositoriesKeys = strings.Split(*repositoriesKeysManual, ",")
	}

	var csvData = [][]string{}
	const (
		NameColumn                   = 0
		AuthorsFollowersColumn       = 4
		AgeColumn                    = 5
		TotalCommitsColumn           = 6
		TotalAdditionsColumn         = 7
		TotalDeletionsColumn         = 8
		TotalCodeChangesColumn       = 9
		MediCommitSizeColumn         = 10
		StargazersColumn             = 11
		ActiveForkersColumn          = 14
		ClosedIssuesPercentageColumn = 17
		TotalPointsColumnIndex       = 18
	)
	headers := []string{
		"Name",
		"URL",
		"Author",
		"Author's followers",
		"Created at",
		"Age in days",
		"Total commits",
		"Total additions",
		"Total deletions",
		"Total code changes",
		"Medium commit size",
		"Stargazers",
		"Forks",
		"Contributors",
		"Active forkers, %",
		"Open issues",
		"Total issues",
		"Closed issues, %",
		"Place",
	}
	for _, rKey := range repositoriesKeys {
		repositoryData := getRepositoryStatistics(rKey, *debug)
		authorLogin := getRepositoryCommits(rKey, *debug)
		authorFollowers := 0
		if authorLogin != "" {
			authorFollowers = getUserFollowers(authorLogin, *debug)
		}
		totalIssues := getRepositoryTotalIssues(rKey, *debug)
		contributors := getRepositoryContributors(rKey, *debug)
		activeForkersPercentage := getActiveForkersPercentage(contributors, repositoryData.Forks)
		closedIssuesPercentage := getClosedIssuesPercentage(repositoryData.OpenIssues, int(totalIssues))
		contributionStatistics := getContributionStatistics(rKey, *debug)
		csvData = append(csvData, []string{
			repositoryData.Name,
			fmt.Sprintf("https://github.com/%s", repositoryData.FullName),
			fmt.Sprintf("%s", func(a string) string {
				if a == "" {
					a = "[Account removed]"
				}
				return a
			}(authorLogin)),
			fmt.Sprintf("%d", authorFollowers),
			fmt.Sprintf("%d/%02d", repositoryData.CreatedAt.Year(), repositoryData.CreatedAt.Month()),
			fmt.Sprintf("%d", int(time.Since(repositoryData.CreatedAt).Seconds()/86400)),
			fmt.Sprintf("%d", contributionStatistics.TotalCommits),
			fmt.Sprintf("%d", contributionStatistics.TotalAdditions),
			fmt.Sprintf("%d", contributionStatistics.TotalDeletions),
			fmt.Sprintf("%d", contributionStatistics.TotalCodeChanges),
			fmt.Sprintf("%d", contributionStatistics.MediumCommitSize),
			fmt.Sprintf("%d", repositoryData.Watchers),
			fmt.Sprintf("%d", repositoryData.Forks),
			fmt.Sprintf("%d", contributors),
			fmt.Sprintf("%.2f", activeForkersPercentage),
			fmt.Sprintf("%d", repositoryData.OpenIssues),
			fmt.Sprintf("%d", totalIssues),
			fmt.Sprintf("%.2f", closedIssuesPercentage),
			"0",
		})
	}

	var csvDataSorted, csvDataTotal [][]string

	// Add points by author's followers
	csvDataSorted = sortSliceByColumnIndexIntDesc(csvData, AuthorsFollowersColumn)
	firstPlaceGreeting(csvDataSorted, "The project made by most notable author is")
	csvDataTotal = addPoints(csvDataSorted, AuthorsFollowersColumn, TotalPointsColumnIndex)

	// Add points by reposotory total populatiry
	csvDataSorted = sortSliceByColumnIndexIntDesc(csvData, StargazersColumn)
	firstPlaceGreeting(csvDataSorted, "The most popular project is")
	csvDataTotal = addPoints(csvDataSorted, StargazersColumn, TotalPointsColumnIndex)

	// Add points by age (we like fresh ideas)
	csvDataSorted = sortSliceByColumnIndexIntAsc(csvData, AgeColumn)
	firstPlaceGreeting(csvDataSorted, "The newest project is")
	csvDataTotal = addPoints(csvDataSorted, AgeColumn, TotalPointsColumnIndex)

	// Add points by active forkers
	csvDataSorted = sortSliceByColumnIndexFloatDesc(csvData, ActiveForkersColumn)
	firstPlaceGreeting(csvDataSorted, "The project with the most active community is")
	csvDataTotal = addPoints(csvDataSorted, ActiveForkersColumn, TotalPointsColumnIndex)

	// Add points by proportion of total and resolved issues
	csvDataSorted = sortSliceByColumnIndexFloatDesc(csvData, ClosedIssuesPercentageColumn)
	firstPlaceGreeting(csvDataSorted, "The project with best errors resolving rate is")
	csvDataTotal = addPoints(csvDataSorted, ClosedIssuesPercentageColumn, TotalPointsColumnIndex)

	// Add points by number of commits
	csvDataSorted = sortSliceByColumnIndexIntDesc(csvData, TotalCommitsColumn)
	firstPlaceGreeting(csvDataSorted, "The project with more commits is")
	csvDataTotal = addPoints(csvDataSorted, TotalCommitsColumn, TotalPointsColumnIndex)

	csvDataTotal = sortSliceByColumnIndexIntAsc(csvDataTotal, TotalPointsColumnIndex)
	firstPlaceGreeting(csvDataSorted, "The best project (taking in account placements in all competitions) is")
	csvDataTotal = assignPlaces(csvData, TotalPointsColumnIndex)

	writeCsv(headers, csvDataTotal)
}

func writeCsv(headers []string, csvData [][]string) {
	file, err := os.Create("result.csv")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(headers)
	if err != nil {
		log.Fatal("Cannot write to file", err)
	}

	for _, value := range csvData {
		err := writer.Write(value)
		if err != nil {
			log.Fatal("Cannot write to file", err)
		}
	}
}
