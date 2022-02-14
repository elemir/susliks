package susliks

import (
    "testing"
)

func TestMaybe(t *testing.T) {
    nothing := Nothing[string]().Get()

    if nothing != "" {
        t.Fatalf("expected to get empty string from Nothing but got '%s'", nothing)
    }

    maybe, ok := Just(10).Unwrap()
    if maybe != 10 || !ok {
        t.Fatalf("expected to get (10, true) from Just(10) but got (%d, %t)", maybe, ok)
    }
}
