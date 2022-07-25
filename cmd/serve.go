package cmd

import (
	"github.com/pestanko/gothy-mini/pkg/cfg"
	"github.com/pestanko/gothy-mini/pkg/rest"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the http server",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := cfg.LoadApplicationConfig(cfg.Vars.Env)
		if err != nil {
			log.Fatal().
				Str("env", cfg.Vars.Env).
				Err(err).
				Msg("unable to load configuration")
			os.Exit(1)
		}

		log.Info().
			Str("addr", config.Server.Addr).
			Str("env", cfg.Vars.Env).
			Msg("server started!")

		if err := http.ListenAndServe(config.Server.Addr, rest.CreateResetServer(config)); err != nil {
			log.Fatal().
				Str("addr", config.Server.Addr).
				Str("env", cfg.Vars.Env).
				Err(err).
				Msg("unable to start an application server")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
