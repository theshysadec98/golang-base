package main

import "fmt"

type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

type SuggestComponent struct {
	trie *TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
	}
}

func (node *TrieNode) AddWord(word string) {
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			node.children[char] = NewTrieNode()
		}
		node = node.children[char]
	}
	node.isWord = true
}

func (node *TrieNode) Search(prefix string) bool {
	for _, char := range prefix {
		if _, ok := node.children[char]; !ok {
			return false
		}
		node = node.children[char]
	}
	return true
}

func (node *TrieNode) Suggest(prefix string) []string {
	suggestions := []string{}
	if node.isWord {
		suggestions = append(suggestions, prefix)
	}
	for key, child := range node.children {
		suggestions = append(suggestions, child.Suggest(prefix+string(key))...)
	}
	return suggestions
}

func NewSuggestComponent() *SuggestComponent {
	return &SuggestComponent{trie: NewTrieNode()}
}

func (component *SuggestComponent) AddWord(word string) {
	component.trie.AddWord(word)
}

func (component *SuggestComponent) Suggest(prefix string) []string {
	return component.trie.Suggest(prefix)
}

func main() {
	component := NewSuggestComponent()

	component.AddWord("Golang")
	component.AddWord("Go")
	component.AddWord("Google")
	component.AddWord("Gorilla")

	suggestions := component.Suggest("")

	for _, suggestion := range suggestions {
		fmt.Println(suggestion)
	}
}
