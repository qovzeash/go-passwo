package cmd

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	"github.com/qovzeash/go-passwo/utils"
	"github.com/spf13/cobra"
)

var size int
var criteria string
var omitSymbols, pinCode bool

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "quick, secure, and reliable password generator.",
	Run: func(cmd *cobra.Command, args []string) {
		if omitSymbols {
			criteria = "noSpecial"
		} else if pinCode {
			criteria = "pinCode"
		} else {
			criteria = "defaultCriteria"
		}

		generatePassword := (utils.GeneratePassword(size, criteria))

		err := clipboard.WriteAll(generatePassword)
		if err != nil {
			fmt.Println("failed to copy to clipboard:", err)

		} else {
			fmt.Println("password copied to clipboard!")
		}
	},
}

func init() {
	rootCmd.Flags().IntVar(&size, "size", 12, "specify the lenght of the password to be generated.")
	rootCmd.Flags().BoolVar(&omitSymbols, "omit-symbols", false, "generate a password without including any special characters.")
	rootCmd.Flags().BoolVar(&pinCode, "pin-code", false, "generate a numeric pin code instead of a full password.")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
