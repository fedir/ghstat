package main

import (
	"encoding/json"
	"regexp"
	"strconv"
)

const (
	regexpPageIndexes = `.*page=(\d+).*page=(\d+).*`
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
