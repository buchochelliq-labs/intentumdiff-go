# Building intentdiff-go

Toolchain: **Go 1.22** (no cgo required on Windows).

```bash
go vet -unsafeptr=false ./...    # the flag covers the one documented FFI pointer conversion
go build ./...
```

Tests exercise a real engine library — point `INTENTDIFF_CORE_LIB` at a built
`intentdiff_rust_core.{dll,so,dylib}` from
[intentdiff-core](https://github.com/buchochelliq-labs/intentdiff-core) (they self-skip
without it):

```bash
INTENTDIFF_CORE_LIB=/path/to/intentdiff_rust_core.dll go test ./... -v
```
