package todoapp

import "time"

type item struct {
	Task string
	Done bool
	CreatedAt time.Time
	CompletedAt time.Time
}

type Todos []items

func (t *Todos) Add(task string) {
	todo := item{
		Task:task
	}
}