// Package btree supports a btree for int or string or a struct that provides a method that conforms to the Val interface.
// Code from github.com/ross-oreto/go-tree
package btree

import (
	"fmt"
)

// Btree represents an AVL tree
type Btree struct {
	root   *Node
	values []Val
	len    int
}

// Val interface to define the compare method used to insert and find values
type Val interface {
	Comp(val Val) int8
}

// Node represents a node in the tree with a value, left and right children, and a height/balance of the node.
type Node struct {
	Value       Val
	left, right *Node
	height      int8
}

// New returns a new btree
func New() *Btree { return new(Btree).Init() }

// Init initializes all values/clears the tree and returns the tree pointer
func (t *Btree) Init() *Btree {
	t.root = nil
	t.values = nil
	t.len = 0
	return t
}

// String returns a string representation of the tree values
func (t *Btree) String() string {
	return fmt.Sprint(t.Values())
}

// Empty returns true if the tree is empty
func (t *Btree) Empty() bool {
	return t.root == nil
}

// NotEmpty returns true if the tree is not empty
func (t *Btree) NotEmpty() bool {
	return t.root != nil
}

func (t *Btree) balance() int8 {
	if t.root != nil {
		return balance(t.root)
	}
	return 0
}

// Insert inserts a new value into the tree and returns the tree pointer
func (t *Btree) Insert(value Val) *Btree {
	added := false
	t.root = insert(t.root, value, &added)
	if added {
		t.len++
	}
	t.values = nil
	return t
}

func insert(n *Node, value Val, added *bool) *Node {
	if n == nil {
		*added = true
		return (&Node{Value: value}).Init()
	}
	c := value.Comp(n.Value)
	if c > 0 {
		n.right = insert(n.right, value, added)
	} else if c < 0 {
		n.left = insert(n.left, value, added)
	} else {
		n.Value = value
		*added = false
		return n
	}

	n.height = n.maxHeight() + 1
	c = balance(n)

	if c > 1 {
		c = value.Comp(n.left.Value)
		if c < 0 {
			return n.rotateRight()
		} else if c > 0 {
			n.left = n.left.rotateLeft()
			return n.rotateRight()
		}
	} else if c < -1 {
		c = value.Comp(n.right.Value)
		if c > 0 {
			return n.rotateLeft()
		} else if c < 0 {
			n.right = n.right.rotateRight()
			return n.rotateLeft()
		}
	}
	return n
}

// InsertAll inserts all the values into the tree and returns the tree pointer
func (t *Btree) InsertAll(values []Val) *Btree {
	for _, v := range values {
		t.Insert(v)
	}
	return t
}

// Contains returns true if the tree contains the specified value
func (t *Btree) Contains(value Val) bool {
	return t.Get(value) != nil
}

// ContainsAny returns true if the tree contains any of the values
func (t *Btree) ContainsAny(values []Val) bool {
	for _, v := range values {
		if t.Contains(v) {
			return true
		}
	}
	return false
}

// ContainsAll returns true if the tree contains all of the values
func (t *Btree) ContainsAll(values []Val) bool {
	for _, v := range values {
		if !t.Contains(v) {
			return false
		}
	}
	return true
}

// Get returns the node value associated with the search value
func (t *Btree) Get(value Val) Val {
	var node *Node
	if t.root != nil {
		node = t.root.get(value)
	}
	if node != nil {
		return node.Value
	}
	return nil
}

// Len return the number of nodes in the tree
func (t *Btree) Len() int {
	return t.len
}

// Head returns the first value in the tree
func (t *Btree) Head() Val {
	if t.root == nil {
		return nil
	}
	var beginning = t.root
	for beginning.left != nil {
		beginning = beginning.left
	}
	if beginning == nil {
		for beginning.right != nil {
			beginning = beginning.right
		}
	}
	if beginning != nil {
		return beginning.Value
	}
	return nil
}

// Tail returns the last value in the tree
func (t *Btree) Tail() Val {
	if t.root == nil {
		return nil
	}
	var beginning = t.root
	for beginning.right != nil {
		beginning = beginning.right
	}
	if beginning == nil {
		for beginning.left != nil {
			beginning = beginning.left
		}
	}
	if beginning != nil {
		return beginning.Value
	}
	return nil
}

// Values returns a slice of all the values in tree in order
func (t *Btree) Values() []Val {
	if t.values == nil {
		t.values = make([]Val, t.len)
		t.Ascend(func(n *Node, i int) bool {
			t.values[i] = n.Value
			return true
		})
	}
	return t.values
}

// Delete deletes the node from the tree associated with the search value
func (t *Btree) Delete(value Val) *Btree {
	deleted := false
	t.root = deleteNode(t.root, value, &deleted)
	if deleted {
		t.len--
	}
	t.values = nil
	return t
}

// DeleteAll deletes the nodes from the tree associated with the search values
func (t *Btree) DeleteAll(values []Val) *Btree {
	for _, v := range values {
		t.Delete(v)
	}
	return t
}

func deleteNode(n *Node, value Val, deleted *bool) *Node {
	if n == nil {
		return n
	}

	c := value.Comp(n.Value)

	if c < 0 {
		n.left = deleteNode(n.left, value, deleted)
	} else if c > 0 {
		n.right = deleteNode(n.right, value, deleted)
	} else {
		if n.left == nil {
			t := n.right
			n.Init()
			return t
		} else if n.right == nil {
			t := n.left
			n.Init()
			return t
		}
		t := n.right.min()
		n.Value = t.Value
		n.right = deleteNode(n.right, t.Value, deleted)
		*deleted = true
	}

	//re-balance
	if n == nil {
		return n
	}
	n.height = n.maxHeight() + 1
	bal := balance(n)
	if bal > 1 {
		if balance(n.left) >= 0 {
			return n.rotateRight()
		}
		n.left = n.left.rotateLeft()
		return n.rotateRight()
	} else if bal < -1 {
		if balance(n.right) <= 0 {
			return n.rotateLeft()
		}
		n.right = n.right.rotateRight()
		return n.rotateLeft()
	}

	return n
}

// Pop deletes the last node from the tree and returns its value
func (t *Btree) Pop() Val {
	value := t.Tail()
	if value != nil {
		t.Delete(value)
	}
	return value
}

// Pull deletes the first node from the tree and returns its value
func (t *Btree) Pull() Val {
	value := t.Head()
	if value != nil {
		t.Delete(value)
	}
	return value
}

// NodeIterator expresses the iterator function used for traversals
type NodeIterator func(n *Node, i int) bool

// Ascend performs an ascending order traversal of the tree calling the iterator function on each node
// the iterator will continue as long as the NodeIterator returns true
func (t *Btree) Ascend(iterator NodeIterator) {
	var i int
	if t.root != nil {
		t.root.iterate(iterator, &i, true)
	}
}

// Descend performs a descending order traversal of the tree using the iterator
// the iterator will continue as long as the NodeIterator returns true
func (t *Btree) Descend(iterator NodeIterator) {
	var i int
	if t.root != nil {
		t.root.rIterate(iterator, &i, true)
	}
}

// Debug prints out useful debug information about the tree for debugging purposes
func (t *Btree) Debug() {
	fmt.Println("----------------------------------------------------------------------------------------------")
	if t.Empty() {
		fmt.Println("tree is empty")
	} else {
		fmt.Println(t.Len(), "elements")
	}

	t.Ascend(func(n *Node, i int) bool {
		if t.root.Value == n.Value {
			fmt.Print("ROOT ** ")
		}
		n.Debug()
		return true
	})
	fmt.Println("----------------------------------------------------------------------------------------------")
}

// Init initializes the values of the node or clears the node and returns the node pointer
func (n *Node) Init() *Node {
	n.height = 1
	n.left = nil
	n.right = nil
	return n
}

// String returns a string representing the node
func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

// Debug prints out useful debug information about the tree node for debugging purposes
func (n *Node) Debug() {
	var children string
	if n.left == nil && n.right == nil {
		children = "no children |"
	} else if n.left != nil && n.right != nil {
		children = fmt.Sprint("left child:", n.left.String(), " right child:", n.right.String())
	} else if n.right != nil {
		children = fmt.Sprint("right child:", n.right.String())
	} else {
		children = fmt.Sprint("left child:", n.left.String())
	}

	fmt.Println(n.String(), "|", "height", n.height, "|", "balance", balance(n), "|", children)
}

func height(n *Node) int8 {
	if n != nil {
		return n.height
	}
	return 0
}

func balance(n *Node) int8 {
	if n == nil {
		return 0
	}
	return height(n.left) - height(n.right)
}

func (n *Node) get(val Val) *Node {
	var node *Node
	c := val.Comp(n.Value)
	if c < 0 {
		if n.left != nil {
			node = n.left.get(val)
		}
	} else if c > 0 {
		if n.right != nil {
			node = n.right.get(val)
		}
	} else {
		node = n
	}
	return node
}

func (n *Node) rotateRight() *Node {
	l := n.left
	// Rotation
	l.right, n.left = n, l.right

	// update heights
	n.height = n.maxHeight() + 1
	l.height = l.maxHeight() + 1

	return l
}

func (n *Node) rotateLeft() *Node {
	r := n.right
	// Rotation
	r.left, n.right = n, r.left

	// update heights
	n.height = n.maxHeight() + 1
	r.height = r.maxHeight() + 1

	return r
}

func (n *Node) iterate(iterator NodeIterator, i *int, cont bool) {
	if n != nil && cont {
		n.left.iterate(iterator, i, cont)
		cont = iterator(n, *i)
		*i++
		n.right.iterate(iterator, i, cont)
	}
}

func (n *Node) rIterate(iterator NodeIterator, i *int, cont bool) {
	if n != nil && cont {
		n.right.iterate(iterator, i, cont)
		cont = iterator(n, *i)
		*i++
		n.left.iterate(iterator, i, cont)
	}
}

func (n *Node) min() *Node {
	current := n
	for current.left != nil {
		current = current.left
	}
	return current
}

func (n *Node) maxHeight() int8 {
	rh := height(n.right)
	lh := height(n.left)
	if rh > lh {
		return rh
	}
	return lh
}

// IntVal represents an integer tree val
type IntVal int

// Comp returns 1 if i > val, -1 if i < val and 0 if i equal to val
func (i IntVal) Comp(val Val) int8 {
	v := val.(IntVal)
	if i > v {
		return 1
	} else if i < v {
		return -1
	} else {
		return 0
	}
}

// StringVal represents an string tree val
type StringVal string

// Comp returns 1 if i > val, -1 if i < val and 0 if i equal to val
func (i StringVal) Comp(val Val) int8 {
	v := val.(StringVal)
	if i > v {
		return 1
	} else if i < v {
		return -1
	} else {
		return 0
	}
}

// UintVal represents an uint tree val
type UintVal uint

// Comp returns 1 if i > val, -1 if i < val and 0 if i equal to val
func (i UintVal) Comp(val Val) int8 {
	v := val.(UintVal)
	if i > v {
		return 1
	} else if i < v {
		return -1
	} else {
		return 0
	}
}

// Float32Val represents an float32 tree val
type Float32Val float32

// Comp returns 1 if i > val, -1 if i < val and 0 if i equal to val
func (i Float32Val) Comp(val Val) int8 {
	v := val.(Float32Val)
	if i > v {
		return 1
	} else if i < v {
		return -1
	} else {
		return 0
	}
}

// Float64Val represents an float64 tree val
type Float64Val float64

// Comp returns 1 if i > val, -1 if i < val and 0 if i equal to val
func (i Float64Val) Comp(val Val) int8 {
	v := val.(Float64Val)
	if i > v {
		return 1
	} else if i < v {
		return -1
	} else {
		return 0
	}
}

// UintptrVal represents a uintptr tree val
type UintptrVal uintptr

// Comp returns 1 if i > val, -1 if i < val and 0 if i equal to val
func (i UintptrVal) Comp(val Val) int8 {
	v := val.(UintptrVal)
	if i > v {
		return 1
	} else if i < v {
		return -1
	} else {
		return 0
	}
}

// RuneVal represents a rune tree val
type RuneVal rune

// Comp returns 1 if i > val, -1 if i < val and 0 if i equal to val
func (i RuneVal) Comp(val Val) int8 {
	v := val.(RuneVal)
	if i > v {
		return 1
	} else if i < v {
		return -1
	} else {
		return 0
	}
}

// ByteVal represents a byte tree val
type ByteVal byte

// Comp returns 1 if i > val, -1 if i < val and 0 if i equal to val
func (i ByteVal) Comp(val Val) int8 {
	v := val.(ByteVal)
	if i > v {
		return 1
	} else if i < v {
		return -1
	} else {
		return 0
	}
}