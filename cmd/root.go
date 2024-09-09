package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tester",
	Short: "tester is a fuzztesting cli application for testing your endpoints",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the fuzztester cli tool. To get help, run fuzztesting-cli")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
