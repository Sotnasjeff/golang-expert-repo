/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Chamado Pré run")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Chamado Post run")
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("Ouve um erro")
	},
}

var category string

func init() {

	rootCmd.AddCommand(categoryCmd)
	//categoryCmd.PersistentFlags().String("name", "", "Name of category") //Global Flag
	//categoryCmd.PersistentFlags().StringVarP(&category, "name", "n", "", "Name of category")
	//categoryCmd.PersistentFlags().StringP("fuck", "f", "Y", "Name of category") //Quando se usa o StringP, você pode passar comando opcional e usar short hand
	//categoryCmd.PersistentFlags().BoolP("exists", "e", false, "Check if category exists")
	//categoryCmd.Flags().String("whatever", "", "Name of whatever") //Local Flag

	//Após o PersistentFlags ou Flags, você passa o tipo da flag, se é booleano, inteiro, string e etc...

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// categoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
