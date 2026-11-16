# learn-logging-observability

My work through boot.dev's
[Learn Logging and Observability in Go](https://www.boot.dev/courses/learn-logging-observability-golang)
(73 lessons, 12 chapters), part of the DevOps path. The course adds logs,
metrics, alerts, profiles and traces to Linko, a small URL shortener that
boot.dev provides already written.

This repository starts from the course's starter,
[bootdotdev/linko-starter](https://github.com/bootdotdev/linko-starter), with
its history kept. Each lesson's changes are committed on top.

> Linko is "intentionally small, a little messy, and realistic enough to
> practice adding logs, metrics, and traces in Go." Some of the odd code is
> deliberate: the course uses observability to find it.

## Run it

```sh
go run .            # http://localhost:8899
go run . -port 9000 -data ./data
```

## Learning goals

The course frames observability as two questions. I hold my production
services to them:

1. **How do I know when something is going wrong?** (detection)
2. **How can I identify the problem as quickly as possible?** (diagnosis)

The course page gives one line per chapter. The goals below restate those
lines; the bullets after the first one in each chapter are my expansion of the
line, not text from the course.

1. **Observability.** The three pillars: logs, metrics, traces.
   - What each pillar answers, and when to use which.
2. **Logging.** Logger setup, dependency injection, buffering, lifecycle.
   - Pass the logger in, don't reach for a global.
   - Flush buffered output before exit.
3. **Structured Logging.** `log/slog` levels, key-value fields, output formats.
   - Text vs JSON handlers, `slog.Attr`, level filtering.
4. **Log Strategies.** Reduce noise, keep signal.
   - Choosing levels, what not to log, sampling.
5. **Logging Errors.** Stack traces, grouped context, rich attributes.
   - Log an error once, at the boundary that handles it.
6. **Logging Context.** Request, user, build and instance fields.
   - `Logger.With`, `context.Context`, correlation IDs.
7. **Log Storage.** Console, files, syslog, rotation.
8. **Log Security.** Filtering, obfuscation, encryption, safe error responses.
   - Redaction with `slog.LogValuer` and `ReplaceAttr`.
9. **Metrics.** Prometheus and Grafana, system and app metrics, dashboards.
   - Counters, gauges, histograms, label cardinality.
10. **Alerting.** Actionable alerts, sensible thresholds, no alert fatigue.
11. **Profiling.** pprof for CPU, memory and goroutines.
12. **Tracing.** OpenTelemetry and Jaeger across services.
    - Span context propagation, linking traces to logs.

## Notes

`notes/chNN.md` holds my notes for each chapter.

## Progress

- [x] 1. Observability
- [x] 2. Logging
- [ ] 3. Structured Logging
- [ ] 4. Log Strategies
- [ ] 5. Logging Errors
- [ ] 6. Logging Context
- [ ] 7. Log Storage
- [ ] 8. Log Security
- [ ] 9. Metrics
- [ ] 10. Alerting
- [ ] 11. Profiling
- [ ] 12. Tracing
