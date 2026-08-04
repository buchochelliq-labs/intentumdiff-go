# Building intentumdiff-go

Toolchain: **Go 1.22** (no cgo required on Windows).

```bash
go vet -unsafeptr=false ./...    # the flag covers the one documented FFI pointer conversion
go build ./...
```

Tests exercise a real engine library — point `INTENTUMDIFF_CORE_LIB` at a built
`intentumdiff_rust_core.{dll,so,dylib}` from
[intentumdiff-core](https://github.com/buchochelliq-labs/intentumdiff-core) (they self-skip
without it):

```bash
INTENTUMDIFF_CORE_LIB=/path/to/intentumdiff_rust_core.dll go test ./... -v
```
