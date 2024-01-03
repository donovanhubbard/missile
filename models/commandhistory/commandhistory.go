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
  text []string
  maxSize int
  width int
}

func New(size int, width int) Model {
  return Model {
    maxSize: size,
    width: width,
  }
}

func (m *Model) AddText(newText string){
  m.text = append(m.text, newText)
}

func (m Model) Init() tea.Cmd {
	return nil
}


func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  return m, nil
}


func (m Model) View() string {
  return lipgloss.NewStyle().
    SetString(strings.Join(m.text, "\n")).
    Foreground(green).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(purple).
    Width(m.width).
		String()
}
