package color

import (
	"os"
)

// Color defines a custom color object which is defined by SGR parameters.
type ColorParame struct {
	params  []attribute
	noColor *bool
}

// noColorIsSet returns true if the environment variable NO_COLOR is set to a non-empty string.
func noColorIsSet() bool {
	return os.Getenv("NO_COLOR") != ""
}

// New returns a newly created color object.
func New(value ...attribute) *ColorParame {
	c := &ColorParame{
		params: make([]attribute, 0),
	}

	if noColorIsSet() {
		c.noColor = boolPtr(true)
	}

	c.Add(value...)
	return c
}

func boolPtr(v bool) *bool {
	return &v
}

// Add is used to chain SGR parameters. Use as many as parameters to combine
// and create custom color objects. Example: Add(color.FgRed, color.Underline).
func (c *ColorParame) Add(value ...attribute) *ColorParame {
	c.params = append(c.params, value...)
	return c
}

// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
// On Windows, users should wrap w with colorable.NewColorable() if w is of
// type *os.File.
// func (c *ColorParame) Fprint(w io.Writer, a ...interface{}) (n int, err error) {
// 	c.SetWriter(w)
// 	defer c.UnsetWriter(w)

// 	return fmt.Fprint(w, a...)
// }

// // Sprint is just like Print, but returns a string instead of printing it.
// func (c *ColorParame) Sprint(a ...interface{}) string {
// 	return c.wrap(fmt.Sprint(a...))
// }

// // wrap wraps the s string with the colors attributes. The string is ready to
// // be printed.
// func (c *ColorParame) wrap(s string) string {
// 	if c.isNoColorSet() {
// 		return s
// 	}

// 	return c.format() + s + c.unformat()
// }
