package storage

type Task struct { //Field Name 	Field Type		Struct Tag
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}