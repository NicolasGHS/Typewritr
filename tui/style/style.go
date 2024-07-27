package style

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	HeaderStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Width(18)

	CurrentLetterStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#e78284"))

	AfterCursorStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#737994"))
)
