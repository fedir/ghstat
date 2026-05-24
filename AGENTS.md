# ghstat — Agent & Developer Guide

## Project overview

CLI tool for multi-criteria statistical comparison of GitHub repositories. Fetches data from the GitHub REST API, scores each repo across ~10 criteria, and writes ranked results to `result.csv`.

## Tech stack

- **Language:** Go 1.26.3
- **Module:** `github.com/fedir/ghstat`
- **Dependencies:** `github.com/tidwall/gjson` (JSON path queries)
- **CI:** GitHub Actions (`.github/workflows/ci.yml`)

## Repository layout

```
ghstat.go           # main: CLI flags, goroutine fan-out
struct.go           # Repository struct with CSV header tags
data.go             # per-repo data fetching (called per goroutine)
competition.go      # scoring and ranking logic
files.go            # CSV output, HTTP cache clearing
github/             # GitHub API client
  repository.go
  contributors.go
  author.go
  limits.go
  statisitcs_contributors.go
  repository_language_sorting.go
httpcache/          # file-based HTTP response cache (SHA-256 keyed)
  httpcache.go
timing/             # Unix timestamp → relative minutes helper
bin/                # shell scripts for common repo group comparisons
test_data/          # cached API responses used by tests
```

## Authentication

Set a single env var before running:

```bash
export GH_TOKEN="your_github_token"
```

The token needs `repo` scope. Without it requests are rate-limited to 60/hour.

## Build & run

```bash
go build
./ghstat                                      # compare default Go frameworks
./ghstat -r angular/angular,vuejs/vue         # custom repo list
./ghstat -f output.csv                        # custom output path
./ghstat -t /tmp/cache                        # custom cache folder
./ghstat -l                                   # check rate limit
./ghstat -cc                                  # clear HTTP cache
./ghstat -ccdr                                # dry-run cache clear
```

## Tests

```bash
go test -race ./...
```

Tests use pre-recorded API responses in `test_data/` — no network required.

## Pre-built comparison scripts

```bash
bash bin/js_frameworks.sh
bash bin/go_frameworks.sh
bash bin/php_frameworks.sh
bash bin/build_all.sh        # runs all categories
```

## HTTP cache

API responses are cached as files under the temp folder (`test_data/` by default), keyed by SHA-256 of the URL. Cache is permanent until manually cleared with `-cc`. This avoids hitting rate limits on repeated runs.

The GitHub `stats/contributors` endpoint returns HTTP 202 on first call (GitHub computes it async). The client retries automatically with a 5s delay.

## Conventional commits

This project uses conventional commits: `feat:`, `fix:`, `refactor:`, `chore:`, `ci:`, `docs:`, `test:`. Keep commit messages as single lines.
