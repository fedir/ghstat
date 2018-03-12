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

func fillCSVData(repository *Repository, totalIssues int64, contributorsNumber int) {
	csvData = append(csvData, []string{
		repository.Name,
		fmt.Sprintf("https://github.com/%s", repository.FullName),
		fmt.Sprintf("%d/%02d", repository.CreatedAt.Year(), repository.CreatedAt.Month()),
		fmt.Sprintf("%d", repository.Watchers),
		fmt.Sprintf("%d", repository.Forks),
		fmt.Sprintf("%d", contributorsNumber),
		fmt.Sprintf("%d", repository.OpenIssues),
		fmt.Sprintf("%d", totalIssues),
	})
}

func main() {
	csvData = append(csvData, []string{"Name", "URL", "Created at", "Watchers", "Forks", "Contributors", "Open Issues", "Total Issues"})
	for _, rKey := range repositoriesKeys {
		repositoryData := getRepositoryStatistics(rKey)
		totalIssues := getRepositoryTotalIssues(rKey)
		contributorsNumber := getRepositoryContributorsNumber(rKey)
		fillCSVData(repositoryData, totalIssues, contributorsNumber)
	}
	writeCsv()
}
