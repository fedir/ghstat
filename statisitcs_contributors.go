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
	//Weeks        []StatsContributorWeek `json:"weeks"`
}

// StatsContributorWeek substructure
type StatsContributorWeek struct {
	Week      string `json:"w"`
	Additions string `json:"a"`
	Deletions string `json:"d"`
	Commits   string `json:"c"`
}

// ContributionStatistics contains multiple statistics about contribution into the repository
type ContributionStatistics struct {
	TotalCommits int
}

func getContributionStatistics(repoKey string, debug bool) ContributionStatistics {
	var cs ContributionStatistics
	cs.TotalCommits = 0
	url := "https://api.github.com/repos/" + repoKey + "/stats/contributors"
	fullResp := MakeCachedHTTPRequest(url, debug)
	jsonResponse, _, _ := ReadResp(fullResp)
	contributionStatistics := make([]StatsContributor, 0)
	json.Unmarshal(jsonResponse, &contributionStatistics)
	//fmt.Printf("%v#", contributionStatistics)
	for _, c := range contributionStatistics {
		//fmt.Printf("%d %s %d\n", i, c.Author.Login, c.TotalCommits)
		cs.TotalCommits += c.TotalCommits
	}
	//fmt.Printf("%d\n", cs.TotalCommits)
	return cs
}
