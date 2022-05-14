package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var encCmd = &cobra.Command{
	Use:   "enc",
	Short: "Encrypt a list of files or directories",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		oracle, password, argSet, err := parseArgs(cmd, args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		verify(password, argSet)
		run(oracle, argSet, true)
	},
}

// verify prompts the user to verify their password
func verify(password string, args map[string]bool) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Attempting to encrypt these files or directories\n{")
	for arg, _ := range args {
		fmt.Println("  " + arg + ",")
	}
	fmt.Print("}\nPlease enter the password to proceed: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	input = strings.TrimSpace(input)
	if input != password {
		fmt.Fprintln(os.Stderr, "password do not match... aborting!")
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(encCmd)
}
