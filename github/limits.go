// Copyright 2018 Fedir RYKHTIK. All rights reserved.
// Use of this source code is governed by the GNU GPL 3.0
// license that can be found in the LICENSE file.

package github

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fedir/ghstat/httpcache"
	"github.com/fedir/ghstat/timing"
)

// CheckAndPrintRateLimit accesses Github's API to check current limits
func CheckAndPrintRateLimit() {
	type RateLimits struct {
		Resources struct {
			Core struct {
				Limit     int `json:"limit"`
				Remaining int `json:"remaining"`
				Reset     int `json:"reset"`
			} `json:"core"`
			Search struct {
				Limit     int `json:"limit"`
				Remaining int `json:"remaining"`
				Reset     int `json:"reset"`
			} `json:"search"`
			GraphQL struct {
				Limit     int `json:"limit"`
				Remaining int `json:"remaining"`
				Reset     int `json:"reset"`
			} `json:"graphql"`
		} `json:"resources"`
		Rate struct {
			Limit     int `json:"limit"`
			Remaining int `json:"remaining"`
			Reset     int `json:"reset"`
		} `json:"rate"`
	}
	url := "https://api.github.com/rate_limit"
	resp, statusCode, err := httpcache.MakeHTTPRequest(url)
	if err != nil {
		log.Fatalf("Error during checking rate limit : %d %v#", statusCode, err)
	}
	jsonResponse, _, _ := httpcache.ReadResp(resp)
	rateLimits := RateLimits{}
	json.Unmarshal(jsonResponse, &rateLimits)
	fmt.Printf("Core: %d/%d (reset in %d minutes)\n", rateLimits.Resources.Core.Remaining, rateLimits.Resources.Core.Limit, timing.GetRelativeTime(rateLimits.Resources.Core.Reset))
	fmt.Printf("Search: %d/%d (reset in %d minutes)\n", rateLimits.Resources.Search.Remaining, rateLimits.Resources.Search.Limit, timing.GetRelativeTime(rateLimits.Resources.Search.Reset))
	fmt.Printf("GraphQL: %d/%d (reset in %d minutes)\n", rateLimits.Resources.GraphQL.Remaining, rateLimits.Resources.GraphQL.Limit, timing.GetRelativeTime(rateLimits.Resources.GraphQL.Reset))
	fmt.Printf("Rate: %d/%d (reset in %d minutes)\n", rateLimits.Rate.Remaining, rateLimits.Rate.Limit, timing.GetRelativeTime(rateLimits.Rate.Reset))
}
