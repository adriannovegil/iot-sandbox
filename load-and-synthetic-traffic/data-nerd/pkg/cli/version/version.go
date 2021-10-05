package version

import (
	"fmt"

	"github.com/spf13/cobra"

	"blacklemur.com/datanerd/pkg/config/system"
)

var ()

// NewCmdVersion prints version
func NewCmdVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "show version statement",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("datanerd", system.GetVersion())
		},
	}
}
