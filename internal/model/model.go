package model

type Signal struct {
	Type     string `json:"type"`
	Reason   string `json:"reason,omitempty"`
	Message  string `json:"message,omitempty"`
	Container string `json:"container,omitempty"`
	ExitCode int    `json:"exitCode,omitempty"`
}

type Incident struct {
	Namespace string   `json:"namespace"`
	Workload  string   `json:"workload"`
	Pod       string   `json:"pod"`
	Signals   []Signal `json:"signals"`
}

type Finding struct {
	Cause       string   `json:"cause"`
	Confidence  float64  `json:"confidence"`
	Evidence    []string `json:"evidence"`
	Remediation []string `json:"remediation,omitempty"`
}
