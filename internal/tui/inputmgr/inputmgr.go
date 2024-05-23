package inputmgr

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func ShowInteractiveArgumentPlaceholdersFiller(placeholders []string) []string {
	ch := make(chan []string)
	defer close(ch)

	p := tea.NewProgram(initialModel(placeholders, ch))

	go func() {
		if _, err := p.Run(); err != nil {
			fmt.Printf("could not start program: %s\n", err)
			os.Exit(1)
		}
	}()

	// Wait for the filled arguments
	filledArg := <-ch
	p.Wait() // wait until tea program finished shutting down (MUST)-<otherwise the tea's session will be stuck in the terminal session after process ended>
	return filledArg

}
