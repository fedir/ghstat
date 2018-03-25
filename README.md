# ghstat

[![Build Status](https://travis-ci.org/fedir/ghstat.svg?branch=master)](https://travis-ci.org/fedir/ghstat)
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

## Go web frameworks rating

* The most popular project is `hugo`
* The newest project is `iris`
* The project with the most active community is `buffalo`
* The project with best errors resolving rate is `iris`
* The project with more commits is `hugo`
* The project made by most notable top contributors is `beego`
* The best project (taking in account placements in all competitions) is `hugo`

[Detailed Go frameworks statistics](https://github.com/fedir/ghstat/blob/master/stats/go_frameworks.csv)

## Haskell frameworks rating

* The most popular project is `yesod`
* The newest project is `magicbane`
* The project with the most active community is `fn`
* The project with best errors resolving rate is `snap-core`
* The project with more commits is `yesod`
* The project made by most notable top contributors is `yesod`
* The best project (taking in account placements in all competitions) is `yesod`

[Detailed Haskell frameworks statistics](https://github.com/fedir/ghstat/blob/master/stats/haskell_frameworks.csv)

## Java frameworks rating

* The most popular project is `playframework`
* The newest project is `bootique`
* The project with the most active community is `ratpack`
* The project with best errors resolving rate is `jooby`
* The project with more commits is `framework`
* The project made by most notable top contributors is `blade`
* The best project (taking in account placements in all competitions) is `ratpack`

[Detailed Java frameworks statistics](https://github.com/fedir/ghstat/blob/master/stats/java_frameworks.csv)

## JS frameworks rating

* The most popular project is `react`
* The newest project is `hyperapp`
* The project with the most active community is `riot`
* The project with best errors resolving rate is `meteor`
* The project with more commits is `meteor`
* The project made by most notable top contributors is `react`
* The best project (taking in account placements in all competitions) is `meteor`

[Detailed JS frameworks statistics](https://github.com/fedir/ghstat/blob/master/stats/js_frameworks.csv)

## PHP frameworks rating

* The most popular project is `symfony`
* The newest project is `framework`
* The project with the most active community is `framework`
* The project with best errors resolving rate is `framework`
* The project with more commits is `symfony`
* The project made by most notable top contributors is `framework`
* The best project (taking in account placements in all competitions) is `framework`

[Detailed PHP frameworks statistics](https://github.com/fedir/ghstat/blob/master/stats/php_frameworks.csv)

## Cross-language frameworks rating

* The most popular project is `react`
* The newest project is `magicbane`
* The project with the most active community is `axiom`
* The project with best errors resolving rate is `iris`
* The project with more commits is `meteor`
* The project made by most notable top contributors is `react`
* The best project (taking in account placements in all competitions) is `meteor`

[Detailed PHP frameworks statistics](https://github.com/fedir/ghstat/blob/master/stats/all_frameworks.csv)
