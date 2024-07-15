package cmd

import (
	"fmt"
	"os"

	"github.com/qovzeash/go-passwo/utils"
	"github.com/spf13/cobra"
)

var size int
var criteria string
var omitSymbols bool

var rootCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates secure passwords automatically.",
	Run: func(cmd *cobra.Command, args []string) {
		if omitSymbols {
			criteria = "noSpecial"
		} else if criteria == "" {
			criteria = "defaultCriteria"
		}
		fmt.Println(utils.GeneratePassword(size, criteria))
	},
}

func init() {
	rootCmd.Flags().IntVar(&size, "size", 12, "size of the password to generate.")
	rootCmd.Flags().BoolVar(&omitSymbols, "omit-symbols", false, "generate password without any special characters.")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
