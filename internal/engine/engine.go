package engine

import (
	"sort"
	"strings"

	"github.com/marcocristofolini/kubernetes-rca-engine/internal/model"
)

func Analyze(i model.Incident) []model.Finding {
	var findings []model.Finding

	for _, s := range i.Signals {
		reason := strings.ToLower(s.Reason)
		message := strings.ToLower(s.Message)

		if reason == "evicted" && strings.Contains(message, "ephemeral-storage") {
			findings = append(findings, model.Finding{
				Cause:      "Node ephemeral storage pressure",
				Confidence: 0.98,
				Evidence:   []string{"Pod was evicted", "Eviction message mentions ephemeral-storage"},
				Remediation: []string{
					"Inspect node ephemeral-storage usage and kubelet eviction signals",
					"Set realistic ephemeral-storage requests and limits",
					"Move unbounded temporary data to an appropriate volume or external store",
				},
			})
		}

		if reason == "oomkilled" || strings.Contains(message, "oomkilled") {
			findings = append(findings, model.Finding{
				Cause:      "Container exceeded its memory limit",
				Confidence: 0.95,
				Evidence:   []string{"Container termination indicates OOMKilled"},
				Remediation: []string{
					"Compare memory working set with requests and limits",
					"Check for memory leaks or load-related growth",
					"Adjust limits only after understanding actual demand",
				},
			})
		}

		if reason == "crashloopbackoff" {
			findings = append(findings, model.Finding{
				Cause:      "Repeated container startup failure",
				Confidence: 0.72,
				Evidence:   []string{"Pod reports CrashLoopBackOff"},
				Remediation: []string{
					"Inspect previous container logs and termination state",
					"Check configuration, dependency reachability and startup probes",
				},
			})
		}
	}

	sort.SliceStable(findings, func(a, b int) bool {
		return findings[a].Confidence > findings[b].Confidence
	})
	return findings
}
