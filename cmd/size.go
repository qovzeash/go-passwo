package cmd

import (
	"fmt"
	"os"

	"github.com/qovzeash/go-passwo/utils"
	"github.com/spf13/cobra"
)

var size int

var sizeCmd = &cobra.Command{
	Use:     "-s",
	Short:   "-s or --size defines the length of the generated password.",
	Aliases: []string{"--size"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(utils.GeneratePassword(size))
	},
}

func init() {
	sizeCmd.Flags().IntVarP(&size, "size", "s", 8, "Length of the generated password")
}

func Execute() {
	if err := sizeCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
