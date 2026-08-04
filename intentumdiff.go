// Package intentumdiff is the thin Go binding for the IntentumDiff engine.
//
// The engine (github.com/buchochelliq-labs/intentumdiff-core) ships a native shared
// library exposing ONE stable C ABI:
//
//	intentumdiff_call(name *char, args_json *char) -> *char   // heap JSON envelope
//	intentumdiff_free(ptr *char)                               // frees the envelope
//
// This binding does zero functional work: it marshals arguments to JSON, calls the
// engine, and maps the {ok, result, error, error_type} envelope to Go values —
// exactly like the Python ctypes binding.
package intentumdiff

import (
	"encoding/json"
	"fmt"
	"io/fs"
)

// Envelope is the C-ABI result envelope.
type envelope struct {
	OK        bool            `json:"ok"`
	Result    json.RawMessage `json:"result"`
	Error     string          `json:"error"`
	ErrorType string          `json:"error_type"`
}

// EngineError is a structured engine failure, carrying the C-ABI error_type slug
// (not_found | value_error | bad_request | internal).
type EngineError struct {
	Type    string
	Message string
}

func (e *EngineError) Error() string { return e.Type + ": " + e.Message }

// Unwrap maps not_found onto fs.ErrNotExist so errors.Is(err, fs.ErrNotExist)
// works, mirroring the Python binding's FileNotFoundError mapping.
func (e *EngineError) Unwrap() error {
	if e.Type == "not_found" {
		return fs.ErrNotExist
	}
	return nil
}

// Engine is a loaded IntentumDiff core library.
type Engine struct {
	lib nativeLib
}

// Load opens the engine shared library (intentumdiff_rust_core.{dll,so,dylib}).
func Load(path string) (*Engine, error) {
	lib, err := openNative(path)
	if err != nil {
		return nil, fmt.Errorf("intentumdiff: load %q: %w", path, err)
	}
	return &Engine{lib: lib}, nil
}

// Call invokes an engine function with positional JSON-marshalled arguments and
// returns the raw JSON result.
func (e *Engine) Call(name string, args ...any) (json.RawMessage, error) {
	if args == nil {
		args = []any{} // a nil slice marshals to JSON null; the ABI wants []
	}
	argsJSON, err := json.Marshal(args)
	if err != nil {
		return nil, fmt.Errorf("intentumdiff: marshal args: %w", err)
	}
	raw, err := e.lib.call(name, string(argsJSON))
	if err != nil {
		return nil, err
	}
	var env envelope
	if err := json.Unmarshal([]byte(raw), &env); err != nil {
		return nil, fmt.Errorf("intentumdiff: bad envelope: %w", err)
	}
	if !env.OK {
		return nil, &EngineError{Type: env.ErrorType, Message: env.Error}
	}
	return env.Result, nil
}

// Version returns the engine version string.
func (e *Engine) Version() (string, error) {
	raw, err := e.Call("version")
	if err != nil {
		return "", err
	}
	var v string
	return v, json.Unmarshal(raw, &v)
}

// SupportsLanguage reports whether the engine's native batch path serves lang.
func (e *Engine) SupportsLanguage(lang string) (bool, error) {
	raw, err := e.Call("supports_language", lang)
	if err != nil {
		return false, err
	}
	var b bool
	return b, json.Unmarshal(raw, &b)
}
