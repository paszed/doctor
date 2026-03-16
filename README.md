# doctor

Doctor is a CLI tool for diagnosing developer environments.

The goal is to help developers quickly identify problems in their local setup such as:

- missing tools
- docker issues
- port conflicts
- environment configuration problems

## Status

This project is currently under active development.

## Project Structure

```
cmd/doctor

internal/
  checks/
  cli/
  config/
  detect/
  fix/
  model/
  report/
  runtime/
  state/
  system/
  ui/

scripts/
examples/
```

## Goal

Doctor aims to provide a simple tool for inspecting and repairing developer environments.
