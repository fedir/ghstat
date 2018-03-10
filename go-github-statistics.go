// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.
package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var repositoriesKeys = []string{
	"astaxie/beego",
	"gobuffalo/buffalo",
	"go-chi/chi",
	"gohugoio/hugo",
	"labstack/echo",
	"revel/revel",
	"gin-gonic/gin",
	"kataras/iris",
}

var data = [][]string{}

var csvContent []string

type Match struct {
	Name       string    `json:"name"`
	FullName   string    `json:"full_name"`
	Watchers   int       `json:"watchers"`
	Forks      int       `json:"forks"`
	OpenIssues int       `json:"open_issues"`
	CreatedAt  time.Time `json:"created_at"`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getRemoteJSON(repoKey string) []byte {
	url := "https://api.github.com/repos/" + repoKey
	r, err := myClient.Get(url)
	if err != nil {
		panic(err.Error())
	}
	defer r.Body.Close()
	jsonResponse, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	return jsonResponse
}

func parseJSON(jsonResponse []byte) *Match {
	result := &Match{}
	err := json.Unmarshal([]byte(jsonResponse), result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func createCsv() {
	file, err := os.Create("result.csv")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(csvContent)
	if err != nil {
		log.Fatal("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func PrepareResult(res *Match) string {
	csvContent = append(csvContent, fmt.Sprintf("%s,%s,%d/%d,%d,%d,%d", res.Name, res.FullName, res.CreatedAt.Year(), res.CreatedAt.Month(), res.Watchers, res.Forks, res.OpenIssues))
	return fmt.Sprintf("\tName: %s\n", res.Name) +
		fmt.Sprintf("\tFull name: %s\n", res.FullName) +
		fmt.Sprintf("\tCreated at: %d/%02d\n", res.CreatedAt.Year(), res.CreatedAt.Month()) +
		fmt.Sprintf("\tStars: %d\n", res.Watchers) +
		fmt.Sprintf("\tForks: %d\n", res.Forks) +
		fmt.Sprintf("\tOpen issues : %d\n", res.OpenIssues)
}

func PrintResult(res *Match) {
	fmt.Println(PrepareResult(res))
}

func PrintRepositoryStatistics(RepoKey string) {
	PrintResult(parseJSON(getRemoteJSON(RepoKey)))
}

func main() {
	for _, rKey := range repositoriesKeys {
		PrintRepositoryStatistics(rKey)
	}
	createCsv()
}
