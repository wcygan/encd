package cmd

import (
	"bufio"
	"fmt"
	"github.com/wcygan/encd/encoder"
	"io"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var encCmd = &cobra.Command{
	Use:   "enc",
	Short: "Encode an image",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("please provide one argument")
			os.Exit(1)
		}

		f := args[0]
		file, err := os.ReadFile(f)
		if err != nil {
			fmt.Println("Could not read " + f)
			os.Exit(1)
		}

		out, err := cmd.Flags().GetString("out")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		password, err := cmd.Flags().GetString("password")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		pwLen := len(password)
		if pwLen != 32 {
			fmt.Println("password must be 32 bytes long (is " + strconv.Itoa(pwLen) + ")")
			os.Exit(1)
		}

		var writer io.Writer
		if out != "" {
			f, err := os.Create(out)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			writer = bufio.NewWriter(f)
		} else {
			writer = os.Stdout
		}

		err = encoder.Encode(file, password, writer)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(encCmd)
	encCmd.Flags().StringP("password", "p", "", "The password used to encode the file")
	encCmd.Flags().StringP("out", "o", "", "The file to write to")
}
