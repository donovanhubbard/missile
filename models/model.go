/*
Copyright © 2023 Donovan Hubbard
*/

package models

import (
	"strings"
	tea "github.com/charmbracelet/bubbletea"
  "github.com/donovanhubbard/missile/models/commandinput"
  "github.com/donovanhubbard/missile/models/commandhistory"
  "github.com/donovanhubbard/missile/models/serverlist"
	"github.com/charmbracelet/lipgloss"
)

var (
	purple = lipgloss.Color("13")
	green  = lipgloss.Color("10")
)

type Model struct {
  commandInput tea.Model
  commandHistory tea.Model
  serverList tea.Model
}

func New(hosts []string) Model {
  width := 35
	commandInput := commandinput.New(width)
  commandHistory := commandhistory.New(10,width+3)
  commandHistory.AddText("Penn")
  commandHistory.AddText("Teller")
  serverList := serverlist.New(hosts)

	return Model{
		commandInput: commandInput,
    commandHistory: commandHistory,
    serverList: serverList,
	}
}

func (m Model) Init() tea.Cmd {
  m.commandInput.Init()
  m.commandHistory.Init()
  m.serverList.Init()
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  var cmds []tea.Cmd
  var cmd tea.Cmd

  commandInput, cmd := m.commandInput.Update(msg)
  cmds = append(cmds, cmd)
  commandHistory, cmd := m.commandHistory.Update(msg)
  cmds = append(cmds, cmd)
  serverList, cmd := m.serverList.Update(msg)
  cmds = append(cmds, cmd)

  m.commandInput = commandInput
  m.commandHistory = commandHistory
  m.serverList = serverList
  return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	defaultStyle := lipgloss.NewStyle().
		Foreground(green).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(purple)
  header := defaultStyle.SetString("missile").String()
  commandHistoryString := m.commandHistory.View()
  commandInputString := m.commandInput.View()
  serverListString := m.serverList.View()
  serverHeader := defaultStyle.SetString("Servers").String()

	text := lipgloss.JoinVertical(lipgloss.Center, header, commandHistoryString, commandInputString, )
  serverListColumn := lipgloss.JoinVertical(lipgloss.Bottom, serverHeader, serverListString)
  text = lipgloss.JoinHorizontal(lipgloss.Top, text, serverListColumn)

	var b strings.Builder
	b.WriteString(text)
	b.WriteString("\n")
	return b.String()
}
