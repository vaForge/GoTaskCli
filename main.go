package main

import (
	"TASKCLI/storage"
	"TASKCLI/ui"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	tasks,err := storage.LoadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n",err)
		os.Exit(1)
	}

	initialModel := ui.NewModel(tasks)
	

	p := tea.NewProgram(initialModel)
	if _,err := p.Run(); err!= nil {
		fmt.Printf("Error :%v\n",err)
		os.Exit(1)
	}
}