// Code generated from NiavkaParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // NiavkaParser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseNiavkaParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseNiavkaParserVisitor) VisitFile(ctx *FileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNiavkaParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNiavkaParserVisitor) VisitProgram_element(ctx *Program_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNiavkaParserVisitor) VisitDiia(ctx *DiiaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNiavkaParserVisitor) VisitArithmetic_mul(ctx *Arithmetic_mulContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNiavkaParserVisitor) VisitCall(ctx *CallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNiavkaParserVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNiavkaParserVisitor) VisitChain(ctx *ChainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNiavkaParserVisitor) VisitArithmetic_add(ctx *Arithmetic_addContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNiavkaParserVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNiavkaParserVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNiavkaParserVisitor) VisitNested(ctx *NestedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNiavkaParserVisitor) VisitAssign(ctx *AssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNiavkaParserVisitor) VisitAssign_value(ctx *Assign_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNiavkaParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNiavkaParserVisitor) VisitType_value(ctx *Type_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNiavkaParserVisitor) VisitArgs(ctx *ArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNiavkaParserVisitor) VisitArg(ctx *ArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNiavkaParserVisitor) VisitParams(ctx *ParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNiavkaParserVisitor) VisitParam(ctx *ParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNiavkaParserVisitor) VisitBody(ctx *BodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNiavkaParserVisitor) VisitBody_element(ctx *Body_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNiavkaParserVisitor) VisitArithmetic_op_mul(ctx *Arithmetic_op_mulContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNiavkaParserVisitor) VisitArithmetic_op_add(ctx *Arithmetic_op_addContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNiavkaParserVisitor) VisitNl(ctx *NlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNiavkaParserVisitor) VisitNls(ctx *NlsContext) interface{} {
	return v.VisitChildren(ctx)
}
