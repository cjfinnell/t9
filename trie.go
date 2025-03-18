package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

const wordsFilePath = "./small-words2.txt"

func loadWords() (*TrieNode, error) {
	wordsFile, err := os.Open(wordsFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open words file: %w", err)
	}
	defer wordsFile.Close()

	node, err := NewWordTrie(wordsFile)
	if err != nil {
		return nil, fmt.Errorf("failed to construct word trie: %w", err)
	}

	return node, nil
}

// TrieNode represents a node in a trie data structure.
type TrieNode struct {
	children map[int]*TrieNode
	words    []string
}

// AddWord adds a word to the trie, recursively creating nodes as needed.
func (n *TrieNode) AddWord(word, wordPart string) {
	if wordPart == "" {
		n.words = append(n.words, word)

		return
	}

	firstRune := rune(wordPart[0])
	runeNum := runeToNum[firstRune]
	remainder := wordPart[1:]

	if _, ok := n.children[runeNum]; !ok {
		n.children[runeNum] = &TrieNode{children: make(map[int]*TrieNode)}
	}

	n.children[runeNum].AddWord(word, remainder)
}

// NextSteps returns the available valid T9 inputs from the current node.
func (n *TrieNode) NextSteps() []int {
	var steps []int
	for step := range n.children {
		steps = append(steps, step)
	}

	sortedSteps := sort.IntSlice(steps)
	sortedSteps.Sort()

	return sortedSteps
}

// Walk traverses the trie one node at a time, returning the next node and the
// words associated with it.
// If the input is not a valid T9 input, the current node and an error is
// returned.
func (n *TrieNode) Walk(t9Input int) (*TrieNode, []string, error) {
	// If invalid input, return current node and an error
	child, ok := n.children[t9Input]
	if !ok {
		return n, nil, ErrInvalidInput
	}

	return child, child.words, nil
}

// NewWordTrie creates a new trie from a file containing words.
func NewWordTrie(wordsSource io.Reader) (*TrieNode, error) {
	root := &TrieNode{children: make(map[int]*TrieNode)}

	count := 0

	scanner := bufio.NewScanner(wordsSource)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word == "" {
			continue
		}

		root.AddWord(word, word)
		count++
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read words file: %w", err)
	}

	return root, nil
}
