package style

import (
	"github.com/charmbracelet/lipgloss"
)

func DocStyle() lipgloss.Style {

	return lipgloss.NewStyle().
		Padding(1, 2, 1, 2)
}
