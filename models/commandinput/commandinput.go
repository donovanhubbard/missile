/*
Copyright © 2023 Donovan Hubbard
*/

package commandinput

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

var (
	purple = lipgloss.Color("13")
	green  = lipgloss.Color("10")
)

type Model struct {
	commandInput textinput.Model
}

func New() Model {
	commandInput := textinput.New()
	// commandInput.Placeholder = "Enter command"
	commandInput.Focus()
	commandInput.CharLimit = 256
  //TODO make this width dynamicly sized
	commandInput.Width = 25
	return Model{
		commandInput: commandInput,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
		m.commandInput, cmd = m.commandInput.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m Model) View() string {
	commandInputString := lipgloss.NewStyle().
		SetString(m.commandInput.View()).
		Foreground(green).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(purple).
		String()

	var b strings.Builder
	b.WriteString(commandInputString)
	b.WriteString("\n")
	return b.String()
}
