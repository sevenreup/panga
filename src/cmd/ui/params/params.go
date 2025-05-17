package params

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sevenreup/panga/src/pkg/engine"
)

type model struct {
	params     []engine.Param
	cursor     int
	inputs     []string
	current    string
	confirming bool
}

func initialModel(params []engine.Param) model {
	return model{
		params:  params,
		cursor:  0,
		inputs:  make([]string, len(params)),
		current: "",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if m.confirming {
				return m, tea.Quit
			}
			m.inputs[m.cursor] = m.current
			m.current = ""
			m.cursor++
			if m.cursor >= len(m.params) {
				m.confirming = true
			}
		case "backspace":
			if len(m.current) > 0 {
				m.current = m.current[:len(m.current)-1]
			}
		default:
			m.current += msg.String()
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.confirming {
		return fmt.Sprintf("Summary: %v\nPress 'q' to quit.\n", m.inputs)
	}
	param := m.params[m.cursor]
	return fmt.Sprintf(
		"%s: %s\nCurrent input: %s\n",
		param.Name,
		param.Description,
		m.current,
	)
}
