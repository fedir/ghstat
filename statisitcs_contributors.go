// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.
package main

import (
	"encoding/json"
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
	TotalCommits     int
	TotalAdditions   int
	TotalDeletions   int
	TotalCodeChanges int
	MediumCommitSize int
}

func getContributionStatistics(repoKey string, debug bool) ContributionStatistics {
	url := "https://api.github.com/repos/" + repoKey + "/stats/contributors"
	fullResp := MakeCachedHTTPRequest(url, debug)
	jsonResponse, _, _ := ReadResp(fullResp)
	cs := extractContributionStatisticsFromJSON(jsonResponse)
	return cs
}

func extractContributionStatisticsFromJSON(jsonResponse []byte) ContributionStatistics {
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
	return cs
}

func calculateMediumCommitSize(totalCommits int, totalCodeChanges int) int {
	return int(float64(totalCodeChanges) / float64(totalCommits))
}
