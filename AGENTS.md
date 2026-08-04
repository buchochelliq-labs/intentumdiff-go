# Agent instructions — intentumdiff-go

The thin Go binding. Zero functional work — JSON args -> intentumdiff_call -> envelope mapping.

## Hard invariants
- No semantics here; new capability = typed helpers over Call, engine work goes upstream.
- nil args marshal as [] (never null); every envelope pointer is freed after copying.
- go vet runs with -unsafeptr=false (one documented foreign-pointer conversion).

Test: INTENTUMDIFF_CORE_LIB=<engine lib> go test ./... (self-skips without it).
Contract: https://github.com/buchochelliq-labs/intentumdiff-core/blob/main/docs/C_ABI.md
