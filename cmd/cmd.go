package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"h-ui/model/constant"
	"os"
	"strconv"
)

var (
	port    string
	version bool
)
var rootCmd = &cobra.Command{
	Use:   "h-ui",
	Short: "just the panel for Hysteria2",
	Long:  "just the panel for Hysteria2",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && args[0] == "server" {
			if err := validateAndSetPort(port); err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			startServer()
		}
		if version {
			fmt.Println("h-ui version", constant.Version)
		}
		fmt.Println("Usage: h-ui [server] [-p <port>] [-v] [-h]")
	},
}

func validateAndSetPort(p string) error {
	if p != "" {
		value, err := strconv.ParseInt(p, 10, 64)
		if err != nil {
			return errors.New("invalid port value")
		}
		if value <= 0 || value > 65535 {
			return errors.New("the port range is between 0-65535")
		}
	}
	return nil
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "", "Set the port for the server")
	rootCmd.PersistentFlags().BoolVarP(&version, "version", "v", false, "Print the version number")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
