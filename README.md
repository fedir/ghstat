# ghstat

[![Build Status](https://travis-ci.org/fedir/ghstat.svg?branch=master)](https://travis-ci.org/fedir/ghstat)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/fedir/ghstat/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/fedir/ghstat/?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/fedir/ghstat)](https://goreportcard.com/report/github.com/fedir/ghstat)
[![Maintainability](https://api.codeclimate.com/v1/badges/572b4413f5c5ebf49e36/maintainability)](https://codeclimate.com/github/fedir/go-github-statistics/maintainability)
[![codecov](https://codecov.io/gh/fedir/ghstat/branch/master/graph/badge.svg)](https://codecov.io/gh/fedir/ghstat)
[![GoDoc](https://godoc.org/github.com/fedir/ghstat?status.svg)](https://godoc.org/github.com/fedir/ghstat)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

Statistical multi-criteria decision-making comparator for selected Github's projects.

Usage example with statistics of Go open source web frameworks, maintained on Github :

    go build
    mkdir tmp
    export GH_USR="your_gh_username" && export GH_PASS="your_gh_api_token"
    ./ghstat

Usage example to compare most famous JS frameworks

    ./ghstat -r angular/angular,facebook/react,vuejs/vue

Usage example to compare most famous PHP frameworks

    ./ghstat -r laravel/framework,symfony/symfony,yiisoft/yii2,bcit-ci/CodeIgniter

After that, `result.csv` file will be created (or updated, if it's already exists) with the statistics of selected repositories.

## Comparaison methodology

At the moment We choosed following metrics, here they are, in alphabetical order :

* Active forkers percentage - more is better
* Age in days - newest is better :)
* Closed issues, % - more is better
* Watchers - more is better
* Total commits - more is better
  * More precisely, it's total commits by existing contributors, commits of deleted accounts, will not be taken in account

## Ratings

[Detailed statistics with ratings made with ghstat](https://github.com/fedir/ghstat/blob/master/ratings.md)