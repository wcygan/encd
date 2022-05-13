package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wcygan/encd/crypto"
	"os"
)

var decCmd = &cobra.Command{
	Use:   "dec",
	Short: "Decrypt a file that is provided as an argument.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file, oracle, writer, err := parseArgs(cmd, args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = crypto.Decrypt(file, oracle, writer)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(decCmd)
}
