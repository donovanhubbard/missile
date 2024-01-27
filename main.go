/*
Copyright © 2023 Donovan Hubbard
*/
package main

import (
  "github.com/donovanhubbard/missile/cmd"
  tea "github.com/charmbracelet/bubbletea"
  "os"
  "fmt"
  "log"
)


func main() {
  f, err := tea.LogToFile("debug.log", "debug")
  if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
  }
  defer f.Close()
  log.Print("Starting program")
	cmd.Execute()
  log.Print("Terminating program")
}
