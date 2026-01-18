package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stderr, "Maintained by Pixel Air IO Inc.")
	fmt.Fprintln(os.Stderr, "Source: github.com/PixelAirIO/metadata-resource")
	fmt.Fprintln(os.Stderr, "The check step does not emit any versions and cannot be used to trigger jobs. Use the resource as a put step in your jobs.")
	fmt.Fprint(os.Stdout, "[]")
}
