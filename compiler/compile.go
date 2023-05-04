package compiler

import (
	"errors"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/value"
	"niavka/parser"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

func CompileProgramNode(program parser.ProgramNode) (*ir.Module, error) {
	m := ir.NewModule()
	puts := m.NewFunc("puts", types.I32, ir.NewParam("", types.NewPointer(types.I8)))
	globalTypes := map[string]types.Type{
		string(Boolean): Boolean.AsIr(),
		string(Int8):    Int8.AsIr(),
		string(Int16):   Int16.AsIr(),
		string(Int32):   Int32.AsIr(),
		string(Int64):   Int64.AsIr(),
		string(UInt8):   UInt8.AsIr(),
		string(UInt16):  UInt16.AsIr(),
		string(UInt32):  UInt32.AsIr(),
		string(UInt64):  UInt64.AsIr(),
		string(Float16): Float16.AsIr(),
		string(Float32): Float32.AsIr(),
		string(Float64): Float64.AsIr(),
		string(Str):     Str.AsIr(),
		string(Void):    Void.AsIr(),
	}
	for name, typ := range globalTypes {
		m.NewTypeDef(name, typ)
	}
	globalScope := &Scope{
		Vars: make(map[string][]Variable),
		Functions: map[string]*ir.Func{
			"друк": puts,
		},
		Types: globalTypes,
	}

	for _, element := range program.Elements {
		switch element := element.(type) {
		case parser.AssignNode:
			err := compileAssign(&element, globalScope, Module{m})
			if err != nil {
				return nil, err
			}
		case parser.DiiaNode:
			//err := compileDiia(&element, globalScope, m, nil)
			//if err != nil {
			//	return nil, err
			//}
			panic("not yet implemented")
		}
	}

	return m, nil
}

type Assignee interface {
	newAssignment(string, value.Value)
}

type Module struct{ *ir.Module }

func (m Module) newAssignment(name string, val value.Value) {
	switch val := val.(type) {
	case constant.Constant:
		m.NewGlobalDef(name, val)
	default:
		panic("not yet implemented")
	}
}

type Block struct{ *ir.Block }

func (b Block) newAssignment(name string, val value.Value) {
	panic("not yet implemented")
}

func compileAssign(assignment *parser.AssignNode, scope *Scope, assignee Assignee) error {
	if scope.GetVar(assignment.Left.Name) != nil {
		return errors.New("variable " + assignment.Left.Name + " already defined")
	}
	if assignment.Type == "" {
		panic("not yet implemented")
	}
	if scope.GetType(assignment.Type) == nil {
		return errors.New("type " + assignment.Type + " not defined")
	}
	typ, err := PrimitiveFromName(assignment.Type)
	if err != nil {
		return err
	}
	variable := Variable{
		Name: assignment.Left.Name,
		Type: typ,
	}
	scope.AddVar(variable)

	assignee.newAssignment(
		assignment.Left.Name,
		constant.NewInt(typ.AsIr().(*types.IntType), int64(assignment.Value.(parser.NumberNode).Value)),
	)

	return nil
}

//func getTypeByName(name string) types.Type {
//	switch name {
//	case "ц8":
//		return types.I8
//	case "ц16":
//		return types.I16
//	case "ц32":
//		return types.I32
//	case "ц64":
//		return types.I64
//	case "д16":
//		return types.Half
//	case "д32":
//		return types.Float
//	case "д64":
//		return types.Double
//	case "логічне":
//		return types.I1
//	case "текст":
//		return &types.VectorType{TypeName: "string", Scalable: true}
//	case "ніщо":
//		return types.Void
//	default:
//		return nil
//	}
//}
//
//func compileCall(call *parser.CallNode, scope *Scope, block *ir.Block) {
//	name := call.Value.(parser.IdentifierNode).Name
//	var funcVariable = scope.GetVar(name)
//	if funcVariable == nil {
//		panic(name + " is not defined")
//	}
//	var args []value.Value
//	for _, arg := range call.Args {
//		switch stringNode := arg.(type) {
//		case parser.StringNode:
//			hello := constant.NewCharArrayFromString(stringNode.Value + "\x00")
//			str := block.Parent.Parent.NewGlobalDef("str", hello)
//			zero := constant.NewInt(types.I64, 0)
//			gep := constant.NewGetElementPtr(hello.Typ, str, zero, zero)
//			args = append(args, gep)
//		}
//	}
//	block.NewCall(funcVariable.Func, args...)
//}
//
//func compileBody(body []interface{}, scope *Scope, block *ir.Block) {
//	for _, element := range body {
//		switch element := element.(type) {
//		case parser.AssignNode:
//			compileAssign(&element, scope, block)
//		case parser.CallNode:
//			compileCall(&element, scope, block)
//		}
//	}
//}
//
//func compileDiia(diia *parser.DiiaNode, scope *Scope, m *ir.Module) {
//	var diiaFuncParams []*ir.Param
//	for _, param := range diia.Params {
//		diiaFuncParams = append(diiaFuncParams, ir.NewParam(param.Name, getTypeByName(param.Type)))
//	}
//
//	name := diia.Name
//	if name == "запуск" {
//		name = "main"
//	}
//
//	diiaFunc := m.NewFunc(name, getTypeByName(diia.ReturnType), diiaFuncParams...)
//
//	diiaFuncBlock := diiaFunc.NewBlock("")
//	compileBody(diia.Body, scope, diiaFuncBlock)
//
//	diiaFuncBlock.NewRet(nil)
//
//	scope.AddVar(Variable{
//		Name: name,
//		Type: diia.ReturnType,
//		Func: diiaFunc,
//	})
//}
//
//func compileAssign(assign *parser.AssignNode, scope *Scope, m *ir.Block) {
//	variable := Variable{
//		Name: assign.Left.Name,
//		Type: assign.Type,
//		Func: scope.GetVar("").Func,
//	}
//
//
//	scope.AddVar(variable)
//}
