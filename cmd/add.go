/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/rozoomcool/task-cli/internal/model"
	"github.com/rozoomcool/task-cli/internal/repository"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			description := args[0]
			if description == "" {
				fmt.Println("Empty description")
				return
			}
			taskRepo := repository.GetTaskRepository()
			newTask := model.Task{
				Description: description,
				Status:      model.TaskStatusToDo,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}
			id, err := taskRepo.Add(&newTask)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Task added successfully (ID: %d)\n", id)
		} else {
			fmt.Println("Add task description")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
