package cmd

import (
	"github.com/wcygan/encd/decode"

	"github.com/spf13/cobra"
)

var decCmd = &cobra.Command{
	Use:   "dec",
	Short: "Decode an image",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO @wcygan: parse a file to decode
		decode.Hello()
	},
}

func init() {
	rootCmd.AddCommand(decCmd)
}
