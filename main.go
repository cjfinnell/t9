package main

import (
	"fmt"
	"log/slog"
	"os"
)

func main() {
	slog.Debug(fmt.Sprintf("loading words from %s", wordsFilePath))

	wordsFile, err := os.Open(wordsFilePath)
	if err != nil {
		slog.With("error", err).Error("failed to open words file")

		os.Exit(1)
	}
	defer wordsFile.Close()

	_, err = NewWordTrie(wordsFile)
	if err != nil {
		slog.With("error", err).Error("failed to construct word trie")

		os.Exit(1)
	}

	slog.Debug("words loaded")
}
