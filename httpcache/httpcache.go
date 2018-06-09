// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.

// Package httpcache provides an interface to HTTP request static cache,
//
// Usage example :
//  func main() {
//    url := "https://api.github.com/repos/astaxie/beego/contributors"
//    body := httpcache.MakeCachedHTTPRequest(url)
//    jsonResp, linkHeader, _ := httpcache.ReadResp(body)
//    fmt.Printf("%s\n%s", jsonResp, linkHeader)
//  }
//
package httpcache

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

var cacheTTL = 3600

var httpClient = &http.Client{Timeout: 10 * time.Second}

const dumpBody = true

// GetFilename gets encoded filename for cache usage
func GetFilename(url string) string {
	encoder := sha256.New()
	encoder.Write([]byte(url))
	return hex.EncodeToString(encoder.Sum(nil))
}

// MakeHTTPRequest makes a non-cacheable request to the external URL
func MakeHTTPRequest(url string) ([]byte, int, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Cannont prepare the HTTP request", err)
	}
	if os.Getenv("GH_USR") != "" && os.Getenv("GH_PASS") != "" {
		req.SetBasicAuth(os.Getenv("GH_USR"), os.Getenv("GH_PASS"))
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
	return body, resp.StatusCode, err
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

// ReadResp :  reads response from the cached HTTP query.
func ReadResp(fullResp []byte) ([]byte, string, error) {
	r := bufio.NewReader(bytes.NewReader(fullResp))
	resp, err := http.ReadResponse(r, nil)
	if err != nil {
		log.Printf("%v\n%s", err, fullResp)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("%v\n%s", err, resp.Body)
	}
	linkHeader := resp.Header.Get("Link")
	return body, linkHeader, err
}

// MakeCachedHTTPRequest makes a cacheable request to the external URL
// If the request was already made once, it will be not done again,
// but read from the file in temporary folder.
// Currently is was tested only for GET queries
func MakeCachedHTTPRequest(url string, tmpFolder string, debug bool) []byte {
	var fullResp []byte
	filename := GetFilename(url)
	filepath := filename
	if tmpFolder != "" {
		filepath = tmpFolder + "/" + filename
	}
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		if debug == true {
			fmt.Println("HTTP query: " + url)
		}
		resp, statusCode, err := MakeHTTPRequest(url)
		if err != nil {
			panic(err)
		}
		if statusCode == 403 {
			log.Fatalf("Looks like the rate limit is exceeded, please try again in 60 minutes. Or make a pull request with authentification feature.")
		} else if statusCode == 202 {
			log.Printf("Server need some time to prepare request. Trying again.")
			time.Sleep(2 * time.Second)
			return MakeCachedHTTPRequest(url, tmpFolder, debug)
		} else if statusCode != 200 {
			log.Fatalf("The status code of URL %s is not OK : %d", url, statusCode)
		}
		saveRespToFile(filepath, resp)
		fullResp = loadRespFromFile(filepath)
	} else {
		if debug == true {
			fmt.Println("Loaded results directly from: " + filepath)
		}
		fullResp = loadRespFromFile(filepath)
	}
	return fullResp
}
