package main

import "testing"

func Test_path_sum_iii(t *testing.T) {
	tests := []struct {
		root      *TreeNode
		targetSum int
		want      int
	}{
		{
			&TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val: -2,
							},
							Right: &TreeNode{
								Val: 1,
							},
						},
						Right: &TreeNode{
							Val: 2,
							Right: &TreeNode{
								Val: 1,
							},
						},
					},
					Right: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
				Right: &TreeNode{
					Val: -3,
					Right: &TreeNode{
						Val: 11,
					},
				},
			},
			8,
			3,
		},
		{
			&TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 11,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 5,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			22,
			3,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := pathSum(tt.root, tt.targetSum); got != tt.want {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}

}
