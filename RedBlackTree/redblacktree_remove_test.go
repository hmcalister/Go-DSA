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

func TestRemoveRoot(t *testing.T) {
	items := []int{3, 4, 2, 5, 1}
	tree := redblacktree.New[int](comparator.DefaultIntegerComparator)
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

func TestRemoveTwoChildNode(t *testing.T) {
	items := []int{1, 3, 2, 4}
	tree := redblacktree.New[int](comparator.DefaultIntegerComparator)
	for _, item := range items {
		tree.Add(item)
	}

	err := tree.Remove(3)
	if err != nil {
		t.Errorf("encountered error (%v) when removing two child node", err)
	}

	node, err := tree.Find(3)
	if node != nil || err == nil {
		t.Errorf("found node that should have been deleted after deleting two child node")
	}
}

func TestRemoveNodeWithOnlyLeftChild(t *testing.T) {
	items := []int{5, 4, 3, 2, 1}
	tree := redblacktree.New[int](comparator.DefaultIntegerComparator)
	for _, item := range items {
		tree.Add(item)
	}

	err := tree.Remove(3)
	if err != nil {
		t.Errorf("encountered error (%v) when removing one child node", err)
	}

	node, err := tree.Find(3)
	if node != nil || err == nil {
		t.Errorf("found node that should have been deleted after deleting one child node")
	}
}

func TestRemoveNodeWithOnlyRightChild(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	tree := redblacktree.New[int](comparator.DefaultIntegerComparator)
	for _, item := range items {
		tree.Add(item)
	}

	err := tree.Remove(3)
	if err != nil {
		t.Errorf("encountered error (%v) when removing one child node", err)
	}

	node, err := tree.Find(3)
	if node != nil || err == nil {
		t.Errorf("found node that should have been deleted after deleting one child node")
	}
}

