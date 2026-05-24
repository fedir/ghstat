# ghstat — Agent & Developer Guide

## Project overview

CLI tool for multi-criteria statistical comparison of GitHub repositories. Combines GitHub REST API data with local git clone analysis, scores each repo across ~10 criteria, and writes ranked results to a CSV file.

## Tech stack

- **Language:** Go 1.26.3
- **Module:** `github.com/fedir/ghstat`
- **Dependencies:** `github.com/tidwall/gjson`, `github.com/joho/godotenv`
- **CI:** GitHub Actions (`.github/workflows/ci.yml`)

## Repository layout

```
ghstat.go           # main: CLI flags, .env loading, goroutine fan-out
struct.go           # Repository struct with CSV header tags
data.go             # per-repo data fetching + hybrid API/local merge
competition.go      # scoring and ranking logic
files.go            # CSV output, HTTP cache clearing
github/             # GitHub API client
  repository.go
  contributors.go
  author.go
  limits.go
  statisitcs_contributors.go
  repository_language_sorting.go
localstat/          # local git clone analysis (authoritative commit history)
  localstat.go
httpcache/          # file-based HTTP response cache (SHA-256 keyed)
  httpcache.go
timing/             # Unix timestamp → relative minutes helper
bin/                # shell scripts for per-category comparisons
stats/              # output CSV files (committed, updated by runs)
test_data/          # cached API responses used by tests (no network)
tmp/projects/       # local git clones (gitignored, can be several GB)
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

## Environment variables

| Variable | Default | Description |
|----------|---------|-------------|
| `GH_TOKEN` | — | GitHub personal access token (required) |
| `GH_HTTP_TIMEOUT` | `30` | HTTP request timeout in seconds |
| `GH_STATS_MAX_RETRIES` | `5` | Max retries when GitHub returns 202 (computing stats) |
| `GH_STATS_RETRY_INTERVAL` | `10` | Seconds between retries for 202 responses |

## Hybrid analysis

Each repository is analysed from two sources, merged into one record:

| Data | Source | Why |
|------|--------|-----|
| Stars, forks, open issues, license | GitHub API | Real-time metadata |
| Author, location, followers | GitHub API | User profile data |
| Closed issues, tags, contributors | GitHub API | GitHub-specific concepts |
| **TotalCommits, Additions, Deletions** | **Local git** | `git rev-list` / `git log --numstat` — authoritative |
| **MediumCommitSize** | **Local git** | Derived from above |
| **AverageContributionPeriod** | **Local git** | Per-author first/last commit date span |
| **ReturningContributors** | **Local git** | Authors active in >4 distinct ISO weeks |
| **CommitsByDay** | **Local git** | More accurate commit count / repo age |

On first run, each repo is fully cloned to `tmp/projects/<owner>_<repo>/`. On subsequent runs the clone is updated (`git fetch origin` + `git reset --hard origin/HEAD`). If the clone fails, API data is kept as-is.

## Common commands

```bash
make build                # compile binary
make test                 # run tests with race detector + coverage.txt
make vet                  # run go vet
make rate-limit           # check GitHub API quota
make cache-clear          # wipe HTTP response cache (preserves clones)
make clone-clear          # remove local git clones in tmp/projects/
make run-go               # Go frameworks → stats/go_frameworks.csv
make run-go-microservices # Go microservice toolkits
make run-rust-crates      # top 25 Rust crates → stats/rust_crates.csv
make run-cncf             # 52 CNCF cloud native projects → stats/cncf_projects.csv
make run-devops           # 40 DevOps tools → stats/devops_tools.csv
make run-all              # all categories via bin/build_all.sh → ratings.md
make clean-data-cms       # remove CMS clones from test_data/projects/
make clean-data-databases # remove database clones
make clean-data-langs     # remove language clones
make clean-data-go        # remove Go framework clones
make clean-data-rust      # remove Rust clones
make clean-data-js        # remove JS framework clones
make clean-data-python    # remove Python framework clones
make clean-data-ruby      # remove Ruby framework clones
make clean-data-java      # remove Java/JVM clones
make clean-data-cncf      # remove CNCF project clones
make clean-data-all       # remove all clones (test_data/ + tmp/projects/)
make clean                # remove binary and tmp/ (preserves clones)
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

Responses cached under `-t` folder (default `test_data/`), keyed by SHA-256 of the URL. Cache is permanent until cleared with `make cache-clear` or `-cc`. Error responses (403/404) are not cached. The `stats/contributors` endpoint returns 202 when GitHub is computing stats; the client retries up to `GH_STATS_MAX_RETRIES` times with `GH_STATS_RETRY_INTERVAL` second delays. Local git stats always override the API result when available, so 202-forever repos are handled correctly.

## Conventional commits

`feat:`, `fix:`, `refactor:`, `chore:`, `ci:`, `docs:`, `test:` — **single-line messages only**. No bullet lists, no multi-line body, no `Co-Authored-By` trailer. One line, under 72 chars, imperative mood. Example: `feat: add CNCF rating script with 404 robustness fixes`.

## Lessons learned

### Clone cache and `-t` flag for new scripts
New `bin/*.sh` scripts must pass `-t tmp` explicitly. Without it the tool defaults to `test_data/` which has no clones for new repo sets. The tool exits after writing only the repos whose clones finished before the others — producing a truncated CSV with no error message. Always verify row count after a run: `wc -l stats/*.csv`.

### test_data/projects disk usage
`test_data/projects/` accumulates full git clones and can grow to 60GB+. Use `make clean-data-<category>` to remove specific sets, or `make clean-data-all` to wipe everything. CNCF and Rust crates use `tmp/projects/` (via `-t tmp`); all other scripts use `test_data/projects/` (default).

### 404 and 403 responses must not crash the tool
`httpcache` returns `[]byte{}` on 404 and 403. All JSON unmarshal callers must guard against empty input — use `if len(jsonResponse) == 0 { return }` before unmarshaling. Slice indexing on unmarshaled results (`commits[0]`, `commits[len-1]`) must also guard against empty slices. `log.Fatal`/`log.Fatalf` in fetch paths kills the entire batch — use `log.Printf` and return zero values instead.

### GitHub search API rate limit (403) with large batches
The `/search/issues` endpoint has a stricter rate limit than the main API (~30 req/min). Running 50+ repos in parallel exhausts it immediately. The tool now logs and skips gracefully instead of fatally crashing. Closed issue counts will be 0 for repos that hit this limit — re-run later with a smaller batch or add a sleep between search calls.
