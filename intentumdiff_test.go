package intentumdiff

import (
	"errors"
	"os"
	"testing"
)

// The tests need a built engine library; set INTENTUMDIFF_CORE_LIB to its path
// (intentumdiff_rust_core.dll / .so / .dylib from intentumdiff-core).
func loadEngine(t *testing.T) *Engine {
	t.Helper()
	path := os.Getenv("INTENTUMDIFF_CORE_LIB")
	if path == "" {
		t.Skip("INTENTUMDIFF_CORE_LIB not set — needs a built intentumdiff-core library")
	}
	e, err := Load(path)
	if err != nil {
		t.Fatalf("load engine: %v", err)
	}
	return e
}

func TestVersionAndLanguageSupport(t *testing.T) {
	e := loadEngine(t)
	v, err := e.Version()
	if err != nil || v == "" {
		t.Fatalf("version: %q err=%v", v, err)
	}
	ok, err := e.SupportsLanguage("python")
	if err != nil || !ok {
		t.Fatalf("supports_language(python): %v err=%v", ok, err)
	}
}

func TestErrorEnvelopeMapping(t *testing.T) {
	e := loadEngine(t)
	// A MISSING required argument is the canonical value_error (bad_request is reserved for malformed args_json itself) (extra args on a
	// zero-arg call are ignored by the ABI's positional arg readers).
	_, err := e.Call("supports_language")
	var engineErr *EngineError
	if !errors.As(err, &engineErr) {
		t.Fatalf("expected EngineError, got %v", err)
	}
	if engineErr.Type != "value_error" {
		t.Fatalf("expected value_error, got %q (%s)", engineErr.Type, engineErr.Message)
	}
}
