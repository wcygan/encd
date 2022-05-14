package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wcygan/encd/crypto"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
)

var rootCmd = &cobra.Command{
	Use:   "encd",
	Short: "A tool to encrypt and decrypt files and directories with passwords",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringP("password", "p", "", "The secret phrase used for encryption and decryption")
}

func run(oracle *crypto.Oracle, paths map[string]bool, isEncrypt bool) {
	for path := range paths {
		err := filepath.WalkDir(path, func(currentPath string, d fs.DirEntry, err error) error {
			// open the file
			f, err := os.Open(currentPath)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return err
			}

			// get the file info
			info, err := f.Stat()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return err
			}

			// skip the encryption if it is a directory
			if !info.IsDir() {
				// get the file contents & obtain an io.Writer
				contents, writer, err := fopen(currentPath)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					return err
				}

				if isEncrypt {
					err = crypto.Encrypt(contents, oracle, writer)
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
						return err
					}
				} else {
					err = crypto.Decrypt(contents, oracle, writer)
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
						return err
					}
				}
			}

			return nil
		})
		if err != nil {
			return
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

	return file, out, nil
}
