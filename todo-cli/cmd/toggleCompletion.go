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

var toggleIndex string

// toggleCompletionCmd represents the toggleCompletion command
var toggleCompletionCmd = &cobra.Command{
	Use:   "toggle",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {

		idx, err := strconv.Atoi(toggleIndex)
		if err != nil {
			fmt.Println("Enter valid index")
			return
		}

		if idx < 0 || idx >= len(handlers.GetAllTodos()) {
			fmt.Println("Index out of bounds")
			return
		}

		handlers.ToggleTodo(idx)
		fmt.Println("Todo status toggled")

	},
}

func init() {
	rootCmd.AddCommand(toggleCompletionCmd)
	toggleCompletionCmd.Flags().StringVarP(&toggleIndex, "todo", "t", "", "Todo index")
	toggleCompletionCmd.MarkFlagRequired("todo")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toggleCompletionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toggleCompletionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
