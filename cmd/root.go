package cmd

import (
	"bufio"
	"fmt"
	"github.com/wcygan/encd/crypto"
	"io"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "encd",
	Short: "A tool to encrypt and decrypt files with passwords",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringP("password", "p", "", "The password used to encode the file")
	rootCmd.PersistentFlags().StringP("out", "o", "", "The file to write to")
}

func parseArgs(cmd *cobra.Command, args []string) ([]byte, *crypto.Oracle, io.Writer, error) {
	password, err := cmd.Flags().GetString("password")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	oracle, err := crypto.NewOracle(password)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f := args[0]
	file, err := os.ReadFile(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	out, err := cmd.Flags().GetString("out")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pwLen := len(password)
	if pwLen < crypto.MinimumPasswordLength {
		fmt.Println("password must be greater than " + strconv.Itoa(crypto.MinimumPasswordLength) + " bytes long (is " + strconv.Itoa(pwLen) + ")")
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

	return file, oracle, writer, nil
}
