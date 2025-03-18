package main

import (
	"log/slog"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTrieNodeWalk(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelWarn)

	testWordsSource := strings.NewReader(`
		bar
		bat
		beam
		bean
		bear
		bird
		cat
		cone
		cow
		dog
		dung
	`)

	root, err := NewWordTrie(testWordsSource)
	require.NoError(t, err)

	type step struct {
		t9    int
		next  []int
		words []string
		err   error
	}

	for _, tc := range []struct {
		name  string
		steps []step
	}{
		{
			name: "happy path",
			steps: []step{
				{t9: 2, next: []int{2, 3, 4, 6}},
				{t9: 2, next: []int{7, 8}},
				{t9: 7, words: []string{"bar"}},
			},
		},
		{
			name: "invalid inputs",
			steps: []step{
				{t9: 2, next: []int{2, 3, 4, 6}},
				{t9: 5, next: []int{2, 3, 4, 6}, err: ErrInvalidInput},
				{t9: 2, next: []int{7, 8}},
				{t9: 5, next: []int{7, 8}, err: ErrInvalidInput},
				{t9: 7, words: []string{"bar"}},
			},
		},
		{
			name: "multiple valid end words",
			steps: []step{
				{t9: 2, next: []int{2, 3, 4, 6}},
				{t9: 3, next: []int{2}},
				{t9: 2, next: []int{6, 7}},
				{t9: 6, words: []string{"beam", "bean"}},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var currentNode, nextNode *TrieNode
			var words []string
			var err error

			currentNode = root

			for i, step := range tc.steps {
				t.Logf("taking step %d", i)

				nextNode, words, err = currentNode.Walk(step.t9)

				assert.Equal(t, step.next, nextNode.NextSteps())
				assert.Equal(t, sort.StringSlice(step.words), sort.StringSlice(words))
				assert.Equal(t, step.err, err)

				currentNode = nextNode
			}
		})
	}
}
