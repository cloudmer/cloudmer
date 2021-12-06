package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "app",
	Short: "app short",
	Long: `app long`,
}

// config path
var configPath string

func init()  {
	rootCmd.AddCommand(httpCmd)
}

// Execute
func Execute()  {
	if err := rootCmd.Execute(); err != nil {

	}
}