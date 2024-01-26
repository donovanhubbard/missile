/*
Copyright © 2023 Donovan Hubbard
*/

package serverlist

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

var (
	purple = lipgloss.Color("13")
	green  = lipgloss.Color("10")
)

type Model struct {
  servers []string
  Height int
}

func New(serverNames []string)Model{
  return Model {
    servers: serverNames,
  }
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
  return m, nil
}

func (m Model) View() string {
  return lipgloss.NewStyle().
    SetString(strings.Join(m.servers, "\n")).
    Foreground(green).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(purple).
    Height(m.Height).
		String()
}

