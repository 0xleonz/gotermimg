package main

import (
	"fmt"
	"gitlab.com/0xleonz/gotermimg/pkg/protocol"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Uso: gotermimg <ruta/a/imagen>\n")
		os.Exit(1)
	}
	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error al abrir el archivo: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error al decodificar la imagen: %v\n", err)
		os.Exit(1)
	}

	renderer := &protocol.IKatRenderer{}

	if !renderer.Supports() {
		fmt.Fprintf(os.Stderr, "Error: La terminal no soporta el protocolo Kitty.\n")
		os.Exit(1)
	}

	// Rendering
	output, err := renderer.Render(img)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error al renderizar la imagen: %v\n", err)
		os.Exit(1)
	}

	_, err = os.Stdout.WriteString(output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error al escribir la salida: %v\n", err)
		os.Exit(1)
	}
}
