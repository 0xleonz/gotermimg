package protocol

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
)

type KittyRenderer struct{}

func (k *KittyRenderer) Supports() bool {
	// TODO: Implement chek for support of the portocol
	return true
}

func (k *KittyRenderer) Render(img image.Image) (string, error) {
	var imgBuf bytes.Buffer

	if err := png.Encode(&imgBuf, img); err != nil {
		return "", fmt.Errorf("error al codificar la imagen a PNG: %w", err)
	}

	base64Data := base64.RawStdEncoding.EncodeToString(imgBuf.Bytes())

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	transferCmd := fmt.Sprintf(
		"\x1b_Gf=100,t=p,s=%d,v=%d,a=t,i=1;%s\a",
		width,
		height,
		base64Data,
	)

	// DrawCmd (a=d)
	// ESC _ G a=d; ESC \
	drawCmd := "\x1b_Ga=d;\x1b\\"

	return transferCmd + drawCmd, nil
}
