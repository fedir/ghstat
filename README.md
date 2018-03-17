# ghstat

[![Build Status](https://travis-ci.org/fedir/ghstat.svg?branch=master)](https://travis-ci.org/fedir/ghstat)
[![Go Report Card](https://goreportcard.com/badge/github.com/fedir/ghstat)](https://goreportcard.com/report/github.com/fedir/ghstat)
[![Maintainability](https://api.codeclimate.com/v1/badges/572b4413f5c5ebf49e36/maintainability)](https://codeclimate.com/github/fedir/go-github-statistics/maintainability)
[![GoDoc](https://godoc.org/github.com/fedir/ghstat?status.svg)](https://godoc.org/github.com/fedir/ghstat)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

Statistical multi-criteria decision-making comparator for selected Github's projects.

Usage example with statistics of Go open source web frameworks, maintained on Github :

    go build
    ./ghstat

Usage example to compare most famous JS frameworks

    ./ghstat -r angular/angular,facebook/react,vuejs/vue

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
* The newest project is `chi`
* The project with the most active community is `buffalo`
* The project with best errors resolving rate is `revel`
* The best project taking in account the placement in multiple section is `hugo`

## JS frameworks rating

* The most popular project is `react`
* The newest project is `angular`
* The project with the most active community is `angular`
* The project with best errors resolving rate is `vue`
* The project with more commits is `angular`
* The best project (taking in account placements in all competitions) is `angular`

## PHP frameworks rating

* The most popular project is `laravel`
* The newest project is `yii2`
* The project with the most active community is `symfony`
* The project with best errors resolving rate is `CodeIgniter`
* The project with more commits is `symfony`
* The best project (taking in account placements in all competitions) is `symfony`