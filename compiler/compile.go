package compiler

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"niavka/parser"
)

func getTypeByName(name string) types.Type {
	if name == "ц32" {
		return types.I32
	} else if name == "ц64" {
		return types.I64
	} else if name == "д32" {
		return types.Float
	} else if name == "д64" {
		return types.Double
	} else if name == "текст" {
		return types.NewPointer(types.I8)
	} else if name == "ніщо" {
		return types.Void
	}

	return nil
}

func compileBody(block *ir.Block, body []interface{}) {
	for _, element := range body {
		switch element.(type) {
		case parser.AssignNode:
			panic("not implemented")
		case parser.CallNode:
			call := element.(parser.CallNode)
			compileCall(block, call)
		}
	}
}

func compileCall(block *ir.Block, call parser.CallNode) {
	var funcToCall *ir.Func
	for _, element := range block.Parent.Parent.Funcs {
		if element.Name() == call.Value.(parser.IdentifierNode).Name {
			funcToCall = element
		}
	}
	var args []value.Value
	for _, arg := range call.Args {
		switch arg.(type) {
		case parser.StringNode:
			stringNode := arg.(parser.StringNode)
			hello := constant.NewCharArrayFromString(stringNode.Value + "\x00")
			str := block.Parent.Parent.NewGlobalDef("str", hello)
			zero := constant.NewInt(types.I64, 0)
			gep := constant.NewGetElementPtr(hello.Typ, str, zero, zero)
			args = append(args, gep)
		}
	}
	block.NewCall(funcToCall, args...)
}

func CompileProgramNode(program parser.ProgramNode) *ir.Module {
	m := ir.NewModule()
	m.NewFunc("puts", types.I32, ir.NewParam("", types.NewPointer(types.I8)))

	for _, element := range program.Elements {
		switch element.(type) {
		case parser.DiiaNode:
			diia := element.(parser.DiiaNode)
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
			compileBody(diiaFuncBlock, diia.Body)
			diiaFuncBlock.NewRet(nil)
		}
	}

	return m
}
