/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"github.com/xh-dev-go/luhnId/luhn"
	"unicode"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate the input luhnId code valid or not",
	Long: `Validate the input luhnId code valid or not
example:
e.g. "luhnId validate --code 9813256329"
validate the luhnId 9813256329 
`,
	Run: func(cmd *cobra.Command, args []string) {
		if *fromClipboard {
			c, err := clipboard.ReadAll()
			if err != nil {
				panic(err)
			}
			code = &c
			fmt.Printf("Read luhnId from %d\n", code)
		}

		if *code == "" {
			panic("Please input code to validate")
		}
		for _, c := range *code {
			if !unicode.IsDigit(c) {
				panic("input code is not all digit")
			}
		}

		var result string
		if luhn.Validate(*code) {
			result = fmt.Sprintf("Validate code")
		} else {
			result = fmt.Sprintf("Invalidate code")
		}

		fmt.Println(result)
	},
}

var code *string
var fromClipboard *bool

func init() {
	rootCmd.AddCommand(validateCmd)
	code = validateCmd.Flags().String("code", "", "Number of digital for the code (including check digit)")
	fromClipboard = validateCmd.Flags().Bool("from-clipboard", false, "get the code from clipboard")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// validateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// validateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
