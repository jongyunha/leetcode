package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_binary_tree_right_side_view(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want []int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
					},
				},
			},
			want: []int{1, 3, 4},
		},
		{
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			want: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := rightSideView(tt.root)
			if !assert.Equal(t, got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
