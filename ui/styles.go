package ui

import "github.com/charmbracelet/lipgloss"

var(
	HeaderStyle = lipgloss.NewStyle().
                    Bold(true).
                    Foreground(lipgloss.Color("#ffed9e")).
					Blink(true).
                    Padding(1, 2).
                    Width(50).
                    MarginBottom(1)

	SeletedStyle = lipgloss.NewStyle().
					Foreground(lipgloss.Color("#c025fd")).
					Bold(true).
					BorderLeft(true).
					BorderStyle(lipgloss.ThickBorder()).
					BorderForeground(lipgloss.Color("#c025fd")).
					PaddingLeft(2)

	NormalStyle = lipgloss.NewStyle().
					Foreground(lipgloss.Color("#EEEEEE")).
					PaddingLeft(2)

	CompletedSyle = lipgloss.NewStyle().
						Strikethrough(true).
						Foreground(lipgloss.Color("#555555")).
						PaddingLeft(2)

	HelpStyle = lipgloss.NewStyle().
					Foreground(lipgloss.Color("#626262")).
					MarginTop(1).
					Border(lipgloss.NormalBorder(),true,false,false,false).
					BorderForeground(lipgloss.Color("#333333"))

	InputBoxStyle = lipgloss.NewStyle().
                    Border(lipgloss.RoundedBorder()).
                    BorderForeground(lipgloss.Color("#ffed9e")).
                    Padding(0, 1).
                    Width(52)

    StatusBarStyle = lipgloss.NewStyle().
                    Foreground(lipgloss.Color("#ffed9e")).
                    Bold(true).
					PaddingLeft(2).
                    MarginBottom(1)

	ComplimentStyle = lipgloss.NewStyle().
					Foreground(lipgloss.Color("FFD700")).
					Bold(true).
					MarginBottom(1)
)