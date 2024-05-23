package cmds

import "github.com/spf13/cobra"

func AddAllCommandsTo(cmd *cobra.Command) {
	for _, command := range []*cobra.Command{AddCmd} {
		cmd.AddCommand(command)
	}
}
