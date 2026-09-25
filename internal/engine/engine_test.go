package engine

import (
	"testing"

	"github.com/marcocristofolini/kubernetes-rca-engine/internal/model"
)

func TestEphemeralStorageEviction(t *testing.T) {
	incident := model.Incident{
		Pod: "api-7b9f",
		Signals: []model.Signal{
			{
				Type:    "event",
				Reason:  "Evicted",
				Message: "The node was low on resource: ephemeral-storage.",
			},
		},
	}

	findings := Analyze(incident)
	if len(findings) == 0 {
		t.Fatal("expected at least one finding")
	}
	if findings[0].Cause != "Node ephemeral storage pressure" {
		t.Fatalf("unexpected cause: %s", findings[0].Cause)
	}
	if findings[0].Confidence < 0.9 {
		t.Fatalf("unexpected confidence: %f", findings[0].Confidence)
	}
}

func TestOOMKilled(t *testing.T) {
	incident := model.Incident{
		Signals: []model.Signal{{Type: "termination", Reason: "OOMKilled"}},
	}
	findings := Analyze(incident)
	if len(findings) != 1 || findings[0].Cause != "Container exceeded its memory limit" {
		t.Fatalf("unexpected findings: %#v", findings)
	}
}
