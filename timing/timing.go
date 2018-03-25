package timing

import "time"

// GetRelativeTime helps to find difference between different two Unix timestamps
func GetRelativeTime(unixTime int) int {
	now := int(time.Now().Unix())
	return int((float64(unixTime) - float64(now)) / 60)
}
