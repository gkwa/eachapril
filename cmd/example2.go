package cmd

import (
	"github.com/gkwa/eachapril/example2"
	"github.com/spf13/cobra"
)

var example2Cmd = &cobra.Command{
	Use:   "example2",
	Short: "Example command for indexing and querying markdown files",
	Run: func(cmd *cobra.Command, args []string) {
		example2.Run()
	},
}

var forceIndex bool

func init() {
	rootCmd.AddCommand(example2Cmd)

	example2Cmd.AddCommand(indexCmd)
	indexCmd.Flags().BoolVarP(&forceIndex, "force", "f", false, "Force re-creation of the index")

	example2Cmd.AddCommand(queryCmd)
}
