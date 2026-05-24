// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.

// ghstat - statistical multi-criteria decision-making comparator for Github's projects.
package main

import (
	"flag"
	"os"
	"strings"
	"sync"

	"github.com/fedir/ghstat/github"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	var (
		clearHTTPCache         = flag.Bool("cc", false, "Clear HTTP cache")
		clearHTTPCacheDryRun   = flag.Bool("ccdr", false, "Clear HTTP cache (dry run)")
		debug                  = flag.Bool("d", false, "Debug mode")
		resultFileSavePath     = flag.String("f", "", "File path where result CSV file will be saved (required)")
		rateLimitCheck         = flag.Bool("l", false, "Rate limit check")
		repositoriesKeysManual = flag.String("r", "", "Repositories keys")
		tmpFolder              = flag.String("t", "test_data", "Temporary folder path")
		repositoriesKeys       = []string{}
	)
	flag.Parse()
	if *clearHTTPCache || *clearHTTPCacheDryRun {
		clearHTTPCacheFolder(*tmpFolder, *clearHTTPCacheDryRun)
		os.Exit(0)
	}
	if *rateLimitCheck {
		github.CheckAndPrintRateLimit()
		os.Exit(0)
	}
	if *repositoriesKeysManual != "" {
		repositoriesKeys = uniqSlice(strings.Split(*repositoriesKeysManual, ","))
	} else {
		repositoriesKeys = uniqSlice([]string{
			"astaxie/beego",
			"gohugoio/hugo",
			"gin-gonic/gin",
			"labstack/echo",
			"revel/revel",
			"gobuffalo/buffalo",
			"go-chi/chi",
			"kataras/iris",
			"zenazn/goji",
			"go-macaron/macaron",
			"go-aah/aah",
		})
	}

	if *resultFileSavePath == "" {
		flag.Usage()
		os.Exit(1)
	}
	if err := os.MkdirAll("stats", 0755); err != nil {
		panic(err)
	}
	var ghData = []Repository{}
	var wg sync.WaitGroup
	wg.Add(len(repositoriesKeys))

	dataChan := make(chan Repository, len(repositoriesKeys))
	for _, rKey := range repositoriesKeys {
		go repositoryData(rKey, *tmpFolder, *debug, dataChan, &wg)
	}
	for range repositoriesKeys {
		ghData = append(ghData, <-dataChan)
	}
	wg.Wait()
	rateAndPrintGreetings(ghData)
	writeCSVStatistics(ghData, *resultFileSavePath)
}

func uniqSlice(s []string) []string {
	unique := make(map[string]bool, len(s))
	us := make([]string, len(unique))
	for _, elem := range s {
		if len(elem) != 0 {
			if !unique[elem] {
				us = append(us, elem)
				unique[elem] = true
			}
		}
	}
	return us
}
