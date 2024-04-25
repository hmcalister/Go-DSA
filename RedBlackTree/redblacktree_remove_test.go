package redblacktree_test

import (
	"math/rand"
	"testing"

	comparator "github.com/hmcalister/Go-DSA/Comparator"
	redblacktree "github.com/hmcalister/Go-DSA/RedBlackTree"
)

func TestRemoveRootAsOnlynode(t *testing.T) {
	items := []int{1}
	tree := redblacktree.New[int](comparator.DefaultIntegerComparator)
	for _, item := range items {
		tree.Add(item)
	}

	err := tree.Remove(1)
	if err != nil {
		t.Errorf("encountered error (%v) when removing root node", err)
	}

	node, err := tree.Find(1)
	if node != nil || err == nil {
		t.Errorf("found node that should have been deleted after deleting root")
	}
}

