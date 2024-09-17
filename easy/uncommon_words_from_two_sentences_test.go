package easy

import "testing"

func Test_uncommon_words_from_two_sentences(t *testing.T) {
	type args struct {
		A string
		B string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test case 1",
			args: args{
				A: "this apple is sweet",
				B: "this apple is sour",
			},
			want: []string{"sweet", "sour"},
		},
		{
			name: "test case 2",
			args: args{
				A: "apple apple",
				B: "banana",
			},
			want: []string{"banana"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uncommonFromSentences(tt.args.A, tt.args.B); !equal(got, tt.want) {
				t.Errorf("uncommon_words_from_two_sentences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[string]int)
	for _, v := range a {
		m[v]++
	}
	for _, v := range b {
		if m[v] == 0 {
			return false
		}
		m[v]--
	}
	return true
}
