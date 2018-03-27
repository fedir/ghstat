// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"regexp"
	"strconv"

	"github.com/fedir/ghstat/httpcache"
)

// Commit structure with selcted data keys for JSON processing
type Commit struct {
	Author struct {
		Login string `json:"login"`
	} `json:"author"`
}

func getRepositoryCommits(repoKey string, tmpFolder string, debug bool) string {
	var total int
	var commitAuthorLogin string
	url := "https://api.github.com/repos/" + repoKey + "/commits"
	fullResp := httpcache.MakeCachedHTTPRequest(url, tmpFolder, debug)
	jsonResponse, linkHeader, _ := httpcache.ReadResp(fullResp)
	var compRegEx = regexp.MustCompile(regexpPageIndexes)
	match := compRegEx.FindStringSubmatch(linkHeader)
	nextPage := 0
	for range compRegEx.SubexpNames() {
		if len(match) == 3 {
			nextPage, _ = strconv.Atoi(match[1])
		}
	}
	if nextPage == 0 {
		commits := make([]Commit, 0)
		json.Unmarshal(jsonResponse, &commits)
		total = len(commits)
		commitAuthorLogin = getCommitAuthorLogin(commits[total-1])
	} else {
		commitAuthorLogin = getRepositoryFirstCommitAuthorLogin(linkHeader, tmpFolder, debug)
	}
	return commitAuthorLogin
}

func getRepositoryFirstCommitAuthorLogin(linkHeader string, tmpFolder string, debug bool) string {
	var commitAuthorLogin string
	compRegExLastURL := regexp.MustCompile(regexpLastPageURL)
	matchLastURL := compRegExLastURL.FindStringSubmatch(linkHeader)
	lastPageURL := matchLastURL[1]
	fullResp := httpcache.MakeCachedHTTPRequest(lastPageURL, tmpFolder, debug)
	jsonResponse, _, _ := httpcache.ReadResp(fullResp)
	commits := make([]Commit, 0)
	json.Unmarshal(jsonResponse, &commits)
	commitsOnLastPage := len(commits)
	commitAuthorLogin = getCommitAuthorLogin(commits[commitsOnLastPage-1])
	return commitAuthorLogin
}

func getCommitAuthorLogin(c Commit) string {
	return c.Author.Login
}

func getUserFollowers(username string, tmpFolder string, debug bool) int {
	var total int
	url := "https://api.github.com/users/" + username + "/followers"
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
	if nextPage == 0 {
		contributors := make([]Contributor, 0)
		json.Unmarshal(jsonResponse, &contributors)
		total = len(contributors)
	} else {
		itemsNumberOnLastPage := getItemsNumberOnLastPage(linkHeader, tmpFolder, debug)
		total = (lastPage-1)*30 + itemsNumberOnLastPage
	}
	return total
}

func getItemsNumberOnLastPage(linkHeader string, tmpFolder string, debug bool) int {
	compRegExLastURL := regexp.MustCompile(regexpLastPageURL)
	matchLastURL := compRegExLastURL.FindStringSubmatch(linkHeader)
	lastPageURL := matchLastURL[1]
	fullResp := httpcache.MakeCachedHTTPRequest(lastPageURL, tmpFolder, debug)
	jsonResponse, _, _ := httpcache.ReadResp(fullResp)
	items := make([]Contributor, 0)
	json.Unmarshal(jsonResponse, &items)
	itemsNumberOnLastPage := len(items)
	return itemsNumberOnLastPage
}
