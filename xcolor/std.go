package xcolor

var (
	Output = NewColorableStdout()

	// Error defines a color supporting writer for os.Stderr.
	Error = NewColorableStderr()
)
