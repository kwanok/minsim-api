package cmd

import (
	"github.com/kwanok/minsim-api/config"
	"github.com/kwanok/minsim-api/server"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "minsim-api",
	Short: "MinSim API",
	Long:  `MinSim API 서버입니다.`,
	Run: func(cmd *cobra.Command, args []string) {
		s := server.NewServer(config.NewConfig(cfgFile))

		if err := s.Start(); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("failed to execute root command: %v", err)
	}
}

var cfgFile string

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
}
