# Architecture

## Target data flow

```mermaid
flowchart TD
    API[Kubernetes API] --> KC[Kubernetes Collector]
    PROM[Prometheus] --> MC[Metrics Collector]
    LOKI[Loki] --> LC[Log Collector]
    TEMPO[Tempo] --> TC[Trace Collector]
    KC --> NORMALIZE[Signal Normalization]
    MC --> NORMALIZE
    LC --> NORMALIZE
    TC --> NORMALIZE
    NORMALIZE --> STORE[Evidence Store]
    STORE --> RULES[Deterministic Detectors]
    STORE --> GRAPH[Correlation Graph]
    RULES --> RANK[Finding Ranker]
    GRAPH --> RANK
    RANK --> REPORT[Explainable RCA Report]
    REPORT --> LLM[Optional LLM Summary]
```

## Boundaries

Collectors are read-only adapters. Normalization converts provider-specific payloads into a small internal model while retaining links to raw evidence.

Deterministic detectors encode high-confidence signatures. The correlation graph models workload, pod, container, node, service and dependency relationships.

The optional AI layer may summarize findings or suggest investigation paths, but must not invent evidence or mutate the cluster.
