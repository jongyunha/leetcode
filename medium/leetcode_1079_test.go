package main

import "testing"

// You have n  tiles, where each tile has one letter tiles[i] printed on it.
//
// Return the number of possible non-empty sequences of letters you can make using the letters printed on those tiles.
//
// Example 1:
//
// Input: tiles = "AAB"
// Output: 8
// Explanation: The possible sequences are "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA".
// Example 2:
//
// Input: tiles = "AAABBC"
// Output: 188
// Example 3:
//
// Input: tiles = "V"
// Output: 1
func TestNumTilePossibilities(t *testing.T) {
	tests := []struct {
		tiles string
		want  int
	}{
		{"AAB", 8},
		{"AAABBC", 188},
		{"V", 1},
	}
	for _, tt := range tests {
		if got := numTilePossibilities(tt.tiles); got != tt.want {
			t.Errorf("numTilePossibilities(%v) = %v, want %v", tt.tiles, got, tt.want)
		}
	}
}
