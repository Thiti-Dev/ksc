package cmds

import (
	"fmt"
	"os"

	"github.com/Thiti-Dev/ksc/internal/constants"
	"github.com/Thiti-Dev/ksc/internal/persistence"
	"github.com/Thiti-Dev/ksc/internal/types"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a command to the ksc",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Add a command")
		if len(args) < 2 {
			fmt.Println("Required command and description to add command > `ksc add <cmd> <desc>`")
			os.Exit(1)
		}

		existingCmds, _ := persistence.LoadFromGob[types.CommandList](constants.PERSISTED_DATA_FILE_NAME)

		var command, description = args[0], args[1]

		existingCmds = append(existingCmds, types.CommandList{
			UUID: uuid.New(),
			Cmd:  command,
			Desc: description,
		})

		persistence.SaveToGob(constants.PERSISTED_DATA_FILE_NAME, existingCmds)

	},
}
