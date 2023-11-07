package cmd

import (
	"github.com/disism/saikan/api-server"
	"github.com/disism/saikan/internal/conf"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long:  ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		if os.Getenv(conf.JWTSecret) == "" {
			log.Fatal("JWT_SECRET environment variable is not set. The application will now exit.")
		}
		if os.Getenv(conf.ServiceEndpoint) == "" {
			log.Fatal("SERVICE_ENDPOINT environment variable is not set. The application will now exit.")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := api_server.Run(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//runCmd.PersistentFlags().String("endpoint", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
