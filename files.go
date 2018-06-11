// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
)

func clearHTTPCacheFolder(tmpFolderPath string, dryRun bool) error {
	d, err := os.Open(tmpFolderPath)
	if err != nil {
		log.Fatalf("Could not open %s", tmpFolderPath)
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		log.Fatalf("Could not read from %s", tmpFolderPath)
	}
	for _, name := range names {
		fp := filepath.Join(tmpFolderPath, name)
		if dryRun {
			fmt.Printf("Deleted %s\n", fp)
		} else {
			err = os.RemoveAll(fp)
			if err != nil {
				log.Fatalf("Could not remove %s", fp)
			}
			fmt.Printf("Deleted %s\n", fp)
		}
	}
	return nil
}

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
		fmt.Sprintf("%s", r.AuthorLocation),
		fmt.Sprintf("%s", r.MainLanguage),
		fmt.Sprintf("%s", r.AllLanguages),
		fmt.Sprintf("%s", r.Description),
		fmt.Sprintf("%d", r.TotalCodeSize),
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
		fmt.Sprintf("%d", r.AverageContributionPeriod),
		fmt.Sprintf("%d", r.MediCommitSize),
		fmt.Sprintf("%d", r.TotalTags),
		fmt.Sprintf("%d", r.Watchers),
		fmt.Sprintf("%d", r.Forks),
		fmt.Sprintf("%d", r.Contributors),
		fmt.Sprintf("%.2f", r.ActiveForkersPercentage),
		fmt.Sprintf("%d", r.ReturningContributors),
		fmt.Sprintf("%d", r.OpenIssues),
		fmt.Sprintf("%d", r.ClosedIssues),
		fmt.Sprintf("%d", r.TotalIssues),
		fmt.Sprintf("%.4f", r.IssueByDay),
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

func writeCsv(csvFilePath string, ghDataCSV [][]string) {
	file, err := os.Create(csvFilePath)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, value := range ghDataCSV {
		err := writer.Write(value)
		if err != nil {
			log.Fatal("Cannot write to file", err)
		}
	}
}
