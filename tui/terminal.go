package tui

import (
	"os"
	"strings"
)

func (m model) termRender() model {
	var c string
	for _, m := range m.terminalMessages {
		if m.termEntity == termEntityUser {
			c += "> " + m.msg + "\n"
		} else if m.termEntity == termEntityBot {
			c += "🤖 " + m.msg + "\n"
		} else {
			c += m.msg + "\n"
		}
	}
	m.terminal.SetContent(c)
	m.input.SetValue("")
	return m
}

func (m model) termProcess() model {
	i := m.input.Value()
	// trim the command
	i = strings.TrimSpace(i)

	switch i {
	case "":
		break
	case "help":
		m.terminalMessages = append(m.terminalMessages,
			terminalMessage{termEntityUser, "help"},
			terminalMessage{termEntityClass, termHelp()},
		)
	case "clear":
		m.terminalMessages = []terminalMessage{}
	case "exit":
		os.Exit(0)
	default:
		m.terminalMessages = append(m.terminalMessages, terminalMessage{termEntityUser, i})
	}
	return m.termRender()
}

func termHelp() string {
	return "Placeholder for showHelper"
}
