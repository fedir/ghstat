// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
)

func writeCSVStatistics(ghData []Repository, csvFilePath string) {
	var csvData [][]string
	csvData = append(csvData, headersFromStructTags())
	for _, r := range ghData {
		csvData = append(csvData, formatRepositoryDataForCSV(r))
	}
	writeCsv(csvFilePath, csvData)
}

func formatRepositoryDataForCSV(r Repository) []string {
	ghProjectData := []string{
		r.Name,
		fmt.Sprintf("%s", r.URL),
		fmt.Sprintf("%s", r.Author),
		fmt.Sprintf("%s", r.Language),
		fmt.Sprintf("%s", r.License),
		fmt.Sprintf("%d", r.AuthorsFollowers),
		fmt.Sprintf("%d", r.Top10ContributorsFollowers),
		fmt.Sprintf("%d/%02d", r.CreatedAt.Year(), r.CreatedAt.Month()),
		fmt.Sprintf("%d", r.Age),
		fmt.Sprintf("%d", r.TotalCommits),
		fmt.Sprintf("%d", r.TotalAdditions),
		fmt.Sprintf("%d", r.TotalDeletions),
		fmt.Sprintf("%d", r.TotalCodeChanges),
		fmt.Sprintf(r.LastCommitDate.Format("2006-01-02 15:04:05")),
		fmt.Sprintf("%.4f", r.CommitsByDay),
		fmt.Sprintf("%d", r.MediCommitSize),
		fmt.Sprintf("%d", r.TotalTags),
		fmt.Sprintf("%d", r.Watchers),
		fmt.Sprintf("%d", r.Forks),
		fmt.Sprintf("%d", r.Contributors),
		fmt.Sprintf("%.2f", r.ActiveForkersPercentage),
		fmt.Sprintf("%d", r.OpenIssues),
		fmt.Sprintf("%d", r.TotalIssues),
		fmt.Sprintf("%.4f", r.IssueByDay),
		fmt.Sprintf("%.2f", r.ClosedIssuesPercentage),
		fmt.Sprintf("%.2f", r.ClosedIssuesPercentage),
		fmt.Sprintf("%d", r.PlacementPopularity),
		fmt.Sprintf("%d", r.PlacementAge),
		fmt.Sprintf("%d", r.PlacementTotalCommits),
		fmt.Sprintf("%d", r.PlacementTotalTags),
		fmt.Sprintf("%d", r.PlacementTop10ContributorsFollowers),
		fmt.Sprintf("%d", r.PlacementClosedIssuesPercentage),
		fmt.Sprintf("%d", r.PlacementCommitsByDay),
		fmt.Sprintf("%d", r.PlacementActiveForkersColumn),
		fmt.Sprintf("%d", r.PlacementOverall),
	}
	return ghProjectData
}

func headersFromStructTags() []string {
	r := new(Repository)
	return r.reflectRepositoryHeaders()
}

func (f *Repository) reflectRepositoryHeaders() []string {
	var headers []string
	val := reflect.ValueOf(f).Elem()
	for i := 0; i < val.NumField(); i++ {
		headers = append(headers, val.Type().Field(i).Tag.Get("header"))
	}
	return headers
}

func writeCsv(csvFilePath string, csvData [][]string) {
	file, err := os.Create(csvFilePath)
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
