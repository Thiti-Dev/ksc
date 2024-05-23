package list

import (
	"fmt"
	"os"

	"github.com/Thiti-Dev/ksc/internal/types"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func ShowInteractiveCommandList(title string, cmds []types.CommandList) selectedUUIDState {

	var items []list.Item

	for _, cmd := range cmds {
		items = append(items, item{
			uuid:  cmd.UUID,
			title: cmd.Cmd,
			desc:  cmd.Desc,
		})
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = title
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
	return selectedUUID.Get()
}
