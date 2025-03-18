package main

import "github.com/charmbracelet/bubbles/help"

// Interface assertion
var _ help.KeyMap = (*keymap)(nil)
