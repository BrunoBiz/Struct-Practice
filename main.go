package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

func main() {
	printSomething(1)
	printSomething("adasdas")
	printSomething(1.50)

	printSomething(add(1, 2))
	printSomething(add(1.5, 2.5))
	printSomething(add("Bru", "no"))

	title, content := getNoteData()
	todoText := getUserInput("Todo list:")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Print(err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Print(err)
		return
	}

	err = outputData(todo)
	if err != nil {

	}

	outputData(userNote)
}

func printSomething(value interface{}) {
	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer:", intVal)
		return
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float65", floatVal)
		return
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println(stringVal)
		return
	}

	// switch value.(type) {
	// case string:
	// 	fmt.Println(value)
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float64:", value)
	// }
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed")
		return err
	}

	fmt.Println("Saving the note succeeded")
	return nil
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func outputData(data outputtable) error {
	data.Display()
	err := saveData(data)

	return err
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getTodoData() string {
	return getUserInput("Todo list:")
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
