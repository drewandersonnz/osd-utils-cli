package servicelog

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// NewCmdServiceLog implements the base account command
func NewCmdServiceLog(streams genericclioptions.IOStreams, flags *genericclioptions.ConfigFlags) *cobra.Command {
	accountCmd := &cobra.Command{
		Use:               "servicelog",
		Short:             "OCM/Hive Service Log",
		Args:              cobra.NoArgs,
		DisableAutoGenTag: true,
		Run:               help,
	}

	accountCmd.AddCommand(newCmdPost(streams, flags))

	return accountCmd
}

func help(cmd *cobra.Command, _ []string) {
	cmd.Help()
}
