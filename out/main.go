package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stderr, "Maintained by Pixel Air IO Inc.")
	fmt.Fprintln(os.Stderr, "Source: github.com/PixelAirIO/metadata-resource")

	buildId := os.Getenv("BUILD_ID")
	if buildId == "" {
		fmt.Fprintln(os.Stderr, "no $BUILD_ID found in env")
		os.Exit(1)
	}
	resp := OutResponse{
		Version: map[string]string{
			"build_id": buildId,
		},
	}

	err := json.NewEncoder(os.Stdout).Encode(resp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "json encoding:", err.Error())
		os.Exit(1)
	}

}

type OutResponse struct {
	Version map[string]string `json:"version"`
}
