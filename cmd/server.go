// +build !runner_only

package cmd

import (
	"git.elenet.me/yuelong.huang/pansible/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here

		viper.SetDefault("host", "0.0.0.0")
		viper.SetDefault("port", 5757)
		viper.SetDefault("ssh_key", "tmp/id")
		viper.SetDefault("jwt_secret", "secret")

		viper.SetDefault("mysql_address", "127.0.0.1:3306")
		viper.SetDefault("mysql_db", "pansible")
		viper.SetDefault("mysql_user", "root")
		viper.SetDefault("mysql_pass", "toor")

		if err := viper.ReadInConfig(); err != nil {
			os.Exit(1)
		}

		api.StartServer()

	},
}

func init() {
	RootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
