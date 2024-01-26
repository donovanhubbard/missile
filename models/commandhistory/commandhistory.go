/*
Copyright © 2023 Donovan Hubbard
*/

package commandhistory

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
  commandText []CommandText
  maxSize int
  Width int
  Height int
}

func New(size int, width int, height int) Model {
  return Model {
    maxSize: size,
    Width: width,
    Height: height,
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
    SetString(m.renderText()).
    Foreground(green).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(purple).
    Width(m.Width).
    Height(m.Height).
		String()
}

func (m Model) renderText() string {
  var sb strings.Builder
  for _, commandText := range m.commandText {
    switch commandText.Type {
    case UserInput:
      sb.WriteString("> ")
    }
    sb.WriteString(commandText.Text)
    sb.WriteString("\n")
  }
  return sb.String()
}

func (m Model) AddCommandText(ct CommandText) Model {
  m.commandText = append(m.commandText,ct)
  return m
}
