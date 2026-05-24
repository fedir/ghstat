package main

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestFormatRepositoryDataForCSVLength(t *testing.T) {
	r := Repository{
		Name:          "owner/repo",
		URL:           "https://github.com/owner/repo",
		CreatedAt:     time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		LastCommitDate: time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC),
	}
	row := formatRepositoryDataForCSV(r)
	headers := headersFromStructTags()

	if len(row) != len(headers) {
		t.Errorf("CSV row length %d != headers length %d", len(row), len(headers))
	}
}

func TestHeadersFromStructTagsNotEmpty(t *testing.T) {
	headers := headersFromStructTags()
	if len(headers) == 0 {
		t.Error("expected non-empty headers")
	}
	for i, h := range headers {
		if h == "" {
			t.Errorf("header at index %d is empty", i)
		}
	}
}

func TestClearHTTPCacheFolderPreservesProjects(t *testing.T) {
	dir := t.TempDir()

	// Create files and a "projects" subdirectory
	if err := os.WriteFile(filepath.Join(dir, "cache1.json"), []byte("{}"), 0644); err != nil {
		t.Fatal(err)
	}
	projectsDir := filepath.Join(dir, "projects")
	if err := os.MkdirAll(projectsDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(projectsDir, "clone.txt"), []byte("keep"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := clearHTTPCacheFolder(dir, false); err != nil {
		t.Fatalf("clearHTTPCacheFolder failed: %v", err)
	}

	// cache1.json should be gone
	if _, err := os.Stat(filepath.Join(dir, "cache1.json")); !os.IsNotExist(err) {
		t.Error("expected cache1.json to be deleted")
	}

	// projects/ and its contents must survive
	if _, err := os.Stat(filepath.Join(projectsDir, "clone.txt")); err != nil {
		t.Errorf("projects/ directory should be preserved: %v", err)
	}
}

func TestClearHTTPCacheFolderDryRun(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "cache2.json"), []byte("{}"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := clearHTTPCacheFolder(dir, true); err != nil {
		t.Fatalf("clearHTTPCacheFolder dry-run failed: %v", err)
	}

	// file must still exist after dry run
	if _, err := os.Stat(filepath.Join(dir, "cache2.json")); err != nil {
		t.Errorf("dry run should not delete files: %v", err)
	}
}

func TestWriteAndReadCSV(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "out.csv")

	data := [][]string{
		{"name", "url", "stars"},
		{"myrepo", "https://github.com/x/y", "999"},
	}
	writeCsv(path, data)

	if _, err := os.Stat(path); err != nil {
		t.Errorf("CSV file not created: %v", err)
	}
}
