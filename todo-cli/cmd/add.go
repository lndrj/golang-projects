/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"todo-cli/handlers"
	"todo-cli/types"

	"github.com/spf13/cobra"
)

var addTodo string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds new task",
	Run: func(cmd *cobra.Command, args []string) {
		todo := types.Todo{
			Title:     addTodo,
			Completed: false,
		}
		handlers.AddTodo(todo)
		fmt.Println("Adding todo", addTodo)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&addTodo, "todo", "t", "", "Todo title")
	addCmd.MarkFlagRequired("todo")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
