package ui

import (
	"TASKCLI/storage"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// Model holds app state
type Model struct {
	Tasks []storage.Task // list of tasks from Json file
	Cursor int 			 // For hovering selected task
}

//Init runs when program starts
//Returning nil coz no initial background I/O
func (m Model) Init() tea.Cmd{
	return nil
}

// Update handles all events, like key presses.
func (m Model) Update(msg tea.Msg) ( tea.Model,tea.Cmd) {
	switch msg := msg.(type){
//  if key pressed
	case tea.KeyMsg : 
		switch msg.String(){
		case "ctrl+c","q" : // key to quit program
			return m,tea.Quit
		}
	}
	return m,nil // return the updated model
}

//View Build the actual Str to terminal
func (m Model) View() string{
	return fmt.Sprintf(
		"GO-TASK!\n There are %d tasks.\nClick 'q' or 'Ctrl+C' to quit.\n",
		len(m.Tasks),
	)
}