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
  "github.com/bradfitz/gomemcache/memcache"
)

var (
	purple = lipgloss.Color("13")
	green  = lipgloss.Color("10")
)

type Model struct {
  commandInput commandinput.Model
  commandHistory commandhistory.Model
  serverListPane serverlist.Model
  mc *memcache.Client
  serverList *memcache.ServerList
  height int
  width int
}

func New(hosts []string) Model {
  width := 35
	commandInput := commandinput.New(width)
  commandHistory := commandhistory.New(10,width+3, 15)
  serverListPane := serverlist.New(hosts)
  mc := memcache.New(hosts...)
  serverList := &memcache.ServerList{}
  serverList.SetServers(hosts...)


	return Model{
		commandInput: commandInput,
    commandHistory: commandHistory,
    serverListPane: serverListPane,
    mc: mc,
    serverList: serverList,
	}
}

func (m Model) Init() tea.Cmd {
  m.commandInput.Init()
  m.commandHistory.Init()
  m.serverListPane.Init()
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
    m.serverListPane.Height = m.height - 10
  case tea.KeyMsg:
    switch msg.String() {
    case "ctrl+c":
      return m, tea.Quit
    case "enter":
      userInput := m.commandInput.Value()

      if userInput == "exit" || userInput == "quit" || userInput == "q" {
        return m, tea.Quit
      }

      m.commandInput.Reset()
      ct := commandhistory.CommandText{ Text: userInput, Type: commandhistory.UserInput }
      m.commandHistory.AddCommandText(ct)
      ct = m.processCommand(userInput)
      m.commandHistory.AddCommandText(ct)
    }
  }

  commandInput, cmd := m.commandInput.Update(msg)
  cmds = append(cmds, cmd)
  commandHistory, cmd := m.commandHistory.Update(msg)
  cmds = append(cmds, cmd)
  serverListPane, cmd := m.serverListPane.Update(msg)
  cmds = append(cmds, cmd)

  m.commandInput = commandInput
  m.commandHistory = commandHistory
  m.serverListPane = serverListPane
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
  serverListPaneString := m.serverListPane.View()
  serverHeader := defaultStyle.SetString("Servers").String()

	text := lipgloss.JoinVertical(lipgloss.Center, header, commandHistoryString, commandInputString, )
  serverListPaneColumn := lipgloss.JoinVertical(lipgloss.Bottom, serverHeader, serverListPaneString)
  text = lipgloss.JoinHorizontal(lipgloss.Top, text, serverListPaneColumn)

	var b strings.Builder
	b.WriteString(text)
	b.WriteString("\n")
	return b.String()
}

func (m Model) processCommand(command string) commandhistory.CommandText {
  var ch commandhistory.CommandText
  words := strings.Fields(command)

  switch words[0]{
  case "set":
    ch = m.processSet(words)
  case "get":
    ch = m.processGet(words)
  default:
    ch = commandhistory.CommandText{Text:"Unrecognized command",Type:commandhistory.FailureResponse}
  }

  return ch
}

func (m *Model) processSet(words []string) commandhistory.CommandText {
  var ch commandhistory.CommandText
  var sb strings.Builder

  if len(words) < 3 {
    return commandhistory.CommandText{Text: "usage: set <key> <value>",Type:commandhistory.FailureResponse}
  }

  key := words[1]
  value := strings.Join(words[2:]," ")
  server, error := m.serverList.PickServer(key)
  if error == nil {
    sb.WriteString(server.String())
  } else {
    sb.WriteString("Unknown")
  }
  error = m.mc.Set(&memcache.Item{Key: words[1], Value: []byte(value)})
  if error == nil {
    sb.WriteString(" OK")
    ch = commandhistory.CommandText{Text:sb.String(),Type:commandhistory.SuccessResponse}
  } else {
    sb.WriteString(" ERROR: ")
    sb.WriteString(error.Error())
    ch = commandhistory.CommandText{Text:sb.String(),Type:commandhistory.FailureResponse}
  }
  return ch
}

func (m *Model) processGet(words []string) commandhistory.CommandText {
  var ch commandhistory.CommandText
  var sb strings.Builder

  if len(words) != 2 {
    return commandhistory.CommandText{Text: "usage: get <key>",Type:commandhistory.FailureResponse}
  }

  key := words[1]

  server, error := m.serverList.PickServer(key)
  if error == nil {
    sb.WriteString(server.String())
  } else {
    sb.WriteString("Unknown")
  }

  item, error := m.mc.Get(key)
  if error == nil {
    sb.WriteString(" ")
    sb.WriteString(string(item.Value))
    ch = commandhistory.CommandText{Text:sb.String(),Type:commandhistory.SuccessResponse}
  } else {
    sb.WriteString(" ERROR: ")
    sb.WriteString(error.Error())
    ch = commandhistory.CommandText{Text:sb.String(),Type:commandhistory.FailureResponse}
  }

  return ch
}
