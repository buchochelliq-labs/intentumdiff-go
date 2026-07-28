# Contributing to intentdiff-go

- The binding stays thin: no semantics, no engine-result post-processing — new engine
  capability lands in intentdiff-core and surfaces here as typed helpers over `Call`.
- The obvious next contribution: the unix loader (`loader_other.go`) via cgo dlopen or purego,
  keeping the portable API untouched.
- `go vet -unsafeptr=false ./... && go test ./...` green (tests against a real engine lib —
  see docs/BUILDING.md).
