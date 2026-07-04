# The 60-Day Curriculum

Build-first. Each day: a small spec, you write ~20 min of code, tutor reviews +
injects theory. Days assume Java/Python fluency — we spend time on what's
*different* in Go, not on what a variable is.

Legend: 🎯 = build-track (feeds the KV store) · 🧩 = standalone drill · 🏁 = checkpoint

---

## Phase 1 — Days 1–12: Go idioms for polyglots
Goal: read/write idiomatic Go fluently. Kill the friction so concurrency (Phase 2)
isn't fighting syntax.

| Day | Topic | You build |
|-----|-------|-----------|
| 1  | Toolchain, modules, `main`, `fmt` | 🧩 Hello + a temperature converter CLI |
| 2  | Types, zero values, `var` vs `:=`, constants, `iota` | 🧩 Enum-like status codes with `iota` |
| 3  | Functions, multiple returns, named returns | 🧩 `divmod`, min/max over a slice |
| 4  | Errors as values, `errors.Is/As`, `fmt.Errorf %w` | 🧩 A parser that returns wrapped errors |
| 5  | Structs, methods, value vs pointer receivers | 🧩 A `Point`/`Rect` type with methods |
| 6  | Slices deep-dive (len/cap, append aliasing traps) | 🧩 Implement a growable ring buffer |
| 7  | Maps, sets-via-maps, ordered iteration | 🧩 Word-frequency counter |
| 8  | Interfaces (structural typing!), `io.Reader/Writer` | 🧩 A `Stringer` + a custom `io.Writer` |
| 9  | `defer`, panic/recover, cleanup patterns | 🧩 Safe file-processing with `defer` |
| 10 | Testing: `testing.T`, table tests, `go test` | 🧩 Full table tests for Day 6's ring buffer |
| 11 | Generics (type params, constraints) | 🧩 Generic `Map/Filter/Reduce` + a `Set[T]` |
| 12 | 🏁 Phase-1 mini-project | 🧩 An in-memory KV store (single-threaded, `map`-backed, with tests) |

## Phase 2 — Days 13–26: Concurrency (Go's crown jewel)
Goal: this is your highest-leverage area as a DS engineer. By the end you think
in goroutines/channels/`context` the way you think in threads/futures today.

| Day | Topic | You build |
|-----|-------|-----------|
| 13 | Goroutines, `sync.WaitGroup` | 🧩 Parallel URL/word fetcher (simulated) |
| 14 | Channels: buffered/unbuffered, direction, close | 🧩 Producer→consumer with a channel |
| 15 | `select`, timeouts, `time.After` | 🧩 A timeout wrapper around a slow op |
| 16 | `context`: cancellation, deadlines, values | 🧩 Cancellable worker via `context` |
| 17 | Mutexes vs channels; `sync.RWMutex`, `sync.Once` | 🧩 Make Day 12's KV store thread-safe |
| 18 | Worker pool pattern | 🧩 Fixed-size pool processing a job queue |
| 19 | Fan-out / fan-in | 🧩 Parallel map-reduce over a dataset |
| 20 | Pipelines + cancellation | 🧩 Multi-stage streaming pipeline |
| 21 | Rate limiting (`time.Ticker`, token bucket) | 🧩 A reusable rate limiter |
| 22 | `sync/atomic`, race detector (`go test -race`) | 🧩 Atomic counter + find a planted race |
| 23 | `errgroup`, structured concurrency | 🧩 Concurrent fetch that fails fast |
| 24 | Deadlocks, leaks, `context` hygiene | 🧩 Fix a set of buggy concurrent snippets |
| 25 | Benchmarking + `pprof` intro | 🧩 Benchmark mutex-KV vs sharded-KV |
| 26 | 🏁 Phase-2 project | 🎯 Concurrent, sharded, thread-safe in-memory KV store |

## Phase 3 — Days 27–45: Distributed KV store components
Goal: turn the in-memory store into a real networked service, one subsystem/day.
This is `kv/` from here on.

| Day | Topic | You build |
|-----|-------|-----------|
| 27 | `net/http` server; JSON encode/decode | 🎯 `GET/PUT/DELETE` HTTP API over the store |
| 28 | Routing, middleware, graceful shutdown | 🎯 Logging + recovery middleware, clean shutdown |
| 29 | Raw TCP server + a simple wire protocol | 🎯 A RESP-lite TCP protocol for the store |
| 30 | 🏁 **Checkpoint 1** (self-grade, see below) | 🎯 Client library for your TCP protocol |
| 31 | Persistence: append-only log (WAL) | 🎯 Write-ahead log; replay on startup |
| 32 | Snapshots + log compaction | 🎯 Periodic snapshot, truncate the WAL |
| 33 | TTL / expiry (background sweeper) | 🎯 Per-key TTL with a janitor goroutine |
| 34 | LRU eviction | 🎯 Bounded store with LRU eviction |
| 35 | Consistent hashing | 🎯 Hash ring for key→node routing |
| 36 | Sharding across nodes (multi-process) | 🎯 Route requests to the owning shard |
| 37 | Replication: leader→follower | 🎯 Async replication of writes |
| 38 | Read repair / eventual consistency basics | 🎯 Handle stale-replica reads |
| 39 | Failure detection: heartbeats | 🎯 Node health tracking |
| 40 | Retries, timeouts, backoff | 🎯 Resilient client calls |
| 41 | Circuit breaker | 🎯 Wrap node calls in a breaker |
| 42 | Metrics (`expvar`/Prometheus-style counters) | 🎯 Ops/latency counters + a `/metrics` endpoint |
| 43 | Structured logging (`log/slog`) | 🎯 Request-scoped structured logs |
| 44 | Config + flags (`flag`, env) | 🎯 Configurable node (ports, peers, data dir) |
| 45 | 🏁 Integration tests for the cluster | 🎯 Spin up N nodes in-test, assert behavior |

## Phase 4 — Days 46–60: Capstone — a small distributed KV cluster
Goal: assemble everything into a runnable multi-node cluster and harden it.

| Day | Topic | You build |
|-----|-------|-----------|
| 46 | Cluster membership (static → gossip-lite) | 🎯 Nodes discover peers |
| 47 | Quorum reads/writes (N/R/W) | 🎯 Tunable consistency |
| 48 | Vector clocks or version numbers | 🎯 Conflict detection |
| 49 | Anti-entropy / hinted handoff | 🎯 Recover missed writes |
| 50 | Leader election (simplified) | 🎯 Pick a coordinator |
| 51 | Backpressure + load shedding | 🎯 Protect nodes under load |
| 52 | Graceful rebalancing on node join/leave | 🎯 Move keys when the ring changes |
| 53 | Chaos testing (inject latency/failures) | 🎯 A fault-injection harness |
| 54 | Profiling + optimizing a hot path | 🎯 Find & fix one real bottleneck |
| 55 | End-to-end benchmark (throughput/latency) | 🎯 A load-gen tool + a results writeup |
| 56 | Dockerize + run a real 3-node cluster | 🎯 `docker compose` cluster |
| 57 | Docs + architecture writeup | 🎯 `kv/ARCHITECTURE.md` |
| 58 | Hardening: edge cases, fuzz tests | 🎯 `go test -fuzz` on the protocol parser |
| 59 | Polish + code review pass | 🎯 Clean up, `go vet`, lint |
| 60 | 🏁 **Checkpoint 2** + experiment writeup | 🎯 Final self-grade + verdict on the experiment |

---

## Checkpoints (the experiment's measurements)

**Day 30 — Checkpoint 1.** Without AI help, in 25 min:
1. Write a thread-safe counter with N goroutines. Does `-race` pass?
2. Explain in the log: value vs pointer receivers; when a channel read blocks;
   what `context` cancellation actually does.
3. Self-grade 1–5 on: reading Go, writing Go unaided, concurrency intuition.

**Day 60 — Checkpoint 2 + verdict.** In 30 min:
1. Add a small new feature to the KV store unaided; time yourself.
2. Re-grade the Day-30 axes. Compare.
3. Write the verdict: *Was learning Go worth it? Did AI-as-tutor accelerate or
   shortcut the learning?* This is the deliverable of the whole experiment.

## Adjustable
This plan is a scaffold, not a contract. If a phase feels too fast/slow we
re-cut it. Missing a day is fine — the experiment is about the trend, not a streak.
