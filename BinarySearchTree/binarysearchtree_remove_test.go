package binarysearchtree_test

import (
	"testing"

	binarysearchtree "github.com/hmcalister/Go-DSA/BinarySearchTree"
	comparator "github.com/hmcalister/Go-DSA/Comparator"
)

func TestRemoveRootAsOnlynode(t *testing.T) {
	items := []int{1}
	tree := binarysearchtree.New[int](comparator.DefaultIntegerComparator)
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

func TestRemoveRoot(t *testing.T) {
	items := []int{3, 4, 2, 5, 1}
	tree := binarysearchtree.New[int](comparator.DefaultIntegerComparator)
	for _, item := range items {
		tree.Add(item)
	}

	err := tree.Remove(3)
	if err != nil {
		t.Errorf("encountered error (%v) when removing root node", err)
	}

	node, err := tree.Find(3)
	if node != nil || err == nil {
		t.Errorf("found node that should have been deleted after deleting root")
	}
}

