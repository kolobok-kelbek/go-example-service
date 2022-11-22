package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"net/http"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "up server",
	Run: func(cmd *cobra.Command, args []string) {
		http.HandleFunc("/hello", getHello)

		_ = http.ListenAndServe(":8081", nil)
	},
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
