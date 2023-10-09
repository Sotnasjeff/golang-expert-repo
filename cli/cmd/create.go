/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"main/cli/internal/database"

	"github.com/spf13/cobra"
)

func newCreateCommand(categoryDB database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "Create",
		Short: "Create a new Category",
		Long:  "Create a new Category",
		RunE:  runCreate(categoryDB),
	}
}

func runCreate(categoryDB database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")

		_, err := categoryDB.CreateCategory(name, description)
		if err != nil {
			return err
		}
		return nil
	}
}

func init() {
	createCmd := newCreateCommand(GetCategory(GetDB()))
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "Name of Category")
	createCmd.Flags().StringP("description", "d", "", "Description of Category")
	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("description")
	createCmd.MarkFlagsRequiredTogether("name", "description")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
