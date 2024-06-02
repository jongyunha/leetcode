package hard

import "testing"

func TestNQueensII(t *testing.T) {
	cases := []struct {
		n      int
		output int
	}{
		{4, 2},
		{1, 1},
	}

	for _, c := range cases {
		x := totalNQueens(c.n)
		if x != c.output {
			t.Fatalf("totalNQueens: n=%d, expected=%d, got=%d", c.n, c.output, x)
		}
	}
}
