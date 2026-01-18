package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type buildMetadata struct {
	BuildId       string         `json:"build_id"`
	BuildName     string         `json:"build_name"`
	JobName       string         `json:"job_name"`
	PipelineName  string         `json:"pipeline_name"`
	InstanceVars  map[string]any `json:"instance_vars"`
	TeamName      string         `json:"team_name"`
	CreatedBy     string         `json:"created_by"`
	ExternalUrl   string         `json:"external_url"`
	BuildUrl      string         `json:"build_url"`
	BuildUrlShort string         `json:"build_url_short"`
}

type InRequest struct {
	Version map[string]string `json:"version"`
}

func main() {
	fmt.Fprintln(os.Stderr, "Maintained by Pixel Air IO Inc.")
	fmt.Fprintln(os.Stderr, "Source: github.com/PixelAirIO/metadata-resource")

	req := InRequest{}
	dc := json.NewDecoder(os.Stdin)
	err := dc.Decode(&req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "json decode:", err.Error())
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "destination path not specified")
		os.Exit(1)
	}
	dest := os.Args[1]

	resp := InResponse{}
	resp.Version = req.Version

	bm := buildMetadata{}
	bm.BuildId = os.Getenv("BUILD_ID")
	resp.Metadata = append(resp.Metadata, MetadataField{
		Name:  "build_id",
		Value: bm.BuildId,
	})
	bm.BuildName = os.Getenv("BUILD_NAME")
	resp.Metadata = append(resp.Metadata, MetadataField{
		Name:  "build_name",
		Value: bm.BuildName,
	})
	bm.JobName = os.Getenv("BUILD_JOB_NAME")
	resp.Metadata = append(resp.Metadata, MetadataField{
		Name:  "job_name",
		Value: bm.JobName,
	})
	bm.PipelineName = os.Getenv("BUILD_PIPELINE_NAME")
	resp.Metadata = append(resp.Metadata, MetadataField{
		Name:  "pipeline_name",
		Value: bm.PipelineName,
	})
	bm.TeamName = os.Getenv("BUILD_TEAM_NAME")
	resp.Metadata = append(resp.Metadata, MetadataField{
		Name:  "team_name",
		Value: bm.TeamName,
	})
	bm.CreatedBy = os.Getenv("BUILD_CREATED_BY")
	resp.Metadata = append(resp.Metadata, MetadataField{
		Name:  "created_by",
		Value: bm.CreatedBy,
	})
	bm.ExternalUrl = os.Getenv("ATC_EXTERNAL_URL")
	resp.Metadata = append(resp.Metadata, MetadataField{
		Name:  "external_url",
		Value: bm.ExternalUrl,
	})
	bm.BuildUrl = os.Getenv("BUILD_URL")
	resp.Metadata = append(resp.Metadata, MetadataField{
		Name:  "build_url",
		Value: bm.BuildUrl,
	})
	bm.BuildUrlShort = os.Getenv("BUILD_URL_SHORT")
	resp.Metadata = append(resp.Metadata, MetadataField{
		Name:  "build_url_short",
		Value: bm.BuildUrlShort,
	})

	instanceVars := os.Getenv("BUILD_PIPELINE_INSTANCE_VARS")
	if instanceVars != "" {
		err = json.Unmarshal([]byte(instanceVars), &bm.InstanceVars)
		resp.Metadata = append(resp.Metadata, MetadataField{
			Name:  "instance_vars",
			Value: instanceVars,
		})
	}

	id, ok := req.Version["build_id"]
	if ok && id != bm.BuildId {
		fmt.Fprintln(os.Stderr, "build_id in version does not match current build ID from env $BUILD_ID")
		os.Exit(1)
	}

	content, err := json.Marshal(bm)
	if err != nil {
		fmt.Fprintln(os.Stderr, "json marshal:", err.Error())
		os.Exit(1)
	}

	err = os.WriteFile(filepath.Join(dest, "build.json"), content, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error writing to build.json:", err.Error())
		os.Exit(1)
	}
	fmt.Fprintln(os.Stderr, "build metadata saved to build.json")

	err = json.NewEncoder(os.Stdout).Encode(resp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "json encoding:", err.Error())
		os.Exit(1)
	}
}

type InResponse struct {
	Version  map[string]string `json:"version"`
	Metadata []MetadataField   `json:"metadata"`
}

type MetadataField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
