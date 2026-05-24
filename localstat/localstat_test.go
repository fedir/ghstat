package localstat

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// initRepo creates a temp git repo, configures identity, and returns its path.
func initRepo(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	run := func(args ...string) {
		t.Helper()
		cmd := exec.Command("git", args...)
		cmd.Dir = dir
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v failed: %s", args, out)
		}
	}
	run("init")
	run("config", "user.email", "test@test.com")
	run("config", "user.name", "Test User")
	return dir
}

// addCommit writes a file with the given content and commits it.
func addCommit(t *testing.T, dir, filename, content, message string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, filename), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	run := func(args ...string) {
		t.Helper()
		cmd := exec.Command("git", args...)
		cmd.Dir = dir
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("git %v failed: %s", args, out)
		}
	}
	run("add", filename)
	run("commit", "-m", message)
}

func TestTotalCommits(t *testing.T) {
	dir := initRepo(t)
	addCommit(t, dir, "a.txt", "hello\n", "first")
	addCommit(t, dir, "b.txt", "world\n", "second")

	n := totalCommits(dir)
	if n != 2 {
		t.Errorf("expected 2 commits, got %d", n)
	}
}

func TestAdditionsDeletions(t *testing.T) {
	dir := initRepo(t)
	// 3 lines added
	addCommit(t, dir, "a.txt", "line1\nline2\nline3\n", "add three lines")
	// overwrite with 1 line: 3 deletions, 1 addition
	addCommit(t, dir, "a.txt", "only\n", "shrink")

	add, del := additionsDeletions(dir)
	// first commit: +3, second: +1 -3
	if add != 4 {
		t.Errorf("expected 4 additions, got %d", add)
	}
	if del != 3 {
		t.Errorf("expected 3 deletions, got %d", del)
	}
}

func TestGetContributionStatisticsEmptyOnMissingDir(t *testing.T) {
	cs := GetContributionStatistics("nonexistent/repo", t.TempDir())
	if cs.TotalCommits != 0 {
		t.Errorf("expected 0 commits for missing repo, got %d", cs.TotalCommits)
	}
}

func TestGetContributionStatisticsBasic(t *testing.T) {
	dir := initRepo(t)
	addCommit(t, dir, "a.txt", "line1\nline2\n", "init")
	addCommit(t, dir, "a.txt", "line1\nline2\nline3\n", "add one line")

	cs := ContributionStatistics{
		TotalCommits: totalCommits(dir),
	}
	cs.TotalAdditions, cs.TotalDeletions = additionsDeletions(dir)
	cs.TotalCodeChanges = cs.TotalAdditions + cs.TotalDeletions

	if cs.TotalCommits != 2 {
		t.Errorf("TotalCommits: expected 2, got %d", cs.TotalCommits)
	}
	if cs.TotalAdditions < 2 {
		t.Errorf("TotalAdditions: expected >= 2, got %d", cs.TotalAdditions)
	}
	if cs.TotalCodeChanges != cs.TotalAdditions+cs.TotalDeletions {
		t.Errorf("TotalCodeChanges mismatch")
	}
}

func TestReturningContributorsBelowThreshold(t *testing.T) {
	dir := initRepo(t)
	// A single contributor with 3 commits — below the >4 weeks threshold
	addCommit(t, dir, "a.txt", "v1\n", "c1")
	addCommit(t, dir, "b.txt", "v2\n", "c2")
	addCommit(t, dir, "c.txt", "v3\n", "c3")

	rc := returningContributors(dir)
	if rc != 0 {
		t.Errorf("expected 0 returning contributors, got %d", rc)
	}
}

func TestRepoDir(t *testing.T) {
	got := repoDir("owner/repo", "/tmp")
	expected := filepath.Join("/tmp", "projects", "owner_repo")
	if got != expected {
		t.Errorf("repoDir: expected %s, got %s", expected, got)
	}
}
