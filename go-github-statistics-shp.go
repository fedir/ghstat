// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.
package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

var cacheTtl = 3600

var tmpFolder = "tmp"

var httpClient = &http.Client{Timeout: 10 * time.Second}

const dumpBody = true

func getFilename(url string) string {
	encoder := sha256.New()
	encoder.Write([]byte(url))
	return hex.EncodeToString(encoder.Sum(nil))
}

func makeHTTPRequest(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Cannont prepare the HTTP request", err)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal("Cannot process the HTTP request", err)
	}
	defer resp.Body.Close()
	body, err := httputil.DumpResponse(resp, dumpBody)
	if err != nil {
		log.Fatal("Cannont dump the body of HTTP response", err)
	}
	return body, err
}

func saveRespToFile(file string, resp []byte) {
	err := ioutil.WriteFile(file, resp, 0644)
	if err != nil {
		panic(err)
	}
}

func loadRespFromFile(file string) []byte {
	resp, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return resp
}

func ReadResp(fullResp []byte) ([]byte, string, error) {
	r := bufio.NewReader(bytes.NewReader(fullResp))
	resp, err := http.ReadResponse(r, nil)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	linkHeader := resp.Header.Get("Link")
	return body, linkHeader, err
}

func MakeCachedHTTPRequest(url string) []byte {
	var fullResp []byte
	filename := getFilename(url)
	filepath := filename
	if tmpFolder != "" {
		filepath = tmpFolder + "/" + filename
	}
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		fmt.Println("HTTP query: " + url)
		resp, err := makeHTTPRequest(url)
		if err != nil {
			panic(err)
		}
		saveRespToFile(filepath, resp)
		fullResp = loadRespFromFile(filepath)
	} else {
		fmt.Println("Loaded results directly from: " + filepath)
		fullResp = loadRespFromFile(filepath)
	}
	return fullResp
}

/*
func main() {
	url := "https://api.github.com/repos/astaxie/beego/contributors"
	body := MakeCachedHTTPRequest(url)
	jsonResp, linkHeader, _ := ReadResp(body)
	fmt.Printf("%s\n%s", jsonResp, linkHeader)
}
*/
