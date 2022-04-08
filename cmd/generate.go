/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"github.com/xh-dev-go/luhnId/luhn"
	"github.com/xh-dev-go/xhUtils/cobraUtils/cobraBool"
	"github.com/xh-dev-go/xhUtils/cobraUtils/cobraInt"
	"github.com/xh-dev-go/xhUtils/cobraUtils/cobraString"
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
		if len(startingDigit.Value()) > 1 {
			panic("starting digit should at most 1 character")
		}
		var prefix string = ""
		if len(startingDigit.Value()) == 1 {
			if !unicode.IsDigit([]rune(startingDigit.Value())[0]) {
				panic("starting digit should be number")
			}
			prefix = startingDigit.Value()
		}

		id, err := luhn.Gen(prefix, digit.Value())
		if err != nil {
			panic(err)
		}
		if outputToClipboard.Value() {
			clipboard.WriteAll(id)
		}
		fmt.Println(id)
	},
}

var digit *cobraInt.CobraInt

var startingDigit *cobraString.CobraString

var outputToClipboard *cobraBool.CobraBool

func init() {
	rootCmd.AddCommand(generateCmd)

	outputToClipboard = cobraBool.NewDefault("to-clipboard", "output result to clipboard", false).Shorthand("c").Bind(generateCmd)
	digit = cobraInt.NewDefault("digit", "Number of digit to be generated", 10).Bind(generateCmd)
	startingDigit = cobraString.NewDefault("starting-digit", "Starting digit", "").Bind(generateCmd)
}
