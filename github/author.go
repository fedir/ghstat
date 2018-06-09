// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.

package github

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/fedir/ghstat/httpcache"
)

// Commit structure with selcted data keys for JSON processing
type Commit struct {
	Author struct {
		Login string `json:"login"`
		Date  string `json:"date"`
	} `json:"author"`
	Commit struct {
		Author struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"author"`
	} `json:"commit"`
}

// User structure with selcted data keys for JSON processing
type User struct {
	Login     string    `json:"login"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// GetRepositoryCommitsData gets information about commits of a repository.
// Currerntly used for author login and last commit date
func GetRepositoryCommitsData(repoKey string, tmpFolder string, debug bool) (string, time.Time) {
	var total int
	var authorLogin string
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
	lastCommitDate := getRepositoryLastCommitDate(jsonResponse)
	if nextPage == 0 {
		commits := make([]Commit, 0)
		json.Unmarshal(jsonResponse, &commits)
		total = len(commits)
		authorLogin = commits[total-1].Author.Login
	} else {
		authorLogin = getRepositoryFirstCommitAuthorLogin(linkHeader, tmpFolder, debug)
	}
	return authorLogin, lastCommitDate
}

func getRepositoryFirstCommitAuthorLogin(linkHeader string, tmpFolder string, debug bool) string {
	jsonResponse := getJSONResponse(linkHeader, tmpFolder, debug)
	commits := make([]Commit, 0)
	json.Unmarshal(jsonResponse, &commits)
	commitsOnLastPage := len(commits)
	commitAuthorLogin := commits[commitsOnLastPage-1].Author.Login
	return commitAuthorLogin
}

func getRepositoryLastCommitDate(jsonResponse []byte) time.Time {
	commits := make([]Commit, 0)
	json.Unmarshal(jsonResponse, &commits)
	return commits[0].Commit.Author.Date
}

// GetUserFollowers gets information about followers of a user
func GetUserFollowers(username string, tmpFolder string, debug bool) int {
	var total int
	url := "https://api.github.com/users/" + username + "/followers"
	fullResp := httpcache.MakeCachedHTTPRequest(url, tmpFolder, debug)
	if len(fullResp) == 0 {
		filename := httpcache.GetFilename(url)
		fmt.Printf("Response for %s is empty, please check %s", url, filename)
	}
	jsonResponse, linkHeader, err := httpcache.ReadResp(fullResp)
	if err != nil {
		fmt.Printf("Problem recovering %s", url)
	}
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

// GetUserData gets information about followers of a user
func GetUserData(username string, tmpFolder string, debug bool) *User {
	url := "https://api.github.com/users/" + username
	fullResp := httpcache.MakeCachedHTTPRequest(url, tmpFolder, debug)
	jsonResponse, _, _ := httpcache.ReadResp(fullResp)
	result := &User{}
	err := json.Unmarshal([]byte(jsonResponse), result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func getItemsNumberOnLastPage(linkHeader string, tmpFolder string, debug bool) int {
	jsonResponse := getJSONResponse(linkHeader, tmpFolder, debug)
	items := make([]Contributor, 0)
	json.Unmarshal(jsonResponse, &items)
	itemsNumberOnLastPage := len(items)
	return itemsNumberOnLastPage
}
