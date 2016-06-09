package color

import (
  "github.com/fatih/color"
)

// type of a func to colorize text
type ColorFunc func (format string, a ...interface{}) string

// list all known color func colorizer
func getColors () []ColorFunc {
  return []ColorFunc{
    color.CyanString,
    color.YellowString,
    color.GreenString,
    color.MagentaString,
    color.RedString,
    color.BlueString,
    color.WhiteString,
  }
}

// store al known colors once
var colorsFunc = getColors()
// store the last color assigned
var currentColor = 0
// pick a new color for a logger
func PickColor () ColorFunc {
  ret := colorsFunc[currentColor]
  currentColor++
  if currentColor>len(colorsFunc) {
    currentColor = 0
  }
  return ret
}
