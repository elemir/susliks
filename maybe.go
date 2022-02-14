// Package susliks provides implementations of several type-parametrized containers as optional value, balanced sorted binary trees, and compacted prefix trees
package susliks

// Maybe encapsulates an optional value
type Maybe[T any] struct {
    contains bool
    value    T
}

// Just wraps value to corresponding Maybe type
func Just[T any](v T) Maybe[T] {
    return Maybe[T]{
        contains: true,
        value:    v,
    }
}

// Nothing returns an empty value of Maybe type
func Nothing[T any]() Maybe[T] {
    return Maybe[T]{
        contains: false,
    }
}

// Unwrap returns an contained value and emptiness
func (m Maybe[T]) Unwrap() (T, bool) {
    return m.value, m.contains
}

// Get return an contained value or zero value if it empty
func (m Maybe[T]) Get() T {
    return m.value
}
