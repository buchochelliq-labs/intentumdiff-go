# intentumdiff-go

[![CI](https://github.com/buchochelliq-labs/intentumdiff-go/actions/workflows/ci.yml/badge.svg)](https://github.com/buchochelliq-labs/intentumdiff-go/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Go 1.22+](https://img.shields.io/badge/go-1.22%2B-00ADD8.svg)](https://go.dev/)

The **thin Go binding** for the IntentumDiff engine
([intentumdiff-core](https://github.com/buchochelliq-labs/intentumdiff-core)). Zero functional
work: arguments marshal to JSON, the engine's stable C ABI (`intentumdiff_call` /
`intentumdiff_free`) does everything, and the `{ok, result, error, error_type}` envelope maps
to Go values (`not_found` unwraps to `fs.ErrNotExist`) — exactly like the Python binding.

```go
engine, _ := intentumdiff.Load(os.Getenv("INTENTUMDIFF_CORE_LIB"))
version, _ := engine.Version()
ok, _ := engine.SupportsLanguage("python")
```

Scaffold status: Windows loader is stdlib-only (no cgo); the unix dlopen seam is stubbed
(`loader_other.go`). Tests need a built engine library: set `INTENTUMDIFF_CORE_LIB` to the
`intentumdiff_rust_core.{dll,so,dylib}` from intentumdiff-core, else they skip.

Migrated concept + provenance: authored fresh for the #82 split; the monorepo
(`buchochelliq-labs/intentumdiff`) remains the archive of record. License: MIT.
