package cmds

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/Thiti-Dev/ksc/internal/cmds/helpers"
	"github.com/Thiti-Dev/ksc/internal/constants"
	"github.com/Thiti-Dev/ksc/internal/persistence"
	"github.com/Thiti-Dev/ksc/internal/tui/inputmgr"
	"github.com/Thiti-Dev/ksc/internal/tui/list"
	"github.com/Thiti-Dev/ksc/internal/types"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func findCmdFromUUID(cmds []types.CommandList, uuid uuid.UUID) (*types.CommandList, bool) {
	for _, cmd := range cmds {
		if cmd.UUID == uuid {
			return &cmd, true
		}
	}

	return nil, false
}

var RootCmd = &cobra.Command{
	Use:   "main",
	Short: "KSC's command hub",
	Run: func(_ *cobra.Command, args []string) {
		loadedCmds, err := persistence.LoadFromGob[types.CommandList](constants.PERSISTED_DATA_FILE_NAME)
		if err != nil {
			// freshy user without saved data yet
			fmt.Println("Doesn't have any commands registered into the ksc yet, consider running `ksc add <cmd> <desc>`")
			os.Exit(1)
		}

		selected_uuid := list.ShowInteractiveCommandList("KSC", loadedCmds)

		// if quit via ctrl + c, tea.quit triggered and we do the safe exit here
		if selected_uuid.Get().GetAsUUID() == uuid.Nil {
			os.Exit(0)
		}

		// No need to do the hashmap as it would be a waste of time also because this is the one time operation
		// O(n) is acceptable
		targetCmd, found := findCmdFromUUID(loadedCmds, selected_uuid.Get().GetAsUUID())
		if !found {
			// this shouldn't be happenning so we panic it
			log.Panic("Selected coomand is not found")
		}
		// Regular expression to match anything between < and >
		re := regexp.MustCompile(`<[^>]*>`)

		// Find all argument placeholder
		argPlaceholders := re.FindAllString(targetCmd.Cmd, -1)

		if len(argPlaceholders) == 0 {
			// just execute the command right away
			fmt.Println("> " + targetCmd.Cmd)
			helpers.ExecuteShellCommand(targetCmd.Cmd)
			os.Exit(0) // after execution, exit the process
		}

		fmt.Println("> " + targetCmd.Cmd)

		filledArgs := inputmgr.ShowInteractiveArgumentPlaceholdersFiller(argPlaceholders)

		executionCommand := targetCmd.Cmd

		for i, placeholder := range argPlaceholders {
			executionCommand = strings.Replace(executionCommand, placeholder, filledArgs[i], 1)
		}

		fmt.Println("> " + executionCommand)
		helpers.ExecuteShellCommand(executionCommand)
	},
}
