// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
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

var csvData = [][]string{}

func writeCsv() {
	file, err := os.Create("result.csv")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, value := range csvData {
		err := writer.Write(value)
		if err != nil {
			log.Fatal("Cannot write to file", err)
		}
	}
}

func fillCSVData(repository *Repository, totalIssues int64, closedIssuesPercentage float64, contributorsNumber int, activeForkersPercentage float64) {
	csvData = append(csvData, []string{
		repository.Name,
		fmt.Sprintf("https://github.com/%s", repository.FullName),
		fmt.Sprintf("%d/%02d", repository.CreatedAt.Year(), repository.CreatedAt.Month()),
		fmt.Sprintf("%d", repository.Watchers),
		fmt.Sprintf("%d", repository.Forks),
		fmt.Sprintf("%d", contributorsNumber),
		fmt.Sprintf("%.2f", activeForkersPercentage),
		fmt.Sprintf("%d", repository.OpenIssues),
		fmt.Sprintf("%d", totalIssues),
		fmt.Sprintf("%.2f", closedIssuesPercentage),
	})
}

func main() {
	csvData = append(csvData, []string{"Name", "URL", "Created at", "Watchers", "Forks", "Contributors", "Active forkers, %", "Open Issues", "Total Issues", "Closed issues, %"})
	for _, rKey := range repositoriesKeys {
		repositoryData := getRepositoryStatistics(rKey)
		totalIssues := getRepositoryTotalIssues(rKey)
		contributorsNumber := getRepositoryContributorsNumber(rKey)
		activeForkersPercentage := getActiveForkersPercentage(contributorsNumber, repositoryData.Forks)
		closedIssuesPercentage := getClosedIssuesPercentage(repositoryData.OpenIssues, int(totalIssues))
		fillCSVData(repositoryData, totalIssues, closedIssuesPercentage, contributorsNumber, activeForkersPercentage)
	}
	writeCsv()
}
