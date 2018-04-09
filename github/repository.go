// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.

package github

import (
	"encoding/json"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/fedir/ghstat/httpcache"
	"github.com/tidwall/gjson"
)

// Repository structure with selcted data keys for JSON processing
type Repository struct {
	Name       string    `json:"name"`
	FullName   string    `json:"full_name"`
	Language   string    `json:"language"`
	Watchers   int       `json:"watchers"`
	Forks      int       `json:"forks"`
	OpenIssues int       `json:"open_issues"`
	CreatedAt  time.Time `json:"created_at"`
	License    struct {
		SPDXID string `json:"spdx_id"`
	} `json:"license"`
}

// Tag structure with selcted data keys for JSON processing
type Tag struct {
	Name string `json:"name"`
}

// GetRepositoryClosedIssues gets number of closed issues of a repository
func GetRepositoryClosedIssues(repoKey string, tmpFolder string, debug bool) int {
	url := "https://api.github.com/search/issues?q=repo:" + repoKey + "+type:issue+state:closed"
	fullResp := httpcache.MakeCachedHTTPRequest(url, tmpFolder, debug)
	jsonResponse, _, _ := httpcache.ReadResp(fullResp)
	closedIssuesResult := gjson.Get(string(jsonResponse), "total_count")
	return int(closedIssuesResult.Int())
}

// GetRepositoryStatistics gets repository common statistics
func GetRepositoryStatistics(RepoKey string, tmpFolder string, debug bool) *Repository {
	return ParseRepositoryData(getRepositoryData(RepoKey, tmpFolder, debug))
}

func getRepositoryData(repoKey string, tmpFolder string, debug bool) []byte {
	url := "https://api.github.com/repos/" + repoKey
	fullResp := httpcache.MakeCachedHTTPRequest(url, tmpFolder, debug)
	jsonResponse, _, _ := httpcache.ReadResp(fullResp)
	return jsonResponse
}

// ParseRepositoryData is used to parse repository common statistics
func ParseRepositoryData(jsonResponse []byte) *Repository {
	result := &Repository{}
	err := json.Unmarshal([]byte(jsonResponse), result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

// GetIssueByDay calculates the rate of issues by day of the repository
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

// GetClosedIssuesPercentage calculates the percentage of closed issues of the repository
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

// GetRepositoryTagsNumber gets information about tags of the repository
func GetRepositoryTagsNumber(repoKey string, tmpFolder string, debug bool) int {
	var totalTags int
	url := "https://api.github.com/repos/" + repoKey + "/tags"
	fullResp := httpcache.MakeCachedHTTPRequest(url, tmpFolder, debug)
	jsonResponse, linkHeader, _ := httpcache.ReadResp(fullResp)
	var compRegEx = regexp.MustCompile(regexpPageIndexes)
	match := compRegEx.FindStringSubmatch(linkHeader)
	nextPage := 0
	lastPage := 0
	for range compRegEx.SubexpNames() {
		if len(match) == 3 {
			nextPage, _ = strconv.Atoi(match[1])
			lastPage, _ = strconv.Atoi(match[2])
		}
	}
	tags := make([]Tag, 0)
	json.Unmarshal(jsonResponse, &tags)
	if nextPage != 0 {
		tagsOnLastPage := getRepositoryTagsNumberLastPage(linkHeader, tmpFolder, debug)
		totalTags = (lastPage-1)*30 + tagsOnLastPage
	} else {
		totalTags = len(tags)
	}
	return totalTags
}

func getRepositoryTagsNumberLastPage(linkHeader string, tmpFolder string, debug bool) int {
	jsonResponse := getJSONResponse(linkHeader, tmpFolder, debug)
	tags := make([]Tag, 0)
	json.Unmarshal(jsonResponse, &tags)
	tagsOnLastPage := len(tags)
	return tagsOnLastPage
}

func getJSONResponse(linkHeader string, tmpFolder string, debug bool) []byte {
	compRegExLastURL := regexp.MustCompile(regexpLastPageURL)
	matchLastURL := compRegExLastURL.FindStringSubmatch(linkHeader)
	fullResp := httpcache.MakeCachedHTTPRequest(matchLastURL[1], tmpFolder, debug)
	jsonResponse, _, _ := httpcache.ReadResp(fullResp)
	return jsonResponse
}
