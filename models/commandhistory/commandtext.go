package commandhistory

import (
  "strings"
  "math"
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

func (ct CommandText) render(width int) []string {
  var sb strings.Builder

  switch ct.Type {
    case UserInput:
      sb.WriteString("> ")
  }

  sb.WriteString(ct.Text)
  result := sb.String()

  return divideIntoLines(result, width)
}

func divideIntoLines(text string, width int) []string{
   var floatText, floatWidth float64
   var linesNeeded, currentLineNumber, currentIndex, endOfLine int
   floatText = float64(len(text))
   floatWidth = float64(width)

   linesNeeded = int(math.Ceil(floatText / floatWidth))

   var lines []string
   lines = make([]string,linesNeeded)
   currentLineNumber = 0
   currentIndex = 0

   for currentLineNumber < linesNeeded {
     if currentIndex + width < len(text) {
       endOfLine = currentIndex + width
     }else {
       endOfLine = len(text) - 1
     }
     lines[currentLineNumber] = text[currentIndex:endOfLine]
     currentLineNumber++
     currentIndex += width
   }
   return lines
}

