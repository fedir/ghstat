// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.

package main

import "time"

// Repository structure with selcted data keys for JSON processing
type Repository struct {
	Name                                string    `header:"Name"`
	URL                                 string    `header:"URL"`
	Author                              string    `header:"Author"`
	AuthorLocation                      string    `header:"Author's location"`
	MainLanguage                        string    `header:"Main language"`
	AllLanguages                        string    `header:"All used languages"`
	TotalCodeSize                       int       `header:"Total code size"`
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
	ActiveForkersPercentage             float64   `header:"Active forkers(%)"`
	OpenIssues                          int       `header:"Open issues"`
	ClosedIssues                        int       `header:"Closed issues"`
	TotalIssues                         int       `header:"Total issues"`
	IssueByDay                          float64   `header:"Issue/day"`
	ClosedIssuesPercentage              float64   `header:"Closed issues(%)"`
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
