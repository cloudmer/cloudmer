package cmd

import (
	"cloudmer/library/config"
	"cloudmer/library/store/mysql"
	"cloudmer/service/http"
	"cloudmer/share"
	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use: "http",
	Short: "http service",
	Long: ``,
	Example: `
app --config file.yaml (default is application/config/http.yaml)
`,
	Run: func(cmd *cobra.Command, args []string) {
		if configPath == "" {
			cmd.Help()
			return
		}
		// 配置项
		share.Viper = config.GetInstance()
		config.Loader(configPath)
		share.Mysql = mysql.StdConfig("mysql").Build()
		// http service run
		http.StdConfig("http").Build()
		return
	},
}

func init()  {
	httpCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "config file (default is application/config/config.yaml)")
}