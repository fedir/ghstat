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
	"reflect"
	"strings"
	"time"

	"github.com/fedir/ghstat/github"
	"github.com/fedir/ghstat/httpcache"
	"github.com/fedir/ghstat/timing"
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

func rateAndPrintGreetings(ghData []Repository) {
	greetings := rateGhData(ghData)
	fmt.Println(greetings)
}

func writeCSVStatistics(ghData []Repository, csvFilePath string) {
	var csvData [][]string
	csvData = append(csvData, headersFromStructTags())
	for _, r := range ghData {
		csvData = append(csvData, formatRepositoryDataForCSV(r))
	}
	writeCsv(csvFilePath, csvData)
}

func formatRepositoryDataForCSV(r Repository) []string {
	ghProjectData := []string{
		r.Name,
		fmt.Sprintf("%s", r.URL),
		fmt.Sprintf("%s", r.Author),
		fmt.Sprintf("%s", r.Language),
		fmt.Sprintf("%s", r.License),
		fmt.Sprintf("%d", r.AuthorsFollowers),
		fmt.Sprintf("%d", r.Top10ContributorsFollowers),
		fmt.Sprintf("%d/%02d", r.CreatedAt.Year(), r.CreatedAt.Month()),
		fmt.Sprintf("%d", r.Age),
		fmt.Sprintf("%d", r.TotalCommits),
		fmt.Sprintf("%d", r.TotalAdditions),
		fmt.Sprintf("%d", r.TotalDeletions),
		fmt.Sprintf("%d", r.TotalCodeChanges),
		fmt.Sprintf(r.LastCommitDate.Format("2006-01-02 15:04:05")),
		fmt.Sprintf("%.4f", r.CommitsByDay),
		fmt.Sprintf("%d", r.MediCommitSize),
		fmt.Sprintf("%d", r.TotalTags),
		fmt.Sprintf("%d", r.Watchers),
		fmt.Sprintf("%d", r.Forks),
		fmt.Sprintf("%d", r.Contributors),
		fmt.Sprintf("%.2f", r.ActiveForkersPercentage),
		fmt.Sprintf("%d", r.OpenIssues),
		fmt.Sprintf("%d", r.TotalIssues),
		fmt.Sprintf("%.4f", r.IssueByDay),
		fmt.Sprintf("%.2f", r.ClosedIssuesPercentage),
		fmt.Sprintf("%.2f", r.ClosedIssuesPercentage),
		fmt.Sprintf("%d", r.PlacementPopularity),
		fmt.Sprintf("%d", r.PlacementAge),
		fmt.Sprintf("%d", r.PlacementTotalCommits),
		fmt.Sprintf("%d", r.PlacementTotalTags),
		fmt.Sprintf("%d", r.PlacementTop10ContributorsFollowers),
		fmt.Sprintf("%d", r.PlacementClosedIssuesPercentage),
		fmt.Sprintf("%d", r.PlacementCommitsByDay),
		fmt.Sprintf("%d", r.PlacementActiveForkersColumn),
		fmt.Sprintf("%d", r.PlacementOverall),
	}
	return ghProjectData
}

func headersFromStructTags() []string {
	r := new(Repository)
	return r.reflectRepositoryHeaders()
}

func (f *Repository) reflectRepositoryHeaders() []string {
	var headers []string
	val := reflect.ValueOf(f).Elem()
	for i := 0; i < val.NumField(); i++ {
		headers = append(headers, val.Type().Field(i).Tag.Get("header"))
	}
	return headers
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

func writeCsv(csvFilePath string, csvData [][]string) {
	file, err := os.Create(csvFilePath)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, value := range csvData {
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
