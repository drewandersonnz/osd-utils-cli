package servicelog

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

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

// PostMessage for receiving inbound messages
type PostMessage struct {
	ClusterUUID  string `json:"cluster_uuid"`
	Description  string `json:"description"`
	InternalOnly bool   `json:"internal_only"`
	Kind         string `json:"kind"`
	ServiceName  string `json:"service_name"`
	Severity     string `json:"severity"`
	Summary      string `json:"summary"`
}

func (o *postOptions) validate(cmd *cobra.Command) error {
	return nil
}

func (o *postOptions) run() error {
	var message PostMessage

	decoded := json.NewDecoder(os.Stdin)
	for {
		err := decoded.Decode(&message)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if err := o.postMessage(message); err != nil {
			return err
		}
	}

	return nil
}

func (o *postOptions) postMessage(message PostMessage) error {

	jMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}

	fmt.Fprintf(o.IOStreams.Out, string(jMessage))
	fmt.Fprintf(o.IOStreams.Out, "\n\n")

	return nil
}
