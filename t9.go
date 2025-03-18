package main

import (
	"bufio"
	"fmt"
	"os"
)

const wordsFilePath = "./small-words.txt"

// rune is a type alias for readability, representing a single Unicode character.
type rune byte

// TrieNode represents a node in a trie data structure.
type TrieNode struct {
	children map[rune]*TrieNode
	words    []string
}

// AddWord adds a word to the trie, recursively creating nodes as needed.
func (n *TrieNode) AddWord(word string) {
	if word == "" {
		return
	}

	firstRune := rune(word[0])
	remainder := word[1:]

	if _, ok := n.children[firstRune]; !ok {
		n.children[firstRune] = &TrieNode{children: make(map[rune]*TrieNode)}
	}

	n.children[firstRune].AddWord(remainder)
}

// NewWordTrie creates a new trie from a file containing words.
func NewWordTrie(filePath string) (*TrieNode, error) {
	wordsFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open words file: %w", err)
	}
	defer wordsFile.Close()

	root := &TrieNode{children: make(map[rune]*TrieNode)}

	scanner := bufio.NewScanner(wordsFile)
	for scanner.Scan() {
		word := scanner.Text()
		root.AddWord(word)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read words file: %w", err)
	}

	return root, nil
}

func main() {
	fmt.Printf("loading words from %s\n", wordsFilePath)

	_, err := NewWordTrie(wordsFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("words loaded")
}
