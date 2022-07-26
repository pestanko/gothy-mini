package cmd

import (
	"fmt"
	"github.com/pestanko/gothy-mini/pkg/security"
	"github.com/rs/zerolog/log"
	"os"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// hashPwdCmd represents the hashPwd command
var hashPwdCmd = &cobra.Command{
	Use:   "hash-pwd",
	Short: "Hash user provided password",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Password: ")
		bytepw, err := term.ReadPassword(syscall.Stdin)
		if err != nil {
			log.Fatal().Err(err).Msg("unable to read from stdin")
			os.Exit(1)
		}
		providedPwd := string(bytepw)

		hasher := security.NewPasswordHasher()
		hash, err := hasher.HashPassword(providedPwd)
		if err != nil {
			log.Fatal().Err(err).Msg("unable to hash the password")
			os.Exit(2)
		}

		fmt.Println("\nPassword Hash: ", hash)
	},
}

func init() {
	toolsCmd.AddCommand(hashPwdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hashPwdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hashPwdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
