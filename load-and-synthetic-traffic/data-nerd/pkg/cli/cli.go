package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"blacklemur.com/datanerd/pkg/cli/run"
	"blacklemur.com/datanerd/pkg/cli/version"
	"blacklemur.com/datanerd/pkg/config/system"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   system.GetAppName(),
	Short: "Data Nerd - Tool to simulate different kinds of data agents, formats and patterns",
	Long: `  Data Nerd is a tool to simulate different kinds of data agents,
	formats and patterns.

  For more information see the official tool documentation.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// init method
func init() {
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// Adding the commands
	rootCmd.AddCommand(version.NewCmdVersion())
	rootCmd.AddCommand(run.NewCmdRun())
}
