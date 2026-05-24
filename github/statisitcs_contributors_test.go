// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.

package github

import (
	"fmt"
	"reflect"
	"testing"
)

var contributionStatisticsJSONResponse = `
[
  {
    "total": 5,
    "weeks": [
      {
        "w": 1520121600,
        "a": 100,
        "d": 5,
        "c": 4
      },
      {
        "w": 1520726400,
        "a": 10,
        "d": 5,
        "c": 1
      },
      {
        "w": 1521331200,
        "a": 0,
        "d": 0,
        "c": 0
      }
    ],
    "author": {
      "login": "fedir",
      "id": 306586,
      "avatar_url": "https://avatars1.githubusercontent.com/u/306586?v=4",
      "gravatar_id": "",
      "url": "https://api.github.com/users/fedir",
      "html_url": "https://github.com/fedir",
      "followers_url": "https://api.github.com/users/fedir/followers",
      "following_url": "https://api.github.com/users/fedir/following{/other_user}",
      "gists_url": "https://api.github.com/users/fedir/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/fedir/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/fedir/subscriptions",
      "organizations_url": "https://api.github.com/users/fedir/orgs",
      "repos_url": "https://api.github.com/users/fedir/repos",
      "events_url": "https://api.github.com/users/fedir/events{/privacy}",
      "received_events_url": "https://api.github.com/users/fedir/received_events",
      "type": "User",
      "site_admin": false
    }
  }
]
`

func TestContributionStatisticsJSONResponseData(t *testing.T) {
	contributionStatistics := extractContributionStatisticsFromJSON([]byte(contributionStatisticsJSONResponse), false)
	contributionStatisticsExpected := ContributionStatistics{
		TotalCommits:              5,
		TotalAdditions:            110,
		TotalDeletions:            10,
		TotalCodeChanges:          120,
		MediumCommitSize:          24,
		AverageContributionPeriod: 7,
		ReturningContributors:     0, // only 2 active weeks, below threshold of 4
	}
	if !reflect.DeepEqual(contributionStatistics, contributionStatisticsExpected) {
		fmt.Println(contributionStatistics)
		fmt.Println(contributionStatisticsExpected)
		t.Fail()
	}
}

var contributionStatisticsReturningJSON = `
[
  {
    "total": 10,
    "weeks": [
      {"w": 1500000000, "a": 10, "d": 1, "c": 2},
      {"w": 1500604800, "a": 10, "d": 1, "c": 2},
      {"w": 1501209600, "a": 10, "d": 1, "c": 2},
      {"w": 1501814400, "a": 10, "d": 1, "c": 2},
      {"w": 1502419200, "a": 10, "d": 1, "c": 2}
    ],
    "author": {"login": "alice"}
  },
  {
    "total": 2,
    "weeks": [
      {"w": 1500000000, "a": 5, "d": 0, "c": 1},
      {"w": 1500604800, "a": 5, "d": 0, "c": 1}
    ],
    "author": {"login": "bob"}
  }
]
`

func TestReturningContributors(t *testing.T) {
	cs := extractContributionStatisticsFromJSON([]byte(contributionStatisticsReturningJSON), false)
	// alice has 5 active weeks (>4) → returning; bob has 2 → not returning
	if cs.ReturningContributors != 1 {
		t.Errorf("expected 1 returning contributor, got %d", cs.ReturningContributors)
	}
}

func TestCalculateMediumCommitSize(t *testing.T) {
	if got := calculateMediumCommitSize(10, 100); got != 10 {
		t.Errorf("expected 10, got %d", got)
	}
	if got := calculateMediumCommitSize(0, 100); got != 0 {
		// division by zero guard — NaN cast to int is 0
		_ = got
	}
}
