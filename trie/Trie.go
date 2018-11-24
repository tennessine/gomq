package trie

import (
	"strings"
)

type Node struct {
	isLastChar bool
	children   [26]*Node
}

type Trie struct {
	root *Node
}

func (t *Trie) Insert(s string) {
	if s == "" {
		return
	}
	str := strings.ToLower(s)
	t.root = t.InsertUtil(t.root, str, 0)
}

func (t *Trie) InsertUtil(curr *Node, str string, index int) *Node {
	if curr == nil {
		curr = new(Node)
	}
	if len(str) == index {
		curr.isLastChar = true
	} else {
		curr.children[str[index]-'a'] = t.InsertUtil(curr.children[str[index]-'a'], str, index+1)
	}
	return curr
}

func (t *Trie) Remove(s string) {
	if s == "" {
		return
	}
	str := strings.ToLower(s)
	t.RemoveUtil(t.root, str, 0)
}

func (t *Trie) RemoveUtil(curr *Node, str string, index int) {
	if curr == nil {
		return
	}
	if len(str) == index {
		if curr.isLastChar {
			curr.isLastChar = false
		}
		return
	}
	t.RemoveUtil(curr.children[str[index]-'a'], str, index+1)
}

func (t *Trie) Find(s string) bool {
	if s == "" {
		return false
	}
	str := strings.ToLower(s)
	return t.FindUtil(t.root, str, 0)
}

func (t *Trie) FindUtil(curr *Node, str string, index int) bool {
	if curr == nil {
		return false
	}
	if len(str) == index {
		return curr.isLastChar
	}
	return t.FindUtil(curr.children[str[index]-'a'], str, index+1)
}
