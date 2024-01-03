/*
Copyright © 2023 Donovan Hubbard
*/

package models

import (
	"strings"
	tea "github.com/charmbracelet/bubbletea"
  "github.com/donovanhubbard/missile/models/commandinput"
  "github.com/donovanhubbard/missile/models/commandhistory"
	"github.com/charmbracelet/lipgloss"
)

var (
	purple = lipgloss.Color("13")
	green  = lipgloss.Color("10")
)

type Model struct {
  commandInput tea.Model
  commandHistory tea.Model
}

func New(hosts []string) Model {
  width := 35
	commandInput := commandinput.New(width)
  commandHistory := commandhistory.New(10,width+3)
  commandHistory.AddText("Penn")
  commandHistory.AddText("Teller")

	return Model{
		commandInput: commandInput,
    commandHistory: commandHistory,
	}
}

func (m Model) Init() tea.Cmd {
  m.commandInput.Init()
  m.commandHistory.Init()
	return m.commandInput.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmds []tea.Cmd
  var cmd tea.Cmd

  commandInput, cmd := m.commandInput.Update(msg)
  cmds = append(cmds, cmd)
  commandHistory, cmd := m.commandHistory.Update(msg)
  cmds = append(cmds, cmd)

  m.commandInput = commandInput
  m.commandHistory = commandHistory
  return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	header := lipgloss.NewStyle().
		SetString("header").
		Foreground(green).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(purple).
		String()
  commandHistoryString := m.commandHistory.View()
  commandInputString := m.commandInput.View()

	text := lipgloss.JoinVertical(lipgloss.Center, header, commandHistoryString, commandInputString, )

	var b strings.Builder
	b.WriteString(text)
	b.WriteString("\n")
	return b.String()
}
