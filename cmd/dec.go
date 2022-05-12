package cmd

import (
	"github.com/wcygan/encd/decoder"

	"github.com/spf13/cobra"
)

var decCmd = &cobra.Command{
	Use:   "dec",
	Short: "Decode an image",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO @wcygan: parse a file to decode
		decoder.Decode()
	},
}

func init() {
	rootCmd.AddCommand(decCmd)
}
