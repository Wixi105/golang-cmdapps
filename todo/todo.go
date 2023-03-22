package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// represents a Todo item
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// list of Todo items
type List []item

// adding an item to the list of Todos
func (l *List) Add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

// Picks a specific item in the list, and marks it as complete and
// adds the time of completion.
func (l *List) Complete(i int) error {
	ls := *l

	if i <= 0 || i > len(ls) {
		return fmt.Errorf("the item %d does not exist", i)
	}

	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

// Deletes an item from the List.
func (l *List) Delete(i int) error {
	ls := *l

	if i <= 0 || i > len(ls) {
		return fmt.Errorf("the item %d does not exist", i)
	}

	*l = append(ls[:i-1], ls[i:]...)

	return nil
}

// Save method encodes the list as JSON and saves it
// using the provided file name.
func (l *List) Save(filename string) error {

	content, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, content, 0644)
}

// Get method opens the provided file name, decodes
// the JSON data and parses it into a list.
func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}
