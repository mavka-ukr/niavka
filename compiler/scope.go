package compiler

import (
	"errors"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

// Scope Represents scope of a code block
type Scope struct {
	Parent    *Scope
	Vars      map[string][]Variable // variables can be shadowed
	Types     map[string]types.Type
	Functions map[string]*ir.Func
}

func (s *Scope) IsGlobal() bool {
	return s.Parent == nil
}

// GetVar from current scope or parent scope
func (s *Scope) GetVar(name string) *Variable {
	if v, ok := s.Vars[name]; ok && len(v) > 0 {
		return &v[len(v)-1]
	}
	if !s.IsGlobal() {
		return s.Parent.GetVar(name)
	}
	return nil
}

// GetType from current scope or parent scope
func (s *Scope) GetType(name string) types.Type {
	if t, ok := s.Types[name]; ok {
		return t
	}
	if !s.IsGlobal() {
		return s.Parent.GetType(name)
	}
	return nil
}

// GetFunction from current scope or parent scope
func (s *Scope) GetFunction(name string) *ir.Func {
	if f, ok := s.Functions[name]; ok {
		return f
	}
	if !s.IsGlobal() {
		return s.Parent.GetFunction(name)
	}
	return nil
}

// AddVar to the current scope
func (s *Scope) AddVar(v Variable) {
	if s.Vars == nil {
		s.Vars = make(map[string][]Variable)
	}
	varsWithName, ok := s.Vars[v.Name]
	if ok {
		s.Vars[v.Name] = append(varsWithName, v)
		return
	}
	varsWithName = make([]Variable, 0, 1)
	varsWithName = append(varsWithName, v)
	s.Vars[v.Name] = varsWithName
}

// SetVarInParent Sets the last variable in parent scope
func (s *Scope) SetVarInParent(v Variable) error {
	if s.IsGlobal() {
		return errors.New("no parent scope")
	}
	varsWithName, ok := s.Parent.Vars[v.Name]
	if !ok || len(varsWithName) == 0 {
		s.Parent.AddVar(v)
		return nil
	}
	varsWithName[len(varsWithName)-1] = v
	return nil
}

func (s *Scope) SetType(name string, t types.Type) {
	if s.Types == nil {
		s.Types = make(map[string]types.Type)
	}
	s.Types[name] = t
}

func (s *Scope) SetFunction(name string, f *ir.Func) {
	if s.Functions == nil {
		s.Functions = make(map[string]*ir.Func)
	}
	s.Functions[name] = f
}
