// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.

// ghstat - statistical multi-criteria decision-making comparator for Github's projects.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fedir/ghstat/github"
)

// Repository structure with selcted data keys for JSON processing
type Repository struct {
	Name                                string    `header:"Name"`
	URL                                 string    `header:"URL"`
	Author                              string    `header:"Author"`
	Language                            string    `header:"Language"`
	License                             string    `header:"License"`
	AuthorsFollowers                    int       `header:"Author's followers"`
	Top10ContributorsFollowers          int       `header:"Top 10 contributors followers"`
	CreatedAt                           time.Time `header:"Created at"`
	Age                                 int       `header:"Age in days"`
	TotalCommits                        int       `header:"Total commits"`
	TotalAdditions                      int       `header:"Total additions"`
	TotalDeletions                      int       `header:"Total deletions"`
	TotalCodeChanges                    int       `header:"Total code changes"`
	LastCommitDate                      time.Time `header:"Last commit date"`
	CommitsByDay                        float64   `header:"Commits/day"`
	MediCommitSize                      int       `header:"Medium commit size"`
	TotalTags                           int       `header:"Total releases"`
	Watchers                            int       `header:"Stargazers"`
	Forks                               int       `header:"Forks"`
	Contributors                        int       `header:"Contributors"`
	ActiveForkersPercentage             float64   `header:"Active forkers, %"`
	OpenIssues                          int       `header:"Open issues"`
	ClosedIssues                        int       `header:"Closed issues"`
	TotalIssues                         int       `header:"Total issues"`
	IssueByDay                          float64   `header:"Issue/day"`
	ClosedIssuesPercentage              float64   `header:"Closed issues, %"`
	PlacementPopularity                 int       `header:"Placement by popularity"`
	PlacementAge                        int       `header:"Placement by age"`
	PlacementTotalCommits               int       `header:"Placement by total commits"`
	PlacementTotalTags                  int       `header:"Placement by total tags"`
	PlacementTop10ContributorsFollowers int       `header:"Placement by top 10 contributors followers"`
	PlacementClosedIssuesPercentage     int       `header:"Placement by closed issues percentage"`
	PlacementCommitsByDay               int       `header:"Placement by commits by day"`
	PlacementActiveForkersColumn        int       `header:"Placement by active forkers column"`
	PlacementOverall                    int       `header:"Placement overall"`
}

func main() {
	var (
		clearHTTPCache         = flag.Bool("cc", false, "Clear HTTP cache")
		clearHTTPCacheDryRun   = flag.Bool("ccdr", false, "Clear HTTP cache (dry run)")
		debug                  = flag.Bool("d", false, "Debug mode")
		resultFileSavePath     = flag.String("f", "", "File path where result CSV file will be saved")
		rateLimitCheck         = flag.Bool("l", false, "Rate limit check")
		repositoriesKeysManual = flag.String("r", "", "Repositories keys")
		tmpFolder              = flag.String("t", "test_data", "Clear HTTP cache (dry run)")
		repositoriesKeys       = []string{}
	)
	flag.Parse()
	if *clearHTTPCache || *clearHTTPCacheDryRun {
		clearHTTPCacheFolder(*tmpFolder, *clearHTTPCacheDryRun)
		os.Exit(0)
	}
	if *rateLimitCheck {
		github.CheckAndPrintRateLimit()
		os.Exit(0)
	}
	if *repositoriesKeysManual != "" {
		repositoriesKeys = uniqSlice(strings.Split(*repositoriesKeysManual, ","))
	} else {
		repositoriesKeys = uniqSlice([]string{
			"astaxie/beego",
			"gohugoio/hugo",
			"gin-gonic/gin",
			"labstack/echo",
			"revel/revel",
			"gobuffalo/buffalo",
			"go-chi/chi",
			"kataras/iris",
			"zenazn/goji",
			"go-macaron/macaron",
		})
	}

	csvFilePath := ""
	if *resultFileSavePath != "" {
		csvFilePath = *resultFileSavePath
	} else {
		csvFilePath = "result.csv"
	}
	var ghData = []Repository{}

	dataChan := make(chan Repository, len(repositoriesKeys))
	for _, rKey := range repositoriesKeys {
		go repositoryData(rKey, *tmpFolder, *debug, dataChan)
	}
	for range repositoriesKeys {
		ghData = append(ghData, <-dataChan)
	}
	rateAndPrintGreetings(ghData)
	writeCSVStatistics(ghData, csvFilePath)
}

func repositoryData(rKey string, tmpFolder string, debug bool, dataChan chan Repository) {

	r := new(Repository)

	repositoryData := github.GetRepositoryStatistics(rKey, tmpFolder, debug)

	r.Name = repositoryData.FullName
	r.URL = fmt.Sprintf("https://github.com/%s", r.Name)
	r.Language = repositoryData.Language
	r.CreatedAt = repositoryData.CreatedAt
	r.Age = int(time.Since(repositoryData.CreatedAt).Seconds() / 86400)
	r.Watchers = repositoryData.Watchers
	r.Forks = repositoryData.Forks
	r.OpenIssues = repositoryData.OpenIssues
	r.License = "[Unknown]"
	if repositoryData.License.SPDXID != "" {
		r.License = repositoryData.License.SPDXID
	}
	r.Author = "[Unknown]"

	r.Author,
		r.LastCommitDate = github.GetRepositoryCommitsData(rKey, tmpFolder, debug)

	r.AuthorsFollowers = 0
	if r.Author != "" {
		r.AuthorsFollowers = github.GetUserFollowers(r.Author, tmpFolder, debug)
	} else {
		r.Author = "[Account removed]"
	}

	r.ClosedIssues = 0
	if repositoryData.HasIssues {
		r.ClosedIssues = github.GetRepositoryClosedIssues(rKey, tmpFolder, debug)
	}
	r.TotalIssues = r.OpenIssues + r.ClosedIssues
	r.Top10ContributorsFollowers,
		r.Contributors = github.GetRepositoryContributors(rKey, tmpFolder, debug)
	r.TotalTags = github.GetRepositoryTagsNumber(rKey, tmpFolder, debug)
	r.ActiveForkersPercentage = github.GetActiveForkersPercentage(r.Contributors, r.Forks)
	r.IssueByDay = github.GetIssueByDay(r.ClosedIssues+r.OpenIssues, r.Age)
	r.ClosedIssuesPercentage = github.GetClosedIssuesPercentage(repositoryData.OpenIssues, int(r.ClosedIssues))

	contributionStatistics := github.GetContributionStatistics(rKey, tmpFolder, debug)
	r.TotalCommits = contributionStatistics.TotalCommits
	r.TotalAdditions = contributionStatistics.TotalAdditions
	r.TotalDeletions = contributionStatistics.TotalDeletions
	r.TotalCodeChanges = contributionStatistics.TotalCodeChanges
	r.MediCommitSize = contributionStatistics.MediumCommitSize

	r.CommitsByDay = github.GetCommitsByDay(contributionStatistics.TotalCommits, r.Age)

	dataChan <- *r
}

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

func uniqSlice(s []string) []string {
	unique := make(map[string]bool, len(s))
	us := make([]string, len(unique))
	for _, elem := range s {
		if len(elem) != 0 {
			if !unique[elem] {
				us = append(us, elem)
				unique[elem] = true
			}
		}
	}
	return us
}
