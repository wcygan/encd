package cmd

import (
	"github.com/wcygan/encd/encode"

	"github.com/spf13/cobra"
)

var encCmd = &cobra.Command{
	Use:   "enc",
	Short: "Encode an image",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO @wcygan: parse a file to encode
		encode.Hello()
	},
}

func init() {
	rootCmd.AddCommand(encCmd)
}
