# ghstat

[![Build Status](https://travis-ci.org/fedir/ghstat.svg?branch=master)](https://travis-ci.org/fedir/ghstat)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/fedir/ghstat/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/fedir/ghstat/?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/fedir/ghstat)](https://goreportcard.com/report/github.com/fedir/ghstat)
[![Maintainability](https://api.codeclimate.com/v1/badges/572b4413f5c5ebf49e36/maintainability)](https://codeclimate.com/github/fedir/go-github-statistics/maintainability)
[![codecov](https://codecov.io/gh/fedir/ghstat/branch/master/graph/badge.svg)](https://codecov.io/gh/fedir/ghstat)
[![GoDoc](https://godoc.org/github.com/fedir/ghstat?status.svg)](https://godoc.org/github.com/fedir/ghstat)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

Statistical multi-criteria decision-making comparator for selected Github's projects.

Project's overview was given on Open Source Summit Europe 2018 "Methodology of Multi-Criteria Comparison and Typology of Open Source Projects" - https://events.linuxfoundation.org/wp-content/uploads/2017/12/Methodology-of-Multi-Criteria-Comparison-and-Typology-of-Open-Source-Project-Fedir-Rykhtik-Stratis-1.pdf

## Getting started

Installation instruction:

* Generate a token for Your GitHub account: https://github.com/settings/tokens
* Select following scope: `repo` and all it's sub-scopes 
* Build the app
* Configure the project with environment variables
* Launch

    go get -u -v github.com/fedir/ghstat
    cd [package location]
    go build
    mkdir tmp
    export GH_USR="your_gh_username" && export GH_PASS="your_gh_api_token"
    ./ghstat

The project contains already some data received from Github API for local testing and debugging, but You could update it in the following way: 

    ./ghstat --cc
    bash bin/build_all.sh

If You have timeouts, You could check the rate limit with :

    ./ghstat -l

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
