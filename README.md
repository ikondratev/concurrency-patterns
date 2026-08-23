# Concurrency Patterns in Go

A learning repository of concurrency patterns in Go.
The goal is to study how goroutines, channels, and synchronization primitives are typically composed: when a pattern fits, how to implement it, and which mistakes it prevents.
This is not a library to import into production code. Each pattern is a self-contained example with an explanation and tests.

## Layout

Each pattern lives in its own directory:

```text
pattern-name/
├── README.md        # problem, idea, when to use, pitfalls
├── pattern.go       # implementation
├── pattern_test.go  # tests and usage examples
└── example/main.go  # optional runnable demo
```

## Running examples

```bash
go test ./...
go test ./worker-pool -v
go run ./worker-pool/example
```

## Example guidelines

- One pattern, one idea. No extra abstraction “for later”.
- Code should be readable, not as short as possible.
- Each example has tests: at least the happy path, plus cancellation or error if that is part of the pattern.
- The pattern README states **when it fits** and **when another pattern is better**.
- Goroutines must not leak: every example shows who closes channels and who waits for completion.

## Further reading

- [Go Concurrency Patterns](https://go.dev/blog/pipelines) — pipelines and cancellation
- [Share Memory by Communicating](https://go.dev/blog/codelab-share)
- [Effective Go: Concurrency](https://go.dev/doc/effective_go#concurrency)
- Katherine Cox-Buday, *Concurrency in Go*
- [golang.org/x/sync](https://pkg.go.dev/golang.org/x/sync) — `errgroup`, `singleflight`, `semaphore`
