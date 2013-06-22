package tree

import (
	"sort"
)

// Any type that implements tree.Interface may be used as a
// max-tree. A tree must be either first Init()'d or build from scratch.
// The tree functions will panic if called inappropriately (as in, you
// cannot call Pop on an empty tree).
//
// This interface embeds sort.Interface, meaning elements
// can be compared and swapped.
//
// Note that Push and Pop are for this package to call. To add or remove
// elements from a tree, use tree.Push and tree.Pop.
type Interface interface {
	sort.Interface
	// Adds a value to the end of the collection.
	Push(val interface{})
	// Removes the value at the end of the collection.
	Pop() interface{}
}
