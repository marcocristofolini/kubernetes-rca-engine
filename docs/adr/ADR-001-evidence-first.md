# ADR-001: Evidence-first RCA

**Status:** Accepted

## Context

AI-assisted incident tools can produce plausible explanations even when evidence is incomplete.

## Decision

The engine builds findings from explicit evidence before any optional LLM step.

Every finding should carry the suspected cause, confidence, supporting evidence and suggested verification or remediation steps.

## Consequences

### Positive
- findings are auditable;
- deterministic detectors are testable;
- the engine works without an external AI provider;
- engineers can challenge individual hypotheses.

### Trade-offs
- more up-front modeling;
- integrations require normalization;
- some incidents remain inconclusive instead of receiving a confident-looking answer.
