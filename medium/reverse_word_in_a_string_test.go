package main

import "testing"

func Test_reverse_word_in_a_string(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "test1",
			s:    "the sky is blue",
			want: "blue is sky the",
		},
		{
			name: "test2",
			s:    "  hello world  ",
			want: "world hello",
		},
		{
			name: "test3",
			s:    "a good   example",
			want: "example good a",
		},
		{
			name: "test4",
			s:    "  Bob    Loves  Alice   ",
			want: "Alice Loves Bob",
		},
		{
			name: "test5",
			s:    "Alice does not even like bob",
			want: "bob like even not does Alice",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseWords(tt.s)
			if got != tt.want {
				t.Errorf("got: %s, want: %s", got, tt.want)
			}
		})
	}

}
