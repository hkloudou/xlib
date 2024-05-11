package color

// Attribute defines a single SGR Code
type attribute int

// Foreground text colors
const (
	fgBlack attribute = iota + 30
	fgRed
	fgGreen
	fgYellow
	fgBlue
	fgMagenta
	fgCyan
	fgWhite
)

// Foreground Hi-Intensity text colors
const (
	fgHiBlack attribute = iota + 90
	fgHiRed
	fgHiGreen
	fgHiYellow
	fgHiBlue
	fgHiMagenta
	fgHiCyan
	fgHiWhite
)

// Background text colors
const (
	bgBlack attribute = iota + 40
	bgRed
	bgGreen
	bgYellow
	bgBlue
	bgMagenta
	bgCyan
	bgWhite
)

// Background Hi-Intensity text colors
const (
	bgHiBlack attribute = iota + 100
	bgHiRed
	bgHiGreen
	bgHiYellow
	bgHiBlue
	bgHiMagenta
	bgHiCyan
	bgHiWhite
)
