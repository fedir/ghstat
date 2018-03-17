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
		debug = flag.Bool("d", false, "Debug mode")
	)
	flag.Parse()

	var csvData = [][]string{}
	const (
		NameColumn                   = 0
		AgeColumn                    = 4
		TotalCommitsColumn           = 5
		StargazersColumn             = 6
		ActiveForkersColumn          = 9
		ClosedIssuesPercentageColumn = 12
		TotalPointsColumnIndex       = 13
	)
	headers := []string{
		"Name",
		"URL",
		"Author",
		"Created at",
		"Age in days",
		"Total commits",
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
		totalIssues := getRepositoryTotalIssues(rKey, *debug)
		contributors := getRepositoryContributors(rKey, *debug)
		activeForkersPercentage := getActiveForkersPercentage(contributors, repositoryData.Forks)
		closedIssuesPercentage := getClosedIssuesPercentage(repositoryData.OpenIssues, int(totalIssues))
		contributionStatistics := getContributionStatistics(rKey, *debug)
		csvData = append(csvData, []string{
			repositoryData.Name,
			fmt.Sprintf("https://github.com/%s", repositoryData.FullName),
			fmt.Sprintf("%d/%02d", repositoryData.CreatedAt.Year(), repositoryData.CreatedAt.Month()),
			fmt.Sprintf("%d", int(time.Since(repositoryData.CreatedAt).Seconds()/86400)),
			fmt.Sprintf("%d", repositoryData.Watchers),
			fmt.Sprintf("%d", repositoryData.Forks),
			fmt.Sprintf("%d", contributors),
			fmt.Sprintf("%.2f", activeForkersPercentage),
			fmt.Sprintf("%d", repositoryData.OpenIssues),
			fmt.Sprintf("%d", totalIssues),
			fmt.Sprintf("%.2f", closedIssuesPercentage),
			"0",
			fmt.Sprintf("%s", func(a string) string {
				if a == "" {
					a = "[Account removed]"
				}
				return a
			}(authorLogin)),
			fmt.Sprintf("%d", contributionStatistics.TotalCommits),
		})
	}

	// Add points by total populatiry
	csvData = addPoints(sortSliceByColumnIndexIntDesc(csvData, StargazersColumn), StargazersColumn, TotalPointsColumnIndex)
	firstPlaceGreeting(csvData, "The most popular project is")

	// Add points by age (we like fresh ideas)
	csvData = addPoints(sortSliceByColumnIndexIntAsc(csvData, AgeColumn), AgeColumn, TotalPointsColumnIndex)
	firstPlaceGreeting(csvData, "The newest project is")

	// Add points by active forkers
	csvData = addPoints(sortSliceByColumnIndexFloatDesc(csvData, ActiveForkersColumn), ActiveForkersColumn, TotalPointsColumnIndex)
	firstPlaceGreeting(csvData, "The project with the most active community is")

	// Add points by proportion of total and resolved issues
	csvData = addPoints(sortSliceByColumnIndexFloatDesc(csvData, ClosedIssuesPercentageColumn), ClosedIssuesPercentageColumn, TotalPointsColumnIndex)
	firstPlaceGreeting(csvData, "The project with best errors resolving rate is")

	// Add points by number of commits
	csvData = addPoints(sortSliceByColumnIndexIntDesc(csvData, TotalCommitsColumn), TotalCommitsColumn, TotalPointsColumnIndex)
	firstPlaceGreeting(csvData, "The project with more commits is")

	csvData = sortSliceByColumnIndexIntAsc(csvData, 11)

	csvData = assignPlaces(csvData, TotalPointsColumnIndex)
	firstPlaceGreeting(csvData, "The best project (taking in account placements in all competitions) is")

	writeCsv(headers, csvData)
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
