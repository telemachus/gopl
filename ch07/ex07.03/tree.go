package main

import (
	"fmt"
	"io"
)

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) inorderPrint(w io.Writer) {
	fmt.Fprintf(w, "{")
	inorderRecursive(w, t)
	fmt.Fprintf(w, " }\n")
}

func inorderRecursive(w io.Writer, t *tree) {
	if t == nil {
		return
	}
	inorderRecursive(w, t.left)
	fmt.Fprintf(w, " %d", t.value)
	inorderRecursive(w, t.right)
}

func (t *tree) preorderPrint(w io.Writer) {
	fmt.Fprintf(w, "{")
	preorderRecursive(w, t)
	fmt.Fprintf(w, " }\n")
}

func preorderRecursive(w io.Writer, t *tree) {
	if t == nil {
		return
	}
	fmt.Fprintf(w, " %d", t.value)
	preorderRecursive(w, t.left)
	preorderRecursive(w, t.right)
}

func (t *tree) postorderPrint(w io.Writer) {
	fmt.Fprintf(w, "{")
	preorderRecursive(w, t)
	fmt.Fprintf(w, " }\n")
}

func postorderRecursive(w io.Writer, t *tree) {
	if t == nil {
		return
	}
	postorderRecursive(w, t.left)
	postorderRecursive(w, t.right)
	fmt.Fprintf(w, " %d", t.value)
}

func morrisPrint(w io.Writer, t *tree) {
	fmt.Fprintf(w, "{")

	current := t
	for current != nil {
		if current.left == nil {
			fmt.Fprintf(w, " %d", current.value)
			current = current.right
		} else {
			pre := current.left
			for pre.right != nil && pre.right != current {
				pre = pre.right
			}
			if pre.right == nil {
				pre.right = current
				current = current.left
			} else {
				pre.right = nil
				fmt.Fprintf(w, " %d", current.value)
				current = current.right
			}
		}
	}

	fmt.Fprintf(w, " }\n")
}

func main() {
	// t := add(nil, 1)
	// t = add(t, 7)
	// t = add(t, 2)
	// t = add(t, 6)
	// t = add(t, 3)
	// t = add(t, 8)
	// t = add(t, 5)
	// t = add(t, 4)

	// fmt.Printf("In order, recursive: ")
	// t.inorderPrint(os.Stdin)

	// fmt.Printf("Morris traversal: ")
	// morrisPrint(os.Stdin, t)
}
