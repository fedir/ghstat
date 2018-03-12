package main

import (
	"encoding/json"
	"regexp"
	"strconv"
)

const (
	regexpPageIndexes = `.*page=(\d+).*page=(\d+).*`
	regexpLastPageURL = `.* rel="next", <(.*)>;.*`
)

// Contributor structure with selcted data keys for JSON processing
type Contributor struct {
	Login string `json:"login"`
}

func getRepositoryContributorsNumber(repoKey string) int {
	var totalContributors int
	url := "https://api.github.com/repos/" + repoKey + "/contributors"
	fullResp := MakeCachedHTTPRequest(url)
	jsonResponse, linkHeader, _ := ReadResp(fullResp)
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
	if nextPage == 0 {
		contributors := make([]Contributor, 0)
		json.Unmarshal(jsonResponse, &contributors)
		totalContributors = len(contributors)
	} else {
		contributorsOnLastPage := getRepositoryContributorsNumberLastPage(linkHeader)
		totalContributors = (lastPage-1)*30 + contributorsOnLastPage
	}
	return totalContributors
}

func getRepositoryContributorsNumberLastPage(linkHeader string) int {
	compRegExLastURL := regexp.MustCompile(regexpLastPageURL)
	matchLastURL := compRegExLastURL.FindStringSubmatch(linkHeader)
	lastPageURL := matchLastURL[1]
	fullResp := MakeCachedHTTPRequest(lastPageURL)
	jsonResponse, linkHeader, _ := ReadResp(fullResp)
	contributors := make([]Contributor, 0)
	json.Unmarshal(jsonResponse, &contributors)
	contributorsOnLastPage := len(contributors)
	return contributorsOnLastPage
}

func getActiveForkersPercentage(contributors int, forkers int) float64 {
	contributorsFloat := float64(contributors)
	forkersFloat := float64(forkers)
	activeForkersPercentage := (contributorsFloat / forkersFloat) * 100
	return activeForkersPercentage
}
