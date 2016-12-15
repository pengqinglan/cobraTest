package main

import (
    "errors"
    "log"

    "github.com/spf13/cobra"
)

func main() {
    var rootCmd = &cobra.Command{
        Use:   "hugo",
        Short: "Hugo is a very fast static site generator",
        Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
        RunE: func(cmd *cobra.Command, args []string) error {
            // Do Stuff Here
            return errors.New("some random error")
        },
    }

    if err := rootCmd.Execute(); err != nil {
        log.Fatal(err)
    }
}