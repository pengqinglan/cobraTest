package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var RootCmd = &cobra.Command{
	Use: "hugo",
	Short: "this is the sort of hugo, fuck you!",
	Long: `asdasd
	asd
	asdddsdasda
	thi is a long long desc on hugo`,
	Run:func(cmd *cobra.Command, args []string) {
		fmt.Println("fuckyou!")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("fuckyou! two")
	},
}


