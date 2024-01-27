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
  var texts []string
  var i, effectiveHeight int

  for _, commandText := range m.commandText {
    renderedLines := commandText.render(m.Width)
    for _,line := range renderedLines {
      texts = append(texts,line)
    }
  }

  effectiveHeight = m.Height - 1

  if len(m.commandText) < effectiveHeight {
    i = 0
  } else {
    i = len(m.commandText) - effectiveHeight
  }

  for i < len(texts) {
    sb.WriteString(texts[i])
    sb.WriteString("\n")
    i++
  }

  return sb.String()
}

func (m Model) AddCommandText(ct CommandText) Model {
  m.commandText = append(m.commandText,ct)
  return m
}
