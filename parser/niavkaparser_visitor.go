// Code generated from NiavkaParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // NiavkaParser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by NiavkaParser.
type NiavkaParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by NiavkaParser#file.
	VisitFile(ctx *FileContext) interface{}

	// Visit a parse tree produced by NiavkaParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by NiavkaParser#program_element.
	VisitProgram_element(ctx *Program_elementContext) interface{}

	// Visit a parse tree produced by NiavkaParser#diia.
	VisitDiia(ctx *DiiaContext) interface{}

	// Visit a parse tree produced by NiavkaParser#arithmetic_mul.
	VisitArithmetic_mul(ctx *Arithmetic_mulContext) interface{}

	// Visit a parse tree produced by NiavkaParser#call.
	VisitCall(ctx *CallContext) interface{}

	// Visit a parse tree produced by NiavkaParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by NiavkaParser#chain.
	VisitChain(ctx *ChainContext) interface{}

	// Visit a parse tree produced by NiavkaParser#arithmetic_add.
	VisitArithmetic_add(ctx *Arithmetic_addContext) interface{}

	// Visit a parse tree produced by NiavkaParser#string.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by NiavkaParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by NiavkaParser#nested.
	VisitNested(ctx *NestedContext) interface{}

	// Visit a parse tree produced by NiavkaParser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by NiavkaParser#assign_value.
	VisitAssign_value(ctx *Assign_valueContext) interface{}

	// Visit a parse tree produced by NiavkaParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by NiavkaParser#type_value.
	VisitType_value(ctx *Type_valueContext) interface{}

	// Visit a parse tree produced by NiavkaParser#args.
	VisitArgs(ctx *ArgsContext) interface{}

	// Visit a parse tree produced by NiavkaParser#arg.
	VisitArg(ctx *ArgContext) interface{}

	// Visit a parse tree produced by NiavkaParser#params.
	VisitParams(ctx *ParamsContext) interface{}

	// Visit a parse tree produced by NiavkaParser#param.
	VisitParam(ctx *ParamContext) interface{}

	// Visit a parse tree produced by NiavkaParser#body.
	VisitBody(ctx *BodyContext) interface{}

	// Visit a parse tree produced by NiavkaParser#body_element.
	VisitBody_element(ctx *Body_elementContext) interface{}

	// Visit a parse tree produced by NiavkaParser#arithmetic_op_mul.
	VisitArithmetic_op_mul(ctx *Arithmetic_op_mulContext) interface{}

	// Visit a parse tree produced by NiavkaParser#arithmetic_op_add.
	VisitArithmetic_op_add(ctx *Arithmetic_op_addContext) interface{}

	// Visit a parse tree produced by NiavkaParser#nl.
	VisitNl(ctx *NlContext) interface{}

	// Visit a parse tree produced by NiavkaParser#nls.
	VisitNls(ctx *NlsContext) interface{}
}
