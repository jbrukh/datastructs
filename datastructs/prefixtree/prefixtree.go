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
	kids  map[string]*node
}

// Create a new PrefixTree.
func New() *PrefixTree {
	return &PrefixTree{
		root: new(node),
	}
}

// Return the number of words contained in this
// prefix tree.
func (this *PrefixTree) Size() int {
	return this.size
}

// Returns true if and only 
func (this *PrefixTree) Contains(word string) bool {
	if word == nil || word == "" {

	}
	ptr := this.root
	for _, letter := range word {

	}
}

func (this *PrefixTree) Put(word string) {
	ptr := this.root
	for _, letter := range word {
		if n, ok := ptr.kids[letter]; ok {
			// 
		}

	}
}
