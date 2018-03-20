// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.
package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/tidwall/gjson"
)

// Repository structure with selcted data keys for JSON processing
type Repository struct {
	Name       string    `json:"name"`
	FullName   string    `json:"full_name"`
	Watchers   int       `json:"watchers"`
	Forks      int       `json:"forks"`
	OpenIssues int       `json:"open_issues"`
	CreatedAt  time.Time `json:"created_at"`
}

func getRepositoryTotalIssues(repoKey string, tmpFolder string, debug bool) int64 {
	url := "https://api.github.com/search/issues?q=repo:" + repoKey + "+type:issue+state:closed"
	fullResp := MakeCachedHTTPRequest(url, tmpFolder, debug)
	jsonResponse, _, _ := ReadResp(fullResp)
	totalIssuesResult := gjson.Get(string(jsonResponse), "total_count")
	//fmt.Printf("%d\n", totalIssuesResult.Int())
	return totalIssuesResult.Int()
}

func getRepositoryData(repoKey string, tmpFolder string, debug bool) []byte {
	url := "https://api.github.com/repos/" + repoKey
	fullResp := MakeCachedHTTPRequest(url, tmpFolder, debug)
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

func getRepositoryStatistics(RepoKey string, tmpFolder string, debug bool) *Repository {
	return parseRepositoryData(getRepositoryData(RepoKey, tmpFolder, debug))
}

func getClosedIssuesPercentage(openIssues int, totalIssues int) float64 {
	openIssuesFloat := float64(openIssues)
	totalIssuesFloat := float64(totalIssues)
	closedIssuesPercentage := 100 - (openIssuesFloat/totalIssuesFloat)*100
	return closedIssuesPercentage
}
