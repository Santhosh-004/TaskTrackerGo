/*
Copyright © 2025 Santhosh G santhoshg0004@gmail.com
*/
package cmd

import (
	"fmt"
	"strconv"
	"tt/enums"
	"tt/utils"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks := utils.ReadJsonFile()
		for _, arg := range args {
			found := false
			for i, _ := range tasks {
				val, err := strconv.Atoi(arg)
				if err != nil {
					fmt.Printf("Invalid task ID: %s\n", arg)
					break
				}
				if tasks[i].Id == val {
					found = true
					if tasks[i].Status == enums.Done.ToString() {
						fmt.Printf("Task with ID %s is already marked as done.\n", arg)
					} else {
						tasks[i].Status = enums.Done.ToString()
						fmt.Printf("Marked task with ID %s as done successfully!\n", arg)
					}
					break
				}
			}
			if !found {
				fmt.Printf("Task with ID %s not found.\n", arg)
			}
		}
		utils.OverwriteJsonFile(tasks)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
