package main

import (
	"fmt"
	"os"

	"github.com/sajanv88/gocli/internal/cli"
	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{Use: "gocli"}
	root.AddCommand(cli.NewCmd())
	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
