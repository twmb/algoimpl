// Package binary has functions for abusing a binary tree.
//
// To use the binary tree functions, start with an empty
// tree.Interface and Insert all elements into it:
//     empty := Ints([]int{})
//     for i := 0; i < 10; i++ {
//       binary.Insert(empty)
//     }
//
package binary

type BinaryTree struct {
	root *node
	size int
}

// Returns a new, empty binary tree.
func New() *BinaryTree {
	return &BinaryTree{size: 0}
}

type node struct {
	parent *node
	left   *node
	right  *node
	value  Comparable
}

// A type that implements the comparable interface can be used in binary trees.
type Comparable interface {
	// Returns -1 if other is less, 0 if they are equal and 1 if other is greater.
	CompareTo(other Comparable) int
}

func walkInOrder(node *node, walked []Comparable) []Comparable {
	if node != nil {
		walked = walkInOrder(node.left, walked)
		walked = append(walked, node.value)
		walked = walkInOrder(node.right, walked)
		return walked
	}
	return nil
}

// Returns an in order Comparable slice of the binary tree.
func (b *BinaryTree) Walk() []Comparable {
	walked := make([]Comparable, 0, b.size)
	return walkInOrder(b.root, walked)
}

func walkPreOrder(node *node, walked []Comparable) []Comparable {
	if node != nil {
		walked = append(walked, node.value)
		walked = walkPreOrder(node.left, walked)
		walked = walkPreOrder(node.right, walked)
		return walked
	}
	return nil
}

// Returns a pre order Comparable slice of the binary tree.
func (b *BinaryTree) WalkPreOrder() []Comparable {
	walked := make([]Comparable, 0, b.size)
	return walkPreOrder(b.root, walked)
}

func walkPostOrder(node *node, walked []Comparable) []Comparable {
	if node != nil {
		walked = walkPostOrder(node.left, walked)
		walked = walkPostOrder(node.right, walked)
		walked = append(walked, node.value)
		return walked
	}
	return nil
}

// Returns a post order Comparable slice of the binary tree.
func (b *BinaryTree) WalkPostOrder() []Comparable {
	walked := make([]Comparable, 0, b.size)
	return walkPostOrder(b.root, walked)
}

// Returns true if the binary tree contains the target Comparable
func (b *BinaryTree) Contains(target Comparable) bool {
	current := b.root
	for current != nil {
		switch current.value.CompareTo(target) {
		case -1:
			current = current.left
		case 0:
			return true
		case 1:
			current = current.right
		}
	}
	return false
}

func minimum(startNode *node) Comparable {
	current := startNode
	for current.left != nil {
		current = current.left
	}
	return current.value
}

// Returns the minimum value of the binary tree.
func (b *BinaryTree) Minimum() Comparable {
	return minimum(b.root)
}

// Returns the maximum value of the binary tree.
func (b *BinaryTree) Maximum() Comparable {
	current := b.root
	for current.right != nil {
		current = current.right
	}
	return current.value
}

// Returns a pointer to the successor value of a target Comparable in a binary tree.
// The pointer will be nil if the tree does not contain the target
// of if there is no successor (i.e., you want the successor to the maximum value)
func (b *BinaryTree) Successor(target Comparable) *Comparable {
	current := b.root
	for current != nil {
		switch current.value.CompareTo(target) {
		case -1:
			current = current.left
		case 0:
			break
		case 1:
			current = current.right
		}
	}
	if current == nil {
		return nil
	}
	if current.right != nil {
		minimum := minimum(current.right)
		return &minimum
	}
	if current == b.root { // if root has no right side, at max, no successor
		return nil
	}
	parent := current.parent
	for parent != nil && current == parent.right {
		current = parent
		parent = current.parent
	}
	return &current.value
}

// Inserts a comparable value into a binary tree.
// Expected running time O(lg n), worst case running time O(n).
func (b *BinaryTree) Insert(value Comparable) {
	newNode := new(node)
	newNode.value = value
	y := (*node)(nil)
	x := b.root
	for x != nil {
		y = x
		if value.CompareTo(x.value) == -1 { // less
			x = x.left
		} else {
			x = x.right
		}
	}
	newNode.parent = y
	if newNode.parent == nil {
		b.root = newNode
	} else if newNode.value.CompareTo(y.value) == -1 {
		y.left = newNode
	} else {
		y.right = newNode
	}
	b.size++
}
