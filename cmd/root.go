/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var filepath string

var Filepath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "textReader",
	Short: "A simple tool to count the number of texts in a file",
	Long: `To start, by passing a valid file path like: textReader -p [YOUR FILE PATH] 
	and choose the options to count the number of certain characters`,

	Run: func(cmd *cobra.Command, args []string) {
		// 取得flag的值，並儲存於全域變數
		Filepath = cmd.Flag("filepath").Value.String()
		fmt.Println(Filepath)
		Read()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// 使用local flag所以不會共享到其他child command
	// 而要共享的話就是用global variable就好
	rootCmd.Flags().StringVarP(&filepath, "filepath", "p", "", "required a filepath")
	rootCmd.MarkFlagRequired("filepath") // 綁定required flag
}
