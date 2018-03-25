// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.
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
		repositoriesKeys = strings.Split(*repositoriesKeysManual, ",")
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
		"Author's followers",
		"Top 10 contributors followers",
		"Created at",
		"Age in days",
		"Total commits",
		"Total additions",
		"Total deletions",
		"Total code changes",
		"Medium commit size",
		"Stargazers",
		"Forks",
		"Contributors",
		"Active forkers, %",
		"Open issues",
		"Total issues",
		"Closed issues, %",
		"Place",
	}
	ghDataColumnIndexes := map[string]int{
		"nameColumn":                       0,
		"authorsFollowersColumn":           3,
		"top10ContributorsFollowersColumn": 4,
		"totalAdditionsColumn":             8,
		"ageColumn":                        6,
		"totalCommitsColumn":               7,
		"totalDeletionsColumn":             9,
		"totalCodeChangesColumn":           10,
		"mediCommitSizeColumn":             11,
		"stargazersColumn":                 12,
		"activeForkersColumn":              15,
		"closedIssuesPercentageColumn":     18,
		"totalPointsColumnIndex":           19,
	}
	for _, rKey := range repositoriesKeys {
		repositoryData := getRepositoryStatistics(rKey, *tmpFolder, *debug)
		authorLogin := getRepositoryCommits(rKey, *tmpFolder, *debug)
		authorFollowers := 0
		if authorLogin != "" {
			authorFollowers = getUserFollowers(authorLogin, *tmpFolder, *debug)
		}
		closedIssues := getRepositoryClosedIssues(rKey, *tmpFolder, *debug)
		topContributorsFollowers, totalContributors := getRepositoryContributors(rKey, *tmpFolder, *debug)
		activeForkersPercentage := getActiveForkersPercentage(totalContributors, repositoryData.Forks)
		closedIssuesPercentage := getClosedIssuesPercentage(repositoryData.OpenIssues, int(closedIssues))
		contributionStatistics := getContributionStatistics(rKey, *tmpFolder, *debug)
		ghData = append(ghData, []string{
			repositoryData.Name,
			fmt.Sprintf("https://github.com/%s", repositoryData.FullName),
			fmt.Sprintf("%s", func(a string) string {
				if a == "" {
					a = "[Account removed]"
				}
				return a
			}(authorLogin)),
			fmt.Sprintf("%d", authorFollowers),
			fmt.Sprintf("%d", topContributorsFollowers),
			fmt.Sprintf("%d/%02d", repositoryData.CreatedAt.Year(), repositoryData.CreatedAt.Month()),
			fmt.Sprintf("%d", int(time.Since(repositoryData.CreatedAt).Seconds()/86400)),
			fmt.Sprintf("%d", contributionStatistics.TotalCommits),
			fmt.Sprintf("%d", contributionStatistics.TotalAdditions),
			fmt.Sprintf("%d", contributionStatistics.TotalDeletions),
			fmt.Sprintf("%d", contributionStatistics.TotalCodeChanges),
			fmt.Sprintf("%d", contributionStatistics.MediumCommitSize),
			fmt.Sprintf("%d", repositoryData.Watchers),
			fmt.Sprintf("%d", repositoryData.Forks),
			fmt.Sprintf("%d", totalContributors),
			fmt.Sprintf("%.2f", activeForkersPercentage),
			fmt.Sprintf("%d", repositoryData.OpenIssues),
			fmt.Sprintf("%d", closedIssues+repositoryData.OpenIssues),
			fmt.Sprintf("%.2f", closedIssuesPercentage),
			"0",
		})
	}
	rateGhData(ghData, ghDataColumnIndexes)
	writeCsv(csvFilePath, headers, ghData)
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
