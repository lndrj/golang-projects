package handlers

import (
	"encoding/json"
	"os"
	"todo-cli/types"
)

var todosFile = "todos.json"

func loadTodos() []types.Todo {
	data, err := os.ReadFile(todosFile)
	if err != nil {
		return []types.Todo{}
	}
	var todos []types.Todo
	json.Unmarshal(data, &todos)
	return todos
}

func saveTodos(todos []types.Todo) {
	data, _ := json.MarshalIndent(todos, "", " ")
	os.WriteFile(todosFile, data, 0644)
}

func AddTodo(todo types.Todo) {

	todos := loadTodos()
	todos = append(todos, todo)
	saveTodos(todos)
}

func GetAllTodos() []types.Todo {
	return loadTodos()
}

func DeleteTodo(index int) {
	todos := GetAllTodos()
	if index >= 0 && index < len(todos) {
		todos = append(todos[:index], todos[index+1:]...)
		saveTodos(todos)
	}
}

func ToggleTodo(index int) {
	todos := loadTodos()
	if index >= 0 && index < len(todos) {
		todos[index].Completed = !todos[index].Completed
		saveTodos(todos)
	}
}
