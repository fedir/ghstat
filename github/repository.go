// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.

package github

import (
	"encoding/json"
	"log"
	"time"

	"github.com/fedir/ghstat/httpcache"
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

func GetRepositoryClosedIssues(repoKey string, tmpFolder string, debug bool) int {
	url := "https://api.github.com/search/issues?q=repo:" + repoKey + "+type:issue+state:closed"
	fullResp := httpcache.MakeCachedHTTPRequest(url, tmpFolder, debug)
	jsonResponse, _, _ := httpcache.ReadResp(fullResp)
	closedIssuesResult := gjson.Get(string(jsonResponse), "total_count")
	//fmt.Printf("%d\n", closedIssuesResult.Int())
	return int(closedIssuesResult.Int())
}

func GetRepositoryStatistics(RepoKey string, tmpFolder string, debug bool) *Repository {
	return ParseRepositoryData(getRepositoryData(RepoKey, tmpFolder, debug))
}

func getRepositoryData(repoKey string, tmpFolder string, debug bool) []byte {
	url := "https://api.github.com/repos/" + repoKey
	fullResp := httpcache.MakeCachedHTTPRequest(url, tmpFolder, debug)
	jsonResponse, _, _ := httpcache.ReadResp(fullResp)
	return jsonResponse
}

func ParseRepositoryData(jsonResponse []byte) *Repository {
	result := &Repository{}
	err := json.Unmarshal([]byte(jsonResponse), result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func GetIssueByDay(totalIssues int, age int) float64 {
	var issueByDay float64
	totalIssuesFloat := float64(totalIssues)
	ageFloat := float64(age)
	if totalIssuesFloat != 0 && ageFloat != 0 {
		issueByDay = totalIssuesFloat / ageFloat
	} else {
		issueByDay = 0
	}
	return issueByDay

}

func GetClosedIssuesPercentage(openIssues int, closedIssues int) float64 {
	var closedIssuesPercentage float64
	openIssuesFloat := float64(openIssues)
	closedIssuesFloat := float64(closedIssues)
	if closedIssuesFloat != 0 && openIssuesFloat != 0 {
		closedIssuesPercentage = closedIssuesFloat / (closedIssuesFloat + openIssuesFloat) * 100
	} else {
		closedIssuesPercentage = 100
	}
	return closedIssuesPercentage
}
