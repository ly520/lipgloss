package lipgloss

import (
	"strings"

	"github.com/muesli/reflow/ansi"
	"github.com/muesli/termenv"
)

// Border contains a series of values which comprise the various parts of a
// border.
type Border struct {
	Top         string
	Bottom      string
	Left        string
	Right       string
	TopLeft     string
	TopRight    string
	BottomRight string
	BottomLeft  string
}

var (
	noBorder = Border{}

	normalBorder = Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "┌",
		TopRight:    "┐",
		BottomLeft:  "└",
		BottomRight: "┘",
	}

	roundedBorder = Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "╰",
		BottomRight: "╯",
	}

	thickBorder = Border{
		Top:         "━",
		Bottom:      "━",
		Left:        "┃",
		Right:       "┃",
		TopLeft:     "┏",
		TopRight:    "┓",
		BottomLeft:  "┗",
		BottomRight: "┛",
	}

	doubleBorder = Border{
		Top:         "═",
		Bottom:      "═",
		Left:        "║",
		Right:       "║",
		TopLeft:     "╔",
		TopRight:    "╗",
		BottomLeft:  "╚",
		BottomRight: "╝",
	}
)

// NormalBorder returns a standard-type border with a normal weight and 90
// degree corners.
func NormalBorder() Border {
	return normalBorder
}

// RoundedBorder returns a border with rounded corners.
func RoundedBorder() Border {
	return roundedBorder
}

// Thick border returns a border that's thicker than the one returned by
// NormalBorder.
func ThickBorder() Border {
	return thickBorder
}

// DoubleBorder returns a border comprised of two thin strokes.
func DoubleBorder() Border {
	return doubleBorder
}

func (s Style) applyBorder(str string) string {
	var (
		topSet    = s.isSet(borderTopKey)
		rightSet  = s.isSet(borderRightKey)
		bottomSet = s.isSet(borderBottomKey)
		leftSet   = s.isSet(borderLeftKey)

		border    = s.getAsBorderStyle(borderStyleKey)
		hasTop    = s.getAsBool(borderTopKey, false)
		hasRight  = s.getAsBool(borderRightKey, false)
		hasBottom = s.getAsBool(borderBottomKey, false)
		hasLeft   = s.getAsBool(borderLeftKey, false)

		topFGColor    = s.getAsColor(borderTopFGColorKey)
		rightFGColor  = s.getAsColor(borderRightFGColorKey)
		bottomFGColor = s.getAsColor(borderBottomFGColorKey)
		leftFGColor   = s.getAsColor(borderLeftFGColorKey)

		topBGColor    = s.getAsColor(borderTopBGColorKey)
		rightBGColor  = s.getAsColor(borderRightBGColorKey)
		bottomBGColor = s.getAsColor(borderBottomBGColorKey)
		leftBGColor   = s.getAsColor(borderLeftBGColorKey)
	)

	// If a border is set and no sides have been specifically turned on or off
	// render borders on all sides.
	// 如果设置了边界，并且没有专门打开或关闭任何边，则在所有边上渲染边界
	if border != noBorder && !(topSet || rightSet || bottomSet || leftSet) {
		hasTop = true
		hasRight = true
		hasBottom = true
		hasLeft = true
	}

	// If no border is set or all borders are been disabled, abort.
	// 如果未设置边框或禁用了所有边框，则中止。
	if border == noBorder || (!hasTop && !hasRight && !hasBottom && !hasLeft) {
		return str
	}

	lines, width := getLines(str)

	if hasLeft {
		width += ansi.PrintableRuneWidth(border.Left)
	}

	// Figure out which corners we should actually be using based on which
	// sides are set to show.
	// 根据要显示的边，找出我们实际应该使用的角。
	if hasTop {
		switch {
		case !hasLeft && !hasRight:
			border.TopLeft = ""
			border.TopRight = ""
		case !hasLeft:
			border.TopLeft = ""
		case !hasRight:
			border.TopRight = ""
		}
	}
	if hasBottom {
		switch {
		case !hasLeft && !hasRight:
			border.BottomLeft = ""
			border.BottomRight = ""
		case !hasLeft:
			border.BottomLeft = ""
		case !hasRight:
			border.BottomRight = ""
		}
	}

	var out strings.Builder

	// Render top
	if hasTop {
		top := renderHorizontalEdge(border.TopLeft, border.Top, border.TopRight, width)
		top = styleBorder(top, topFGColor, topBGColor)
		out.WriteString(top)
		out.WriteRune('\n')
	}

	// Render sides
	for i, l := range lines {
		if hasLeft {
			out.WriteString(styleBorder(border.Left, leftFGColor, leftBGColor))
		}
		out.WriteString(l)
		if hasRight {
			out.WriteString(styleBorder(border.Right, rightFGColor, rightBGColor))
		}
		if i < len(lines)-1 {
			out.WriteRune('\n')
		}
	}

	// Render bottom
	if hasBottom {
		bottom := renderHorizontalEdge(border.BottomLeft, border.Bottom, border.BottomRight, width)
		bottom = styleBorder(bottom, bottomFGColor, bottomBGColor)
		out.WriteRune('\n')
		out.WriteString(bottom)
	}

	return out.String()
}

// Render the horizontal (top or bottom) portion of a border.
// 渲染边框的水平（顶部或底部）部分。
func renderHorizontalEdge(left, middle, right string, width int) string {
	if width < 1 {
		return ""
	}

	if middle == "" {
		middle = " "
	}

	leftWidth := ansi.PrintableRuneWidth(left)
	midWidth := ansi.PrintableRuneWidth(middle)
	rightWidth := ansi.PrintableRuneWidth(right)

	out := strings.Builder{}
	out.WriteString(left)
	for i := leftWidth + rightWidth; i < width+rightWidth; i += midWidth {
		out.WriteString(middle)
	}
	out.WriteString(right)

	return out.String()
}

// Apply foreground and background styling to a border.
// 渲染边框的前景和背景色
func styleBorder(border string, fg, bg TerminalColor) string {
	if fg == noColor && bg == noColor {
		return border
	}

	var style = termenv.Style{}

	if fg != noColor {
		style = style.Foreground(ColorProfile().Color(fg.value()))
	}
	if bg != noColor {
		style = style.Background(ColorProfile().Color(bg.value()))
	}

	return style.Styled(border)
}
