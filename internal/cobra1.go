package internal

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func StudyCobra() {
	var cmdPull = &cobra.Command{
		Use:   "pull [OPTIONS] NAME[:TAG|@DIGEST]",
		Short: "Pull an image or a repository from a registry",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Pull: " + strings.Join(args, " "))
		},
	}

	var rootCmd = &cobra.Command{Use: "docker"}
	rootCmd.AddCommand(cmdPull)
	rootCmd.Execute()
}