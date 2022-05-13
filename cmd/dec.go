package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wcygan/encd/crypto"
	"os"
)

var decCmd = &cobra.Command{
	Use:   "dec",
	Short: "Decode an image",
	Run: func(cmd *cobra.Command, args []string) {
		file, password, writer, err := parseArgs(cmd, args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = crypto.Decode(file, password, writer)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(decCmd)
}
