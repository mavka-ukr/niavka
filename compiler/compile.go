package compiler

import (
	"niavka/parser"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func getTypeByName(name string) types.Type {
	switch name {
	case "ц8":
		return types.I8
	case "ц16":
		return types.I16
	case "ц32":
		return types.I32
	case "ц64":
		return types.I64
	case "д16":
		return types.Half
	case "д32":
		return types.Float
	case "д64":
		return types.Double
	case "логічне":
		return types.I1
	case "текст":
		return types.I8Ptr
	case "ніщо":
		return types.Void
	default:
		return nil
	}
}

func compileCall(call *parser.CallNode, scope *Scope, block *ir.Block) {
	name := call.Value.(parser.IdentifierNode).Name
	var funcVariable = scope.Get(name)
	if funcVariable == nil {
		panic(name + " is not defined")
	}
	var args []value.Value
	for _, arg := range call.Args {
		switch stringNode := arg.(type) {
		case parser.StringNode:
			hello := constant.NewCharArrayFromString(stringNode.Value + "\x00")
			str := block.Parent.Parent.NewGlobalDef("str", hello)
			zero := constant.NewInt(types.I64, 0)
			gep := constant.NewGetElementPtr(hello.Typ, str, zero, zero)
			args = append(args, gep)
		}
	}
	block.NewCall(funcVariable.Func, args...)
}

func compileBody(body []interface{}, scope *Scope, block *ir.Block) {
	for _, element := range body {
		switch el := element.(type) {
		case parser.AssignNode:
			panic("not implemented")
		case parser.CallNode:
			compileCall(&el, scope, block)
		}
	}
}

func compileDiia(diia *parser.DiiaNode, scope *Scope, m *ir.Module) {
	var diiaFuncParams []*ir.Param
	for _, param := range diia.Params {
		diiaFuncParams = append(diiaFuncParams, ir.NewParam(param.Name, getTypeByName(param.Type)))
	}

	name := diia.Name
	if name == "запуск" {
		name = "main"
	}

	diiaFunc := m.NewFunc(name, getTypeByName(diia.ReturnType), diiaFuncParams...)

	diiaFuncBlock := diiaFunc.NewBlock("")
	compileBody(diia.Body, scope, diiaFuncBlock)

	diiaFuncBlock.NewRet(nil)

	scope.Set(name, Variable{
		Name: name,
		Type: diia.ReturnType,
		Func: diiaFunc,
	})
}

func CompileProgramNode(program parser.ProgramNode) *ir.Module {
	m := ir.NewModule()
	puts := m.NewFunc("puts", types.I32, ir.NewParam("", types.NewPointer(types.I8)))

	scope := &Scope{
		Parent: nil,
		Vars:   make(map[string]Variable),
	}

	scope.Set("друк", Variable{
		Name: "друк",
		Type: "Дія",
		Func: puts,
	})

	for _, element := range program.Elements {
		switch diia := element.(type) {
		case parser.DiiaNode:
			compileDiia(&diia, scope, m)
		}
	}

	return m
}
