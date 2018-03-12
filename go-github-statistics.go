// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.
package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/tidwall/gjson"
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

// Repository structure with selcted data keys
type Repository struct {
	Name       string    `json:"name"`
	FullName   string    `json:"full_name"`
	Watchers   int       `json:"watchers"`
	Forks      int       `json:"forks"`
	OpenIssues int       `json:"open_issues"`
	CreatedAt  time.Time `json:"created_at"`
}

type Contributor struct {
	Login string `json:"login"`
}

func getRepositoryContributorsNumber(repoKey string) int {
	var totalContributors int
	url := "https://api.github.com/repos/" + repoKey + "/contributors"
	fullResp := MakeCachedHTTPRequest(url)
	jsonResponse, linkHeader, _ := ReadResp(fullResp)
	var compRegEx = regexp.MustCompile(`.*page=(\d+).*page=(\d+).*`)
	match := compRegEx.FindStringSubmatch(linkHeader)
	nextPage := 0
	lastPage := 0
	for range compRegEx.SubexpNames() {
		if len(match) == 3 {
			nextPage, _ = strconv.Atoi(match[1])
			lastPage, _ = strconv.Atoi(match[2])
		}
	}
	//fmt.Printf("%d / %d\n", nextPage, lastPage)
	if nextPage == 0 {
		contributors := make([]Contributor, 0)
		json.Unmarshal(jsonResponse, &contributors)
		totalContributors = len(contributors)
	} else {
		// TODO :
		// get with regexps last URL
		// make an additional query to the last page
		// count contributors, add to the base
		totalContributors = (lastPage - 1) * 30
		totalContributors++
	}
	return totalContributors
}

func getRepositoryTotalIssues(repoKey string) int64 {
	url := "https://api.github.com/search/issues?q=repo:" + repoKey + "+type:issue+state:closed"
	fullResp := MakeCachedHTTPRequest(url)
	jsonResponse, _, _ := ReadResp(fullResp)
	totalIssuesResult := gjson.Get(string(jsonResponse), "total_count")
	//fmt.Printf("%d\n", totalIssuesResult.Int())
	return totalIssuesResult.Int()
}

func getRepositoryData(repoKey string) []byte {
	url := "https://api.github.com/repos/" + repoKey
	fullResp := MakeCachedHTTPRequest(url)
	jsonResponse, _, _ := ReadResp(fullResp)
	return jsonResponse
}

func parseRepositoryData(jsonResponse []byte) *Repository {
	result := &Repository{}
	err := json.Unmarshal([]byte(jsonResponse), result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

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
		repository.FullName,
		fmt.Sprintf("%d/%02d", repository.CreatedAt.Year(), repository.CreatedAt.Month()),
		fmt.Sprintf("%d", repository.Watchers),
		fmt.Sprintf("%d", repository.Forks),
		fmt.Sprintf("%d", repository.OpenIssues),
		fmt.Sprintf("%d", totalIssues),
		fmt.Sprintf("%d", contributorsNumber),
	})
}

func getRepositoryStatistics(RepoKey string) *Repository {
	return parseRepositoryData(getRepositoryData(RepoKey))
}

func main() {
	csvData = append(csvData, []string{"Name", "Full name", "Created at", "Watchers", "Forks", "Open Issues", "Total Issues", "Total contributors"})
	for _, rKey := range repositoriesKeys {
		repositoryData := getRepositoryStatistics(rKey)
		totalIssues := getRepositoryTotalIssues(rKey)
		contributorsNumber := getRepositoryContributorsNumber(rKey)
		fillCSVData(repositoryData, totalIssues, contributorsNumber)
	}
	writeCsv()
}
