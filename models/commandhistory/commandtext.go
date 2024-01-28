package commandhistory

import (
  "github.com/muesli/reflow/wordwrap"
  "github.com/muesli/reflow/wrap"
  "github.com/charmbracelet/lipgloss"
  "strings"
)

const (
  UserInput = "UserInput"
  SuccessResponse = "SuccessResponse"
  FailureResponse = "FailureResponse"
)

var (
  red = lipgloss.Color("9")
  yellow = lipgloss.Color("227")
)

type CommandText struct {
  Text string
  Host string
  Type string
}

func (ct CommandText) render(width int) []string {
  var sb strings.Builder
  var currentStyle lipgloss.Style

  userInputStyle := lipgloss.NewStyle().Foreground(green)
  failureResponseStyle := lipgloss.NewStyle().Foreground(red)
  successResponseStyle := lipgloss.NewStyle().Foreground(yellow)

  switch ct.Type {
    case UserInput:
      sb.WriteString("> ")
      currentStyle = userInputStyle
    case FailureResponse:
      currentStyle = failureResponseStyle
    case SuccessResponse:
      currentStyle = successResponseStyle
  }

  sb.WriteString(ct.Text)
  result := sb.String()

  //wordwrap will break up by words
  //wrap will break up by letters ignoring spaces.
  //We want to split by words first then by characters
  wordWrapped := strings.Split(wordwrap.String(result, width), "\n")
  var unformattedLines []string
  for _, line := range wordWrapped {
    unformattedLines = append(unformattedLines, strings.Split(wrap.String(line,width),"\n")...)
  }

  formattedLines := make([]string, len(unformattedLines))

  for i, line := range unformattedLines {
    formattedLines[i] = currentStyle.Copy().SetString(line).String()
  }
  return formattedLines
}

