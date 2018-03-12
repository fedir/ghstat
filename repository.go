package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/tidwall/gjson"
)

// Repository structure with selcted data keys for JSON processing
type Repository struct {
	Name       string    `json:"name"`
	FullName   string    `json:"full_name"`
	Watchers   int       `json:"watchers"`
	Forks      int       `json:"forks"`
	OpenIssues int       `json:"open_issues"`
	CreatedAt  time.Time `json:"created_at"`
}

func getRepositoryTotalIssues(repoKey string) int64 {
	url := "https://api.github.com/search/issues?q=repo:" + repoKey + "+type:issue+state:closed"
	fullResp := MakeCachedHTTPRequest(url)
	jsonResponse, _, _ := ReadResp(fullResp)
	totalIssuesResult := gjson.Get(string(jsonResponse), "total_count")
	//fmt.Printf("%d\n", totalIssuesResult.Int())
	return totalIssuesResult.Int()
}

func getRepositoryData(repoKey string) []byte {
	url := "https://api.github.com/repos/" + repoKey
	fullResp := MakeCachedHTTPRequest(url)
	jsonResponse, _, _ := ReadResp(fullResp)
	return jsonResponse
}

func parseRepositoryData(jsonResponse []byte) *Repository {
	result := &Repository{}
	err := json.Unmarshal([]byte(jsonResponse), result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func getRepositoryStatistics(RepoKey string) *Repository {
	return parseRepositoryData(getRepositoryData(RepoKey))
}

func getActiveForkersPercentage(contributors int, forkers int) float64 {
	contributorsFloat := float64(contributors)
	forkersFloat := float64(forkers)
	activeForkersPercentage := (contributorsFloat / forkersFloat) * 100
	return activeForkersPercentage
}
