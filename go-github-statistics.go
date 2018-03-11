// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.
package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

var repositoriesKeys = []string{
	"astaxie/beego",
	"gohugoio/hugo",
	"gin-gonic/gin",
	"labstack/echo",
	"revel/revel",
	"gobuffalo/buffalo",
	"go-chi/chi",
	"kataras/iris",
}

var csvData = [][]string{}

// Repository structure with selcted data keys
type Repository struct {
	Name       string    `json:"name"`
	FullName   string    `json:"full_name"`
	Watchers   int       `json:"watchers"`
	Forks      int       `json:"forks"`
	OpenIssues int       `json:"open_issues"`
	CreatedAt  time.Time `json:"created_at"`
}

func getRemoteJSON(repoKey string) []byte {
	url := "https://api.github.com/repos/" + repoKey
	fullResp := MakeCachedHTTPRequest(url)
	jsonResponse, _, _ := ReadResp(fullResp)
	return jsonResponse
}

func parseJSON(jsonResponse []byte) *Repository {
	result := &Repository{}
	err := json.Unmarshal([]byte(jsonResponse), result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func writeCsv() {
	file, err := os.Create("result.csv")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, value := range csvData {
		err := writer.Write(value)
		if err != nil {
			log.Fatal("Cannot write to file", err)
		}
	}
}

func fillCSVData(res *Repository) {
	csvData = append(csvData, []string{
		res.Name,
		res.FullName,
		fmt.Sprintf("%d/%02d", res.CreatedAt.Year(), res.CreatedAt.Month()),
		fmt.Sprintf("%d", res.Watchers),
		fmt.Sprintf("%d", res.Forks),
		fmt.Sprintf("%d", res.OpenIssues)},
	)
}

func getRepositoryStatistics(RepoKey string) *Repository {
	return parseJSON(getRemoteJSON(RepoKey))
}

func main() {
	csvData = append(csvData, []string{"Name", "Full name", "Created at", "Watchers", "Forks", "Open Issues"})
	for _, rKey := range repositoriesKeys {
		repositoryData := getRepositoryStatistics(rKey)
		fillCSVData(repositoryData)
	}
	writeCsv()
}
