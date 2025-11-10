package protocol

import (
	"image"
)

// All protocols should check this
type Renderer interface {
	// Checks support of the terminal
	Supports() bool
	// Takes an images and returns a list of bytes
	Render(img image.Image) (string, error)
}
