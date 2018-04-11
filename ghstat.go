// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.

// ghstat - statistical multi-criteria decision-making comparator for Github's projects.
package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fedir/ghstat/github"
	"github.com/fedir/ghstat/httpcache"
	"github.com/fedir/ghstat/timing"
)

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
		checkAndPrintRateLimit()
		os.Exit(0)
	}
	if *rateLimitCheck {
		checkAndPrintRateLimit()
		os.Exit(0)
	}
	if *repositoriesKeysManual != "" {
		repositoriesKeys = uniqSlice(strings.Split(*repositoriesKeysManual, ","))
	} else {
		repositoriesKeys = []string{
			"astaxie/beego",
			"gohugoio/hugo",
			"gin-gonic/gin",
			"labstack/echo",
			"revel/revel",
			"gobuffalo/buffalo",
			"go-chi/chi",
			"kataras/iris",
			"zenazn/goji",
		}
	}

	csvFilePath := ""
	if *resultFileSavePath != "" {
		csvFilePath = *resultFileSavePath
	} else {
		csvFilePath = "result.csv"
	}
	var ghData = [][]string{}
	headers := []string{
		"Name",
		"URL",
		"Author",
		"Language",
		"License",
		"Author's followers",
		"Top 10 contributors followers",
		"Created at",
		"Age in days",
		"Total commits",
		"Total additions",
		"Total deletions",
		"Total code changes",
		"Last commit date",
		"Commits/day",
		"Medium commit size",
		"Total releases",
		"Stargazers",
		"Forks",
		"Contributors",
		"Active forkers, %",
		"Open issues",
		"Total issues",
		"Issue/day",
		"Closed issues, %",
		"Place",
	}
	ghDataColumnIndexes := map[string]int{
		"nameColumn":                       0,
		"urlColumn":                        1,
		"authorColumn":                     2,
		"languageColumn":                   3,
		"licenseColumn":                    4,
		"authorsFollowersColumn":           5,
		"top10ContributorsFollowersColumn": 6,
		"ageColumn":                        8,
		"totalCommitsColumn":               9,
		"totalAdditionsColumn":             10,
		"totalDeletionsColumn":             11,
		"totalCodeChangesColumn":           12,
		"lastCommitDateColumn":             13,
		"commitsByDayColumn":               14,
		"mediCommitSizeColumn":             15,
		"totalTagsColumn":                  16,
		"stargazersColumn":                 17,
		"activeForkersColumn":              20,
		"issuesByDayColumn":                23,
		"closedIssuesPercentageColumn":     24,
		"totalPointsColumnIndex":           25,
	}
	dataChan := make(chan []string, len(repositoriesKeys))
	for _, rKey := range repositoriesKeys {
		go fillRepositoryStatistics(rKey, *tmpFolder, *debug, dataChan)
	}
	for range repositoriesKeys {
		ghData = append(ghData, <-dataChan)
	}
	greetings := rateGhData(ghData, ghDataColumnIndexes)
	fmt.Println(greetings)
	writeCsv(csvFilePath, headers, ghData)
}

func fillRepositoryStatistics(rKey string, tmpFolder string, debug bool, dataChan chan []string) {
	repositoryData := github.GetRepositoryStatistics(rKey, tmpFolder, debug)
	repositoryAge := int(time.Since(repositoryData.CreatedAt).Seconds() / 86400)
	authorLogin, lastCommitDate := github.GetRepositoryCommitsData(rKey, tmpFolder, debug)
	authorFollowers := 0
	if authorLogin != "" {
		authorFollowers = github.GetUserFollowers(authorLogin, tmpFolder, debug)
	}

	closedIssues := 0
	if repositoryData.HasIssues {
		closedIssues = github.GetRepositoryClosedIssues(rKey, tmpFolder, debug)
	}
	topContributorsFollowers, totalContributors := github.GetRepositoryContributors(rKey, tmpFolder, debug)
	totalTags := github.GetRepositoryTagsNumber(rKey, tmpFolder, debug)
	activeForkersPercentage := github.GetActiveForkersPercentage(totalContributors, repositoryData.Forks)
	issueByDay := github.GetIssueByDay(closedIssues+repositoryData.OpenIssues, repositoryAge)
	closedIssuesPercentage := github.GetClosedIssuesPercentage(repositoryData.OpenIssues, int(closedIssues))
	contributionStatistics := github.GetContributionStatistics(rKey, tmpFolder, debug)
	commitsByDay := github.GetCommitsByDay(contributionStatistics.TotalCommits, repositoryAge)
	ghProjectData := []string{
		repositoryData.FullName,
		fmt.Sprintf("https://github.com/%s", repositoryData.FullName),
		fmt.Sprintf("%s", repositoryData.Language),
		fmt.Sprintf("%s", func(a string) string {
			if a == "" {
				a = "[Account removed]"
			}
			return a
		}(authorLogin)),
		fmt.Sprintf("%s", func(l string) string {
			if l == "" {
				l = "[Unknown]"
			}
			return l
		}(repositoryData.License.SPDXID)),
		fmt.Sprintf("%d", authorFollowers),
		fmt.Sprintf("%d", topContributorsFollowers),
		fmt.Sprintf("%d/%02d", repositoryData.CreatedAt.Year(), repositoryData.CreatedAt.Month()),
		fmt.Sprintf("%d", repositoryAge),
		fmt.Sprintf("%d", contributionStatistics.TotalCommits),
		fmt.Sprintf("%d", contributionStatistics.TotalAdditions),
		fmt.Sprintf("%d", contributionStatistics.TotalDeletions),
		fmt.Sprintf("%d", contributionStatistics.TotalCodeChanges),
		fmt.Sprintf(lastCommitDate.Format("2006-01-02 15:04:05")),
		fmt.Sprintf("%.4f", commitsByDay),
		fmt.Sprintf("%d", contributionStatistics.MediumCommitSize),
		fmt.Sprintf("%d", totalTags),
		fmt.Sprintf("%d", repositoryData.Watchers),
		fmt.Sprintf("%d", repositoryData.Forks),
		fmt.Sprintf("%d", totalContributors),
		fmt.Sprintf("%.2f", activeForkersPercentage),
		fmt.Sprintf("%d", repositoryData.OpenIssues),
		fmt.Sprintf("%d", closedIssues+repositoryData.OpenIssues),
		fmt.Sprintf("%.4f", issueByDay),
		fmt.Sprintf("%.2f", closedIssuesPercentage),
		"0",
	}
	dataChan <- ghProjectData
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

func checkAndPrintRateLimit() {
	type RateLimits struct {
		Resources struct {
			Core struct {
				Limit     int `json:"limit"`
				Remaining int `json:"remaining"`
				Reset     int `json:"reset"`
			} `json:"core"`
			Search struct {
				Limit     int `json:"limit"`
				Remaining int `json:"remaining"`
				Reset     int `json:"reset"`
			} `json:"search"`
			GraphQL struct {
				Limit     int `json:"limit"`
				Remaining int `json:"remaining"`
				Reset     int `json:"reset"`
			} `json:"graphql"`
		} `json:"resources"`
		Rate struct {
			Limit     int `json:"limit"`
			Remaining int `json:"remaining"`
			Reset     int `json:"reset"`
		} `json:"rate"`
	}
	url := "https://api.github.com/rate_limit"
	resp, statusCode, err := httpcache.MakeHTTPRequest(url)
	if err != nil {
		log.Fatalf("Error during checking rate limit : %d %v#", statusCode, err)
	}
	jsonResponse, _, _ := httpcache.ReadResp(resp)
	rateLimits := RateLimits{}
	json.Unmarshal(jsonResponse, &rateLimits)
	fmt.Printf("Core: %d/%d (reset in %d minutes)\n", rateLimits.Resources.Core.Remaining, rateLimits.Resources.Core.Limit, timing.GetRelativeTime(rateLimits.Resources.Core.Reset))
	fmt.Printf("Search: %d/%d (reset in %d minutes)\n", rateLimits.Resources.Search.Remaining, rateLimits.Resources.Search.Limit, timing.GetRelativeTime(rateLimits.Resources.Search.Reset))
	fmt.Printf("GraphQL: %d/%d (reset in %d minutes)\n", rateLimits.Resources.GraphQL.Remaining, rateLimits.Resources.GraphQL.Limit, timing.GetRelativeTime(rateLimits.Resources.GraphQL.Reset))
	fmt.Printf("Rate: %d/%d (reset in %d minutes)\n", rateLimits.Rate.Remaining, rateLimits.Rate.Limit, timing.GetRelativeTime(rateLimits.Rate.Reset))
}

func writeCsv(csvFilePath string, headers []string, ghData [][]string) {
	file, err := os.Create(csvFilePath)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	err = writer.Write(headers)
	if err != nil {
		log.Fatal("Cannot write to file", err)
	}
	for _, value := range ghData {
		err := writer.Write(value)
		if err != nil {
			log.Fatal("Cannot write to file", err)
		}
	}
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
