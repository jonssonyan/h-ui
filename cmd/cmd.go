package cmd

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"h-ui/model/constant"
	"h-ui/service"
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
		if version {
			fmt.Println("h-ui version ", constant.Version)
			return
		}
		if port != "" {
			if err := validateAndSetPort(port); err != nil {
				logrus.Errorf(err.Error())
				return
			}
		}
	},
}

func validateAndSetPort(p string) error {
	value, err := strconv.ParseInt(p, 10, 64)
	if err != nil {
		return errors.New("invalid port value")
	}
	if value < 0 || value > 65535 {
		return errors.New("the port range is between 0-65535")
	}
	if err := service.UpdateConfig(constant.HUIWebPort, p); err != nil {
		return err
	}
	return nil
}

func InitCmd() {
	rootCmd.PersistentFlags().StringVarP(&port, "port", "p", "", "Set the port for the server")
	rootCmd.PersistentFlags().BoolVarP(&version, "version", "v", false, "Print the version number")

	if err := rootCmd.Execute(); err != nil {
		logrus.Errorf("execute root command error: %v", err)
	}
}
