package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"h-ui/model/constant"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long:  "Show version",
	Run:   runVersionCmd,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func runVersionCmd(cmd *cobra.Command, args []string) {
	fmt.Println("h-ui version", constant.Version)
}
