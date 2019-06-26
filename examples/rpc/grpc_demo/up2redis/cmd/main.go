package main

import (
	up2redis2 "github.com/aaronize/grandonion/examples/rpc/grpc_demo/up2redis"
	"github.com/spf13/cobra"
	"log"
)

var (
	configPath string
	defaultConfigPath = "./config.json"
)

var rootCmd = cobra.Command{
	Use: "up2redis",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		path := defaultConfigPath
		if configPath != "" {
			path = configPath
		}
		log.Println("Using config ", path)
		server := up2redis2.Server{Config: &up2redis2.Config{}}
		if err := server.Config.LoadConfig(path); err != nil {
			log.Fatalf("start server error. %s", err.Error())
		}
		log.Printf("CONFIGURES %v", server.Config)
		if err := server.Start(); err != nil {
			log.Fatalf("start server error. %s", err.Error())
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("start server error. %s", err.Error())
	}
}

func init() {
	rootCmd.Flags().StringVarP(&configPath, "config", "c", "", "specify config file path.")
}
