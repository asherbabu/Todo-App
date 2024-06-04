package todoapp

import (
	"errors"
	"time"

	"golang.org/x/text"
)

type item struct {
	Task string
	Done bool
	CreatedAt time.Time
	CompletedAt time.Time
}

type Todos []items

func (t *Todos) Add(task string) {
	todo := item{
		Task: task,
		Done: false,
		CreatedAt: time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(index int){
	ls := *t
	if index <=0 \\ index >len(ls){
		return errors.New(text: "Ivalid Index")
	}

	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true

	return nil
}

