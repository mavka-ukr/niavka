package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"strconv"
)

type NiavkaVisitor struct {
	BaseNiavkaParserVisitor
}

func (v *NiavkaVisitor) Visit(tree antlr.Tree) interface{} {
	switch val := tree.(type) {
	case *FileContext:
		return v.VisitFile(val)
	case *ProgramContext:
		return v.VisitProgram(val)
	case *Program_elementContext:
		return v.VisitProgram_element(val)
	case *DiiaContext:
		return v.VisitDiia(val)
	case *NumberContext:
		return v.VisitNumber(val)
	case *StringContext:
		return v.VisitString(val)
	case *IdContext:
		return v.VisitId(val)
	case *ChainContext:
		return v.VisitChain(val)
	case *CallContext:
		return v.VisitCall(val)
	case *NestedContext:
		return v.VisitNested(val)
	case *Arithmetic_mulContext:
		return v.VisitArithmetic_mul(val)
	case *Arithmetic_addContext:
		return v.VisitArithmetic_add(val)
	case *AssignContext:
		return v.VisitAssign(val)
	case *Assign_valueContext:
		return v.VisitAssign_value(val)
	case *IdentifierContext:
		return v.VisitIdentifier(val)
	case *Type_valueContext:
		return v.VisitType_value(val)
	case *ArgsContext:
		return v.VisitArgs(val)
	case *ArgContext:
		return v.VisitArg(val)
	case *ParamsContext:
		return v.VisitParams(val)
	case *ParamContext:
		return v.VisitParam(val)
	case *BodyContext:
		return v.VisitBody(val)
	case *Body_elementContext:
		return v.VisitBody_element(val)
	default:
		return nil
	}
}

func (v *NiavkaVisitor) VisitFile(ctx IFileContext) ProgramNode {
	return v.VisitProgram(ctx.GetF_program())
}

func (v *NiavkaVisitor) VisitProgram(ctx IProgramContext) ProgramNode {
	programNode := ProgramNode{}
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

func (v *NiavkaVisitor) VisitProgram_element(ctx IProgram_elementContext) interface{} {
	return v.Visit(ctx.GetChild(0))
}

func (v *NiavkaVisitor) VisitDiia(ctx IDiiaContext) DiiaNode {
	var params []ParamNode
	paramsContext := ctx.GetD_params()
	if paramsContext != nil {
		params = v.VisitParams(ctx.GetD_params())
	} else {
		params = []ParamNode{}
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
	return DiiaNode{
		Name:       ctx.GetD_name().GetText(),
		Params:     params,
		ReturnType: returnType,
		Body:       body,
	}
}

func (v *NiavkaVisitor) VisitNumber(ctx *NumberContext) NumberNode {
	value, _ := strconv.Atoi(ctx.GetText())
	return NumberNode{Value: value}
}

func (v *NiavkaVisitor) VisitString(ctx *StringContext) StringNode {
	text := ctx.GetText()
	return StringNode{Value: text[1 : len(text)-1]}
}

func (v *NiavkaVisitor) VisitId(ctx *IdContext) interface{} {
	return v.Visit(ctx.GetChild(0))
}

func (v *NiavkaVisitor) VisitChain(ctx *ChainContext) ChainNode {
	return ChainNode{
		Left:  v.Visit(ctx.GetC_left()),
		Right: v.Visit(ctx.GetC_right()),
	}
}

func (v *NiavkaVisitor) VisitCall(ctx *CallContext) CallNode {
	var args []interface{}
	argsContext := ctx.GetC_args()
	if argsContext != nil {
		args = v.VisitArgs(argsContext)
	}
	return CallNode{
		Value: v.Visit(ctx.GetC_value()),
		Args:  args,
	}
}

func (v *NiavkaVisitor) VisitNested(ctx *NestedContext) interface{} {
	return v.Visit(ctx.GetN_value())
}

func (v *NiavkaVisitor) VisitArithmetic_mul(ctx *Arithmetic_mulContext) ArithmeticNode {
	return ArithmeticNode{
		Left:  v.Visit(ctx.GetA_left()),
		Right: v.Visit(ctx.GetA_right()),
		Op:    ctx.GetA_operation().GetText(),
	}
}

func (v *NiavkaVisitor) VisitArithmetic_add(ctx *Arithmetic_addContext) ArithmeticNode {
	return ArithmeticNode{
		Left:  v.Visit(ctx.GetA_left()),
		Right: v.Visit(ctx.GetA_right()),
		Op:    ctx.GetA_operation().GetText(),
	}
}

func (v *NiavkaVisitor) VisitAssign(ctx IAssignContext) AssignNode {
	return AssignNode{
		Left:  v.Visit(ctx.GetA_left()),
		Type:  v.VisitType_value(ctx.GetA_type()),
		Value: v.Visit(ctx.GetA_value()),
	}
}

func (v *NiavkaVisitor) VisitAssign_value(ctx IAssign_valueContext) interface{} {
	return v.Visit(ctx.GetChild(0))
}

func (v *NiavkaVisitor) VisitIdentifier(ctx IIdentifierContext) IdentifierNode {
	return IdentifierNode{Name: ctx.GetText()}
}

func (v *NiavkaVisitor) VisitType_value(ctx IType_valueContext) string {
	return ctx.GetTv_single().GetText()
}

func (v *NiavkaVisitor) VisitArgs(ctx IArgsContext) []interface{} {
	var args []interface{}
	for _, arg := range ctx.AllArg() {
		a := v.Visit(arg)
		if a != nil {
			args = append(args, a)
		}
	}
	return args
}

func (v *NiavkaVisitor) VisitArg(ctx IArgContext) interface{} {
	return v.Visit(ctx.GetA_value())
}

func (v *NiavkaVisitor) VisitParams(ctx IParamsContext) []ParamNode {
	var params []ParamNode
	for _, param := range ctx.AllParam() {
		p := v.Visit(param)
		if p != nil {
			params = append(params, p.(ParamNode))
		}
	}
	return params
}

func (v *NiavkaVisitor) VisitParam(ctx IParamContext) interface{} {
	return ParamNode{
		Name: ctx.GetP_name().GetText(),
		Type: v.VisitType_value(ctx.GetP_type()),
	}
}

func (v *NiavkaVisitor) VisitBody(ctx IBodyContext) []interface{} {
	var body []interface{}
	for _, element := range ctx.AllBody_element() {
		el := v.Visit(element)
		if el != nil {
			body = append(body, el)
		}
	}
	return body
}

func (v *NiavkaVisitor) VisitBody_element(ctx IBody_elementContext) interface{} {
	return v.Visit(ctx.GetChild(0))
}
