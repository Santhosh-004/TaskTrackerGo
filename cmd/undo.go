/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"tt/enums"
	"tt/utils"

	"github.com/spf13/cobra"
)

// undoCmd represents the undo command
var undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks := utils.ReadJsonFile()
		for _, arg := range args {
			val, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Invalid task ID: %s\n", arg)
				continue
			}

			if tasks[val-1].Status == enums.Done.ToString() {
				tasks[val-1].Status = enums.ToDo.ToString()
				fmt.Printf("Reverted task with ID %d to To-Do successfully!\n", val)
				break
			} else {
				fmt.Printf("Task with ID %d is already in To-Do status.\n", val)
			}

		}
		utils.OverwriteJsonFile(tasks)
	},
}

func init() {
	rootCmd.AddCommand(undoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// undoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// undoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
