package main

import tea "github.com/charmbracelet/bubbletea"

// Interface assertion
var _ tea.Model = (*model)(nil)
