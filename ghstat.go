// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.

// ghstat - statistical multi-criteria decision-making comparator for Github's projects.
package main

import (
	"flag"
	"fmt"
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
			"gin-gonic/gin",
			"gofiber/fiber",
			"labstack/echo",
			"go-chi/chi",
			"beego/beego",
			"gohugoio/hugo",
			"gobuffalo/buffalo",
			"revel/revel",
			"kataras/iris",
			"go-macaron/macaron",
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

	total := len(repositoriesKeys)
	fmt.Printf("Fetching %d repositories...\n", total)
	dataChan := make(chan Repository, total)
	for _, rKey := range repositoriesKeys {
		go repositoryData(rKey, *tmpFolder, *debug, dataChan, &wg)
	}
	for i := 1; i <= total; i++ {
		r := <-dataChan
		ghData = append(ghData, r)
		fmt.Printf("[%d/%d] %s\n", i, total, r.Name)
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
