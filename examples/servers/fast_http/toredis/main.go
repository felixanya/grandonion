package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"toredis/server"
)

var configPath string
const DefaultConfigPath = "./config.json"

var rootCmd = cobra.Command{
	Use: "toredis",
	Short: "",
	Long: ``,
	RunE: Run,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("cmd error,", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&configPath, "config", "c", "", "specify config file path.")
}

func Run(cmd *cobra.Command, args []string) error {
	conf := DefaultConfigPath
	if configPath != "" {
		conf = configPath
	}

	s := &server.Server{
		Config: &server.Config{},
	}
	if err := s.Config.LoadConfig(conf); err != nil {
		fmt.Println("Load config error,", err)
		os.Exit(1)
	}
	fmt.Printf("Using config %s\n", conf)
	fmt.Println("listen:", s.Config.Listen)
	fmt.Println("log:", s.Config.Logger.MaxAge)
	fmt.Println("cluster:", s.Config.Redis.Cluster)

	if err := s.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return nil
}
