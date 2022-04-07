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

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate luhn id",
	Long: `Generate a luhn id by configuration 
example: 
e.g. "luhnId generate -starting-digit 2 -c" => 
generate luhnId starting with 2 and totally 10 digit (including check digit) and copy to clipboard
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(*startingDigit) > 1 {
			panic("starting digit should at most 1 character")
		}
		var prefix string = ""
		if len(*startingDigit) == 1 {
			if !unicode.IsDigit([]rune(*startingDigit)[0]) {
				panic("starting digit should be number")
			}
			prefix = *startingDigit
		}

		if len(prefix) != 0 {
			d := *digit - 1
			digit = &d
		}

		if *digit <= 0 {
			panic("digit should not be zero or less")
		}

		id := luhn.Gen(prefix, *digit)
		if *outputToClipboard {
			clipboard.WriteAll(id)
		}
		fmt.Println(id)
	},
}

var digit *int
var startingDigit *string
var outputToClipboard *bool

func init() {
	rootCmd.AddCommand(generateCmd)

	outputToClipboard = generateCmd.Flags().BoolP("to-clipboard", "c", false, "output result to clipboard")
	digit = generateCmd.Flags().Int("digit", 10, "Number of digit to be generated")
	startingDigit = generateCmd.Flags().String("starting-digit", "", "Starting digit")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
