package main

import (
	"testing"
)

func makeRepo(name string, watchers, age, commits, tags, followers int, closedPct, commitsByDay, activeForkPct float64, returningContribs int) Repository {
	return Repository{
		Name:                      name,
		Watchers:                  watchers,
		Age:                       age,
		TotalCommits:              commits,
		TotalTags:                 tags,
		Top10ContributorsFollowers: followers,
		ClosedIssuesPercentage:    closedPct,
		CommitsByDay:              commitsByDay,
		ActiveForkersPercentage:   activeForkPct,
		ReturningContributors:     returningContribs,
	}
}

func TestRateGhDataWinnerIsConsistent(t *testing.T) {
	repos := []Repository{
		makeRepo("alpha", 10000, 100, 5000, 50, 20000, 80.0, 5.0, 30.0, 10),
		makeRepo("beta", 500, 500, 100, 5, 500, 20.0, 0.2, 5.0, 1),
		makeRepo("gamma", 2000, 200, 1000, 20, 5000, 50.0, 1.0, 15.0, 3),
	}

	result := rateGhData(repos)

	if result == "" {
		t.Error("rateGhData returned empty result")
	}

	// alpha dominates all criteria — must win overall
	var winner Repository
	for _, r := range repos {
		if r.PlacementOverall == 1 {
			winner = r
			break
		}
	}
	if winner.Name != "alpha" {
		t.Errorf("expected alpha to win overall, got %s", winner.Name)
	}
}

func TestRateGhDataAllPlacementsAssigned(t *testing.T) {
	repos := []Repository{
		makeRepo("a", 9000, 100, 4000, 40, 15000, 75.0, 4.0, 25.0, 8),
		makeRepo("b", 6000, 150, 3000, 30, 10000, 60.0, 3.0, 20.0, 6),
		makeRepo("c", 3000, 200, 2000, 20, 5000, 45.0, 2.0, 15.0, 4),
	}

	rateGhData(repos)

	placements := make(map[int]bool)
	for _, r := range repos {
		if r.PlacementOverall == 0 {
			t.Errorf("repo %s has PlacementOverall=0 (unassigned)", r.Name)
		}
		placements[r.PlacementOverall] = true
	}
	// All three placements 1, 2, 3 must be present
	for _, p := range []int{1, 2, 3} {
		if !placements[p] {
			t.Errorf("placement %d not assigned to any repo", p)
		}
	}
}

func TestRateGhDataSingleRepo(t *testing.T) {
	repos := []Repository{
		makeRepo("solo", 1000, 300, 500, 10, 2000, 50.0, 1.5, 10.0, 2),
	}

	rateGhData(repos)

	if repos[0].PlacementOverall != 1 {
		t.Errorf("single repo must have PlacementOverall=1, got %d", repos[0].PlacementOverall)
	}
}
