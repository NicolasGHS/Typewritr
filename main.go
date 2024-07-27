package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/NicolasGHS/Typewritr/tui/model"
)

func main() {
	p := tea.NewProgram(model.InitialModel())
	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting program: %v", err)
		os.Exit(1)
	}
}