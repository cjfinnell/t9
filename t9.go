package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const wordsFilePath = "./small-words.txt"

var runeToNum = map[rune]int{
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

// rune is a type alias for readability, representing a single Unicode character.
type rune byte

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

// NewWordTrie creates a new trie from a file containing words.
func NewWordTrie(wordsSource io.Reader) (*TrieNode, error) {
	root := &TrieNode{children: make(map[int]*TrieNode)}

	scanner := bufio.NewScanner(wordsSource)
	for scanner.Scan() {
		word := scanner.Text()
		root.AddWord(word, "")
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read words file: %w", err)
	}

	return root, nil
}

func main() {
	fmt.Printf("loading words from %s\n", wordsFilePath)
	wordsFile, err := os.Open(wordsFilePath)
	if err != nil {
		panic(fmt.Errorf("failed to open words file: %w", err))
	}
	defer wordsFile.Close()

	_, err = NewWordTrie(wordsFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("words loaded")
}
