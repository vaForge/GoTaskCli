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

		case "up","k" :
			if m.Cursor > 0 {
				m.Cursor--
			}

		case "down","j":
			if m.Cursor < len(m.Tasks) -1 {
				m.Cursor++
			}

		case "m", "M","enter" :
			if len(m.Tasks) > 0 {
				m.Tasks[m.Cursor].Completed = !m.Tasks[m.Cursor].Completed
				storage.SaveTasks(m.Tasks)
			}

		
		case "d","D" : 
			if len(m.Tasks)>0{
				m.Tasks = append(m.Tasks[:m.Cursor],m.Tasks[m.Cursor+1:]...)

				if m.Cursor >= len(m.Tasks)&& m.Cursor > 0 {
					m.Cursor--
				}
				storage.SaveTasks(m.Tasks)
			}


		case "c","C" : 
			var incompleteTasks []storage.Task
			for _,t := range m.Tasks{
				if !t.Completed {
					incompleteTasks = append(incompleteTasks, t)
				}
			}
			m.Tasks = incompleteTasks
			m.Cursor = 0
			storage.SaveTasks(m.Tasks)
		}
		
	}
	return m,nil // return the updated model
}

//View Build the actual Str to terminal
func (m Model) View() string{
	s := HeaderStyle.Render("GoTaskCLI - Get Things Done") + "\n"
	//Render Task List
	if len(m.Tasks) == 0 {
		s += "No Tasks yet, Wanna get some finised? Press 'N' to add \n"
	}else{
		for i,task := range m.Tasks{
			cursor := " "
			if m.Cursor == i { 
				cursor = "> "
			}
			checked := " "
			if task.Completed {
				checked = "x"
			}
			taskStr := fmt.Sprintf("%s[%s]%s",cursor,checked,task.Title)

			if task.Completed {
				s += CompletedSyle.Render(taskStr) + "\n"
			}else if m.Cursor == i{
				s += SeletedStyle.Render(taskStr) + "\n"
			} else{
				s += NormalStyle.Render(taskStr) + "\n"
			}
		}
	}

	helpMsg := "↑/↓: Navigate • N: New • E: Edit • M: Mark/Unmark • D: Delete • C: Clear • Q: Quit"
	s += HelpStyle.Render(helpMsg) + "\n"

	return s
}