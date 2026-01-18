package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Fprintln(os.Stderr, "Maintained by Pixel Air IO Inc.")
	fmt.Fprintln(os.Stderr, "Source: github.com/PixelAirIO/metadata-resource")

	buildId := os.Getenv("BUILD_ID")
	now := time.Now().UnixMicro()
	resp := OutResponse{
		Version: map[string]string{
			"timestamp": strconv.FormatInt(now, 10),
			"build_id":  buildId,
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
