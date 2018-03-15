// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.
package main

import (
	"encoding/csv"
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
	"kataras/iris",
}

func main() {
	var totalPointsColumnIndex = 11
	var csvData = [][]string{}
	headers := []string{
		"Name",
		"URL",
		"Created at",
		"Age in days",
		"Watchers",
		"Forks",
		"Contributors",
		"Active forkers, %",
		"Open issues",
		"Total issues",
		"Closed issues, %",
		"Place",
	}
	for _, rKey := range repositoriesKeys {
		repositoryData := getRepositoryStatistics(rKey)
		totalIssues := getRepositoryTotalIssues(rKey)
		contributors := getRepositoryContributors(rKey)
		activeForkersPercentage := getActiveForkersPercentage(contributors, repositoryData.Forks)
		closedIssuesPercentage := getClosedIssuesPercentage(repositoryData.OpenIssues, int(totalIssues))
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
		})
	}

	// Add points by active forkers
	csvData = addPoints(sortSliceByColumnIndexFloatDesc(csvData, 7), 7, totalPointsColumnIndex)
	// Add points by proportion of total and resolved issues
	csvData = addPoints(sortSliceByColumnIndexFloatDesc(csvData, 10), 10, totalPointsColumnIndex)
	// Add points by age (we like fresh ideas)
	csvData = addPoints(sortSliceByColumnIndexIntAsc(csvData, 3), 3, totalPointsColumnIndex)
	// Add points by total populatiry
	csvData = addPoints(sortSliceByColumnIndexIntDesc(csvData, 4), 4, totalPointsColumnIndex)

	csvData = sortSliceByColumnIndexFloatDesc(csvData, 11)

	csvData = assignPlaces(csvData, totalPointsColumnIndex)

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
