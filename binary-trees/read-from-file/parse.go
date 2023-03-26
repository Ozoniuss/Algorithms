package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ParseTree parses a binary tree from a file, assuming the position of each
// node in a tree is known.
func ParseTree(in string) (*TreeNode, error) {

	f, err := os.Open(in)
	if err != nil {
		return nil, err
	}

	// Create root node
	var t = &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}

	// Stores the nodes that have already been read.
	var existing = make(map[int]*TreeNode)

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		content := strings.Split(line, " ")
		if len(content) != 2 {
			return nil, errors.New("invalid format")
		}

		pos, err := strconv.Atoi(content[0])
		if err != nil {
			return nil, fmt.Errorf("invalid position (must be int): %s", content[0])
		}
		value, err := strconv.Atoi(content[1])
		if err != nil {
			return nil, fmt.Errorf("invalid node value (must be int): %s", content[1])
		}

		// Add root node to the existing values.
		if pos == 1 {
			t.Val = value
			existing[1] = t
		} else {
			parentNode, ok := existing[pos/2]
			if !ok {
				return nil, fmt.Errorf("parent not found for node at position %d (supposed to be at pos %d)", pos, pos/2)
			}
			currentNode := &TreeNode{
				Val:   value,
				Left:  nil,
				Right: nil,
			}
			if pos%2 == 0 {
				parentNode.Left = currentNode
			} else {
				parentNode.Right = currentNode
			}
			existing[pos] = currentNode
		}
	}
	return t, nil
}
