package main

import (
	"fmt"
	"os"


	tea "github.com/charmbracelet/bubbletea"
	"https://github.com/NicolasGHS/Typewritr/tui/model"

)


type model struct {
    words    []string           
    cursor   int                
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
	return model {
		words: []string{"Test", "Hello", "World"},
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
			}
	}
	return m, nil
}

func (m model) View() string {
	s := "Typewritr\n"

	for _, word := range m.words {
		s += fmt.Sprintf("%s ", word)
	}

	s += "\nPress q to quit.\n"

	return s
}

