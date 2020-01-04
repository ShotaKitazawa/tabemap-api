package main

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ShotaKitazawa/tabemap-api/external"
)

var rootCmd = &cobra.Command{
	Use:           "tabemap-api",
	Short:         "tabemap API Server",
	SilenceErrors: true,
	SilenceUsage:  true,
	Run: func(cmd *cobra.Command, args []string) {
		// Set context
		ctx := context.Background()
		ctx = context.WithValue(ctx, "bind-host",
			fmt.Sprintf("%s:%d",
				viper.GetString("bind-address"),
				viper.GetUint("bind-port"),
			),
		)
		ctx = context.WithValue(ctx, "db-type", viper.GetString("db-type"))
		ctx = context.WithValue(ctx, "mysql-uri",
			fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
				viper.GetString("db-user"),
				viper.GetString("db-password"),
				viper.GetString("db-host"),
				viper.GetUint("db-port"),
				viper.GetString("db-table-name"),
			))
		ctx = context.WithValue(ctx, "sqlite-uri", viper.GetString("db-filepath"))

		// Run
		external.Run(ctx)
	},
}

// set up Cobra/Viper
func init() {

	var cfgFile string
	cobra.OnInitialize(func() {
		if cfgFile != "" {
			viper.SetConfigFile(cfgFile)
		}
		viper.AutomaticEnv()
		viper.ReadInConfig()
	})

	rootCmd.PersistentFlags().StringP("bind-address", "", "0.0.0.0", "Bind address")
	rootCmd.PersistentFlags().UintP("bind-port", "", 8080, "Bind port")
	rootCmd.PersistentFlags().StringP("db-type", "", "mysql", "kind of DB (c.f. mysql, sqlite3)")
	rootCmd.PersistentFlags().StringP("db-user", "", "root", "User to connect DB")
	rootCmd.PersistentFlags().StringP("db-password", "", "password", "Password to connect DB")
	rootCmd.PersistentFlags().StringP("db-host", "", "127.0.0.1", "Host to connect DB")
	rootCmd.PersistentFlags().UintP("db-port", "", 3306, "Port to connect DB")
	rootCmd.PersistentFlags().StringP("db-table-name", "", "sample", "DB table name")
	rootCmd.PersistentFlags().StringP("db-filepath", "", "tabemap.sqlite3", "db filepath for sqlite3")
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "f", "", "Path to config file")
	viper.BindPFlag("bind-address", rootCmd.PersistentFlags().Lookup("bind-address"))
	viper.BindPFlag("bind-port", rootCmd.PersistentFlags().Lookup("bind-port"))
	viper.BindPFlag("db-type", rootCmd.PersistentFlags().Lookup("db-type"))
	viper.BindPFlag("db-user", rootCmd.PersistentFlags().Lookup("db-user"))
	viper.BindPFlag("db-password", rootCmd.PersistentFlags().Lookup("db-password"))
	viper.BindPFlag("db-host", rootCmd.PersistentFlags().Lookup("db-host"))
	viper.BindPFlag("db-port", rootCmd.PersistentFlags().Lookup("db-port"))
	viper.BindPFlag("db-table-name", rootCmd.PersistentFlags().Lookup("db-table-name"))
}

// entrypoint
func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
