package inmemory

import (
	"sync"
	"log"
)

type SortedSet struct {
	name string
	tree    *AVLNode
	mu      sync.Mutex
}

func NewSortedSet(name string) *SortedSet {
	return &SortedSet{name: name, tree: nil}
}

func (ss *SortedSet) addNode(value string, score uint32) {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	ss.tree = insert(ss.tree, score, value)
}

func (ss *SortedSet) getNode(value string) (uint32, string) {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	nnode, err := getNode(ss.tree, value)
	if err != nil{
		log.Println("sortedset: failed to take the node.", err)
		return 0, ""
	}
	return nnode.Score, nnode.Value
}

func (ss *SortedSet) deleteNode(value string, score uint32) {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	ss.tree = deleteNode(ss.tree, score)
}

func (ss *SortedSet) getMaxScoreNode() (string, uint32) {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	maxNode := findMax(ss.tree)
	return maxNode.Value, maxNode.Score
}

func (ss *SortedSet) getMinScoreNode() (string, uint32) {
	ss.mu.Lock()
	defer ss.mu.Unlock()
	minNode := findMin(ss.tree)
	return minNode.Value, minNode.Score
}
