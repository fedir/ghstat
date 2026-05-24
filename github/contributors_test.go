package github

import "testing"

func TestGetActiveForkersPercentage(t *testing.T) {
	tests := []struct {
		contributors int
		forkers      int
		expected     float64
	}{
		{10, 100, 10.0},
		{50, 100, 50.0},
		{100, 100, 100.0},
		{0, 100, 0.0},
	}
	for _, tc := range tests {
		got := GetActiveForkersPercentage(tc.contributors, tc.forkers)
		if got != tc.expected {
			t.Errorf("GetActiveForkersPercentage(%d, %d) = %.2f, want %.2f",
				tc.contributors, tc.forkers, got, tc.expected)
		}
	}
}

func TestGetCommitsByDay(t *testing.T) {
	tests := []struct {
		commits int
		age     int
		want    float64
	}{
		{365, 365, 1.0},
		{0, 365, 0.0},
		{100, 0, 0.0},
		{730, 365, 2.0},
	}
	for _, tc := range tests {
		got := GetCommitsByDay(tc.commits, tc.age)
		if got != tc.want {
			t.Errorf("GetCommitsByDay(%d, %d) = %f, want %f", tc.commits, tc.age, got, tc.want)
		}
	}
}
