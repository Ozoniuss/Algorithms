package llist

import (
	"errors"
	"testing"
)

func TestString(t *testing.T) {
	var tests = []struct {
		name     string
		input    *ListNode
		expected string
	}{
		{
			name:     "Testing nil",
			input:    nil,
			expected: "nil",
		},
		{
			name: "Testing one element",
			input: &ListNode{
				Val:  1,
				Next: nil,
			},
			expected: "1",
		},
		{
			name: "Testing two elements",
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
			expected: "1 -> 2",
		},
		{
			name: "Testing three elements",
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
			expected: "1 -> 2 -> 3",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			out := tt.input.String()
			if out != tt.expected {
				t.Errorf("got %s, expected %s", out, tt.expected)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	var tests = []struct {
		name     string
		l1, l2   *ListNode
		expected bool
	}{
		{
			name:     "Testing nil comparison",
			l1:       nil,
			l2:       nil,
			expected: true,
		},
		{
			name: "Testing one nil list",
			l1:   nil,
			l2: &ListNode{
				Val:  1,
				Next: nil,
			},
			expected: false,
		},
		{
			name: "Testing the other nil list",
			l1: &ListNode{
				Val:  1,
				Next: nil,
			},
			l2:       nil,
			expected: false,
		},
		{
			name: "Testing equal lists with more elements",
			l1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
			l2: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
			expected: true,
		},
		{
			name: "Testing lists with different number of elements",
			l1: &ListNode{
				Val:  1,
				Next: nil,
			},
			l2: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
			expected: false,
		},
		{
			name: "Testing the same number of elements but different values",
			l1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
			l2: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if Equal(tt.l1, tt.l2) != tt.expected {
				t.Error("Failed comparison")
			}
		})
	}
}

func TestFromStringArray(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected *ListNode
		isError  bool
	}{
		{
			name:     "Testing nil",
			input:    "",
			expected: nil,
			isError:  false,
		},
		{
			name:     "Testing empty array",
			input:    "[]",
			expected: nil,
			isError:  false,
		},
		{
			name:     "Testing one character",
			input:    "[",
			expected: nil,
			isError:  true,
		},
		{
			name:  "Testing one element",
			input: "[1004]",
			expected: &ListNode{
				Val:  1004,
				Next: nil,
			},
			isError: false,
		},
		{
			name:  "Testing negative element",
			input: "[-1004]",
			expected: &ListNode{
				Val:  -1004,
				Next: nil,
			},
			isError: false,
		},
		{
			name:     "Testing non-number element",
			input:    "[1a04]",
			expected: nil,
			isError:  true,
		},
		{
			name:  "Testing multiple elements",
			input: "[100,200]",
			expected: &ListNode{
				Val: 100,
				Next: &ListNode{
					Val:  200,
					Next: nil,
				},
			},
			isError: false,
		},
		{
			name:     "Invalid end comma",
			input:    "[100,200,]",
			expected: nil,
			isError:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			l, err := FromStringArray(tt.input)
			if tt.isError {
				var newerr ErrInvalidInputArray
				if !errors.As(err, &newerr) {
					t.Error("Did not receive error")
				}
			} else {
				if err != nil {
					t.Errorf("Got unexpected error: %s", err.Error())
				}
			}
			if !Equal(tt.expected, l) {
				t.Errorf("Got a different output than expected: %s", l.String())
			}
		})
	}
}
