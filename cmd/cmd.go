package cmd

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"h-ui/model/constant"
	"h-ui/service"
	"strconv"
)

var rootCmd = &cobra.Command{
	Use:   "hui",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		p, _ := cmd.Flags().GetString("p")
		if p != "" {
			if err := validateAndSetPort(p); err != nil {
				logrus.Errorf(err.Error())
				panic(err)
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
	rootCmd.PersistentFlags().StringP("p", "p", "", "web port")
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
