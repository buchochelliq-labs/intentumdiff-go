# intentumdiff-go architecture — the thin Go binding

Zero functional work: arguments marshal to a JSON array, the engine's
[C ABI](https://github.com/buchochelliq-labs/intentumdiff-core/blob/main/docs/C_ABI.md) does
everything, and the `{ok, result, error, error_type}` envelope maps to Go values.

- `intentumdiff.go` — the portable API: `Load(path)`, `Engine.Call(name, args...)` (a nil args
  slice marshals as `[]`, never `null`), typed helpers (`Version`, `SupportsLanguage`), and
  `EngineError` carrying the `error_type` slug (`not_found` unwraps to `fs.ErrNotExist`).
- `loader_windows.go` — stdlib-only loader (`syscall.LoadDLL` + raw proc calls; one documented
  foreign-pointer conversion, vetted with `-unsafeptr=false`). No cgo.
- `loader_other.go` — the unix dlopen seam (stubbed; lands with cgo or purego).

Contract facts the tests pin: missing positional arg → `value_error`; malformed `args_json`
→ `bad_request`; extra args on zero-arg calls are ignored; every envelope pointer is released
via `intentumdiff_free` after copying.
