// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.

package github

import (
	"encoding/json"
	"fmt"

	"github.com/fedir/ghstat/httpcache"
)

// StatsContributor contains statistical data for contribution
type StatsContributor struct {
	Author struct {
		Login string `json:"login"`
	} `json:"author"`
	TotalCommits int `json:"total"`
	Weeks        []struct {
		Week      int `json:"w"`
		Additions int `json:"a"`
		Deletions int `json:"d"`
		Commits   int `json:"c"`
	} `json:"weeks"`
}

// ContributionStatistics contains multiple statistics about contribution into the repository
type ContributionStatistics struct {
	TotalCommits              int
	TotalAdditions            int
	TotalDeletions            int
	TotalCodeChanges          int
	MediumCommitSize          int
	AverageContributionPeriod int
}

// GetContributionStatistics gets detailsed statistics about contributors of the repository
func GetContributionStatistics(repoKey string, tmpFolder string, debug bool) ContributionStatistics {
	url := "https://api.github.com/repos/" + repoKey + "/stats/contributors"
	fullResp := httpcache.MakeCachedHTTPRequest(url, tmpFolder, debug)
	jsonResponse, _, _ := httpcache.ReadResp(fullResp)
	cs := extractContributionStatisticsFromJSON(jsonResponse, debug)
	if debug {
		fmt.Printf("ACP for %s: %d days\n", repoKey, cs.AverageContributionPeriod)
	}
	return cs
}

// GetCommitsByDay calculates the rate of commits by day of the repository
func GetCommitsByDay(totalCommits int, repositoryAge int) float64 {
	var commitsByDay float64
	totalCommitsFloat := float64(totalCommits)
	repositoryAgeFloat := float64(repositoryAge)
	if totalCommitsFloat != 0 && repositoryAgeFloat != 0 {
		commitsByDay = totalCommitsFloat / repositoryAgeFloat
	} else {
		commitsByDay = 0
	}
	return commitsByDay
}

func extractContributionStatisticsFromJSON(jsonResponse []byte, debug bool) ContributionStatistics {
	var cs ContributionStatistics
	cs.TotalCommits = 0
	cs.TotalAdditions = 0
	cs.TotalDeletions = 0
	cs.TotalCodeChanges = 0
	contributionStatistics := make([]StatsContributor, 0)
	json.Unmarshal(jsonResponse, &contributionStatistics)
	for _, c := range contributionStatistics {
		cs.TotalCommits += c.TotalCommits
		for _, cw := range c.Weeks {
			cs.TotalAdditions += cw.Additions
			cs.TotalDeletions += cw.Deletions
			cs.TotalCodeChanges += cw.Additions
			cs.TotalCodeChanges += cw.Deletions
		}
	}
	cs.MediumCommitSize = calculateMediumCommitSize(cs.TotalCommits, cs.TotalCodeChanges)
	cs.AverageContributionPeriod = calculateAverageContributionPeriod(contributionStatistics, debug)
	return cs
}

func calculateMediumCommitSize(totalCommits int, totalCodeChanges int) int {
	return int(float64(totalCodeChanges) / float64(totalCommits))
}

func calculateAverageContributionPeriod(cs []StatsContributor, debug bool) int {
	type contributorContibutionPeriod struct {
		first  int
		last   int
		period int
	}
	acp := 0
	totalAcp := 0
	nContirbutors := len(cs)
	for _, c := range cs {
		ccp := contributorContibutionPeriod{
			first:  0,
			last:   0,
			period: 0,
		}
		for _, cw := range c.Weeks {
			if cw.Additions > 0 || cw.Commits > 0 || cw.Deletions > 0 {
				if ccp.first == 0 {
					ccp.first = cw.Week
				} else if cw.Week < ccp.first {
					ccp.first = cw.Week
				}
				if cw.Week > ccp.last {
					ccp.last = cw.Week
				}
			}
		}
		ccp.period = (ccp.last - ccp.first) / 86400
		totalAcp += ccp.period
		if debug {
			fmt.Printf("%s, %d, %d, %d\n", c.Author, ccp.first, ccp.last, ccp.period)
		}
	}
	acp = int(float64(totalAcp) / float64(nContirbutors))
	return acp
}
