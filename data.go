// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"time"

	"github.com/fedir/ghstat/github"
)

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
	r.License = "[Custom license]"
	if repositoryData.License.SPDXID != "" {
		r.License = repositoryData.License.SPDXID
	}
	r.Author = "[No GitHub account detected]"

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
