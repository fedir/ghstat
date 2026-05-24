# ghstat — Agent & Developer Guide

## Project overview

CLI tool for multi-criteria statistical comparison of GitHub repositories. Fetches data from the GitHub REST API, scores each repo across ~10 criteria, and writes ranked results to a CSV file.

## Tech stack

- **Language:** Go 1.26.3
- **Module:** `github.com/fedir/ghstat`
- **Dependencies:** `github.com/tidwall/gjson`, `github.com/joho/godotenv`
- **CI:** GitHub Actions (`.github/workflows/ci.yml`)

## Repository layout

```
ghstat.go           # main: CLI flags, .env loading, goroutine fan-out
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
bin/                # shell scripts for per-category comparisons
stats/              # output CSV files (committed, updated by runs)
test_data/          # cached API responses used by tests (no network)
Makefile            # all developer commands
.env                # local secrets (gitignored)
.env.sample         # template for .env
```

## Authentication

Copy `.env.sample` to `.env` and set your token:

```bash
cp .env.sample .env
# GH_TOKEN=your_github_token_here
```

The token needs `repo` scope. The app loads `.env` automatically on startup. A shell environment variable takes precedence over `.env`.

## Common commands

```bash
make build                # compile binary
make test                 # run tests with race detector + coverage.txt
make vet                  # run go vet
make rate-limit           # check GitHub API quota
make cache-clear          # wipe tmp/ cache
make run-go               # Go frameworks → stats/go_frameworks.csv
make run-go-microservices # Go microservice toolkits
make run-all              # all categories via bin/build_all.sh
make clean                # remove binary and tmp/
make help                 # list all targets
```

Manual run:

```bash
./ghstat -r owner/repo1,owner/repo2 -f stats/output.csv -t tmp
```

## Output

- `-f` is required — no default output path
- `stats/` is created automatically if missing
- Results are written as CSV; headers come from `header` struct tags on `Repository`

## HTTP cache

Responses cached under `-t` folder (default `test_data/`), keyed by SHA-256 of the URL. Cache is permanent until cleared with `make cache-clear` or `-cc`. Do not cache error responses — 403/404 are handled before writing to disk. The `stats/contributors` endpoint returns 202 on first call; the client retries with 5s delay automatically.

## Conventional commits

`feat:`, `fix:`, `refactor:`, `chore:`, `ci:`, `docs:`, `test:` — single-line messages only.
