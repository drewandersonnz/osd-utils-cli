package servicelog

import (
	"fmt"

	"github.com/spf13/cobra"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

// newCmdPost implements the Console command which Consoles the specified account cr
func newCmdPost(streams genericclioptions.IOStreams, flags *genericclioptions.ConfigFlags) *cobra.Command {
	ops := newPostOptions(streams, flags)
	consoleCmd := &cobra.Command{
		Use:               "post",
		Short:             "Post a servicelog entry",
		Args:              cobra.NoArgs,
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(ops.validate(cmd))
			cmdutil.CheckErr(ops.run())
		},
	}

	consoleCmd.Flags().BoolVarP(&ops.verbose, "verbose", "v", false, "Verbose output")

	return consoleCmd
}

// consoleOptions defines the struct for running Console command
type postOptions struct {
	verbose bool

	genericclioptions.IOStreams
}

func newPostOptions(streams genericclioptions.IOStreams, flags *genericclioptions.ConfigFlags) *postOptions {
	return &postOptions{
		IOStreams: streams,
	}
}

func (o *postOptions) validate(cmd *cobra.Command) error {
	return nil
}

func (o *postOptions) run() error {
	fmt.Fprintf(o.IOStreams.Out, "Hi\n\n")

	return nil
}
