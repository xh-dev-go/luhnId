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
	"github.com/xh-dev-go/xhUtils/cobraUtils/cobraString"
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
		if fromClipboard.Value() {
			c, err := clipboard.ReadAll()
			if err != nil {
				panic(err)
			}
			code.SetValue(&c)
			fmt.Printf("Read luhnId from %s\n", code.Value())
		}

		if code.Value() == "" {
			panic("Please input code to validate")
		}
		for _, c := range code.Value() {
			if !unicode.IsDigit(c) {
				panic("input code is not all digit")
			}
		}

		var result string
		if luhn.Validate(code.Value()) {
			result = fmt.Sprintf("Valid code")
		} else {
			result = fmt.Sprintf("Invalid code")
		}

		fmt.Println(result)
	},
}

var code *cobraString.CobraString
var fromClipboard *cobraBool.CobraBool

func init() {
	rootCmd.AddCommand(validateCmd)
	code = cobraString.NewDefault("code", "Number of digital for the code (including check digit)", "").Bind(validateCmd)
	fromClipboard = cobraBool.NewDefault("from-clipboard", "get the code from clipboard", false).Bind(validateCmd)
}
