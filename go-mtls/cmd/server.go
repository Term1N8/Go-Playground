package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	_ "nate.com/go-mtls/checks"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(server)

}

var server = &cobra.Command{
	Use:   "server",
	Short: "run teh server",
	Long:  `run teh server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Startup server...")
			http.HandleFunc("/hello", helloHandler)

	// Listen to port 8080 and wait
	log.Fatal(http.ListenAndServe(":8080", nil))
	},
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Write "Hello, world!" to the response body
	io.WriteString(w, "Hello, world!\n")
}