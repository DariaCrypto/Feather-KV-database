package inmemory

import (
	"errors"
)

type AVLNode struct {
	Value string
	Score uint32
	Depth int
	Left  *AVLNode
	Right *AVLNode
}

func newNode(key uint32, value string) *AVLNode {
	node := &AVLNode{Score: key, Value: value}
	node.Left = nil
	node.Right = nil
	node.Depth = 1
	return node
}

func getBalance(node *AVLNode) int {
	return getDepth(node.Left) - getDepth(node.Right)
}

func rightRotate(selectedNode *AVLNode) *AVLNode {
	LeftNode := selectedNode.Left
	RightNode := LeftNode.Right

	LeftNode.Right = selectedNode
	selectedNode.Left = RightNode

	selectedNode.Depth = max(getDepth(selectedNode.Left), getDepth(selectedNode.Right)) + 1
	LeftNode.Depth = max(getDepth(LeftNode.Left), getDepth(LeftNode.Right)) + 1

	return LeftNode
}

func leftRotate(selectedNode *AVLNode) *AVLNode {
	RightNode := selectedNode.Right
	LeftNode := RightNode.Left

	RightNode.Left = selectedNode
	selectedNode.Right = LeftNode

	selectedNode.Depth = max(getDepth(selectedNode.Left), getDepth(selectedNode.Right)) + 1
	RightNode.Depth = max(getDepth(LeftNode), getDepth(selectedNode)) + 1

	return RightNode
}

func insert(root *AVLNode, Score uint32, Value string) *AVLNode {
	if root == nil {
		return newNode(Score, Value)
	}
	if Score == root.Score && Value == root.Value {
		return root
	}

	if Score < root.Score {
		root.Left = insert(root.Left, Score, Value)
	} else if Score > root.Score {
		root.Right = insert(root.Right, Score, Value)
	} else {
		return root
	}

	balance := getBalance(root)
	root.Depth = 1 + max(getDepth(root.Left), getDepth(root.Right))

	if balance > 1 {
		if Score < root.Left.Score {
			return rightRotate(root)
		} else if Score > root.Left.Score {
			root.Left = leftRotate(root.Left)
			return rightRotate(root)
		}
	}

	if balance < -1 {
		if Score > root.Right.Score {
			return leftRotate(root)
		} else if Score < root.Right.Score {
			root.Right = rightRotate(root.Right)
			return leftRotate(root)
		}
	}

	return root
}

func nodeWithMinimumValue(node *AVLNode) *AVLNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func getNode(root *AVLNode, Value string) (*AVLNode, error) {
	if root == nil {
		return nil, nil
	}

	if Value == root.Value {
		return root, nil
	}

	rightNode, _ := getNode(root.Right, Value)
	if rightNode != nil {
		return rightNode, nil
	}

	leftNode, _ := getNode(root.Left, Value)
	if leftNode != nil {
		return leftNode, nil
	}

	return nil, errors.New("avl-tree: node not found in tree.")
}

func deleteNode(root *AVLNode, key uint32) *AVLNode {
	if root == nil {
		return root
	}
	if key < root.Score {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Score {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		temp := nodeWithMinimumValue(root.Right)
		root.Score = temp.Score
		root.Right = deleteNode(root.Right, temp.Score)
	}

	root.Depth = 1 + max(getDepth(root.Left), getDepth(root.Right))
	balanceFactor := getBalance(root)

	if balanceFactor > 1 {
		if getBalance(root.Left) >= 0 {
			return rightRotate(root)
		} else {
			root.Left = leftRotate(root.Left)
			return rightRotate(root)
		}
	}
	if balanceFactor < -1 {
		if getBalance(root.Right) <= 0 {
			return leftRotate(root)
		} else {
			root.Right = rightRotate(root.Right)
			return leftRotate(root)
		}
	}
	return root
}

func findMax(root *AVLNode) *AVLNode {
	if root == nil {
		return nil
	}

	res := root
	lres := findMax(root.Left)
	rres := findMax(root.Right)

	if lres != nil && lres.Score > res.Score {
		res = lres
	}
	if rres != nil && rres.Score > res.Score {
		res = rres
	}

	return res
}

func findMin(root *AVLNode) *AVLNode {
	if root == nil {
		return nil
	}

	res := root
	lres := findMin(root.Left)
	rres := findMin(root.Right)

	if lres != nil && lres.Score < res.Score {
		res = lres
	}
	if rres != nil && rres.Score < res.Score {
		res = rres
	}

	return res
}

func getDepth(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return node.Depth
}
