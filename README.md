# intentdiff-go

[![CI](https://github.com/buchochelliq-labs/intentdiff-go/actions/workflows/ci.yml/badge.svg)](https://github.com/buchochelliq-labs/intentdiff-go/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Go 1.22+](https://img.shields.io/badge/go-1.22%2B-00ADD8.svg)](https://go.dev/)

The **thin Go binding** for the IntentDiff engine
([intentdiff-core](https://github.com/buchochelliq-labs/intentdiff-core)). Zero functional
work: arguments marshal to JSON, the engine's stable C ABI (`intentdiff_call` /
`intentdiff_free`) does everything, and the `{ok, result, error, error_type}` envelope maps
to Go values (`not_found` unwraps to `fs.ErrNotExist`) — exactly like the Python binding.

```go
engine, _ := intentdiff.Load(os.Getenv("INTENTDIFF_CORE_LIB"))
version, _ := engine.Version()
ok, _ := engine.SupportsLanguage("python")
```

Scaffold status: Windows loader is stdlib-only (no cgo); the unix dlopen seam is stubbed
(`loader_other.go`). Tests need a built engine library: set `INTENTDIFF_CORE_LIB` to the
`intentdiff_rust_core.{dll,so,dylib}` from intentdiff-core, else they skip.

Migrated concept + provenance: authored fresh for the #82 split; the monorepo
(`buchochelliq-labs/intentdiff`) remains the archive of record. License: MIT.
