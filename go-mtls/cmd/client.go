package cmd

import (
	"fmt"
	
	"nate.com/go-mtls/checks"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(client)

}

var client = &cobra.Command{
	Use:   "client",
	Short: "run teh client",
	Long:  `run teh client`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Startup client...")
		for _, check := range checks.All {
			if check.Run() {
				fmt.Println("Success: ", check.Name)
			} else {
				fmt.Println("Fail: ", check.Name)
			}
		}
	},
}
