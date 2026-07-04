# Learning Go in 60 Days — an experiment

**Hypothesis:** Is it worth an experienced engineer's time to learn a new
language — using AI as the tutor? This repo is the experiment log.

- **Learner:** backend engineer, fluent in Java & Python, works in distributed systems. Zero Go.
- **Budget:** 20–30 min/day, 60 days.
- **Method:** build-first. Every concept lands in a running program the same day.
- **Build track:** a distributed key-value store, grown piece by piece.
- **Tutor:** Claude (theory injected just-in-time, per task).

## How a day works

1. Open `curriculum/PROGRESS.md`, find today's task.
2. Read the day's spec (in `curriculum/`), open the starter file under `dayNN/`.
3. Code it (~20 min). Run it. Commit.
4. Ask the tutor to review + inject the theory that task needed.

## The rule that makes the experiment honest

**You write the code. AI reviews and explains.** If the AI writes it, you're
measuring the AI, not your learning. Get stuck → ask for a hint, not the answer.

## Measuring the experiment

- **Commit history** = time-stamped proof of consistency.
- `curriculum/PROGRESS.md` = daily one-line "what clicked / what confused me."
- Day 30 and Day 60: a short self-graded checkpoint (in the curriculum).

## Layout

```
curriculum/
  CURRICULUM.md   full 60-day plan
  PROGRESS.md     your daily log (you edit this)
day01/, day02/... one folder per day, self-contained programs
kv/               the distributed KV store, grown from ~Day 27
```

## Running anything

```
go run ./day01           # run a day's program
go test ./...            # run all tests
```
