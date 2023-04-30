package compiler

import "github.com/llir/llvm/ir"

type Variable struct {
	Name string
	Type string
	Func *ir.Func
}

type Scope struct {
	Parent *Scope
	Vars   map[string]Variable
}

func (s Scope) Get(name string) *Variable {
	if _, ok := s.Vars[name]; ok {
		v := s.Vars[name]
		return &v
	}
	if s.Parent != nil {
		return s.Parent.Get(name)
	}
	return nil
}

func (s Scope) Set(name string, v Variable) {
	s.Vars[name] = v
}

func (s Scope) SetParent(name string, v Variable) {
	if s.Parent == nil {
		panic("no parent scope")
	}
	s.Parent.Set(name, v)
}
