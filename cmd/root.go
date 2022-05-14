package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wcygan/encd/crypto"
	"io"
	"os"
	"strconv"
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
}

func run(oracle *crypto.Oracle, args map[string]bool, isEncrypt bool) {
	for arg, _ := range args {
		// open the file
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		// get the file info
		info, err := f.Stat()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		// skip if it is a directory
		// todo @wcygan: make this recursive and split it into a new function
		if !info.IsDir() {
			// get the file contents & obtain an io.Writer
			contents, writer, err := fopen(arg)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}

			if isEncrypt {
				err = crypto.Encrypt(contents, oracle, writer)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				}
			} else {
				err = crypto.Decrypt(contents, oracle, writer)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					continue
				}
			}
		} else {
			fmt.Fprintln(os.Stdout, "skipping directory: "+arg)
		}
	}
}

func parseArgs(cmd *cobra.Command, args []string) (*crypto.Oracle, string, map[string]bool, error) {
	password, err := cmd.Flags().GetString("password")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pwLen := len(password)
	if pwLen < crypto.MinimumPasswordLength {
		_, err := fmt.Fprintln(os.Stderr, "Password must be at least "+strconv.Itoa(crypto.MinimumPasswordLength)+" characters long (is "+strconv.Itoa(pwLen)+")")
		if err != nil {
			println("error checking password")
			os.Exit(1)
		}
		os.Exit(1)
	}

	oracle, err := crypto.NewOracle(password)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// de-duplicate the args, we could surprise the user if we encrypt the same file multiple times in a row
	argSet := make(map[string]bool)
	for _, arg := range args {
		argSet[arg] = true
	}

	return oracle, password, argSet, nil
}

// fopen reads a file then returns its contents & an io.Writer to overwrite the file
func fopen(filename string) ([]byte, io.Writer, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}

	out, err := os.Create(filename)
	if err != nil {
		return nil, nil, err
	}

	var writer io.Writer
	writer = bufio.NewWriter(out)

	return file, writer, nil
}
