package localstat

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

// ContributionStatistics holds git-derived contribution data.
type ContributionStatistics struct {
	TotalCommits              int
	TotalAdditions            int
	TotalDeletions            int
	TotalCodeChanges          int
	MediumCommitSize          int
	AverageContributionPeriod int
	ReturningContributors     int
}

func repoDir(repoKey, tmpFolder string) string {
	safe := strings.ReplaceAll(repoKey, "/", "_")
	return filepath.Join(tmpFolder, "projects", safe)
}

// EnsureCloned clones the repo on first run, or fetches and resets to origin/HEAD on subsequent runs.
func EnsureCloned(repoKey, tmpFolder string) (string, error) {
	dir := repoDir(repoKey, tmpFolder)
	if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
		log.Printf("%-40s updating local clone...", repoKey)
		if out, err := gitCmd(dir, "fetch", "origin"); err != nil {
			return "", fmt.Errorf("git fetch failed for %s: %s: %w", repoKey, out, err)
		}
		if out, err := gitCmd(dir, "reset", "--hard", "origin/HEAD"); err != nil {
			return "", fmt.Errorf("git reset failed for %s: %s: %w", repoKey, out, err)
		}
		return dir, nil
	}
	log.Printf("%-40s cloning locally (full)...", repoKey)
	if err := os.MkdirAll(filepath.Dir(dir), 0755); err != nil {
		return "", err
	}
	url := "https://github.com/" + repoKey + ".git"
	if out, err := exec.Command("git", "clone", url, dir).CombinedOutput(); err != nil {
		return "", fmt.Errorf("git clone failed for %s: %s: %w", repoKey, out, err)
	}
	return dir, nil
}

// GetContributionStatistics computes contribution stats from a local git clone.
func GetContributionStatistics(repoKey, tmpFolder string) ContributionStatistics {
	dir, err := EnsureCloned(repoKey, tmpFolder)
	if err != nil {
		log.Printf("%-40s local analysis failed: %v", repoKey, err)
		return ContributionStatistics{}
	}
	var cs ContributionStatistics
	cs.TotalCommits = totalCommits(dir)
	cs.TotalAdditions, cs.TotalDeletions = additionsDeletions(dir)
	cs.TotalCodeChanges = cs.TotalAdditions + cs.TotalDeletions
	if cs.TotalCommits > 0 {
		cs.MediumCommitSize = cs.TotalCodeChanges / cs.TotalCommits
	}
	cs.AverageContributionPeriod = averageContributionPeriod(dir)
	cs.ReturningContributors = returningContributors(dir)
	return cs
}

func totalCommits(dir string) int {
	out, err := gitCmd(dir, "rev-list", "--count", "HEAD")
	if err != nil {
		return 0
	}
	n, _ := strconv.Atoi(strings.TrimSpace(out))
	return n
}

func additionsDeletions(dir string) (int, int) {
	out, err := gitCmd(dir, "log", "--numstat", "--format=")
	if err != nil {
		return 0, 0
	}
	var add, del int
	for _, line := range strings.Split(out, "\n") {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		a, _ := strconv.Atoi(fields[0])
		d, _ := strconv.Atoi(fields[1])
		add += a
		del += d
	}
	return add, del
}

func averageContributionPeriod(dir string) int {
	out, err := gitCmd(dir, "log", "--format=%aN|%aI", "--no-merges")
	if err != nil {
		return 0
	}
	// map[author] -> []time.Time
	authorDates := make(map[string][]time.Time)
	for _, line := range strings.Split(out, "\n") {
		parts := strings.SplitN(line, "|", 2)
		if len(parts) != 2 {
			continue
		}
		author := strings.TrimSpace(parts[0])
		t, err := time.Parse(time.RFC3339, strings.TrimSpace(parts[1]))
		if err != nil {
			continue
		}
		authorDates[author] = append(authorDates[author], t)
	}
	if len(authorDates) == 0 {
		return 0
	}
	totalDays := 0
	for _, dates := range authorDates {
		if len(dates) < 2 {
			continue
		}
		sort.Slice(dates, func(i, j int) bool { return dates[i].Before(dates[j]) })
		days := int(dates[len(dates)-1].Sub(dates[0]).Hours() / 24)
		totalDays += days
	}
	return totalDays / len(authorDates)
}

func returningContributors(dir string) int {
	out, err := gitCmd(dir, "log", "--format=%aN|%aI", "--no-merges")
	if err != nil {
		return 0
	}
	// map[author] -> set of ISO week numbers
	authorWeeks := make(map[string]map[int]struct{})
	for _, line := range strings.Split(out, "\n") {
		parts := strings.SplitN(line, "|", 2)
		if len(parts) != 2 {
			continue
		}
		author := strings.TrimSpace(parts[0])
		t, err := time.Parse(time.RFC3339, strings.TrimSpace(parts[1]))
		if err != nil {
			continue
		}
		if authorWeeks[author] == nil {
			authorWeeks[author] = make(map[int]struct{})
		}
		year, isoWeek := t.ISOWeek()
		authorWeeks[author][year*100+isoWeek] = struct{}{}
	}
	count := 0
	for _, weeks := range authorWeeks {
		if len(weeks) > 4 {
			count++
		}
	}
	return count
}

func gitCmd(dir string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	out, err := cmd.Output()
	return string(out), err
}
