package main

import (
	"fmt"
	"os"

	// "https://github.com/NicolasGHS/Typewritr/tui/model"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	words  []string
	cursor int
	typed  string
}

// TODO: make model + view folder and divide functions

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel() model {
	return model{
		words:  []string{"Test", "Hello", "World"},
		cursor: 0,
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		default:
			// typed key to the typed string
			if m.cursor < len(m.words[0]) {
				currentWord := m.words[0]
				if msg.String() == string(currentWord[m.cursor]) {
					m.typed += msg.String()
					m.cursor++
				}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Width(18)

	currentLetterStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#e78284")) 

	header := headerStyle.Render("Typewritr\n")
	s := header + "\n"


	if m.cursor < len(m.words[0]) {
		currentWord := m.words[0]
		beforeCursor := currentWord[:m.cursor]
		currentLetter := currentWord[m.cursor : m.cursor+1]
		afterCursor := currentWord[m.cursor+1:]

		s += fmt.Sprintf("%s%s%s", beforeCursor, currentLetterStyle.Render(currentLetter), afterCursor)
	} else {
		s += m.words[0]
	}

	s += "\nPress q to quit.\n"
	s += fmt.Sprintf("\nYou typed: %s\n", m.typed)

	if m.cursor < len(m.words[0]) {
		currentWord := m.words[0]
		s += fmt.Sprintf("\nNext letter to type: %c\n", currentWord[m.cursor])
	} else {
		s += "\nYou've typed the word correctly!\n"
	}


	return s
}
