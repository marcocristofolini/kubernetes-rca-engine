package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/marcocristofolini/kubernetes-rca-engine/internal/engine"
	"github.com/marcocristofolini/kubernetes-rca-engine/internal/model"
)

func main() {
	if len(os.Args) != 3 || os.Args[1] != "analyze" {
		fmt.Fprintln(os.Stderr, "usage: krca analyze <incident.json>")
		os.Exit(2)
	}

	data, err := os.ReadFile(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "read incident: %v\n", err)
		os.Exit(1)
	}

	var incident model.Incident
	if err := json.Unmarshal(data, &incident); err != nil {
		fmt.Fprintf(os.Stderr, "parse incident: %v\n", err)
		os.Exit(1)
	}

	result := struct {
		Incident model.Incident  `json:"incident"`
		Findings []model.Finding `json:"findings"`
	}{incident, engine.Analyze(incident)}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(result); err != nil {
		fmt.Fprintf(os.Stderr, "encode result: %v\n", err)
		os.Exit(1)
	}
}
