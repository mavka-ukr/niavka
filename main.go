package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"niavka/parser"
	"os"
	"strconv"
)

type NiavkaVisitor struct {
	parser.BaseNiavkaParserVisitor
}

func (v *NiavkaVisitor) Visit(tree antlr.Tree) interface{} {
	switch val := tree.(type) {
	case *parser.FileContext:
		return v.VisitFile(val)
	case *parser.ProgramContext:
		return v.VisitProgram(val)
	case *parser.Program_elementContext:
		return v.VisitProgram_element(val)
	case *parser.DiiaContext:
		return v.VisitDiia(val)
	case *parser.NumberContext:
		return v.VisitNumber(val)
	case *parser.StringContext:
		return v.VisitString(val)
	case *parser.IdContext:
		return v.VisitId(val)
	case *parser.ChainContext:
		return v.VisitChain(val)
	case *parser.CallContext:
		return v.VisitCall(val)
	case *parser.NestedContext:
		return v.VisitNested(val)
	case *parser.Arithmetic_mulContext:
		return v.VisitArithmetic_mul(val)
	case *parser.Arithmetic_addContext:
		return v.VisitArithmetic_add(val)
	case *parser.AssignContext:
		return v.VisitAssign(val)
	case *parser.Assign_valueContext:
		return v.VisitAssign_value(val)
	case *parser.IdentifierContext:
		return v.VisitIdentifier(val)
	case *parser.Type_valueContext:
		return v.VisitType_value(val)
	case *parser.ArgsContext:
		return v.VisitArgs(val)
	case *parser.ArgContext:
		return v.VisitArg(val)
	case *parser.ParamsContext:
		return v.VisitParams(val)
	case *parser.ParamContext:
		return v.VisitParam(val)
	case *parser.BodyContext:
		return v.VisitBody(val)
	case *parser.Body_elementContext:
		return v.VisitBody_element(val)
	default:
		return nil
	}
}

func (v *NiavkaVisitor) VisitFile(ctx parser.IFileContext) parser.ProgramNode {
	return v.VisitProgram(ctx.GetF_program())
}

func (v *NiavkaVisitor) VisitProgram(ctx parser.IProgramContext) parser.ProgramNode {
	programNode := parser.ProgramNode{}
	var elements []interface{}
	for _, element := range ctx.AllProgram_element() {
		programElementNode := v.Visit(element)
		if programElementNode != nil {
			elements = append(elements, programElementNode)
		}
	}
	programNode.Elements = elements
	return programNode
}

func (v *NiavkaVisitor) VisitProgram_element(ctx parser.IProgram_elementContext) interface{} {
	return v.Visit(ctx.GetChild(0))
}

func (v *NiavkaVisitor) VisitDiia(ctx parser.IDiiaContext) parser.DiiaNode {
	var params []parser.ParamNode
	paramsContext := ctx.GetD_params()
	if paramsContext != nil {
		params = v.VisitParams(ctx.GetD_params())
	} else {
		params = []parser.ParamNode{}
	}

	returnType := "ніщо"
	returnTypeContext := ctx.GetD_type()
	if returnTypeContext != nil {
		returnType = returnTypeContext.GetText()
	}

	var body []interface{}
	bodyContext := ctx.GetD_body()
	if bodyContext != nil {
		body = v.VisitBody(bodyContext)
	} else {
		body = []interface{}{}
	}
	return parser.DiiaNode{
		Name:       ctx.GetD_name().GetText(),
		Params:     params,
		ReturnType: returnType,
		Body:       body,
	}
}

func (v *NiavkaVisitor) VisitNumber(ctx *parser.NumberContext) parser.NumberNode {
	value, _ := strconv.Atoi(ctx.GetText())
	return parser.NumberNode{Value: value}
}

func (v *NiavkaVisitor) VisitString(ctx *parser.StringContext) parser.StringNode {
	text := ctx.GetText()
	return parser.StringNode{Value: text[1 : len(text)-1]}
}

func (v *NiavkaVisitor) VisitId(ctx *parser.IdContext) interface{} {
	return v.Visit(ctx.GetChild(0))
}

func (v *NiavkaVisitor) VisitChain(ctx *parser.ChainContext) parser.ChainNode {
	return parser.ChainNode{
		Left:  v.Visit(ctx.GetC_left()),
		Right: v.Visit(ctx.GetC_right()),
	}
}

func (v *NiavkaVisitor) VisitCall(ctx *parser.CallContext) parser.CallNode {
	var args []interface{}
	argsContext := ctx.GetC_args()
	if argsContext != nil {
		args = v.VisitArgs(argsContext)
	}
	return parser.CallNode{
		Value: v.Visit(ctx.GetC_value()),
		Args:  args,
	}
}

func (v *NiavkaVisitor) VisitNested(ctx *parser.NestedContext) interface{} {
	return v.Visit(ctx.GetN_value())
}

func (v *NiavkaVisitor) VisitArithmetic_mul(ctx *parser.Arithmetic_mulContext) parser.ArithmeticNode {
	return parser.ArithmeticNode{
		Left:  v.Visit(ctx.GetA_left()),
		Right: v.Visit(ctx.GetA_right()),
		Op:    ctx.GetA_operation().GetText(),
	}
}

func (v *NiavkaVisitor) VisitArithmetic_add(ctx *parser.Arithmetic_addContext) parser.ArithmeticNode {
	return parser.ArithmeticNode{
		Left:  v.Visit(ctx.GetA_left()),
		Right: v.Visit(ctx.GetA_right()),
		Op:    ctx.GetA_operation().GetText(),
	}
}

func (v *NiavkaVisitor) VisitAssign(ctx parser.IAssignContext) parser.AssignNode {
	return parser.AssignNode{
		Left:  v.Visit(ctx.GetA_left()),
		Type:  v.VisitType_value(ctx.GetA_type()),
		Value: v.Visit(ctx.GetA_value()),
	}
}

func (v *NiavkaVisitor) VisitAssign_value(ctx parser.IAssign_valueContext) interface{} {
	return v.Visit(ctx.GetChild(0))
}

func (v *NiavkaVisitor) VisitIdentifier(ctx parser.IIdentifierContext) parser.IdentifierNode {
	return parser.IdentifierNode{Name: ctx.GetText()}
}

func (v *NiavkaVisitor) VisitType_value(ctx parser.IType_valueContext) string {
	return ctx.GetTv_single().GetText()
}

func (v *NiavkaVisitor) VisitArgs(ctx parser.IArgsContext) []interface{} {
	var args []interface{}
	for _, arg := range ctx.AllArg() {
		a := v.Visit(arg)
		if a != nil {
			args = append(args, a)
		}
	}
	return args
}

func (v *NiavkaVisitor) VisitArg(ctx parser.IArgContext) interface{} {
	return v.Visit(ctx.GetA_value())
}

func (v *NiavkaVisitor) VisitParams(ctx parser.IParamsContext) []parser.ParamNode {
	var params []parser.ParamNode
	for _, param := range ctx.AllParam() {
		p := v.Visit(param)
		if p != nil {
			params = append(params, p.(parser.ParamNode))
		}
	}
	return params
}

func (v *NiavkaVisitor) VisitParam(ctx parser.IParamContext) interface{} {
	return parser.ParamNode{
		Name: ctx.GetP_name().GetText(),
		Type: v.VisitType_value(ctx.GetP_type()),
	}
}

func (v *NiavkaVisitor) VisitBody(ctx parser.IBodyContext) []interface{} {
	var body []interface{}
	for _, element := range ctx.AllBody_element() {
		el := v.Visit(element)
		if el != nil {
			body = append(body, el)
		}
	}
	return body
}

func (v *NiavkaVisitor) VisitBody_element(ctx parser.IBody_elementContext) interface{} {
	return v.Visit(ctx.GetChild(0))
}

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

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewNiavkaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewNiavkaParser(stream)
	p.BuildParseTrees = true
	tree := p.File()
	visitor := NiavkaVisitor{}
	var result = visitor.Visit(tree).(parser.ProgramNode)

	m := ir.NewModule()
	m.NewFunc("puts", types.I32, ir.NewParam("", types.NewPointer(types.I8)))

	for _, element := range result.Elements {
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

	fmt.Println(m)
}
