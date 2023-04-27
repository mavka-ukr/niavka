// Code generated from NiavkaParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // NiavkaParser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// NiavkaParserListener is a complete listener for a parse tree produced by NiavkaParser.
type NiavkaParserListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterProgram_element is called when entering the program_element production.
	EnterProgram_element(c *Program_elementContext)

	// EnterDiia is called when entering the diia production.
	EnterDiia(c *DiiaContext)

	// EnterArithmetic_mul is called when entering the arithmetic_mul production.
	EnterArithmetic_mul(c *Arithmetic_mulContext)

	// EnterCall is called when entering the call production.
	EnterCall(c *CallContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterChain is called when entering the chain production.
	EnterChain(c *ChainContext)

	// EnterArithmetic_add is called when entering the arithmetic_add production.
	EnterArithmetic_add(c *Arithmetic_addContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterNested is called when entering the nested production.
	EnterNested(c *NestedContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterAssign_value is called when entering the assign_value production.
	EnterAssign_value(c *Assign_valueContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterType_value is called when entering the type_value production.
	EnterType_value(c *Type_valueContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterArg is called when entering the arg production.
	EnterArg(c *ArgContext)

	// EnterParams is called when entering the params production.
	EnterParams(c *ParamsContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterBody is called when entering the body production.
	EnterBody(c *BodyContext)

	// EnterBody_element is called when entering the body_element production.
	EnterBody_element(c *Body_elementContext)

	// EnterArithmetic_op_mul is called when entering the arithmetic_op_mul production.
	EnterArithmetic_op_mul(c *Arithmetic_op_mulContext)

	// EnterArithmetic_op_add is called when entering the arithmetic_op_add production.
	EnterArithmetic_op_add(c *Arithmetic_op_addContext)

	// EnterNl is called when entering the nl production.
	EnterNl(c *NlContext)

	// EnterNls is called when entering the nls production.
	EnterNls(c *NlsContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitProgram_element is called when exiting the program_element production.
	ExitProgram_element(c *Program_elementContext)

	// ExitDiia is called when exiting the diia production.
	ExitDiia(c *DiiaContext)

	// ExitArithmetic_mul is called when exiting the arithmetic_mul production.
	ExitArithmetic_mul(c *Arithmetic_mulContext)

	// ExitCall is called when exiting the call production.
	ExitCall(c *CallContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitChain is called when exiting the chain production.
	ExitChain(c *ChainContext)

	// ExitArithmetic_add is called when exiting the arithmetic_add production.
	ExitArithmetic_add(c *Arithmetic_addContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitNested is called when exiting the nested production.
	ExitNested(c *NestedContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitAssign_value is called when exiting the assign_value production.
	ExitAssign_value(c *Assign_valueContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitType_value is called when exiting the type_value production.
	ExitType_value(c *Type_valueContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitArg is called when exiting the arg production.
	ExitArg(c *ArgContext)

	// ExitParams is called when exiting the params production.
	ExitParams(c *ParamsContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitBody is called when exiting the body production.
	ExitBody(c *BodyContext)

	// ExitBody_element is called when exiting the body_element production.
	ExitBody_element(c *Body_elementContext)

	// ExitArithmetic_op_mul is called when exiting the arithmetic_op_mul production.
	ExitArithmetic_op_mul(c *Arithmetic_op_mulContext)

	// ExitArithmetic_op_add is called when exiting the arithmetic_op_add production.
	ExitArithmetic_op_add(c *Arithmetic_op_addContext)

	// ExitNl is called when exiting the nl production.
	ExitNl(c *NlContext)

	// ExitNls is called when exiting the nls production.
	ExitNls(c *NlsContext)
}
