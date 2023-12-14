/*
Copyright © 2023 Donovan Hubbard
*/

package models

import (
	"strings"
	tea "github.com/charmbracelet/bubbletea"
  "github.com/donovanhubbard/missile/models/commandinput"
	"github.com/charmbracelet/lipgloss"
)

var (
	purple = lipgloss.Color("13")
	green  = lipgloss.Color("10")
)

type Model struct {
  commandInput tea.Model
}

func New(hosts []string) Model {
	commandInput := commandinput.New()
	return Model{
		commandInput: commandInput,
	}
}

func (m Model) Init() tea.Cmd {
	return m.commandInput.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  commandInput, commandInputCmd := m.commandInput.Update(msg)

  m.commandInput = commandInput
  return m, commandInputCmd
}

func (m Model) View() string {
	header := lipgloss.NewStyle().
		SetString("header").
		Foreground(green).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(purple).
		String()
  commandInputString := m.commandInput.View()

	text := lipgloss.JoinVertical(lipgloss.Center, header, commandInputString)

	var b strings.Builder
	b.WriteString(text)
	b.WriteString("\n")
	return b.String()
}
