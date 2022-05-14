package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var decCmd = &cobra.Command{
	Use:   "dec",
	Short: "Decrypt a file that is provided as an argument.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		oracle, _, argSet, err := parseArgs(cmd, args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		run(oracle, argSet, false)
	},
}

func init() {
	rootCmd.AddCommand(decCmd)
}
