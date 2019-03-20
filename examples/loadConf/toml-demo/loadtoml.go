package toml_demo

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var confFile string

func main() {
	var RootCmd = &cobra.Command{
		Use:   "SKProxy Server",
		Short: "SKProxy Server",
		Long:  "SKProxy Server",
		Run: func(cmd *cobra.Command, args []string) {
			serviceMap := viper.GetStringMap("service")

			fmt.Println(">>>service map>>>:", serviceMap)

			fmt.Println("")

			serviceRedis := viper.GetStringMap("redis")
			fmt.Println(">>>service redis>>>:", serviceRedis)
		},
	}

	cobra.OnInitialize(loadConf)
	RootCmd.PersistentFlags().StringVarP(&confFile, "config", "c", "", "config file")

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func loadConf() {
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}

	if confFile != "" {
		viper.SetConfigFile(confFile)
	} else {
		fmt.Println("配置文件路径为空！")
		return
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println(err.Error())
		return
	}
}
