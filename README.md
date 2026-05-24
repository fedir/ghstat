# ghstat

[![CI](https://github.com/fedir/ghstat/actions/workflows/ci.yml/badge.svg)](https://github.com/fedir/ghstat/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/fedir/ghstat)](https://goreportcard.com/report/github.com/fedir/ghstat)
[![codecov](https://codecov.io/gh/fedir/ghstat/branch/master/graph/badge.svg)](https://codecov.io/gh/fedir/ghstat)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

Statistical multi-criteria decision-making comparator for GitHub projects.

Project overview was presented at Open Source Summit Europe 2018 — ["Methodology of Multi-Criteria Comparison and Typology of Open Source Projects"](https://events.linuxfoundation.org/wp-content/uploads/2017/12/Methodology-of-Multi-Criteria-Comparison-and-Typology-of-Open-Source-Project-Fedir-Rykhtik-Stratis-1.pdf).

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
make cache-clear          # wipe cached responses
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

## Comparison methodology

Each repository is scored across these criteria (more is better unless noted):

- **Stargazers** — popularity
- **Age** — newest is better
- **Total commits** — activity
- **Closed issues %** — maintenance quality
- **Commits/day** — development pace
- **Top 10 contributors followers** — community notability
- **Active forkers %** — engagement
- **Returning contributors** — project retention
- **Average contribution period** — contributor loyalty
- **Total releases** — release cadence

A final overall placement is computed by summing individual rankings.

## Ratings

[Detailed statistics with ratings](https://github.com/fedir/ghstat/blob/master/ratings.md)
