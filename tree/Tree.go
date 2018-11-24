package tree

import (
	"fmt"
	"gekongfei.com/tennessine/gomq/stack"
	"math"
)

type Node struct {
	value       int
	left, right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) PrintPreOrder() {
	printPreOrder(t.root)
	fmt.Println()
}

func printPreOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.value, " ")
	printPreOrder(node.left)
	printPreOrder(node.right)
}

func (t *Tree) PrintPostOrder() {
	printPostOrder(t.root)
	fmt.Println()
}

func printPostOrder(node *Node) {
	if node == nil {
		return
	}
	printPostOrder(node.left)
	printPostOrder(node.right)
	fmt.Print(node.value, " ")
}
func (t *Tree) PrintInOrder() {
	printInOrder(t.root)
	fmt.Println()
}

func printInOrder(node *Node) {
	if node == nil {
		return
	}
	printInOrder(node.left)
	fmt.Print(node.value, " ")
	printInOrder(node.right)
}

func (t *Tree) TreeDepth() int {
	return treeDepth(t.root)
}

func treeDepth(node *Node) int {
	if node == nil {
		return 0
	}
	lDepth := treeDepth(node.left)
	rDepth := treeDepth(node.right)
	if lDepth > rDepth {
		return rDepth + 1
	}
	return rDepth + 1
}

func (t *Tree) NthPreOrder(index int) {
	var counter int
	nthPreOrder(t.root, index, &counter)
}

func nthPreOrder(node *Node, index int, counter *int) {
	if node != nil {
		*counter++
		if *counter == index {
			fmt.Println(node.value)
		}
		nthPreOrder(node.left, index, counter)
		nthPreOrder(node.right, index, counter)
	}
}

func (t *Tree) NthPostOrder(index int) {
	var counter int
	nthPostOrder(t.root, index, &counter)
}

func nthPostOrder(node *Node, index int, counter *int) {
	if node != nil {
		nthPostOrder(node.left, index, counter)
		nthPostOrder(node.right, index, counter)
		*counter++
		if *counter == index {
			fmt.Println(node.value)
		}
	}
}

func (t *Tree) NthInOrder(index int) {
	var counter int
	nthInCounter(t.root, index, &counter)
}

func nthInCounter(node *Node, index int, counter *int) {
	if node != nil {
		nthInCounter(node.left, index, counter)
		*counter++
		if *counter == index {
			fmt.Println(node.value)
		}
		nthInCounter(node.right, index, counter)
	}
}

func (t *Tree) CopyTree() *Tree {
	tree := new(Tree)
	tree.root = copyTree(t.root)
	return tree
}

func copyTree(node *Node) *Node {
	var temp *Node
	if node != nil {
		temp = new(Node)
		temp.value = node.value
		temp.left = copyTree(node.left)
		temp.right = copyTree(node.right)
		return temp
	}
	return nil
}

func (t *Tree) CopyMirrorTree() *Tree {
	tree := new(Tree)
	tree.root = copyMirrorTree(t.root)
	return tree
}

func copyMirrorTree(node *Node) *Node {
	var temp *Node
	if node != nil {
		temp = new(Node)
		temp.value = node.value
		temp.right = copyMirrorTree(node.left)
		temp.left = copyMirrorTree(node.right)
		return temp
	}
	return nil
}

func (t *Tree) NumNodes() int {
	return numNodes(t.root)
}

func numNodes(node *Node) int {
	if node == nil {
		return 0
	}
	return 1 + numNodes(node.left) + numNodes(node.right)
}

func (t *Tree) NumLeafNodes() int {
	return numLeafNodes(t.root)
}

func numLeafNodes(node *Node) int {
	if node == nil {
		return 0
	}
	if node.left == nil && node.right == nil {
		return 1
	}
	return numLeafNodes(node.right) + numLeafNodes(node.left)
}

func (t *Tree) IsEqual(t2 *Tree) bool {
	return isEqual(t.root, t2.root)
}

func isEqual(node1 *Node, node2 *Node) bool {
	if node1 == nil && node2 == nil {
		return true
	} else if node1 == nil || node2 == nil {
		return false
	} else {
		return (node1.value == node2.value) && isEqual(node1.left, node2.left) && isEqual(node1.right, node2.right)
	}
}

func (t *Tree) PrintAllPath() {
	s := new(stack.Stack)
	printAllPath(t.root, s)
}

func printAllPath(node *Node, s *stack.Stack) {
	if node == nil {
		return
	}
	s.Push(node.value)
	if node.left == nil && node.right == nil {
		s.Print()
		s.Pop()
		return
	}

	printAllPath(node.right, s)
	printAllPath(node.left, s)
	s.Pop()
}

func (t *Tree) FindMaxBT() int {
	return findMaxBT(t.root)
}

func findMaxBT(node *Node) int {
	if node == nil {
		return math.MinInt32
	}
	max := node.value
	left := findMaxBT(node.left)
	right := findMaxBT(node.right)

	if left > max {
		max = left
	}

	if right > max {
		max = right
	}
	return max
}

func (t *Tree) SearchBT(value int) bool {
	return searchBT(t.root, value)
}

func searchBT(node *Node, value int) bool {
	var left, right bool
	if node == nil {
		return false
	}
	if node.value == value {
		return true
	}
	left = searchBT(node.left, value)
	if left {
		return true
	}
	right = searchBT(node.right, value)
	if right {
		return true
	}
	return false
}

func (t *Tree) NumFullNodesBT() int {
	return numFullNodesBT(t.root)
}

func numFullNodesBT(node *Node) int {
	var count int
	if node == nil {
		return 0
	}
	count = numFullNodesBT(node.right) + numFullNodesBT(node.left)
	if node.right != nil && node.left != nil {
		count++
	}
	return count
}

func (t *Tree) MaxLengthPathBT() int {
	return maxLengthPathBT(t.root)
}

func maxLengthPathBT(node *Node) int {
	var max, leftPath, rightPath, leftMax, rightMax int
	if node == nil {
		return 0
	}
	leftPath = treeDepth(node.left)
	rightPath = treeDepth(node.right)
	max = leftPath + rightPath + 1
	leftMax = maxLengthPathBT(node.left)
	rightMax = maxLengthPathBT(node.right)
	if leftMax > max {
		max = leftMax
	}
	if rightMax > max {
		max = rightMax
	}
	return max
}

func (t *Tree) SumAllBT() int {
	return sumAllBT(t.root)
}

func sumAllBT(node *Node) int {
	var sum, leftSum, rightSum int
	if node == nil {
		return 0
	}
	rightSum = sumAllBT(node.right)
	leftSum = sumAllBT(node.left)
	sum = rightSum + leftSum + node.value
	return sum
}

func CreateBinaryTree(arr []int) *Tree {
	t := new(Tree)
	size := len(arr)
	t.root = createBinaryTreeUtil(arr, 0, size-1)
	return t
}

func createBinaryTreeUtil(arr []int, start int, end int) *Node {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	node := new(Node)
	node.value = arr[mid]
	node.left = createBinaryTreeUtil(arr, start, mid-1)
	node.right = createBinaryTreeUtil(arr, mid+1, end)
	return node
}
