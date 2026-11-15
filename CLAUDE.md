# learn-logging-observability: working conventions

## Roles

- The user writes every change. Claude teaches and reviews: explains concepts, reviews code, finds bugs. Claude does not write lesson solutions unless asked.
- The user writes Go professionally (embedded Linux, IoT telematics, concurrent data processing, AWS IoT, RabbitMQ backends). Skip beginner explanations. Explain what is new or easy to get wrong: runtime behavior, the memory model, idioms, how the course's answer differs from production code, and recent Go version changes.
- The user holds every production service to the course's two questions: "How do I know when something is going wrong?" (detection) and "How can I identify the problem as quickly as possible?" (diagnosis). In each lesson's production note, say which question the lesson's tool answers and where it falls short.
- The user already uses `log/slog` with handlers and `With` at work. Keep chapters 2 to 8 short on basics. Spend time on what is new: custom `slog.Handler` implementations, `ReplaceAttr` and `LogValuer` redaction, syslog and rotation on embedded targets (journald, limited flash), Prometheus metric types and label cardinality, alert design, pprof, and OpenTelemetry trace propagation and log correlation.

## Course

- Learn Logging and Observability in Go, boot.dev: https://www.boot.dev/courses/learn-logging-observability-golang (73 lessons, 12 chapters).
- Mode: **course with a starter project** (see the `bootdev` skill's `course-mode.md`, section "Course with a starter project"). The lessons are a course, but most of them change one program, Linko, a URL shortener that boot.dev wrote. The user works locally and checks each lesson with the boot.dev CLI.
- Part of the boot.dev DevOps path, done in Go.

## Layout

- This repository is the starter, https://github.com/bootdotdev/linko-starter, with its history kept. Remote `upstream` points at it. `origin` is `github.com/zoumas/learn-logging-observability`.
- Module: `boot.dev/linko`, as the starter ships it. Do not rename it: the starter reverted a rename once ("revert module URL"), and boot.dev's checks may depend on it. Go 1.27.
- The Linko code lives at the root (`main.go`, `server.go`, `handlers.go`, `internal/store`). There are no `chNN/lNN/` directories. The lesson number lives in the commit message.
- `notes/chNN.md`: notes for each chapter. They feed the notes page at the end of the course.
- `data/` is the runtime store and is gitignored by the starter.

## Lesson kinds

| Kind | What arrives | What to do |
|---|---|---|
| Linko change | Lesson URL, the user's uncommitted diff | Full review of the diff, notes, `feat(N.M): ...` with the code and the notes |
| CLI check | Lesson URL, `bootdev run <id>` | Reproduce what the check needs (server running, a request, an output line), notes, `docs(N.M): ...` |
| Quiz | Lesson URL, screenshots of each question and the chosen answer | Record each question and correct answer, notes, `docs(N.M): ...` |

## Lesson workflow

1. The user pastes a lesson URL. Fetch it for the requirements. The test panel (the `bootdev run` command and what it checks), expected output and quiz options are behind the login. Ask the user to paste the test panel when the exact check matters.
2. Read the user's work: `git status --short`, `git diff`, and the whole of every changed file. Run the gates and `golangci-lint run`.
3. Run Linko. Start it on a scratch data directory (`go run . -data <scratch dir>`), exercise the endpoints the lesson touches with `curl`, and stop it with `POST /admin/shutdown` or SIGINT. Read what it logs. For logging lessons the output *is* the feature, so quote it.
4. Run the race detector on the running server (`go build -race`), not only on the tests: most of Linko's concurrency is in the HTTP server and shutdown path.
5. Review with the `go-review` skill. Report correctness first, then spec mismatches, then idiom, then naming. Add the production note, tied to the user's domain where it helps.
6. Verify every claim in a note by running it before writing it down. Start a measured claim with `Measured:`; label anything else "reasoned, not measured" or cite the doc. When an earlier note is wrong, say so plainly and fix it.
7. Append the lesson to `notes/chNN.md` as `## N.M <Lesson Title>`.
8. Commit after the gates pass. Last lesson of a chapter: tick the chapter in the README's Progress list in the same commit.
9. The user runs `bootdev run -s <id>` (submit). Claude may run `bootdev run <id>` (check only) to verify.

Pace: review and commit every lesson on its own. Keep reviews of trivial lessons short.

## The starter's deliberate bugs

The course says: "There are a few strange bits of code in the Linko starter repo that you may be tempted to 'fix.' Don't! They're intentional bugs and oddities that we'll be using good observability to find and diagnose!"

- Do not report the starter's own bugs as review findings, and do not suggest fixing them, unless the current lesson is about that bug.
- When a lesson exposes one, record in the notes how the observability tool found it.
- Baseline measured on the starter at `36bba29` (2026-11-13), so later reviews can tell starter issues from new ones:
  - `gofmt -l .` and `go vet ./...` are clean.
  - `golangci-lint run` reports 6 errcheck issues (unchecked `io.WriteString`, `json.Encoder.Encode`, `f.Close`).
  - From 2.2, `.golangci.yml` enables `forbidigo` (`^fmt\.Print.*$`). It adds 1 issue on starter code: `internal/store/store.go:109`. Baseline is 7 issues.
  - `go build -race`, run, then `POST /admin/shutdown`: data race on `serverErr` in `main.go` (written by the server goroutine at line 36, read by `run` at line 47).
  - `auth.go` has a redundant `fmt.Printf("error validating password: %v\n", err)` added on purpose (commits `2f42479` "add redundant log statement" and `15f9dad` "make redundant log line more obvious").

## Git

- The user pushes. Claude commits only.
- Never amend or rewrite a commit that is on `origin/main`. Check with `git branch -r --contains <sha>`. Fix a pushed note with a new commit.
- Never push to `upstream`. To pick up a starter fix: `git fetch upstream`, then merge it, never rebase pushed lesson commits.

## Tooling

- Go 1.27. golangci-lint must be v2.15.0 or newer and gopls must be built with Go 1.27, or they reject newer export data and syntax (`export data version 5 is greater than maximum supported version 4`). Check with `go version $(which gopls)` when the editor and CLI disagree.
- Boot.dev CLI (`bootdev`, v1.32.6 checked). `bootdev run <id>` checks; `bootdev run -s <id>` submits. Run it from the repository root, with Linko running when the lesson says so.
- Third-party modules (Prometheus client, OpenTelemetry) arrive with the lessons that need them. Add them with `go get` in the lesson commit and keep `go.sum` committed.
- Prometheus, Grafana and Jaeger: run locally only (containers or packages), never against a work account or work endpoint. Docker is not installed on every machine; check before a lesson needs it.

## Progress

Chapter 1 done. `git log --oneline` and `notes/chNN.md` are the source of truth.

## Tests

- The starter has no tests. boot.dev's checks run through the CLI. Copy in any test files the course provides.
- Do not suggest extra tests. Write tests only when the user and Claude are the ones writing them.

## Gates before every commit

- `gofmt -l .` prints nothing
- `go vet ./...`
- `go test -race ./...`
- `go build` succeeds and Linko serves `GET /` with 200
- `go mod tidy` leaves `go.mod` and `go.sum` unchanged

## Commits

- `<type>(<chapter>.<lesson>): <what changed>`, for example `feat(2.3): inject the logger into server`, or `docs(1.3): add quiz notes`.
- Use the body only for a non-obvious reason.

## End of course

1. Publish the chapter notes as one shareable Artifact page with the `bootdev` skill's `tools/notes_page/build.py`, and serve the standalone build from `docs/index.html` with GitHub Pages.
2. Update the `bootdev` skill's `course-mode.md` and `devops-go-playbook.md` with what this course taught about the workflow.
