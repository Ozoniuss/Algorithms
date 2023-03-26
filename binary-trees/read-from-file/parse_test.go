package main

import "testing"

func TestParse(t *testing.T) {
	root, err := ParseTree("tree.txt")
	if err != nil {
		t.Fatal("should have no error")
	}
	if root.Val != 1 {
		t.Fatalf("root should have value 1, has val %d", root.Val)
	}
	if root.Left != nil {
		t.Fatal("root has no left subtree")
	}
	if root.Right == nil {
		t.Fatalf("root should have right subtree")
	}

	right := root.Right
	if right.Val != 2 {
		t.Fatalf("right subtree should have val 2, has val %d", right.Val)
	}

	if right.Right != nil {
		t.Fatal("right has no right subtree")
	}

	if right.Left == nil {
		t.Fatal("right should have left subtree")
	}

	rightleft := right.Left
	if rightleft.Val != 3 {
		t.Fatalf("rightleft should have value 3, has val %d", rightleft.Val)
	}
}
