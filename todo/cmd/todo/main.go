package main

import (
	"fmt"
	"os"
	"strings"

	"cmdapps/todo"
)

const todoFilename = ".todo.json"

func main() {
	l := &todo.List{}

	if err := l.Get(todoFilename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}
		//Concantenate all provided items with a space and 
		// add to the list as an item.
	default:
		//Concantenate all arguments with a space.
		item := strings.Join(os.Args[1:], " ")
		//Add the task
		l.Add(item)
		//Save the new list
		if err := l.Save(todoFilename); err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}

	}
}
