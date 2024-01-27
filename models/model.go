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
  commandInput commandinput.Model
  commandHistory commandhistory.Model
  serverList serverlist.Model
  height int
  width int
}

func New(hosts []string) Model {
  width := 35
	commandInput := commandinput.New(width)
  commandHistory := commandhistory.New(10,width+3, 15)
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

  switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
    m.commandInput.Width = m.width - 30
    m.commandHistory.Width = m.width - 30
    m.commandHistory.Height = m.height - 10
    m.serverList.Height = m.height - 10
  case tea.KeyMsg:
    switch msg.String() {
    case "ctrl+c":
      return m, tea.Quit
    case "enter":
      userInput := m.commandInput.Value()

      if userInput == "exit" || userInput == "quit" {
        return m, tea.Quit
      }

      m.commandInput = m.commandInput.Reset()
      ct := commandhistory.CommandText{ Text: userInput, Type: commandhistory.UserInput }
      m.commandHistory = m.commandHistory.AddCommandText(ct)
      ct = m.processCommand(userInput)
      m.commandHistory = m.commandHistory.AddCommandText(ct)
    }
  }

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

func (m Model) processCommand(command string) commandhistory.CommandText {
  return commandhistory.CommandText{Text:"Unrecognized command",Type:commandhistory.FailureResponse}
}
