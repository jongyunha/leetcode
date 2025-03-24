package main

import "testing"

// You are given a positive integer days representing the total number of days an employee is available for work (starting from day 1). You are also given a 2D array meetings of size n where, meetings[i] = [start_i, end_i] represents the starting and ending days of meeting i (inclusive).
//
// Return the count of days when the employee is available for work but no meetings are scheduled.
//
// Note: The meetings may overlap.
//
// Example 1:
//
// Input: days = 10, meetings = [[5,7],[1,3],[9,10]]
//
// Output: 2
//
// Explanation:
//
// There is no meeting scheduled on the 4th and 8th days.
//
// Example 2:
//
// Input: days = 5, meetings = [[2,4],[1,3]]
//
// Output: 1
//
// Explanation:
//
// There is no meeting scheduled on the 5th day.
//
// Example 3:
//
// Input: days = 6, meetings = [[1,6]]
//
// Output: 0
//
// Explanation:
//
// Meetings are scheduled for all working days.
//
// Constraints:
//
// 1 <= days <= 109
// 1 <= meetings.length <= 105
// meetings[i].length == 2
// 1 <= meetings[i][0] <= meetings[i][1] <= days
func TestCountDays(t *testing.T) {
	tests := []struct {
		days     int
		meetings [][]int
		want     int
	}{
		{10, [][]int{{5, 7}, {1, 3}, {9, 10}}, 2},
		{5, [][]int{{2, 4}, {1, 3}}, 1},
		{6, [][]int{{1, 6}}, 0},
		{4, [][]int{{2, 3}, {1, 2}, {2, 3}, {2, 4}, {1, 2}, {1, 3}}, 0},
		{14, [][]int{{2, 3}, {3, 4}, {12, 14}, {8, 10}}, 5},
		{8, [][]int{{3, 4}, {4, 8}, {2, 5}, {3, 8}}, 1},
	}
	for _, tt := range tests {
		if got := countDays(tt.days, tt.meetings); got != tt.want {
			t.Errorf("countDays(%v, %v) = %v, want %v", tt.days, tt.meetings, got, tt.want)
		}
	}
}
