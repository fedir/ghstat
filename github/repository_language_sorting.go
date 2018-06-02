package github

import (
	"fmt"
	"sort"
	"strings"
)

// LanguageList contains all languages
type LanguageList []Language

// Language with size
type Language struct {
	Language string
	Size     float64
}

func repositoryLanguagesBySize(languages map[string]interface{}) string {
	languagesStats := make([]string, 0, len(languages))
	var totalBytes float64

	// Count total bytes
	for key := range languages {
		totalBytes = totalBytes + languages[key].(float64)
	}
	ll := make(LanguageList, len(languages))
	i := 0
	for k, v := range languages {
		ll[i] = Language{k, v.(float64)}
		i++
	}
	sort.Sort(sort.Reverse(ll))
	for i := 0; i < len(ll); i++ {
		languagePercentage := ll[i].Size / totalBytes * 100
		languagesStats = append(languagesStats, fmt.Sprintf("%s(%.2f)", ll[i].Language, languagePercentage))
	}
	languagesStatsResult := strings.Join(languagesStats, ",")
	return languagesStatsResult
}

func (l LanguageList) Len() int           { return len(l) }
func (l LanguageList) Less(i, j int) bool { return l[i].Size < l[j].Size }
func (l LanguageList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
