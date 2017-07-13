package merge_two_binary_trees

import (
	"reflect"
	"testing"
)

type testCase struct {
	t1, t2, v *TreeNode
}

var (
	t1 = &TreeNode{Val: 1}
	t2 = &TreeNode{Val: 2}
	v1 = &TreeNode{Val: t1.Val + t2.Val}
	t3 = &TreeNode{Val: 1, Left: t2}
	t4 = &TreeNode{Val: 1, Left: t1, Right: t2}
	v2 = &TreeNode{
		Val:   2,
		Left:  v1,
		Right: t2,
	}

	testCases = []testCase{
		testCase{nil, nil, nil},
		testCase{t1, nil, t1},
		testCase{nil, t1, t1},
		testCase{t1, t2, v1},
		testCase{t3, t4, v2},
	}
)

func TestMergeTwoBinaryTrees(t *testing.T) {
	for _, test := range testCases {
		if v := mergeTrees(test.t1, test.t2); !reflect.DeepEqual(test.v, v) {
			t.Errorf("expect %v, got %v\n", test.v, v)
		}
	}
}
