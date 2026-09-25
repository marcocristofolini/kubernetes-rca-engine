# Changelog

All notable changes to this project are documented here.

## [0.1.0] - 2026-09-25

### Added
- evidence-first incident model;
- deterministic RCA rules for ephemeral-storage eviction, OOMKilled and CrashLoopBackOff;
- `krca analyze` CLI;
- unit tests and example incident payload;
- architecture documentation and ADR;
- GitHub Actions CI;
- CodeQL security scanning;
- Dependabot dependency automation.

### Design
Version 0.1.0 intentionally keeps collectors separate from the diagnostic engine. This allows evidence correlation and confidence behavior to remain testable before adding Kubernetes, Prometheus, Loki and Tempo integrations.
