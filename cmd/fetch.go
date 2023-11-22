package cmd

import (
	"os"

	"github.com/oryon-cloud/oryon/config"
	"github.com/oryon-cloud/oryon/logger"
	"github.com/spf13/cobra"
)

var FetchCmd = &cobra.Command{
	Use:   "fetch [url]",
	Short: "Fetch Oryon Application",
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.Flags().StringVarP(&configFile, "config", "c", "", "config file (default is $HOME/.oryon.toml)")
		// check args
		if len(args) == 0 {
			logger.Logger.Error("Please specify the module name")
			os.Exit(1)
		}

		// read config file
		if configFile != "" {
			err := config.ParseConfig(configFile)
			if err != nil {
				logger.Logger.Error("Error reading config file", err)
				os.Exit(1)
			}
		}
	},
	// Run: func(cmd *cobra.Command, args []string) {
	// 	for _, name := range args {
	// 		mm := pkm.PackageManager{NameUrl: name, Force: Force}
	// 		if err := mm.Fetch(); err != nil {
	// 			logger.Logger.Error("Error fetching module", err)
	// 			os.Exit(1)
	// 		}
	// 	}
	// },
}

func init() {

}
