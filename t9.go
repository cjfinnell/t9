package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"sort"
	"strings"
)

const wordsFilePath = "./small-words.txt"

var (
	ErrInvalidInput = errors.New("invalid input")

	runeToNum = map[rune]int{
		'a': 2,
		'b': 2,
		'c': 2,
		'd': 3,
		'e': 3,
		'f': 3,
		'g': 4,
		'h': 4,
		'i': 4,
		'j': 5,
		'k': 5,
		'l': 5,
		'm': 6,
		'n': 6,
		'o': 6,
		'p': 7,
		'q': 7,
		'r': 7,
		's': 7,
		't': 8,
		'u': 8,
		'v': 8,
		'w': 9,
		'x': 9,
		'y': 9,
		'z': 9,
	}
)

// rune is a type alias for readability, representing a single Unicode character.
type rune byte

// TrieNode represents a node in a trie data structure.
type TrieNode struct {
	children map[int]*TrieNode
	words    []string
}

// AddWord adds a word to the trie, recursively creating nodes as needed.
func (n *TrieNode) AddWord(word, wordPart string) {
	slog.Debug(fmt.Sprintf("adding word %s\tpart %s", word, wordPart))

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

	slog.Info("loading words into trie")

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

	slog.Info(fmt.Sprintf("loaded %d words into trie", count))

	return root, nil
}
