package compiler

import "github.com/llir/llvm/ir"

// Variable Represents variable in scope
type Variable struct {
	Name string
	Type string
	Func *ir.Func
}

// Scope Represents scope of a code block
type Scope struct {
	Parent *Scope
	Vars   map[string]Variable
}

// Get Gets variable from current scope or parent scope
func (s *Scope) Get(name string) *Variable {
	if _, ok := s.Vars[name]; ok {
		v := s.Vars[name]
		return &v
	}
	if s.Parent != nil {
		return s.Parent.Get(name)
	}
	return nil
}

// Set Sets variable in current scope
func (s *Scope) Set(name string, v Variable) {
	s.Vars[name] = v
}

// SetInParent Sets variable in parent scope
func (s *Scope) SetInParent(name string, v Variable) {
	if s.Parent == nil {
		panic("no parent scope")
	}
	s.Parent.Set(name, v)
}
