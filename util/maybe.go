package util

import (
	"errors"
)

type Maybe[T any] struct {
	Value *T
}

func NewMaybe[T any](value T) Maybe[T] {
	var m Maybe[T]
	m.Set(value)
	return m
}

func (m *Maybe[T]) Set(value T) {
	m.Value = &value
}

func (m *Maybe[T]) Absent() bool {
	return !m.Present()
}

func (m *Maybe[T]) Present() bool {
	return m.Value != nil
}

func (m *Maybe[T]) Get() (T, error) {
	if m.Absent() {
		var zero T
		return zero, errors.New("value is absent")
	}
	return *m.Value, nil
}

func (m *Maybe[T]) GetOrDefault(defaultValue T) T {
	if m.Absent() {
		return defaultValue
	}
	return *m.Value
}

func (m *Maybe[T]) GetOrElse(f func() T) T {
	if m.Absent() {
		return f()
	}
	return *m.Value
}

func (m *Maybe[T]) Or(other Maybe[T]) Maybe[T] {
	if m.Absent() {
		return other
	}
	return *m
}

func (m *Maybe[T]) OrElse(f func() Maybe[T]) Maybe[T] {
	if m.Absent() {
		return f()
	}
	return *m
}

func (m *Maybe[T]) OrPanic() T {
	if m.Absent() {
		panic("value is absent")
	}
	return *m.Value
}

func (m *Maybe[T]) IfPresent(f func(T)) {
	if m.Present() {
		f(*m.Value)
	}
}

func (m *Maybe[T]) IfAbsent(f func()) {
	if m.Absent() {
		f()
	}
}

func (m *Maybe[T]) Filter(f func(T) bool) Maybe[T] {
	if m.Absent() {
		return *m
	}
	if f(*m.Value) {
		return *m
	}
	return Maybe[T]{}
}
