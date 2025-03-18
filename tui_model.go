package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	textInput textinput.Model
	help      help.Model
	keymap    keymap

	node *TrieNode
}

func initialModel() model {
	node, err := loadWords()
	if err != nil {
		log.Fatal("failed to load words: ", err)
	}

	ti := textinput.New()
	ti.Prompt = "input: "
	ti.Placeholder = "repository"
	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	ti.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 20
	ti.ShowSuggestions = true

	h := help.New()

	km := keymap{}

	return model{textInput: ti, help: h, keymap: km, node: node}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

		if keyIsValidT9(msg.String()) {
			m.textInput, cmd = m.textInput.Update(msg)
			return m, cmd
		}
	}

	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"T9 Text Input Emulator:\n\n  %s\n\n%s\n\n",
		m.textInput.View(),
		m.help.View(m.keymap),
	)
}
