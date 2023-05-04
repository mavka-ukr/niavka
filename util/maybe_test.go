package util

import (
	"testing"
)

func TestNewMaybe(t *testing.T) {
	m := NewMaybe(10)
	if m.Absent() {
		t.Error("should be present")
	}
	if v := m.OrPanic(); v != 10 {
		t.Error("should be 10")
	}
}

func TestMaybe_Set(t *testing.T) {
	var m Maybe[int]
	m.Set(5)
	if v, err := m.Get(); v != 5 || err != nil {
		t.Error(err)
	}
}

func TestMaybe_OrPanic(t *testing.T) {
	var m Maybe[int]
	m.Set(5)
	if v := m.OrPanic(); v != 5 {
		t.Error("value is not 5")
	}
}

func TestMaybe_OrPanic2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("should panic")
		}
	}()
	var m Maybe[int]
	m.OrPanic()
}

func TestMaybe_Or(t *testing.T) {
	var m Maybe[int]
	m.Set(5)
	var m2 Maybe[int]
	m2.Set(6)
	if v := m.Or(m2); v != m {
		t.Error("should be m")
	}
}

func TestMaybe_Or2(t *testing.T) {
	var m Maybe[int]
	var m2 Maybe[int]
	m2.Set(6)
	if v := m.Or(m2); v != m2 {
		t.Error("should be m2")
	}
}

func TestMaybe_IfPresent(t *testing.T) {
	var m Maybe[int]
	m.Set(5)
	passed := false
	m.IfPresent(func(v int) {
		if v != 5 {
			t.Error("should be 5")
		}
		passed = true
	})
	if !passed {
		t.Fail()
	}
}

func TestMaybe_IfPresent2(t *testing.T) {
	var m Maybe[int]
	passed := true
	m.IfPresent(func(v int) {
		passed = false
	})
	if !passed {
		t.Fail()
	}
}

func TestMaybe_IfAbsent(t *testing.T) {
	var m Maybe[int]
	passed := false
	m.IfAbsent(func() {
		passed = true
	})
	if !passed {
		t.Fail()
	}
}

func TestMaybe_IfAbsent2(t *testing.T) {
	var m Maybe[int]
	m.Set(5)
	passed := true
	m.IfAbsent(func() {
		passed = false
	})
	if !passed {
		t.Fail()
	}
}
