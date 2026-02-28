package ui

import "github.com/charmbracelet/lipgloss"

var (

	HeaderStyle = lipgloss.NewStyle().
					Bold(true).
					Foreground(lipgloss.Color("#FAFAFA")).
					Background(lipgloss.Color("#7D56F4")).
					Padding(0,1).
					MarginBottom(1)

	SeletedStyle = lipgloss.NewStyle().
					Foreground(lipgloss.Color("#01FAC6")).
					Bold(true)

	NormalStyle = lipgloss.NewStyle().
					Foreground(lipgloss.Color("#EEEEEE"))

	CompletedSyle = lipgloss.NewStyle().
						Strikethrough(true).
						Foreground(lipgloss.Color("#888888"))

	HelpStyle = lipgloss.NewStyle().
					Foreground(lipgloss.Color("#626262")).
					MarginTop(1)
)