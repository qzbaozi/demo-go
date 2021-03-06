package main

import (
	"demo-go/pkg/command/node"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var versionFlag bool
var mainCmd = &cobra.Command{
	Use: "command",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if versionFlag {
			//        version.Print()
		} else {
			cmd.HelpFunc()(cmd, args)
		}
		fmt.Println(args)
	},
}

func main() {
	mainCmd.AddCommand(node.Cmd())
	if mainCmd.Execute() != nil {
		os.Exit(1)
	}
}
