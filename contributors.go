// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.
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

func getRepositoryContributors(repoKey string, tmpFolder string, debug bool) (int, int) {
	//var topContributors []string
	var totalContributors int
	var topContributorsFollowers = 0
	url := "https://api.github.com/repos/" + repoKey + "/contributors"
	fullResp := MakeCachedHTTPRequest(url, tmpFolder, debug)
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

	contributors := make([]Contributor, 0)
	json.Unmarshal(jsonResponse, &contributors)
	i := 0
	for _, contributor := range contributors {
		//fmt.Printf("%s %s", contributor, index)
		//topContributors = append(topContributors, contributor.Login)
		topContributorsFollowers = topContributorsFollowers + getContributorFollowers(contributor.Login, tmpFolder, debug)
		i++
		if i == 10 {
			goto TOTAL_CONTRIBUTORS
		}
	}
TOTAL_CONTRIBUTORS:
	if nextPage != 0 {
		contributorsOnLastPage := getRepositoryContributorsNumberLastPage(linkHeader, tmpFolder, debug)
		totalContributors = (lastPage-1)*30 + contributorsOnLastPage
	} else {
		totalContributors = len(contributors)
	}
	return topContributorsFollowers, totalContributors
}

func getContributorFollowers(login string, tmpFolder string, debug bool) int {
	totalUsers := 0
	url := "https://api.github.com/users/" + login + "/followers"
	fullResp := MakeCachedHTTPRequest(url, tmpFolder, debug)
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
	contributors := make([]Contributor, 0)
	json.Unmarshal(jsonResponse, &contributors)
	if nextPage != 0 {
		contributorsOnLastPage := getRepositoryContributorsNumberLastPage(linkHeader, tmpFolder, debug)
		totalUsers = (lastPage-1)*30 + contributorsOnLastPage
	} else {
		totalUsers = len(contributors)
	}
	return totalUsers
}

func getRepositoryContributorsNumberLastPage(linkHeader string, tmpFolder string, debug bool) int {
	compRegExLastURL := regexp.MustCompile(regexpLastPageURL)
	matchLastURL := compRegExLastURL.FindStringSubmatch(linkHeader)
	lastPageURL := matchLastURL[1]
	fullResp := MakeCachedHTTPRequest(lastPageURL, tmpFolder, debug)
	jsonResponse, _, _ := ReadResp(fullResp)
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
