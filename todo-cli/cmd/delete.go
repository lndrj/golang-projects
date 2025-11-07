/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"todo-cli/handlers"

	"github.com/spf13/cobra"
)

var deleteIndex string

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes Todo",
	Run: func(cmd *cobra.Command, args []string) {

		idx, err := strconv.Atoi(deleteIndex)
		if err != nil {
			fmt.Println("Enter valid index")
			return
		}

		if idx < 0 || idx >= len(handlers.GetAllTodos()) {
			fmt.Println("Index out of bounds")
			return
		}

		handlers.DeleteTodo(idx)
		fmt.Println("Todo deleted")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringVarP(&deleteIndex, "todo", "t", "", "Todo title")
	deleteCmd.MarkFlagRequired("todo")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
