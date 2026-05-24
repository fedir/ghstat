# ghstat

[![CI](https://github.com/fedir/ghstat/actions/workflows/ci.yml/badge.svg)](https://github.com/fedir/ghstat/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/fedir/ghstat)](https://goreportcard.com/report/github.com/fedir/ghstat)
[![codecov](https://codecov.io/gh/fedir/ghstat/branch/master/graph/badge.svg)](https://codecov.io/gh/fedir/ghstat)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

Statistical multi-criteria decision-making comparator for GitHub projects. Combines GitHub REST API data with local git clone analysis for accurate historical commit statistics.

## Background

The tool was built out of a frustration familiar to most engineers: choosing between open source dependencies by star count alone. Stars measure marketing, not health.

The methodology behind ghstat was presented at **Linux Foundation Open Source Summit Europe 2018 in Edinburgh, UK** — ["Methodology of Multi-Criteria Comparison and Typology of Open Source Projects"](https://events19.linuxfoundation.org/wp-content/uploads/2017/12/Methodology-of-Multi-Criteria-Comparison-and-Typology-of-Open-Source-Project-Fedir-Rykhtik-Stratis-1.pdf). The core argument: open source project selection should be treated like engineering — systematic, multi-dimensional, reproducible. Commit velocity, contributor retention, issue resolution, and code churn together tell a story that no single metric can.

A companion blog post from 2018 walking through the analysis on Python frameworks: [Using ghstat for open source project statistics and ratings](https://fedir.github.io/development/2018/04/02/using-ghstat-open-source-projects-statistics-and-ratings).

Seven years later the tool has been rebuilt with hybrid analysis (GitHub API + local git clone) and run across ~200 repositories in 12 categories. The methodology holds.

## Getting started

**1. Generate a GitHub token**

Go to https://github.com/settings/tokens and create a token with `repo` scope.

**2. Clone and configure**

```bash
git clone https://github.com/fedir/ghstat
cd ghstat
cp .env.sample .env
# edit .env and set your token
```

**3. Build and run**

```bash
make build
make run-go
```

Output is written to `stats/go_frameworks.csv`.

## Usage

```bash
make help                 # list all available commands
make rate-limit           # check GitHub API quota
make cache-clear          # wipe HTTP response cache (preserves local clones)
make clone-clear          # remove local git clones in tmp/projects/
make run-go               # compare Go frameworks
make run-go-microservices # compare Go microservice toolkits
make run-all              # run all comparisons
make test                 # run tests with coverage
```

Custom comparison:

```bash
./ghstat -r angular/angular,facebook/react,vuejs/vue -f stats/js.csv -t tmp
```

## Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-r` | Go frameworks | Comma-separated list of `owner/repo` |
| `-f` | *(required)* | Output CSV file path |
| `-t` | `test_data` | Cache folder |
| `-l` | | Check GitHub rate limit |
| `-cc` | | Clear HTTP cache |
| `-ccdr` | | Dry-run cache clear |
| `-d` | | Debug mode |

## How it works

Each repository is analysed from two sources:

- **GitHub API** — real-time data: stars, forks, issues, license, author profile, closed issues, tags, contributors
- **Local git clone** — authoritative history: commit count, additions/deletions, commit size, contribution period, returning contributors

On first run repos are cloned to `tmp/projects/`. On subsequent runs the clones are updated. Local stats override API stats when available, so repositories where GitHub's stats API returns 202 (inactive repos) still get accurate data.

## Comparison methodology

Each repository is scored across these criteria (more is better unless noted):

- **Stargazers** — popularity
- **Age** — newest is better
- **Total commits** — activity (from local git)
- **Closed issues %** — maintenance quality
- **Commits/day** — development pace (from local git)
- **Top 10 contributors followers** — community notability
- **Active forkers %** — engagement
- **Returning contributors** — project retention (from local git)
- **Average contribution period** — contributor loyalty (from local git)
- **Total releases** — release cadence

A final overall placement is computed by summing individual rankings.

## Ratings

[Detailed statistics with ratings](https://github.com/fedir/ghstat/blob/master/ratings.md)
