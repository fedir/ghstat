# ghstat

[![Build Status](https://travis-ci.org/fedir/ghstat.svg?branch=master)](https://travis-ci.org/fedir/ghstat)
[![Go Report Card](https://goreportcard.com/badge/github.com/fedir/ghstat)](https://goreportcard.com/report/github.com/fedir/ghstat)
[![Maintainability](https://api.codeclimate.com/v1/badges/572b4413f5c5ebf49e36/maintainability)](https://codeclimate.com/github/fedir/go-github-statistics/maintainability)
[![GoDoc](https://godoc.org/github.com/fedir/ghstat?status.svg)](https://godoc.org/github.com/fedir/ghstat)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

Statistical multi-criteria decision-making comparator for selected Github's projects.

Usage example with statistics of Go open source web frameworks, maintained on Github :

    go build
    mkdir tmp
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

* The project made by most notable author is `beego`
* The most popular project is `hugo`
* The newest project is `chi`
* The project with the most active community is `buffalo`
* The project with best errors resolving rate is `chi`
* The project with more commits is `hugo`
* The best project (taking in account placements in all competitions) is `hugo`

Details : https://github.com/fedir/ghstat/blob/master/stats/go_frameworks.csv

## JS frameworks rating

* The project made by most notable author is `angular`
* The most popular project is `react`
* The newest project is `hyperapp`
* The project with the most active community is `riot`
* The project with best errors resolving rate is `meteor`
* The project with more commits is `meteor`
* The best project (taking in account placements in all competitions) is `meteor`

Details : https://github.com/fedir/ghstat/blob/master/stats/js_frameworks.csv

## PHP frameworks rating

* The project made by most notable author is `framework`
* The most popular project is `symfony`
* The newest project is `framework`
* The project with the most active community is `framework`
* The project with best errors resolving rate is `framework`
* The project with more commits is `symfony`
* The best project (taking in account placements in all competitions) is `framework`

Details : https://github.com/fedir/ghstat/blob/master/stats/php_frameworks.csv

## Java frameworks rating

* The project made by most notable author is `playframework`
* The most popular project is `playframework`
* The newest project is `bootique`
* The project with the most active community is `ratpack`
* The project with best errors resolving rate is `jooby`
* The project with more commits is `framework`
* The best project (taking in account placements in all competitions) is `framework`

https://github.com/fedir/ghstat/blob/master/stats/java_frameworks.csv