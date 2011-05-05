package prefixtree

// PrefixTree is a datastructure which is used to
// store strings efficiently and is searchable by
// prefix.
type PrefixTree struct {
	root *node
	size int
}

// node is a node in the prefix tree.
type node struct {
	valid bool
	kids  map[int]*node
}

// Create a new PrefixTree.
func New() *PrefixTree {
	return &PrefixTree{
		root: newNode(),
	}
}

// newNode creates a new prefix tree node.
func newNode() *node {
	return &node{
		kids: make(map[int]*node),
	}
}

// Return the number of words contained in this PrefixTree.
func (this *PrefixTree) Size() int {
	return this.size
}

// Returns true if and only if the tree contains the given
// word.
func (this *PrefixTree) Contains(word string) bool {
	n, ok := this.getPrefix(word)
	return ok && n.valid
}

// Returns true if and only if the PrefixTree contains the
// given prefix.
func (this *PrefixTree) ContainsPrefix(prefix string) bool {
	_, ok := this.getPrefix(prefix)
	return ok
}

func (this *PrefixTree) getPrefix(prefix string) (n *node, ok bool) {
	ptr := this.root
	for _, symbol := range prefix {
		ptr, ok = ptr.kids[symbol]
		if !ok {
			break
		}
	}
	return ptr, ok
}

// getChild will return the child of the given node for the
// given symbol. If no such child exists, it is created.
func getChild(n *node, symbol int) *node {
	ptr, ok := n.kids[symbol]
	if !ok {
		ptr = newNode()
		n.kids[symbol] = ptr
	}
	return ptr
}

// Put inserts a word into the PrefixTree.
func (this *PrefixTree) Put(word string) {
	if len(word) < 1 {
		return
	}
	ptr := this.root
	for _, symbol := range word {
		ptr = getChild(ptr, symbol)
	}
	ptr.valid = true
	this.size++
}
