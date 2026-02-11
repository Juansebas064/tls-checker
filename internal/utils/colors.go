package utils

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
)

// Colors for main components
var (
	ColorPrimary = tcell.Color69
	ColorText    = tcell.ColorWhite
	ColorOk      = tcell.ColorGreen
	ColorWarning = tcell.ColorYellow
	ColorError   = tcell.ColorRed
	ColorLabel   = tcell.ColorNames["cadetblue"]
)

// Colors for details output
const (
	TagPrimary = "#5f87ff"
	TagText    = "white"
	TagOk      = "green"
	TagWarning = "yellow"
	TagBad     = "orange"
	TagError   = "red"
	TagLabel   = "#87afaf"
)

// Returns the color according to grade for color tags
func gradeColor(grade string) string {
	switch {
	case strings.HasPrefix(grade, "A"):
		return TagOk
	case grade == "B":
		return TagWarning
	case grade == "C" || grade == "D":
		return TagBad
	case grade == "F" || grade == "T" || grade == "M":
		return TagError
	default:
		return TagText
	}
}

// Returns the color for suiteStrength based on its values
func suiteStrengthColor(bits int) string {
	switch {
	case bits >= 128:
		return fmt.Sprintf("[%s]", TagOk)
	case bits >= 112:
		return fmt.Sprintf("[%s]", TagWarning)
	default:
		return fmt.Sprintf("[%s]", TagError)
	}
}

