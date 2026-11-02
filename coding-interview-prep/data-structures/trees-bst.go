package main

import "fmt"

// TreeNode represents a node in the Binary Search Tree
type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

// BST represents the Binary Search Tree
type BST struct {
	root *TreeNode
}

// Insert adds a new value to the BST
func (bst *BST) Insert(value int) {
	newNode := &TreeNode{value: value}
	if bst.root == nil {
		bst.root = newNode
		return
	}

	current := bst.root
	for {
		if value < current.value {
			if current.left == nil {
				current.left = newNode
				return
			}
			current = current.left
		} else if value > current.value {
			if current.right == nil {
				current.right = newNode
				return
			}
			current = current.right
		} else {
			// Value already exists, do nothing or update
			return
		}
	}
}

// Search finds a value in the BST
func (bst *BST) Search(value int) bool {
	current := bst.root
	for current != nil {
		if value == current.value {
			return true
		} else if value < current.value {
			current = current.left
		} else {
			current = current.right
		}
	}
	return false
}

// InOrderTraversal performs an in-order traversal of the BST
func (bst *BST) InOrderTraversal() {
	fmt.Print("In-order Traversal: ")
	inOrder(bst.root)
	fmt.Println()
}

func inOrder(node *TreeNode) {
	if node != nil {
		inOrder(node.left)
		fmt.Printf("%d ", node.value)
		inOrder(node.right)
	}
}

// PreOrderTraversal performs a pre-order traversal of the BST
func (bst *BST) PreOrderTraversal() {
	fmt.Print("Pre-order Traversal: ")
	preOrder(bst.root)
	fmt.Println()
}

func preOrder(node *TreeNode) {
	if node != nil {
		fmt.Printf("%d ", node.value)
		preOrder(node.left)
		preOrder(node.right)
	}
}

// PostOrderTraversal performs a post-order traversal of the BST
func (bst *BST) PostOrderTraversal() {
	fmt.Print("Post-order Traversal: ")
	postOrder(bst.root)
	fmt.Println()
}

func postOrder(node *TreeNode) {
	if node != nil {
		postOrder(node.left)
		postOrder(node.right)
		fmt.Printf("%d ", node.value)
	}
}

// Delete removes a value from the BST
func (bst *BST) Delete(value int) {
	bst.root = deleteNode(bst.root, value)
}

func deleteNode(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return nil
	}

	if value < node.value {
		node.left = deleteNode(node.left, value)
	} else if value > node.value {
		node.right = deleteNode(node.right, value)
	} else {
		// Node with only one child or no child
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}

		// Node with two children: Get the in-order successor (smallest in the right subtree)
		node.value = minValue(node.right)
		// Delete the in-order successor
		node.right = deleteNode(node.right, node.value)
	}
	return node
}

func minValue(node *TreeNode) int {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current.value
}

func main() {
	bst := BST{}
	fmt.Println("Inserting values: 50, 30, 70, 20, 40, 60, 80")
	bst.Insert(50)
	bst.Insert(30)
	bst.Insert(70)
	bst.Insert(20)
	bst.Insert(40)
	bst.Insert(60)
	bst.Insert(80)

	bst.InOrderTraversal()  // Expected: 20 30 40 50 60 70 80
	bst.PreOrderTraversal() // Expected: 50 30 20 40 70 60 80
	bst.PostOrderTraversal() // Expected: 20 40 30 60 80 70 50

	fmt.Println("Searching for 40:", bst.Search(40))   // Expected: true
	fmt.Println("Searching for 90:", bst.Search(90))   // Expected: false

	fmt.Println("Deleting 20 (leaf node)")
	bst.Delete(20)
	bst.InOrderTraversal() // Expected: 30 40 50 60 70 80

	fmt.Println("Deleting 70 (node with one child)")
	bst.Delete(70)
	bst.InOrderTraversal() // Expected: 30 40 50 60 80

	fmt.Println("Deleting 50 (root node with two children)")
	bst.Delete(50)
	bst.InOrderTraversal() // Expected: 30 40 60 80 (or 40 60 80 if 40 was chosen as successor)
}
