package hard

import "testing"

func Test_regular_expression_matching(t *testing.T) {
	tests := []struct {
		s    string
		p    string
		want bool
	}{
		{"aa", "a", false},
		{"aa", "a*", true},
		{"cb", ".a", false},
		{"adceb", ".*a*b", true},
		{"acdcb", "a.*c.b", false},
		{"aab", "c*a*b", true},
	}
	for _, tt := range tests {
		t.Run(tt.s+" "+tt.p, func(t *testing.T) {
			if got := isMatch2(tt.s, tt.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
