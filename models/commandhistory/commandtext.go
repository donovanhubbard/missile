package commandhistory

import (
  "strings"
  "log"
)

const (
  UserInput = "UserInput"
  SuccessResponse = "SuccessResponse"
  FailureResponse = "FailureResponse"
)

type CommandText struct {
  Text string
  Host string
  Type string
}

func (ct CommandText) render() string {
  var sb strings.Builder

  switch ct.Type {
    case UserInput:
      sb.WriteString("> ")
  }

  sb.WriteString(ct.Text)
  result := sb.String()
  log.Print("CommandText.render() " + result)
  return sb.String()
}

