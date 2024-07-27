package model

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbletea"
	"github.com/NicolasGHS/Typewritr/tui/style"
)

type model struct {
	words     []string
	cursor    int
	typed     string
	wordIndex int
	timer time.Duration
	
}

func InitialModel() model {
	return model{
		words:  []string{"Test", "Hello", "World"},
		cursor: 0,
		timer: 10 * time.Second,
	}
}

func tick() tea.Msg {
	time.Sleep(1 * time.Second)
	return tickMsg{}
}

type tickMsg struct{}

func (m model) Init() tea.Cmd {
	return tick 
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
			currentWord := m.words[m.wordIndex]
			if m.cursor < len(currentWord) && msg.String() == string(currentWord[m.cursor]) {
				m.typed += msg.String()
				m.cursor++
				if m.cursor == len(currentWord) {
					m.cursor = 0
					m.wordIndex++
					if m.wordIndex >= len(m.words) {
						m.words = []string{}
						m.typed = ""
					}	
				}
			}
		}
	case tickMsg:
		m.timer -= 1 * time.Second
		if m.timer <= 0 {
			return m, tea.Quit
		}
		return m, tick
	}
	return m, nil
}

func (m model) View() string {
	headerStyle := style.HeaderStyle
	currentLetterStyle := style.CurrentLetterStyle
	afterCursorStyle := style.AfterCursorStyle	

	header := headerStyle.Render("Typewritr\n")
	s := header + "\n"

	s += fmt.Sprintf("Timer: %v\n", m.timer)


	if len(m.words) == 0 {
		s += "\nCongratulations! You've typed all the words correctly!\n"
	} else {
		for i, word := range m.words {
			if i < m.wordIndex {
				s += word + " "
			} else if i == m.wordIndex {
				beforeCursor := word[:m.cursor]
				currentLetter := word[m.cursor : m.cursor+1]
				afterCursor := word[m.cursor+1:]
				s += fmt.Sprintf("%s%s%s ", beforeCursor, currentLetterStyle.Render(currentLetter), afterCursorStyle.Render(afterCursor))
			} else {
				s += afterCursorStyle.Render(word) + " "
			}
		}
	}

	
	s += "\nPress q to quit.\n"
	

	return s
}

