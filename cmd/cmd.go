package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"h-ui/model/constant"
	"h-ui/util"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "h-ui",
	Short: "just the panel for Hysteria2",
	Long:  "just the panel for Hysteria2.",
	Run:   run,
}

var (
	version bool
	port    string
)

func init() {
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "Show version")
	rootCmd.Flags().StringVarP(&port, "port", "p", "", "The port of the web server")
}

func run(cmd *cobra.Command, args []string) {
	if version {
		fmt.Println("h-ui version", constant.Version)
		return
	}

	if err := util.VerifyPort(port); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for {
		if err := runServer(port); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
