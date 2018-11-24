package hashtable

import "fmt"

type Node struct {
	value int
	next  *Node
}

type HashTable struct {
	listArray []*Node
	tableSize int
}

func (ht *HashTable) Init() {
	ht.tableSize = 101
	ht.listArray = make([]*Node, ht.tableSize)

	for i := 0; i < ht.tableSize; i++ {
		ht.listArray[i] = nil
	}
}

func (ht *HashTable) ComputeHash(key int) int {
	return key % ht.tableSize
}

func (ht *HashTable) Add(value int) {
	hash := ht.ComputeHash(value)
	temp := new(Node)
	temp.value = value
	temp.next = ht.listArray[hash]
	ht.listArray[hash] = temp
}

func (ht *HashTable) Remove(value int) bool {
	hash := ht.ComputeHash(value)
	var nextNode, head *Node
	head = ht.listArray[hash]
	if head != nil && head.value == value {
		ht.listArray[hash] = head.next
		return true
	}
	for head != nil {
		nextNode = head.next
		if nextNode != nil && nextNode.value == value {
			head.next = nextNode.next
			return true
		}
		head = nextNode
	}
	return false
}

func (ht *HashTable) Print() {
	for i := 0; i < ht.tableSize; i++ {
		head := ht.listArray[i]
		if head != nil {
			fmt.Printf("\nValues at index :: %d are ::", i)
		}
		for head != nil {
			fmt.Print(head.value, " ")
			head = head.next
		}
	}
	fmt.Println()
}

func (ht *HashTable) Find(value int) bool {
	hash := ht.ComputeHash(value)
	head := ht.listArray[hash]
	for head != nil {
		if head.value == value {
			return true
		}
		head = head.next
	}
	return false
}
