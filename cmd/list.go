/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/rozoomcool/task-cli/internal/model"
	"github.com/rozoomcool/task-cli/internal/service"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		status := ""
		if len(args) > 0 {
			status = args[0]
		}

		tasksService := service.NewTaskService()

		tasks, err := tasksService.ListTasks(model.TaskStatus(status))
		if err != nil {
			fmt.Println(err)
			return
		}
		DDMMYYYYhhmm := "2006-01-02 15:04"
		for _, v := range tasks {
			fmt.Printf("%v | %v\t| %v\t| %v | %v\n", v.Id, v.Description, v.Status, v.CreatedAt.Local().Format(DDMMYYYYhhmm), v.UpdatedAt.Local().Format(DDMMYYYYhhmm))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
