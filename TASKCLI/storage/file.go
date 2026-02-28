package storage

import (
	"encoding/json" //why ? for Marshal() [Go Slice(struct)->Json] & Unmarshal(json->Go Slice)
	"errors"
	"os"
)

const fileName = "tasks.json"
// This function reads JSON file and returns a slice 
func LoadTasks() ([]Task , error){
	file,err := os.ReadFile(fileName)
	if err != nil{
		if errors.Is(err,os.ErrNotExist) { //if file doesn't exist ,return empty{}
			return []Task{},nil
		}
		return nil,err
	}
	var tasks []Task
	err = json.Unmarshal(file,&tasks)
	if err != nil {
		return nil , err
	}
	return tasks,nil
}

func SaveTasks(tasks []Task) error{
	data,err := json.MarshalIndent(tasks," ","  ") //Formatted n readable
	if err != nil {
		return err
	}
	return os.WriteFile(fileName,data,0644) //0644 = rw- r-- r--
}

