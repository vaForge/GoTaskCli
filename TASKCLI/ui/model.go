package ui

import (
	"TASKCLI/storage"
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

//App state defines what view the user is currently looking at
type AppState int

const (
	Statelist AppState = iota
	StateNew
	StateEdit
)
// Model holds app state
type Model struct {
	Tasks []storage.Task // list of tasks from Json file
	Cursor int 			 // For hovering selected task
	State AppState
	Input textinput.Model
}

func NewModel(tasks []storage.Task) Model{
	ti := textinput.New()
	ti.Placeholder = "Enter New Task : "
	ti.CharLimit = 100
	ti.Width = 50

	return Model{
			Tasks : tasks,
			Cursor : 0,
			State : Statelist,
			Input : ti,
	}
}

//Init runs when program starts
//Returning nil coz no initial background I/O
func (m Model) Init() tea.Cmd{
	return textinput.Blink
}

// Update handles all events, like key presses.
func (m Model) Update(msg tea.Msg) ( tea.Model,tea.Cmd) {
	var cmd tea.Cmd

	// ---- IF WE ARE TYPING A TASK -----------
	if m.State == StateNew || m.State == StateEdit{
		switch msg := msg.(type){
		case tea.KeyMsg : 
			switch msg.String(){
			case "enter" :  //Saving a task
				val := m.Input.Value()
				if val!= "" {
					if m.State == StateNew{ 
						// Create a simple ID and append
						newID := 1
						if len(m.Tasks) > 0 {
							newID = m.Tasks[len(m.Tasks)-1].ID + 1 //assign last ID
						}
						m.Tasks = append(m.Tasks, storage.Task{ID:newID,Title: val,Completed:false})
					}else if m.State == StateEdit{
						m.Tasks[m.Cursor].Title = val
					}
					storage.SaveTasks (m.Tasks)
				}
				m.Input.Reset()
				m.State = Statelist
				return m,nil

			case "esc" : 
				m.Input.Reset()
				m.State = Statelist
				return m,nil
			}
		}

		// Route all other keys to text input component
		m.Input,cmd = m.Input.Update(msg)
		return m,cmd
	}

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

		//Trigger Input States 
		case "n","N" :
			m.State = StateNew
			m.Input.Focus()
			return m,textinput.Blink

		case "e","E" :
			if len(m.Tasks) > 0 {
				m.State = StateEdit
				m.Input.SetValue(m.Tasks[m.Cursor].Title)
				m.Input.Focus()
				return m,textinput.Blink
			}

		}
		
	}
	return m,nil // return the updated model
}

//View Build the actual Str to terminal
func (m Model) View() string{
	//If we are typing, only show the input box
	if m.State == StateNew || m.State == StateEdit{
		headerMsg := "Add New Task"
		if m.State == StateEdit{
			headerMsg = "Edit Task"
		}
		return fmt.Sprintf(
			"%s\n\n%s\n\n%s",
			HeaderStyle.Render(headerMsg),
			m.Input.View(),
			HelpStyle.Render("(esc to cancel • enter to save)"),
		)
	}

	s := HeaderStyle.Render("GoTaskCLI - Get Things Done") + "\n\n"

	allCompleted := len(m.Tasks) > 0 
	for _,t := range(m.Tasks){
		if !t.Completed {
			allCompleted = false
		}
	}
	//Render Task List
	if allCompleted{
		complimentStyle := lipgloss.NewStyle().
								Foreground(lipgloss.Color("FFD700")).
								Bold(true).MarginBottom(1)
		s += complimentStyle.Render("GREAT WORK! All tasks are clear") + "\n"
	}
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

	helpMsg := "↑/↓: Navigate • N: New • E: Edit • M: Mark/Unmark \n• D: Delete • C: Clear • Q: Quit"
	s += HelpStyle.Render(helpMsg) + "\n"

	return s
}